package constants

const (
	PaladinsURL string = "https://api.paladins.com/paladinsapi.svc"

	// DateFormat yyyyMMdd
	DateFormat string = "20060102"

	// TimeFormat yyyyMMddHHmmss
	TimeFormat string = "20060102150405"

	// PortalIDs (Gaming Platforms)
	HiRez, Steam, PS4, XBOX, Switch, Discord, Epic = "1", "5", "9", "10", "22", "25", "28"

	// League Tiers
	Bronze5, Bronze4, Bronze3, Bronze2, Bronze1           = "1", "2", "3", "4", "5"
	Silver5, Silver4, Silver3, Silver2, Silver1           = "6", "7", "8", "9", "10"
	Gold5, Gold4, Gold3, Gold2, Gold1                     = "11", "12", "13", "14", "15"
	Platinum5, Platinum4, Platinum3, Platinum2, Platinum1 = "16", "17", "18", "19", "20"
	Diamond5, Diamond4, Diamond3, Diamond2, Diamond1      = "21", "22", "23", "24", "25"
	Masters, Grandmaster                                  = "26", "27"

	// Queue Ids
	Unknown                = "0"
	CasualSiege            = "424"
	TeamDeathmatch         = "469"
	Onslaught              = "452"
	RankedKeyboard         = "486"
	RankedController       = "428"
	ShootingRange          = "434"
	TrainingSiege          = "425"
	TrainingTeamDeathmatch = "470"
	TrainingOnslaught      = "453"
	TestMaps               = "445"
)
