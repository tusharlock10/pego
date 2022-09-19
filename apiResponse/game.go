package apiResponse

// Item stores data related to an in game Item
type Item struct {
	Description       string `json:"Description"`
	DeviceName        string `json:"DeviceName"`
	IconID            uint   `json:"IconId"`
	ItemID            uint   `json:"ItemId"`
	Price             uint   `json:"Price"`
	ShortDesc         string `json:"ShortDesc"`
	ChampionID        uint   `json:"champion_id"`
	ItemIconURL       string `json:"itemIcon_URL"`
	ItemType          string `json:"item_type"`
	RechargeSeconds   uint   `json:"recharge_seconds"`
	RetMsg            string `json:"ret_msg"`
	TalentRewardLevel uint   `json:"talent_reward_level"`
}
