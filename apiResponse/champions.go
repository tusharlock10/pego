package apiResponse

// ChampionAbility stores data related to a Champion Ability
type ChampionAbility struct {
	Description     string `json:"Description"`
	AbilityID       uint   `json:"Id"`
	Summary         string `json:"Summary"`
	URL             string `json:"URL"`
	DamageType      string `json:"damageType"`
	RechargeSeconds uint   `json:"rechargeSeconds"`
}

// ChampionCard stores data related to a Card for a Champion
type ChampionCard struct {
	ActiveFlagActivationSchedule string `json:"active_flag_activation_schedule"`
	ActiveFlagLTI                string `json:"active_flag_lti"`
	CardDescription              string `json:"card_description"`
	CardID1                      uint   `json:"card_id1"`
	CardID2                      uint   `json:"card_id2"`
	CardName                     string `json:"card_name"`
	CardNameEnglish              string `json:"card_name_english"`
	ChampionCardIconURL          string `json:"championIcon_URL"`
	ChampionTalentURL            string `json:"championTalent_URL"`
	ChampionID                   uint   `json:"champion_id"`
	ChampionName                 string `json:"champion_name"`
	Exclusive                    string `json:"exclusive"`
	Rank                         uint   `json:"rank"`
	Rarity                       string `json:"rarity"`
	RechargeSeconds              uint   `json:"recharge_seconds"`
	RetMsg                       string `json:"ret_msg"`
}

// Champion stores data related to an in game Champion
type Champion struct {
	Ability1               ChampionAbility `json:"Ability_1"`
	Ability2               ChampionAbility `json:"Ability_2"`
	Ability3               ChampionAbility `json:"Ability_3"`
	Ability4               ChampionAbility `json:"Ability_4"`
	Ability5               ChampionAbility `json:"Ability_5"`
	ChampionAbility1URL    string          `json:"ChampionAbility1_URL"`
	ChampionAbility2URL    string          `json:"ChampionAbility2_URL"`
	ChampionAbility3URL    string          `json:"ChampionAbility3_URL"`
	ChampionAbility4URL    string          `json:"ChampionAbility4_URL"`
	ChampionAbility5URL    string          `json:"ChampionAbility5_URL"`
	ChampionAbilityCardURL string          `json:"ChampionCard_URL"`
	ChampionAbilityIconURL string          `json:"ChampionIcon_URL"`
	Health                 uint            `json:"Health"`
	Lore                   string          `json:"Lore"`
	Name                   string          `json:"Name"`
	NameEnglish            string          `json:"Name_English"`
	OnFreeRotation         string          `json:"OnFreeRotation"`
	OnFreeWeeklyRotation   string          `json:"OnFreeWeeklyRotation"`
	Roles                  string          `json:"Roles"`
	Speed                  uint            `json:"Speed"`
	Title                  string          `json:"Title"`
	AbilityDescription1    string          `json:"abilityDescription1"`
	AbilityDescription2    string          `json:"abilityDescription2"`
	AbilityDescription3    string          `json:"abilityDescription3"`
	AbilityDescription4    string          `json:"abilityDescription4"`
	AbilityDescription5    string          `json:"abilityDescription5"`
	ChampionID             uint            `json:"id"`
	LatestChampion         string          `json:"latestChampion"`
	RetMsg                 string          `json:"ret_msg"`
}

// ChampionRank stores data related to player stats for a Champion
type ChampionRank struct {
	Assists     uint   `json:"Assists"`
	Deaths      uint   `json:"Deaths"`
	Gold        uint   `json:"Gold"`
	Kills       uint   `json:"Kills"`
	LastPlayed  string `json:"LastPlayed"`
	Losses      uint   `json:"Losses"`
	MinionKills uint   `json:"MinionKills"`
	Minutes     uint   `json:"Minutes"`
	Rank        uint   `json:"Rank"`
	Wins        uint   `json:"Wins"`
	Worshippers uint   `json:"Worshippers"`
	Champion    string `json:"champion"`
	ChampionID  string `json:"champion_id"`
	PlayerID    string `json:"player_id"`
	RetMsg      string `json:"ret_msg"`
}
