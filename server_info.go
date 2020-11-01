package voteapi

import (
	"encoding/json"
)

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
	MaxPlayers string `json:"max_players"`
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

type ServerInfo struct {
	ID         int
	Name       string
	Address    string
	Port       int
	Private    bool
	Password   bool
	Location   string
	HostName   string
	Players    int
	MaxPlayers int
	Version    string
	Platform   string
	Uptime     int
	Score      int
	Rank       int
	Votes      int
	Favorited  int
	Comments   int
	URL        string
	LastCheck  string
	LastOnline string
}

func serverInfoFromBytes(data []byte) (ServerInfo, error) {
	var resp responseServerInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return ServerInfo{}, err
	}

	return ServerInfo{
		ID:         parseInt(resp.ID),
		Name:       resp.Name,
		Address:    resp.Address,
		Port:       parseInt(resp.Port),
		Private:    resp.Private == "1",
		Password:   resp.Password == "1",
		Location:   resp.Location,
		HostName:   resp.HostName,
		Players:    parseInt(resp.Players),
		MaxPlayers: parseInt(resp.MaxPlayers),
		Version:    resp.Version,
		Platform:   resp.Platform,
		Uptime:     parseInt(resp.Uptime),
		Score:      parseInt(resp.Score),
		Rank:       parseInt(resp.Rank),
		Votes:      parseInt(resp.Votes),
		Favorited:  parseInt(resp.Favorited),
		Comments:   parseInt(resp.Comments),
		URL:        resp.URL,
		LastCheck:  resp.LastCheck,
		LastOnline: resp.LastOnline,
	}, nil
}
