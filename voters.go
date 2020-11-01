package voteapi

import (
	"encoding/json"
)

type MonthFilter string

const (
	MonthFilterCurrent  MonthFilter = "current"
	MonthFilterPrevious MonthFilter = "previous"
)

type responseVoters struct {
	Name    string          `json:"name"`
	Address string          `json:"address"`
	Port    string          `json:"port"`
	Month   string          `json:"month"`
	Voters  []responseVoter `json:"voters"`
}

type responseVoter struct {
	Nickname string `json:"nickname"`
	Votes    string `json:"votes"`
}

type Voters struct {
	Name    string
	Address string
	Port    int
	Month   string
	Voters  []Voter
}

type Voter struct {
	Nickname string
	Votes    int
}

func votersFromBytes(data []byte) (Voters, error) {
	var resp responseVoters

	if err := json.Unmarshal(data, &resp); err != nil {
		return Voters{}, err
	}

	var voters []Voter
	for _, v := range resp.Voters {
		voters = append(voters, Voter{
			Nickname: v.Nickname,
			Votes:    parseInt(v.Votes),
		})
	}
	return Voters{
		Name:    resp.Name,
		Address: resp.Address,
		Port:    parseInt(resp.Port),
		Month:   resp.Month,
		Voters:  voters,
	}, nil
}
