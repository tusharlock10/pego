package apiResponse

// PlayerLoadout stores data related to a Player's Champion Loadout
type PlayerLoadout struct {
	ChampionID   uint          `json:"ChampionId"`
	ChampionName string        `json:"ChampionName"`
	DeckID       uint          `json:"DeckId"`
	DeckName     string        `json:"DeckName"`
	LoadoutItems []LoadoutItem `json:"LoadoutItems"`
	PlayerID     uint          `json:"playerId"`
	PlayerName   string        `json:"playerName"`
	RetMsg       string        `json:"ret_msg"`
}

// LoadoutItem stores data related to an Item inside a Loadout
type LoadoutItem struct {
	ItemID   uint   `json:"ItemId"`
	ItemName string `json:"ItemName"`
	Points   uint   `json:"Points"`
}

// BountyItem stores data related to a Bounty Item in Paladins
type BountyItem struct {
	Active          string `json:"active"`
	BountyItemID1   uint   `json:"bounty_item_id1"`
	BountyItemID2   uint   `json:"bounty_item_id2"`
	BountyItemName  string `json:"bounty_item_name"`
	ChampionID      uint   `json:"champion_id"`
	ChampionName    string `json:"champion_name"`
	FinalPrice      string `json:"final_price"`
	InitialPrice    string `json:"initial_price"`
	RetMsg          string `json:"ret_msg"`
	SaleEndDatetime string `json:"sale_end_datetime"`
	SaleType        string `json:"sale_type"`
}
