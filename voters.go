package vote

import "encoding/json"

type MonthFilter string

const (
	MonthFilterCurrent  MonthFilter = "current"
	MonthFilterPrevious MonthFilter = "previous"
)

type Voter struct {
	Name  string
	Votes int
}

func readVoters(data []byte) ([]Voter, error) {
	var resp responseVoters
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	voters := make([]Voter, len(resp.Voters))
	for i, v := range resp.Voters {
		voters[i] = Voter{
			Name:  v.Nickname,
			Votes: parseInt(v.Votes),
		}
	}
	return voters, nil
}

type responseVoters struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    string `json:"port"`
	Month   string `json:"month"`
	Voters  []struct {
		Nickname string `json:"nickname"`
		Votes    string `json:"votes"`
	} `json:"voters"`
}
