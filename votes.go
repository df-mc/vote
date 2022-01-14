package vote

import (
	"encoding/json"
	"time"
)

type Vote struct {
	Time    time.Time
	Name    string
	Claimed bool
}

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
