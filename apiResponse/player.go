package apiResponse

// Player stores data related to a Player account
type Ranked struct {
	Leaves   uint   `json:"Leaves"`
	Losses   uint   `json:"Losses"`
	Name     string `json:"Name"`
	Points   uint   `json:"Points"`
	PrevRank uint   `json:"PrevRank"`
	Rank     uint   `json:"Rank"`
	Season   uint   `json:"Season"`
	Tier     uint   `json:"Tier"`
	Trend    uint   `json:"Trend"`
	Wins     uint   `json:"Wins"`
}

type Player struct {
	ActivePlayerID        uint   `json:"ActivePlayerId"`
	AvatarID              uint   `json:"AvatarId"`
	AvatarURL             string `json:"AvatarURL"`
	CreatedDatetime       string `json:"Created_Datetime"`
	HoursPlayed           uint   `json:"HoursPlayed"`
	PlayerID              uint   `json:"Id"`
	LastLoginDatetime     string `json:"Last_Login_Datetime"`
	Leaves                uint   `json:"Leaves"`
	Level                 uint   `json:"Level"`
	LoadingFrame          string `json:"LoadingFrame"`
	Losses                uint   `json:"Losses"`
	MasteryLevel          uint   `json:"MasteryLevel"`
	MinutesPlayed         uint   `json:"MinutesPlayed"`
	Name                  string `json:"Name"`
	PersonalStatusMessage string `json:"Personal_Status_Message"`
	Platform              string `json:"Platform"`
	RankedController      Ranked `json:"RankedController"`
	RankedKBM             Ranked `json:"RankedKBM"`
	Region                string `json:"Region"`
	Title                 string `json:"Title"`
	TotalAchievements     uint   `json:"Total_Achievements"`
	TotalWorshippers      uint   `json:"Total_Worshippers"`
	TotalXP               uint   `json:"Total_XP"`
	Wins                  uint   `json:"Wins"`
	HirezPlayerName       string `json:"hz_player_name"`
	RetMsg                string `json:"ret_msg"`
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

// Friend stores data related to a Player's Friend
type Friend struct {
	AccountID   string `json:"account_id"`
	FriendFlags string `json:"friend_flags"`
	Name        string `json:"name"`
	PlayerID    string `json:"player_id"`
	PortalID    uint   `json:"portal_id,string"`
	RetMsg      string `json:"ret_msg"`
	Status      string `json:"status"`
}

// PlayerStatus stores data related to a Players current Status
type PlayerStatus struct {
	Match                 uint   `json:"Match"`
	MatchQueueID          uint   `json:"match_queue_id"`
	PersonalStatusMessage string `json:"personal_status_message"`
	RetMsg                string `json:"ret_msg"`
	Status                uint   `json:"status"`
	StatusString          string `json:"status_string"`
}

// PlayerStatus stores data related to a Players current Status
type PlayerQueueStat struct {
	Assists    uint   `json:"Assists"`
	Champion   string `json:"Champion"`
	ChampionID uint   `json:"ChampionId"`
	Deaths     uint   `json:"Deaths"`
	Gold       uint   `json:"Gold"`
	Kills      uint   `json:"Kills"`
	LastPlayed string `json:"LastPlayed"`
	Losses     uint   `json:"Losses"`
	Matches    uint   `json:"Matches"`
	Minutes    uint   `json:"Minutes"`
	Queue      string `json:"Queue"`
	Wins       uint   `json:"Wins"`
	PlayerID   string `json:"player_id"`
	RetMsg     string `json:"ret_msg"`
}
