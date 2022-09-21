package utilities

import "github.com/tusharlock10/pego/constants"

var RankMapping = map[uint]string{
	constants.Qualifying:  "Qualifying",
	constants.Bronze5:     "Bronze V",
	constants.Bronze4:     "Bronze IV",
	constants.Bronze3:     "Bronze III",
	constants.Bronze2:     "Bronze II",
	constants.Bronze1:     "Bronze I",
	constants.Silver5:     "Silver V",
	constants.Silver4:     "Silver IV",
	constants.Silver3:     "Silver III",
	constants.Silver2:     "Silver II",
	constants.Silver1:     "Silver I",
	constants.Gold5:       "Gold V",
	constants.Gold4:       "Gold IV",
	constants.Gold3:       "Gold III",
	constants.Gold2:       "Gold II",
	constants.Gold1:       "Gold I",
	constants.Platinum5:   "Platinum V",
	constants.Platinum4:   "Platinum IV",
	constants.Platinum3:   "Platinum III",
	constants.Platinum2:   "Platinum II",
	constants.Platinum1:   "Platinum I",
	constants.Diamond5:    "Diamond V",
	constants.Diamond4:    "Diamond IV",
	constants.Diamond3:    "Diamond III",
	constants.Diamond2:    "Diamond II",
	constants.Diamond1:    "Diamond I",
	constants.Masters:     "Master",
	constants.Grandmaster: "Grandmaster",
}

var QueueMapping = map[uint]string{
	constants.Unknown:                "Unknown",
	constants.CasualSiege:            "Casual Siege",
	constants.TeamDeathmatch:         "Team Deathmatch",
	constants.Onslaught:              "Onslaught",
	constants.RankedKeyboard:         "Ranked Keyboard",
	constants.RankedController:       "Ranked Controller",
	constants.ShootingRange:          "Shooting Range",
	constants.TrainingSiege:          "Training Siege",
	constants.TrainingTeamDeathmatch: "Training Team Deathmatch",
	constants.TrainingOnslaught:      "Training Onslaught",
	constants.TestMaps:               "Test Maps",
}

var PortalMapping = map[uint]string{
	constants.HiRez:   "HiRez",
	constants.Steam:   "Steam",
	constants.PS4:     "PS4",
	constants.XBOX:    "XBOX",
	constants.Switch:  "Switch",
	constants.Discord: "Discord",
	constants.Epic:    "Epic",
}

// Get the Queue from QueueID
func GetQueueName(queueID uint) string {
	queue := QueueMapping[queueID]
	if queue == "" {
		return QueueMapping[0]
	}
	return queue
}

// Get the Rank from RankID
func GetRankName(rank uint) string {
	rankName := RankMapping[rank]
	if rankName == "" {
		return RankMapping[0]
	}
	return rankName
}

// Get the Portal from PortalID
func GetPortalName(portalID uint) string {
	portalName := PortalMapping[portalID]
	if portalName == "" {
		return "Unknown"
	}
	return portalName
}
