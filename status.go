package vote

import (
	"encoding/json"
)

// ServerStatus contains information on the server associated with the voting key specified in the call to NewClient.
// It has status and statistics relating to the voting website and server itself.
type ServerStatus struct {
	Players    int
	MaxPlayers int
	Version    string
	Uptime     int
	Score      int
	Rank       int
	Votes      int
	Favourited int
	Comments   int
	LastOnline string
}

// readServerStatus reads a ServerStatus from the JSON encoded data slice passed.
func readServerStatus(data []byte) (ServerStatus, error) {
	var resp responseServerInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return ServerStatus{}, err
	}

	return ServerStatus{
		Players:    parseInt(resp.Players),
		MaxPlayers: parseInt(resp.MaxPlayers),
		Version:    resp.Version,
		Uptime:     parseInt(resp.Uptime),
		Score:      parseInt(resp.Score),
		Rank:       parseInt(resp.Rank),
		Votes:      parseInt(resp.Votes),
		Favourited: parseInt(resp.Favorited),
		Comments:   parseInt(resp.Comments),
		LastOnline: resp.LastOnline,
	}, nil
}

type responseServerInfo struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Port       string `json:"port"`
	Private    string `json:"private"`
	Password   string `json:"password"`
	Location   string `json:"location"`
	HostName   string `json:"host_name"`
	Players    string `json:"players"`
	MaxPlayers string `json:"maxplayers"`
	Version    string `json:"version"`
	Platform   string `json:"platform"`
	Uptime     string `json:"uptime"`
	Score      string `json:"score"`
	Rank       string `json:"rank"`
	Votes      string `json:"votes"`
	Favorited  string `json:"favorited"`
	Comments   string `json:"comments"`
	URL        string `json:"url"`
	LastCheck  string `json:"last_check"`
	LastOnline string `json:"last_online"`
}
