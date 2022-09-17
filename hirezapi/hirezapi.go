package hirezapi

import "github.com/tusharlock10/pego/models"

// HiRezAPI is a collection of endpoint methods for interacting with the Hi-Rez API
// Definitions here are taken from the publicly available "Smite API Developer Guide" PDF
type HiRezAPI interface {
	// CreateSession is a required step to Authenticate the developerId/signature for further API use.
	CreateSession() error
	// GetHirezServerStatus returns UP/DOWN status for the primary game/platform environments. Data is cached once a minute.
	GetHiRezServerStatus() ([]models.HiRezServerStatus, error)
	// GetDataUsed returns API Developer daily usage limits and the current status against those limits.
	GetDataUsed() ([]models.DataUsed, error)
	// GetPlayer returns league and other high level data for a particular player.
	GetPlayer(player string) ([]models.Player, error)
	// GetFriends returns the player's friends
	GetFriends(player string) ([]models.Friend, error)
	// GetMatchHistory returns a list of recent matches and high level match statistics for a particular player
	GetMatchHistory(player string) ([]models.Match, error)
	// GetPlayerStatus returns a player status. 0 - offline, 1 - lobby, 2 - champion select, 3 - in game, 4 - online, 5 - unknown
	GetPlayerStatus(player string) ([]models.PlayerStatus, error)
	// SearchPlayers returns playerID values for all names and/or gamerTags containing searchPlayer
	SearchPlayers(searchPlayer string) ([]models.SearchPlayer, error)
	// GetItems returns all items and their various attributes.
	GetItems() ([]models.Item, error)
	// GetMatchDetails returns the statistics for a particular completed match.
	GetMatchDetails(matchID string) ([]models.MatchPlayer, error)
	// GetMatchDetailsBatch returns the statistics for a particular set of completed matches
	GetMatchDetailsBatch(matchIDs []string) ([]models.MatchPlayer, error)
	// GetMatchPlayerDetails returns player information for a live match.
	GetMatchPlayerDetails(matchID string) ([]models.LiveMatchPlayer, error)
	/*GetMatchIDsByQueue lists all MatchIDs for a particular match queue.
	- queueID can be referenced by constants defined in this package (eg, hirezapi.LiveSiege).
	- date must be formatted/formattable by hirezapi.DateFormat (yyyyMMdd).
	- hour may be "0" - "23" and optionally may contain a ten minute window separated by a comma (eg, "6,30").
	- hour may also be "-1" to fetch the whole day, but may stall/fail due to the amount of data.
	*/
	GetMatchIDsByQueue(queueID, date, hour string) ([]models.Match, error)
	// GetPlayerBatch returns league and other high level data for a particular list of players. [20 max]
	GetPlayerBatch(playerIDs []string) ([]models.Player, error)
	// GetChampionRanks returns the rank and worshipper values for each Champion a player has played.
	GetChampionRanks(player string) ([]models.ChampionRank, error)
	// GetChampions returns all Champions and their various attributes.
	GetChampions() ([]models.Champion, error)
	// GetPlayerLoadouts returns deck loadouts per Champion
	GetPlayerLoadouts(player string) ([]models.PlayerLoadout, error)
	// GetChampionCards returns all Champion cards
	GetChampionCards(champID string) ([]models.ChampionCard, error)
	// GetBountyItems returns daily Bounty Item history for the past 6 months.
	GetBountyItems() ([]models.BountyItem, error)
}
