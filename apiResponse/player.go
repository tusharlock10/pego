package apiResponse

// Player stores data related to a Player account
type Player struct {
	ActivePlayerID               int64            `json:"ActivePlayerId"`
	AvatarID                     int64            `json:"AvatarID"`
	AvatarURL                    string           `json:"Avatar_URL"`
	CreatedDatetime              string           `json:"Created_Datetime"`
	HoursPlayed                  int64            `json:"HoursPlayed"`
	ID                           int64            `json:"Id"`
	LastLoginDatetime            string           `json:"Last_Login_Datetime"`
	Leaves                       int64            `json:"Leaves"`
	Level                        int64            `json:"Level"`
	LoadingFrame                 string           `json:"LoadingFrame"`
	Losses                       int64            `json:"Losses"`
	MasteryLevel                 int              `json:"MasteryLevel"`
	MergedPlayers                []MergedPlayer   `json:"MergedPlayers"`
	MinutesPlayed                int64            `json:"MinutesPlayed"`
	Name                         string           `json:"Name"`
	PersonalStatusMessage        string           `json:"Personal_Status_Message"`
	Platform                     string           `json:"Platform"`
	RankedStatConquest           float32          `json:"Rank_Stat_Conquest"`
	RankedStatConquestController float32          `json:"Rank_Stat_Conquest_Controller"`
	RankedStatDuel               float32          `json:"Rank_Stat_Duel"`
	RankedStatDuelController     float32          `json:"Rank_Stat_Duel_Controller"`
	RankedStatJoust              float32          `json:"Rank_Stat_Joust"`
	RankedStatJoustController    float32          `json:"Rank_Stat_Joust_Controller"`
	RankedConquest               PlayerRankedInfo `json:"RankedConquest"`
	RankedConquestController     PlayerRankedInfo `json:"RankedConquestController"`
	RankedDuel                   PlayerRankedInfo `json:"RankedDuel"`
	RankedDuelController         PlayerRankedInfo `json:"RankedDuelController"`
	RankedJoust                  PlayerRankedInfo `json:"RankedJoust"`
	RankedJoustController        PlayerRankedInfo `json:"RankedJoustController"`
	RankedKBM                    PlayerRankedInfo `json:"RankedKBM"`
	Region                       string           `json:"Region"`
	TeamID                       int64            `json:"TeamId"`
	TeamName                     string           `json:"Team_Name"`
	TierConquest                 int              `json:"Tier_Conquest"`
	TierDuel                     int              `json:"Tier_Duel"`
	TierJoust                    int              `json:"Tier_Joust"`
	TierRankedController         int              `json:"Tier_RankedController"`
	TierRankedKBM                int              `json:"Tier_RankedKBM"`
	Title                        string           `json:"Title"`
	TotalAchievements            int64            `json:"Total_Achievements"`
	TotalWorshippers             int64            `json:"Total_Worshippers"`
	TotalXP                      int64            `json:"Total_XP"`
	Wins                         int64            `json:"Wins"`
	HZGamerTag                   string           `json:"hz_gamer_tag"`
	HZPlayerName                 string           `json:"hz_player_name"`
	RetMsg                       string           `json:"ret_msg"`
}

// MergedPlayer stores data related to merged identites of the same Player on different Platforms
type MergedPlayer struct {
	MergeDatetime string `json:"merge_datetime"`
	PlayerID      string `json:"playerId"`
	PortalID      string `json:"portalId"`
}

// SearchPlayer stores data related to a searched player
type SearchPlayer struct {
	Name         string `json:"Name"`
	HZPlayerName string `json:"hz_player_name"`
	PlayerID     string `json:"player_id"`
	PortalID     string `json:"portal_id"`
	PrivacyFlag  string `json:"privacy_flag"`
	RetMsg       string `json:"ret_msg"`
}

// PlayerRankedInfo stores data related to Ranked Info for a Player
type PlayerRankedInfo struct {
	Leaves       int64   `json:"Leaves"`
	Losses       int64   `json:"Losses"`
	Name         string  `json:"Name"`
	Points       int64   `json:"Points"`
	PrevRank     int64   `json:"PrevRank"`
	Rank         int64   `json:"Rank"`
	RankStat     float32 `json:"Rank_Stat"`
	RankVariance int64   `json:"Rank_Variance"`
	Season       int64   `json:"Season"`
	Tier         int64   `json:"Tier"`
	Trend        int64   `json:"Trend"`
	Wins         int64   `json:"Wins"`
	PlayerID     string  `json:"player_id"`
	RetMsg       string  `json:"ret_msg"`
}

// Friend stores data related to a Player's Friend
type Friend struct {
	AccountID   string `json:"account_id"`
	AvatarURL   string `json:"avatar_url"`
	FriendFlags string `json:"friend_flags"`
	Name        string `json:"name"`
	PlayerID    string `json:"player_id"`
	PortalID    string `json:"portal_id"`
	RetMsg      string `json:"ret_msg"`
	Status      string `json:"status"`
}

// PlayerStatus stores data related to a Players current Status
type PlayerStatus struct {
	Match                 int64  `json:"Match"`
	MatchQueueID          int64  `json:"match_queue_id"`
	PersonalStatusMessage string `json:"personal_status_message"`
	RetMsg                string `json:"ret_msg"`
	Status                int    `json:"status"`
	StatusString          string `json:"status_string"`
}
