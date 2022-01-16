package vote

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	// defaultBaseURL is the default URL used as voting website.
	defaultBaseURL = "https://minecraftpocket-servers.com/api/"
)

// Client is a client used to access the API of a voting website. Client has methods that may be used to get
// information about votes, voters and general status and statistics of the server.
type Client struct {
	// BaseURL is the URL that points to the API of the voting website. By default, this field is set to
	// https://minecraftpocket-servers.com/api/.
	BaseURL *url.URL

	key string
	c   *http.Client
}

// NewClient creates a new Client struct which is used to fetch information from minecraftpocket-servers.com.
func NewClient(key string) *Client {
	base, err := url.Parse(defaultBaseURL)
	if err != nil {
		// Should never happen.
		panic(err)
	}
	return &Client{BaseURL: base, key: key, c: &http.Client{}}
}

// Voted checks if the username provided has voted within the last 24 hours. It returns true if a player with the
// username voted regardless of whether the vote was claimed.
func (c *Client) Voted(username string) (bool, error) {
	data, err := c.get(map[string]string{
		"object":   "votes",
		"element":  "claim",
		"username": username,
	})
	if err != nil {
		return false, err
	}
	return parseInt(string(data)) != 0, nil
}

// Claim attempts to claim the vote of the provided username. Claim returns false both if the player with that name has
// not voted or if it has already had its vote claimed.
func (c *Client) Claim(name string) (bool, error) {
	data, err := c.get(map[string]string{
		"action":   "post",
		"object":   "votes",
		"element":  "claim",
		"username": name,
	})
	if err != nil {
		return false, err
	}
	return string(data) == "1", nil
}

// Voters returns up to n top voters of the current month. The limit n can be anywhere in the range of 1-1000.
func (c *Client) Voters(n int) ([]Voter, error) {
	return c.voters("current", n)
}

// VotersPreviousMonth returns up to n top voters of the previous month. The limit n can be anywhere in the range of
// 1-1000.
func (c *Client) VotersPreviousMonth(n int) ([]Voter, error) {
	return c.voters("previous", n)
}

// voters returns up to n top voters of the month passed ("current" or "previous"). The limit n can be anywhere in the
// range of 1-1000.
func (c *Client) voters(month string, n int) ([]Voter, error) {
	if n < 1 || n > 1000 {
		return nil, errors.New("n must be in the range of 1-1000")
	}
	data, err := c.get(map[string]string{
		"object":  "servers",
		"element": "voters",
		"month":   month,
		"format":  "json",
		"limit":   strconv.Itoa(n),
	})
	if err != nil {
		return nil, err
	}
	return readVoters(data)
}

// Votes returns the n most recent votes and the information about them for the current month. The limit n can be
// anywhere in the range of 1-1000.
func (c *Client) Votes(n int) ([]Vote, error) {
	if n < 1 || n > 1000 {
		return nil, errors.New("n must be in the range of 1-1000")
	}
	data, err := c.get(map[string]string{
		"object":  "servers",
		"element": "votes",
		"format":  "json",
		"limit":   strconv.Itoa(n),
	})
	if err != nil {
		return nil, err
	}
	return readVotes(data)
}

// ServerStatus returns the status of the server connected to the voting key passed. It contains information such as
// the vote count, comment count and total score.
func (c *Client) ServerStatus() (ServerStatus, error) {
	data, err := c.get(map[string]string{
		"object":  "servers",
		"element": "detail",
	})
	if err != nil {
		return ServerStatus{}, err
	}
	return readServerStatus(data)
}

// get performs a GET request with the key and query strings passed added to the base URL of the Client. The method
// returns the body of the response as a byte slice or an error if the http request was unsuccessful.
func (c *Client) get(queryStrings map[string]string) ([]byte, error) {
	parsed, err := c.BaseURL.Parse("?key=" + c.key)
	if err != nil {
		// Should never happen.
		panic(err)
	}
	for k, v := range queryStrings {
		parsed, err = parsed.Parse("&" + k + "=" + v)
		if err != nil {
			// Should never happen.
			panic(err)
		}
	}
	req, err := http.NewRequest(http.MethodGet, parsed.String(), nil)
	if err != nil {
		// Should never happen.
		panic(err)
	}
	res, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	return data, err
}

// parseInt parses an integer from s and returns it, or 0 if no valid integer could be found.
func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
