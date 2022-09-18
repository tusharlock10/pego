package hirezapi

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/tusharlock10/pego/apiResponse"
)

// GetPlayerBatch returns league and other high level data for a particular list of players. [20 max]
func (a *APIClient) GetPlayerBatch(playerIDs []string) ([]apiResponse.Player, error) {
	if len(playerIDs) > 20 {
		return nil, fmt.Errorf("per API docs, the list of playerIDs should contain no more than 20")
	}
	resp, err := a.MakeRequest("getplayerbatch", strings.Join(playerIDs, ","))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []apiResponse.Player
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetChampionRanks returns the rank and worshipper values for each Champion a player has played.
func (a *APIClient) GetChampionRanks(player string) ([]apiResponse.ChampionRank, error) {
	resp, err := a.MakeRequest("getchampionranks", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []apiResponse.ChampionRank
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetChampions returns all Champions and their various attributes.
func (a *APIClient) GetChampions() ([]apiResponse.Champion, error) {
	resp, err := a.MakeRequest("getchampions", "1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []apiResponse.Champion
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetPlayerLoadouts returns deck loadouts per Champion
func (a *APIClient) GetPlayerLoadouts(player string) ([]apiResponse.PlayerLoadout, error) {
	path := fmt.Sprintf("%s/%s", player, "1")
	resp, err := a.MakeRequest("getplayerloadouts", path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []apiResponse.PlayerLoadout
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetChampionCards returns all Champion cards.
func (a *APIClient) GetChampionCards(champID string) ([]apiResponse.ChampionCard, error) {
	path := fmt.Sprintf("%s/%s", champID, "1")
	resp, err := a.MakeRequest("getchampioncards", path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []apiResponse.ChampionCard
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetBountyItems returns daily Bounty Item history for the past 6 months.
func (a *APIClient) GetBountyItems() ([]apiResponse.BountyItem, error) {
	resp, err := a.MakeRequest("getbountyitems", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []apiResponse.BountyItem
	err = json.Unmarshal(body, &output)
	return output, err
}
