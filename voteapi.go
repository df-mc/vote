package voteapi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type VoteResponse int

const (
	PlayerNotVoted VoteResponse = iota
	PlayerVoteUnclaimed
	PlayerVoteClaimed
)

type VoteAPI struct {
	key string
}

// New creates a new VoteAPI struct which is used to fetch information from mincraftpocket-servers.com.
func New(key string) VoteAPI {
	return VoteAPI{key: key}
}

// Voted checks if the username provided has voted within the last 24 hours. The different responses can be
// found above.
func (v VoteAPI) Voted(username string) (VoteResponse, error) {
	data, err := v.get(map[string]string{
		"object":   "votes",
		"element":  "claim",
		"username": username,
	})
	if err != nil {
		return PlayerNotVoted, err
	}
	return VoteResponse(parseInt(string(data))), nil
}

// ClaimVote attempts to claim the vote of the provided username. The returned values depend on whether or
// not the claim was successful.
func (v VoteAPI) ClaimVote(username string) (bool, error) {
	data, err := v.get(map[string]string{
		"action":   "post",
		"object":   "votes",
		"element":  "claim",
		"username": username,
	})
	if err != nil {
		return false, err
	}
	return string(data) == "1", nil
}

// Voters returns the top voters of the month provided. The limit can be anywhere in the range of 1-1000.
func (v VoteAPI) Voters(month MonthFilter, limit int) (Voters, error) {
	if limit < 1 || limit > 1000 {
		return Voters{}, errors.New("limit must be in the range of 1-1000")
	}
	data, err := v.get(map[string]string{
		"object":  "servers",
		"element": "voters",
		"month":   string(month),
		"format":  "json",
		"limit":   strconv.Itoa(limit),
	})
	if err != nil {
		return Voters{}, err
	}
	return votersFromBytes(data)
}

// Votes returns the most recent votes and the information about them for the current month. THe limit can be
// anywhere in the range of 1-1000.
func (v VoteAPI) Votes(limit int) (Votes, error) {
	if limit < 1 || limit > 1000 {
		return Votes{}, errors.New("limit must be in the range of 1-1000")
	}
	data, err := v.get(map[string]string{
		"object":  "servers",
		"element": "votes",
		"format":  "json",
		"limit":   strconv.Itoa(limit),
	})
	if err != nil {
		return Votes{}, err
	}
	return allVotesFromBytes(data)
}

// ServerInfo returns the statistics and information about the server.
func (v VoteAPI) ServerInfo() (ServerInfo, error) {
	data, err := v.get(map[string]string{
		"object":  "servers",
		"element": "detail",
	})
	if err != nil {
		return ServerInfo{}, err
	}
	return serverInfoFromBytes(data)
}

func (v VoteAPI) get(params map[string]string) ([]byte, error) {
	url := fmt.Sprintf("https://minecraftpocket-servers.com/api/?key=%s", v.key)
	for k, v := range params {
		url += fmt.Sprintf("&%s=%s", k, v)
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}
