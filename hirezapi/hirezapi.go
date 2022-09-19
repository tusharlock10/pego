package hirezapi

import "github.com/tusharlock10/pego/apiResponse"

// HiRezAPI is a collection of endpoint methods for interacting with the Hi-Rez API
// Definitions here are taken from the publicly available "Smite API Developer Guide" PDF
type HiRezAPI interface {
	// CreateSession is a required step to Authenticate the developerId/signature for further API use.
	CreateSession() error
	// GetHirezServerStatus returns UP/DOWN status for the primary game/platform environments. Data is cached once a minute.
	GetHiRezServerStatus() ([]apiResponse.HiRezServerStatus, error)
	// GetDataUsed returns API Developer daily usage limits and the current status against those limits.
	GetDataUsed() ([]apiResponse.DataUsed, error)
	// GetPlayer returns league and other high level data for a particular player.
	GetPlayer(playerID uint) ([]apiResponse.Player, error)
	// GetFriends returns the player's friends
	GetFriends(playerID uint) ([]apiResponse.Friend, error)
	// GetPlayerMatchHistory returns a list of recent matches and high level match statistics for a particular player
	GetPlayerMatchHistory(playerID uint) ([]apiResponse.PlayerMatchHistory, error)
	// GetPlayerStatus returns a player status. 0 - offline, 1 - lobby, 2 - champion select, 3 - in game, 4 - online, 5 - unknown
	GetPlayerStatus(playerID uint) ([]apiResponse.PlayerStatus, error)
	// SearchPlayers returns playerID values for all names and/or gamerTags containing searchPlayer
	SearchPlayers(searchPlayer string) ([]apiResponse.SearchPlayer, error)
	// GetItems returns all items and their various attributes.
	GetItems() ([]apiResponse.Item, error)
	// GetMatchDetails returns the statistics for a particular completed match.
	GetMatchDetails(matchID uint) ([]apiResponse.MatchPlayerDetail, error)
	// GetMatchDetailsBatch returns the statistics for a particular set of completed matches
	GetMatchDetailsBatch(matchIDs []uint) ([]apiResponse.MatchPlayerDetail, error)
	// GetActiveMatchDetails returns player information for a live match.
	GetActiveMatchDetails(matchID uint) ([]apiResponse.ActiveMatchDetail, error)
	/*GetMatchIDsByQueue lists all MatchIDs for a particular match queue.
	- queueID can be referenced by constants
	- date must be formatted/formattable by hirezapi.DateFormat (yyyyMMdd).
	- hour may be "0" - "23" and optionally may contain a ten minute window separated by a comma (eg, "6,30").
	- hour may also be "-1" to fetch the whole day, but may stall/fail due to the amount of data.
	*/
	GetMatchIDsByQueue(queueID uint, date, hour string) ([]apiResponse.MatchIdsByQueue, error)
	// GetPlayerBatch returns league and other high level data for a particular list of players. [20 max]
	GetPlayerBatch(playerIDs []uint) ([]apiResponse.Player, error)
	// GetChampionRanks returns the rank and worshipper values for each Champion a player has played.
	GetChampionRanks(playerID uint) ([]apiResponse.ChampionRank, error)
	// GetChampions returns all Champions and their various attributes.
	GetChampions() ([]apiResponse.Champion, error)
	// GetPlayerLoadouts returns deck loadouts per Champion
	GetPlayerLoadouts(playerID uint) ([]apiResponse.PlayerLoadout, error)
	// GetChampionCards returns all Champion cards
	GetChampionCards(championID uint) ([]apiResponse.ChampionCard, error)
	// GetBountyItems returns daily Bounty Item history for the past 6 months.
	GetBountyItems() ([]apiResponse.BountyItem, error)
}
