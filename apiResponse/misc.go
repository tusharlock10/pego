package apiResponse

// Season stores data related to a game Season
type Season struct {
	Complete          bool   `json:"complete"`
	ID                int64  `json:"id"`
	LeagueDescription string `json:"league_description"`
	Name              string `json:"name"`
	RetMsg            string `json:"ret_msg"`
	Round             int64  `json:"round"`
	Season            int64  `json:"season"`
}

// TopMatch stores data related to a recent Top Match
type TopMatch struct {
	Ban1              string `json:"Ban1"`
	Ban1ID            int64  `json:"Ban1Id"`
	Ban2              string `json:"Ban2"`
	Ban2ID            int64  `json:"Ban2Id"`
	EntryDatetime     string `json:"Entry_Datetime"`
	LiveSpectators    int64  `json:"LiveSpectators"`
	Match             int64  `json:"Match"`
	MatchTime         int64  `json:"Match_Time"`
	OfflineSpectators int64  `json:"OfflineSpectators"`
	Queue             string `json:"Queue"`
	RecordingFinished string `json:"RecordingFinished"`
	RecordingStarted  string `json:"RecordingStarted"`
	Team1AvgLevel     int64  `json:"Team1_AvgLevel"`
	Team1Gold         int64  `json:"Team1_Gold"`
	Team1Kills        int64  `json:"Team1_Kills"`
	Team1Score        int64  `json:"Team1_Score"`
	Team2AvgLevel     int64  `json:"Team2_AvgLevel"`
	Team2Gold         int64  `json:"Team2_Gold"`
	Team2Kills        int64  `json:"Team2_Kills"`
	Team2Score        int64  `json:"Team2_Score"`
	WinningTeam       int64  `json:"WinningTeam"`
	RetMsg            string `json:"ret_msg"`
}

// VerstionInfo stores data related to the version of a game patch
type VersionInfo struct {
	RetMsg        string `json:"ret_msg"`
	VersionString string `json:"version_string"`
}
