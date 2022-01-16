package vote

import "encoding/json"

// Voter represents a user that voted for a server at least once in the month specified.
type Voter struct {
	// Name is the username of the voter as specified when voting on the voting website.
	Name string
	// Votes is the amount of times the Voter voted in the month specified.
	Votes int
}

// readVoters reads a Voter slice from the JSON encoded data slice passed.
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
