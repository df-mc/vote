// Package vote implements a simple interface to the voting API as exposed by minecraftpocket-servers.com.
// `vote.NewClient()` may be called to create a new `*vote.Client` with a key that points to a specific server. This key
// can be found in the settings. `*vote.Client` exports several methods to get the server status, top voters and most
// recent votes.
package vote
