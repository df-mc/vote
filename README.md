# vote
Go library for accessing the minecraftpocket-servers.com voting API

## Getting started
The vote library may be imported using `go get`:
```
go get github.com/df-mc/vote
```

## Usage
[![Go Reference](https://pkg.go.dev/badge/github.com/df-mc/vote.svg)](https://pkg.go.dev/github.com/df-mc/vote)

Usage of the vote package relies on a `*vote.Client` that may be constructed and used as such:

```go
// var key string

c := vote.NewClient(key)
voted, err := c.Voted("Steve")
if err != nil {
	panic(err)
}
if !voted {
	fmt.Println("Steve has not yet voted")
	return
}
claimed, err := c.Claim("Steve")
if err != nil {
	panic(err)
}
if !claimed {
	fmt.Println("Steve has already claimed his vote")
	return
}
fmt.Println("Steve claimed his vote successfully")
```
The errors returned by methods on `*vote.Client` only return HTTP errors.

## Contact
[![Discord Banner 2](https://discordapp.com/api/guilds/623638955262345216/widget.png?style=banner2)](https://discord.gg/U4kFWHhTNR)
