package apiResponse

// Match stores data related to an in game Match
type MatchIdsByQueue struct {
	ActiveFlag    string `json:"Active_Flag"`
	EntryDatetime string `json:"Entry_Datetime"`
	Match         uint   `json:"Match,string"`
	Region        string `json:"Region"`
	RetMsg        string `json:"ret_msg"`
}

// MatchPlayerDetail stores data related to a Player present in a Match
type MatchPlayerDetail struct {
	AccountLevel        uint   `json:"Account_Level"`
	ActiveID1           uint   `json:"ActiveId1"`
	ActiveID2           uint   `json:"ActiveId2"`
	ActiveID3           uint   `json:"ActiveId3"`
	ActiveID4           uint   `json:"ActiveId4"`
	ActiveLevel1        uint   `json:"ActiveLevel1"`
	ActiveLevel2        uint   `json:"ActiveLevel2"`
	ActiveLevel3        uint   `json:"ActiveLevel3"`
	ActiveLevel4        uint   `json:"ActiveLevel4"`
	ActivePlayerID      string `json:"ActivePlayerId"`
	Assists             uint   `json:"Assists"`
	BanID1              uint   `json:"BanId1"`
	BanID2              uint   `json:"BanId2"`
	BanID3              uint   `json:"BanId3"`
	BanID4              uint   `json:"BanId4"`
	BanID5              uint   `json:"BanId5"`
	BanID6              uint   `json:"BanId6"`
	ChampionID          uint   `json:"ChampionId"`
	DamageDoneInHand    uint   `json:"Damage_Done_In_Hand"`
	DamageDonePhysical  uint   `json:"Damage_Done_Physical"`
	DamageMitigated     uint   `json:"Damage_Mitigated"`
	DamagePlayer        uint   `json:"Damage_Player"`
	DamageTaken         uint   `json:"Damage_Taken"`
	DamageTakenPhysical uint   `json:"Damage_Taken_Physical"`
	Deaths              uint   `json:"Deaths"`
	EntryDatetime       string `json:"Entry_Datetime"`
	GoldEarned          uint   `json:"Gold_Earned"`
	GoldPerMinute       uint   `json:"Gold_Per_Minute"`
	Healing             uint   `json:"Healing"`
	HealingPlayerSelf   uint   `json:"Healing_Player_Self"`
	ItemID1             uint   `json:"ItemId1"`
	ItemID2             uint   `json:"ItemId2"`
	ItemID3             uint   `json:"ItemId3"`
	ItemID4             uint   `json:"ItemId4"`
	ItemID5             uint   `json:"ItemId5"`
	ItemID6             uint   `json:"ItemId6"`
	ItemLevel1          uint   `json:"ItemLevel1"`
	ItemLevel2          uint   `json:"ItemLevel2"`
	ItemLevel3          uint   `json:"ItemLevel3"`
	ItemLevel4          uint   `json:"ItemLevel4"`
	ItemLevel5          uint   `json:"ItemLevel5"`
	ItemLevel6          uint   `json:"ItemLevel6"`
	ItemActive1         string `json:"Item_Active_1"`
	ItemActive2         string `json:"Item_Active_2"`
	ItemActive3         string `json:"Item_Active_3"`
	ItemActive4         string `json:"Item_Active_4"`
	ItemPurch1          string `json:"Item_Purch_1"`
	ItemPurch2          string `json:"Item_Purch_2"`
	ItemPurch3          string `json:"Item_Purch_3"`
	ItemPurch4          string `json:"Item_Purch_4"`
	ItemPurch5          string `json:"Item_Purch_5"`
	ItemPurch6          string `json:"Item_Purch_6"`
	KillingSpree        uint   `json:"Killing_Spree"`
	KillsBot            uint   `json:"Kills_Bot"`
	LeagueLosses        uint   `json:"League_Losses"`
	LeaguePoints        uint   `json:"League_Points"`
	LeagueTier          uint   `json:"League_Tier"`
	LeagueWins          uint   `json:"League_Wins"`
	MapGame             string `json:"Map_Game"`
	MasteryLevel        uint   `json:"Mastery_Level"`
	Match               uint   `json:"Match"`
	MatchDuration       uint   `json:"Match_Duration"`
	Minutes             uint   `json:"Minutes"`
	MultiKillMax        uint   `json:"Multi_kill_Max"`
	ObjectiveAssists    uint   `json:"Objective_Assists"`
	PartyID             uint   `json:"PartyId"`
	Platform            string `json:"Platform"`
	ReferenceName       string `json:"Reference_Name"`
	Region              string `json:"Region"`
	Skin                string `json:"Skin"`
	SkinID              uint   `json:"SkinId"`
	TaskForce           uint   `json:"TaskForce"`
	Team1Score          uint   `json:"Team1Score"`
	Team2Score          uint   `json:"Team2Score"`
	TeamID              uint   `json:"TeamId"`
	TimeInMatchSeconds  uint   `json:"Time_In_Match_Seconds"`
	WinStatus           string `json:"Win_Status"`
	WinningTaskForce    uint   `json:"Winning_TaskForce"`
	HasReplay           string `json:"hasReplay"`
	MatchQueueID        uint   `json:"match_queue_id"`
	Name                string `json:"name"`
	PlayerID            string `json:"playerId"`
	PlayerName          string `json:"playerName"`
	PlayerPortalID      uint   `json:"playerPortalId"`
	PlayerPortalUserID  string `json:"playerPortalUserId"`
	RetMsg              string `json:"ret_msg"`
}

// ActiveMatchDetail stores data related to a Player in a currently Live Match
type ActiveMatchDetail struct {
	AccountChampionsPlayed uint   `json:"Account_Champions_Played"`
	AccountLevel           uint   `json:"Account_Level"`
	ChampionID             uint   `json:"ChampionId"`
	ChampionLevel          uint   `json:"ChampionLevel"`
	ChampionName           string `json:"ChampionName"`
	MasteryLevel           uint   `json:"Mastery_Level"`
	Match                  uint   `json:"Match"`
	Queue                  uint   `json:"Queue,string"`
	Tier                   uint   `json:"Tier"`
	MapGame                string `json:"mapGame"`
	PlayerCreated          string `json:"playerCreated"`
	PlayerID               uint   `json:"playerId,string"`
	PlayerName             string `json:"playerName"`
	PlayerRegion           string `json:"playerRegion"`
	RetMsg                 string `json:"ret_msg"`
	TaskForce              uint   `json:"taskForce"`
	TierLosses             uint   `json:"tierLosses"`
	TierWins               uint   `json:"tierWins"`
}

type PlayerMatchHistory struct {
	ActiveID1           uint   `json:"ActiveId1"`
	ActiveID2           uint   `json:"ActiveId2"`
	ActiveID3           uint   `json:"ActiveId3"`
	ActiveID4           uint   `json:"ActiveId4"`
	ActiveLevel1        uint   `json:"ActiveLevel1"`
	ActiveLevel2        uint   `json:"ActiveLevel2"`
	ActiveLevel3        uint   `json:"ActiveLevel3"`
	ActiveLevel4        uint   `json:"ActiveLevel4"`
	Active1             string `json:"Active_1"`
	Active2             string `json:"Active_2"`
	Active3             string `json:"Active_3"`
	Active4             string `json:"Active_4"`
	Assists             uint   `json:"Assists"`
	Champion            string `json:"Champion"`
	ChampionID          uint   `json:"ChampionId"`
	Damage              uint   `json:"Damage"`
	DamageDoneInHand    uint   `json:"Damage_Done_In_Hand"`
	DamageMitigated     uint   `json:"Damage_Mitigated"`
	DamageTaken         uint   `json:"Damage_Taken"`
	DamageTakenPhysical uint   `json:"Damage_Taken_Physical"`
	Deaths              uint   `json:"Deaths"`
	Gold                uint   `json:"Gold"`
	Healing             uint   `json:"Healing"`
	HealingPlayerSelf   uint   `json:"Healing_Player_Self"`
	ItemID1             uint   `json:"ItemId1"`
	ItemID2             uint   `json:"ItemId2"`
	ItemID3             uint   `json:"ItemId3"`
	ItemID4             uint   `json:"ItemId4"`
	ItemID5             uint   `json:"ItemId5"`
	ItemID6             uint   `json:"ItemId6"`
	ItemLevel1          uint   `json:"ItemLevel1"`
	ItemLevel2          uint   `json:"ItemLevel2"`
	ItemLevel3          uint   `json:"ItemLevel3"`
	ItemLevel4          uint   `json:"ItemLevel4"`
	ItemLevel5          uint   `json:"ItemLevel5"`
	ItemLevel6          uint   `json:"ItemLevel6"`
	Item1               string `json:"Item_1"`
	Item2               string `json:"Item_2"`
	Item3               string `json:"Item_3"`
	Item4               string `json:"Item_4"`
	Item5               string `json:"Item_5"`
	Item6               string `json:"Item_6"`
	Killing_Spree       uint   `json:"Killing_Spree"`
	Kills               uint   `json:"Kills"`
	Level               uint   `json:"Level"`
	MapGame             string `json:"Map_Game"`
	Match               uint   `json:"Match"`
	MatchQueueID        uint   `json:"Match_Queue_Id"`
	MatchTime           string `json:"Match_Time"`
	Minutes             uint   `json:"Minutes"`
	MultiKillMax        uint   `json:"Multi_kill_Max"`
	ObjectiveAssists    uint   `json:"Objective_Assists"`
	Queue               string `json:"Queue"`
	Region              string `json:"Region"`
	Skin                string `json:"Skin"`
	SkinID              uint   `json:"SkinId"`
	TaskForce           uint   `json:"TaskForce"`
	Team1Score          uint   `json:"Team1Score"`
	Team2Score          uint   `json:"Team2Score"`
	TimeInMatchSeconds  uint   `json:"Time_In_Match_Seconds"`
	WardsPlaced         uint   `json:"Wards_Placed"`
	WinStatus           string `json:"Win_Status"`
	WinningTaskForce    uint   `json:"Winning_TaskForce"`
	PlayerID            uint   `json:"playerId"`
	PlayerName          string `json:"playerName"`
	RetMsg              string `json:"ret_msg"`
}
