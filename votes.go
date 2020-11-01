package voteapi

import (
	"encoding/json"
	"time"
)

type responseVotes struct {
	Name    string             `json:"name"`
	Address string             `json:"address"`
	Port    string             `json:"port"`
	Month   string             `json:"month"`
	Votes   []responseVoteInfo `json:"votes"`
}

type responseVoteInfo struct {
	Date      string `json:"date"`
	Timestamp int    `json:"timestamp"`
	Nickname  string `json:"nickname"`
	SteamID   string `json:"steam_id"`
	Claimed   string `json:"claimed"`
}

type Votes struct {
	Name    string
	Address string
	Port    int
	Month   string
	Votes   []VoteInfo
}

type VoteInfo struct {
	Date      time.Time
	Timestamp int
	Nickname  string
	SteamID   string
	Claimed   bool
}

func allVotesFromBytes(data []byte) (Votes, error) {
	var resp responseVotes

	if err := json.Unmarshal(data, &resp); err != nil {
		return Votes{}, err
	}

	var votes []VoteInfo
	for _, v := range resp.Votes {
		votes = append(votes, VoteInfo{
			Date:      time.Unix(int64(v.Timestamp), 0),
			Timestamp: v.Timestamp,
			Nickname:  v.Nickname,
			SteamID:   v.SteamID,
			Claimed:   v.Claimed == "1",
		})
	}
	return Votes{
		Name:    resp.Name,
		Address: resp.Address,
		Port:    parseInt(resp.Port),
		Month:   resp.Month,
		Votes:   votes,
	}, nil
}
