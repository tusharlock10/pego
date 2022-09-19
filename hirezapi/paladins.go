package hirezapi

import (
	"fmt"

	"github.com/tusharlock10/pego/apiResponse"
)

// GetPlayerBatch returns league and other high level data for a particular list of players
func (a *APIClient) GetPlayerBatch(playerIDs []uint) ([]apiResponse.Player, error) {
	joinedPlayerIDs := ""
	for i, playerID := range playerIDs {
		separator := ","
		if i == 0 {
			separator = ""
		}
		joinedPlayerIDs = joinedPlayerIDs + separator + fmt.Sprint(playerID)
	}
	var output []apiResponse.Player
	err := a.MakeRequest("getplayerbatch", joinedPlayerIDs, &output)
	return output, err
}

// GetChampionRanks returns the rank and worshipper values for each Champion a player has played.
func (a *APIClient) GetChampionRanks(playerID uint) ([]apiResponse.ChampionRank, error) {
	var output []apiResponse.ChampionRank
	err := a.MakeRequest("getchampionranks", fmt.Sprint(playerID), &output)
	return output, err
}

// GetChampions returns all Champions and their various attributes.
func (a *APIClient) GetChampions() ([]apiResponse.Champion, error) {
	var output []apiResponse.Champion
	err := a.MakeRequest("getchampions", "1", &output)
	return output, err
}

// GetPlayerLoadouts returns deck loadouts per Champion
func (a *APIClient) GetPlayerLoadouts(playerID uint) ([]apiResponse.PlayerLoadout, error) {
	var output []apiResponse.PlayerLoadout
	path := fmt.Sprintf("%s/%s", fmt.Sprint(playerID), "1")
	err := a.MakeRequest("getplayerloadouts", path, &output)
	return output, err
}

// GetChampionCards returns all Champion cards.
func (a *APIClient) GetChampionCards(championID uint) ([]apiResponse.ChampionCard, error) {
	var output []apiResponse.ChampionCard
	path := fmt.Sprintf("%s/%s", fmt.Sprint(championID), "1")
	err := a.MakeRequest("getchampioncards", path, &output)
	return output, err
}

// GetBountyItems returns daily Bounty Item history for the past 6 months.
func (a *APIClient) GetBountyItems() ([]apiResponse.BountyItem, error) {
	var output []apiResponse.BountyItem
	err := a.MakeRequest("getbountyitems", "", &output)
	return output, err
}
