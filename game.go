// Package goriot : Currently up to date as of game-v1.3
package goriot

import (
	"encoding/json"
	"strconv"
)

// GRecentGamesDTO : A struct containing recent games information
type GRecentGamesDTO struct {
	Games      []GGameDTO `json:"games"`      // Collection of recent games played (max 10).
	SummonerID int64      `json:"summonerId"` // Summoner ID.
}

// GGameDTO : A struct containing game information
type GGameDTO struct {
	ChampionID    int          `json:"championId"`    // Champion ID associated with game.
	CreateDate    int64        `json:"createDate"`    // Date that end game data was recorded, specified as epoch milliseconds.
	FellowPlayers []GPlayerDTO `json:"fellowPlayers"` // Other players associated with the game.
	GameID        int64        `json:"gameId"`        // Game ID.
	GameMode      string       `json:"gameMode"`      // Game mode. (Legal values: CLASSIC, ODIN, ARAM, TUTORIAL, ONEFORALL, ASCENSION, FIRSTBLOOD, KINGPORO)
	GameType      string       `json:"gameType"`      // Game type. (Legal values: CUSTOM_GAME, MATCHED_GAME, TUTORIAL_GAME)
	Invalid       bool         `json:"invalid"`       // Invalid flag.
	IPEarned      int          `json:"ipEarned"`      // IP Earned.
	Level         int          `json:"level"`         // Level.
	MapID         int          `json:"mapId"`         // Map ID.
	Spell1        int          `json:"spell1"`        // ID of first summoner spell.
	Spell2        int          `json:"spell2"`        // ID of second summoner spell.
	Stats         GRawStatsDTO `json:"stats"`         // Statistics associated with the game for this summoner.
	SubType       string       `json:"subType"`       // Game sub-type. (Legal values: NONE, NORMAL, BOT, RANKED_SOLO_5x5, RANKED_PREMADE_3x3, RANKED_PREMADE_5x5, ODIN_UNRANKED, RANKED_TEAM_3x3, RANKED_TEAM_5x5, NORMAL_3x3, BOT_3x3, CAP_5x5, ARAM_UNRANKED_5x5, ONEFORALL_5x5, FIRSTBLOOD_1x1, FIRSTBLOOD_2x2, SR_6x6, URF, URF_BOT, NIGHTMARE_BOT, ASCENSION, HEXAKILL, KING_PORO, COUNTER_PICK, BILGEWATER)
	TeamID        int          `json:"teamId"`        // Team ID associated with game. Team ID 100 is blue team. Team ID 200 is purple team.
}

// GPlayerDTO : A struct containing player information
type GPlayerDTO struct {
	ChampionID int   `json:"championId"` // Champion id associated with player.
	SummonerID int64 `json:"summonerId"` // Summoner id associated with player.
	TeamID     int   `json:"teamId"`     // Team id associated with player.
}

// GRawStatsDTO : A struct containing raw game stat information
type GRawStatsDTO struct {
	Assists                         int  `json:"assists"`
	BarracksKilled                  int  `json:"barracksKilled"` // Number of enemy inhibitors killed.
	BountyLevel                     int  `json:"bountyLevel"`
	ChampionsKilled                 int  `json:"championsKilled"`
	CombatPlayerScore               int  `json:"combatPlayerScore"`
	ConsumablesPurchased            int  `json:"consumablesPurchased"`
	DamageDealtPlayer               int  `json:"damageDealtPlayer"`
	DoubleKills                     int  `json:"doubleKills"`
	FirstBlood                      int  `json:"firstBlood"`
	Gold                            int  `json:"gold"`
	GoldEarned                      int  `json:"goldEarned"`
	GoldSpent                       int  `json:"goldSpent"`
	Item0                           int  `json:"item0"`
	Item1                           int  `json:"item1"`
	Item2                           int  `json:"item2"`
	Item3                           int  `json:"item3"`
	Item4                           int  `json:"item4"`
	Item5                           int  `json:"item5"`
	Item6                           int  `json:"item6"`
	ItemsPurchased                  int  `json:"itemsPurchased"`
	KillingSprees                   int  `json:"killingSprees"`
	LargestCriticalStrike           int  `json:"largestCriticalStrike"`
	LargestKillingSpree             int  `json:"largestKillingSpree"`
	LargestMultiKill                int  `json:"largestMultiKill"`
	LegendaryItemsCreated           int  `json:"legendaryItemsCreated"` // Number of tier 3 items built.
	Level                           int  `json:"level"`
	MagicDamageDealtPlayer          int  `json:"magicDamageDealtPlayer"`
	MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
	MagicDamageTaken                int  `json:"magicDamageTaken"`
	MinionsDenied                   int  `json:"minionsDenied"`
	MinionsKilled                   int  `json:"minionsKilled"`
	NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
	NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
	NeutralMinionsKilledYourJungle  int  `json:"neutralMinionsKilledYourJungle"`
	NexusKilled                     bool `json:"nexusKilled"` // Flag specifying if the summoner got the killing blow on the nexus.
	NodeCapture                     int  `json:"nodeCapture"`
	NodeCaptureAssist               int  `json:"nodeCaptureAssist"`
	NodeNeutralize                  int  `json:"nodeNeutralize"`
	NodeNeutralizeAssist            int  `json:"nodeNeutralizeAssist"`
	NumDeaths                       int  `json:"numDeaths"`
	NumItemsBought                  int  `json:"numItemsBought"`
	ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
	PentaKills                      int  `json:"pentaKills"`
	PhysicalDamageDealtPlayer       int  `json:"physicalDamageDealtPlayer"`
	PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
	PlayerPosition                  int  `json:"playerPosition"` // Player position (Legal values: TOP(1), MIDDLE(2), JUNGLE(3), BOT(4))
	PlayerRole                      int  `json:"playerRole"`     // Player role (Legal values: DUO(1), SUPPORT(2), CARRY(3), SOLO(4))
	PlayerScore0                    int  `json:"playerScore0"`
	PlayerScore1                    int  `json:"playerScore1"`
	PlayerScore2                    int  `json:"playerScore2"`
	PlayerScore3                    int  `json:"playerScore3"`
	PlayerScore4                    int  `json:"playerScore4"`
	PlayerScore5                    int  `json:"playerScore5"`
	PlayerScore6                    int  `json:"playerScore6"`
	PlayerScore7                    int  `json:"playerScore7"`
	PlayerScore8                    int  `json:"playerScore8"`
	PlayerScore9                    int  `json:"playerScore9"`
	QuadraKills                     int  `json:"quadraKills"`
	SightWardsBought                int  `json:"sightWardsBought"`
	Spell1Cast                      int  `json:"spell1Cast"` // Number of times first champion spell was cast.
	Spell2Cast                      int  `json:"spell2Cast"` // Number of times second champion spell was cast.
	Spell3Cast                      int  `json:"spell3Cast"` // Number of times third champion spell was cast.
	Spell4Cast                      int  `json:"spell4Cast"` // Number of times fourth champion spell was cast.
	SummonSpell1Cast                int  `json:"summonSpell1Cast"`
	SummonSpell2Cast                int  `json:"summonSpell2Cast"`
	SuperMonsterKilled              int  `json:"superMonsterKilled"`
	Team                            int  `json:"team"`
	TeamObjective                   int  `json:"teamObjective"`
	TimePlayed                      int  `json:"timePlayed"`
	TotalDamageDealt                int  `json:"totalDamageDealt"`
	TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
	TotalDamageTaken                int  `json:"totalDamageTaken"`
	TotalHeal                       int  `json:"totalHeal"`
	TotalPlayerScore                int  `json:"totalPlayerScore"`
	TotalScoreRank                  int  `json:"totalScoreRank"`
	TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
	TotalUnitsHealed                int  `json:"totalUnitsHealed"`
	TripleKills                     int  `json:"tripleKills"`
	TrueDamageDealtPlayer           int  `json:"trueDamageDealtPlayer"`
	TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                 int  `json:"trueDamageTaken"`
	TurretsKilled                   int  `json:"turretsKilled"`
	UnrealKills                     int  `json:"unrealKills"`
	VictoryPointTotal               int  `json:"victoryPointTotal"`
	VisionWardsBought               int  `json:"visionWardsBought"`
	WardKilled                      int  `json:"wardKilled"`
	WardPlaced                      int  `json:"wardPlaced"`
	Win                             bool `json:"win"` // Flag specifying whether or not this game was won.
}

// GetRecentGames : Gets a list of recent games for the summoner passed
func (c *RiotClient) GetRecentGames(region string, summonerID int64) (GRecentGamesDTO, error) {
	var recentGamesDTO GRecentGamesDTO
	// Performs the http request on Riots API to retrieve the current games information
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.3/game/by-summoner/"+
		strconv.FormatInt(summonerID, 10)+"/recent", region, nil)
	if err != nil {
		return recentGamesDTO, err
	}

	// Unmarshals the response JSON into our FeaturedGames struct
	if err := json.Unmarshal(resBody, &recentGamesDTO); err != nil {
		return recentGamesDTO, err
	}
	return recentGamesDTO, nil
}
