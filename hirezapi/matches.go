package hirezapi

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/tusharlock10/pego/models"
)

// GetMatchDetails returns the statistics for a particular completed match.
func (a *APIClient) GetMatchDetails(matchID string) ([]models.MatchPlayer, error) {
	resp, err := a.MakeRequest("getmatchdetails", matchID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.MatchPlayer
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetMatchDetailsBatch returns the statistics for a particular set of completed matches. (limit batch query to 5-10 matchIDs)
func (a *APIClient) GetMatchDetailsBatch(matchIDs []string) ([]models.MatchPlayer, error) {
	if len(matchIDs) > 10 {
		return nil, fmt.Errorf("per API docs, the list of matchIDs should contain no more than 10")
	}
	resp, err := a.MakeRequest("getmatchdetailsbatch", strings.Join(matchIDs, ","))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.MatchPlayer
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetMatchPlayerDetails returns player information for a live match.
func (a *APIClient) GetMatchPlayerDetails(matchID string) ([]models.LiveMatchPlayer, error) {
	resp, err := a.MakeRequest("getmatchplayerdetails", matchID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.LiveMatchPlayer
	err = json.Unmarshal(body, &output)
	return output, err
}

/*
GetMatchIDsByQueue lists all MatchIDs for a particular match queue.
- queueID can be referened by constants defined in this package (eg, hirezapi.ConquestRanked).
- date must be formatted/formattable by hirezapi.DateFormat (yyyyMMdd).
- hour may be "0" - "23" and optionally may contain a ten minute window separated by a comma (eg, "6,30").
- hour may also be "-1" to fetch the whole day, but may stall/fail due to the amount of data.
*/
func (a *APIClient) GetMatchIDsByQueue(queueID, date, hour string) ([]models.Match, error) {
	path := fmt.Sprintf("%s/%s/%s", queueID, date, hour)
	resp, err := a.MakeRequest("getmatchidsbyqueue", path)
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
