package vote

import (
	"encoding/json"
	"time"
)

// Vote represents a single vote on a voting website.
type Vote struct {
	// Time is the time at which the vote was carried out.
	Time time.Time
	// Name is the username of the voter that submitted this vote.
	Name string
	// Claimed specifies if the Vote was claimed. Votes can be claimed a single time after submitting, up to 24 hours
	// after having submitted it.
	Claimed bool
}

// readVotes reads a Vote slice from the JSON encoded data slice passed.
func readVotes(data []byte) ([]Vote, error) {
	var resp responseVotes
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	votes := make([]Vote, len(resp.Votes))
	for i, v := range resp.Votes {
		votes[i] = Vote{
			Time:    time.Unix(int64(v.Timestamp), 0),
			Name:    v.Nickname,
			Claimed: v.Claimed == "1",
		}
	}
	return votes, nil
}

type responseVotes struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    string `json:"port"`
	Month   string `json:"month"`
	Votes   []struct {
		Date      string `json:"date"`
		Timestamp int    `json:"timestamp"`
		Nickname  string `json:"nickname"`
		SteamID   string `json:"steam_id"`
		Claimed   string `json:"claimed"`
	} `json:"votes"`
}
