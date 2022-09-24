package hirezapi

import (
	"fmt"

	"github.com/tusharlock10/pego/apiResponse"
)

// GetPlayer returns league and other high level data for a particular player.
func (a *APIClient) GetPlayer(playerID uint) ([]apiResponse.Player, error) {
	var output []apiResponse.Player
	err := a.MakeRequest("getplayer", fmt.Sprint(playerID), &output)
	return output, err
}

// GetFriends returns Smite Usernames of each of the player's friends. [PC Only]
func (a *APIClient) GetFriends(playerID uint) ([]apiResponse.Friend, error) {
	var output []apiResponse.Friend
	err := a.MakeRequest("getfriends", fmt.Sprint(playerID), &output)
	return output, err
}

// GetPlayerMatchHistory returns a list of recent matches and high level match statistics for a particular player
func (a *APIClient) GetPlayerMatchHistory(playerID uint) ([]apiResponse.PlayerMatchHistory, error) {
	var output []apiResponse.PlayerMatchHistory
	err := a.MakeRequest("getmatchhistory", fmt.Sprint(playerID), &output)
	return output, err
}

// GetPlayerStatus returns a player status.
func (a *APIClient) GetPlayerStatus(playerID uint) ([]apiResponse.PlayerStatus, error) {
	var output []apiResponse.PlayerStatus
	err := a.MakeRequest("getplayerstatus", fmt.Sprint(playerID), &output)
	return output, err
}

// SearchPlayers returns playerID values for all names and/or gamerTags containing searchPlayer
func (a *APIClient) SearchPlayers(searchPlayer string) ([]apiResponse.SearchPlayer, error) {
	var output []apiResponse.SearchPlayer
	err := a.MakeRequest("searchplayers", searchPlayer, &output)
	return output, err
}
