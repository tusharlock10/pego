package hirezapi

import (
	"fmt"

	"github.com/tusharlock10/pego/apiResponse"
)

// GetMatchDetails returns the statistics for a particular completed match.
func (a *APIClient) GetMatchDetails(matchID uint) ([]apiResponse.MatchPlayerDetail, error) {
	var output []apiResponse.MatchPlayerDetail
	err := a.MakeRequest("getmatchdetails", fmt.Sprint(matchID), &output)
	return output, err
}

// GetMatchDetailsBatch returns the statistics for a particular set of completed matches
func (a *APIClient) GetMatchDetailsBatch(matchIDs []uint) ([]apiResponse.MatchPlayerDetail, error) {
	joinedMatchIDs := ""
	for i, matchID := range matchIDs {
		separator := ","
		if i == 0 {
			separator = ""
		}
		joinedMatchIDs = joinedMatchIDs + separator + fmt.Sprint(matchID)
	}
	var output []apiResponse.MatchPlayerDetail
	err := a.MakeRequest("getmatchdetailsbatch", joinedMatchIDs, &output)
	return output, err
}

// GetActiveMatchDetails returns player information for an active match.
func (a *APIClient) GetActiveMatchDetails(matchID uint) ([]apiResponse.ActiveMatchDetail, error) {
	var output []apiResponse.ActiveMatchDetail
	err := a.MakeRequest("getmatchplayerdetails", fmt.Sprint(matchID), &output)
	return output, err
}

/*
GetMatchIDsByQueue lists all MatchIDs for a particular match queue.
- queueID can be referened by constants
- date must be formatted/formattable by hirezapi.DateFormat (yyyyMMdd).
- hour may be "0" - "23" and optionally may contain a ten minute window separated by a comma (eg, "6,30").
- hour may also be "-1" to fetch the whole day, but may stall/fail due to the amount of data.
*/
func (a *APIClient) GetMatchIDsByQueue(queueID uint, date, hour string) ([]apiResponse.MatchIdsByQueue, error) {
	var output []apiResponse.MatchIdsByQueue
	path := fmt.Sprintf("%s/%s/%s", fmt.Sprint(queueID), date, hour)
	err := a.MakeRequest("getmatchidsbyqueue", path, &output)
	return output, err
}
