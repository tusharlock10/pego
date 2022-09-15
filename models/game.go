package models

// Item stores data related to an in game Item
type Item struct {
	ActiveFlag      string          `json:"ActiveFlag"`
	ChildItemID     int64           `json:"ChildItemId"`
	DeviceName      string          `json:"DeviceName"`
	IconID          int64           `json:"IconId"`
	ItemDescription ItemDescription `json:"ItemDescription"`
	ItemID          int64           `json:"ItemId"`
	ItemTier        int             `json:"ItemTier"`
	Price           int64           `json:"Price"`
	RestrictedRoles string          `json:"RestrictedRoles"`
	RootItemID      int64           `json:"RootItemId"`
	ShortDesc       string          `json:"ShortDesc"`
	StartingItem    bool            `json:"StartingItem"`
	Type            string          `json:"Type"`
	ItemIconURL     string          `json:"itemIcon_URL"`
	RetMsg          string          `json:"ret_msg"`
}

// ItemDescription stores data related to an Item's Description
type ItemDescription struct {
	Description          string                     `json:"Description"`
	MenuItems            []ItemDescriptionValueItem `json:"Menuitems"`
	SecondaryDescription string                     `json:"SecondaryDescription"`
}

// ItemDescriptionValueItem stores data related to the Value of an Item stat
type ItemDescriptionValueItem struct {
	Description string `json:"Description"`
	Value       string `json:"value"`
}
