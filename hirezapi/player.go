package hirezapi

import (
	"encoding/json"
	"io"

	"github.com/tusharlock10/pego/models"
)

// GetPlayer returns league and other high level data for a particular player.
func (a *APIClient) GetPlayer(player string) ([]models.Player, error) {
	resp, err := a.MakeRequest("getplayer", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.Player
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetFriends returns Smite Usernames of each of the player's friends. [PC Only]
func (a *APIClient) GetFriends(player string) ([]models.Friend, error) {
	resp, err := a.MakeRequest("getfriends", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.Friend
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetMatchHistory returns a list of recent matches and high level match statistics for a particular player
func (a *APIClient) GetMatchHistory(player string) ([]models.Match, error) {
	resp, err := a.MakeRequest("getmatchhistory", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.Match
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetPlayerStatus returns a player status.
func (a *APIClient) GetPlayerStatus(player string) ([]models.PlayerStatus, error) {
	resp, err := a.MakeRequest("getplayerstatus", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.PlayerStatus
	err = json.Unmarshal(body, &output)
	return output, err
}

// SearchPlayers returns playerID values for all names and/or gamerTags containing searchPlayer
func (a *APIClient) SearchPlayers(searchPlayer string) ([]models.SearchPlayer, error) {
	resp, err := a.MakeRequest("searchplayers", searchPlayer)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.SearchPlayer
	err = json.Unmarshal(body, &output)
	return output, err
}
