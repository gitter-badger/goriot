// Package goriot : Currently up to date as of stats-v1.3
package goriot

import (
	"encoding/json"
	"net/url"
)

const (
// Season3 = "SEASON3"
// Season2014 = "SEASON2014"
// Season2015 = "SEASON2015"
// Season2016 = "SEASON2016"
)

// SRankedStatsDTO : A struct containing ranked stats information
type SRankedStatsDTO struct {
	Champions  []SChampionStatsDTO `json:"champions"`  // Collection of aggregated stats summarized by champion
	ModifyDate int64               `json:"modifyDate"` // Date stats were last modified specified as epoch milliseconds
	SummonerID int64               `json:"summonerId"` // Summoner ID
}

// SChampionStatsDTO : A struct containing a collection of champion stats information
type SChampionStatsDTO struct {
	ID    int                 `json:"id"`    // Champion ID
	Stats SAggregatedStatsDTO `json:"stats"` // Aggregated stats associated with the champion
}

// SPlayerStatsSummaryListDTO : A struct containing a collection of player stats summary information
type SPlayerStatsSummaryListDTO struct {
	PlayerStatSummaries []SPlayerStatsSummaryDTO `json:"playerStatSummaries"` // Collection of player stats summaries associated with the summoner
	SummonerID          int64                    `json:"summonerId"`          // Summoner ID
}

// SPlayerStatsSummaryDTO : A struct containing player stats summary information
type SPlayerStatsSummaryDTO struct {
	AggregatedStats       SAggregatedStatsDTO `json:"aggregatedStats"`       // Aggregated stats
	Losses                int                 `json:"losses"`                // Number of losses for this queue type. Returned for ranked queue types only
	ModifyDate            int64               `json:"modifyDate"`            // Date stats were last modified specified as epoch milliseconds
	PlayerStatSummaryType string              `json:"playerStatSummaryType"` // Player stats summary type
	Wins                  int                 `json:"wins"`                  //	Number of wins for this queue type
}

// SAggregatedStatsDTO : A struct containing aggregated stats information
type SAggregatedStatsDTO struct {
	AverageAssists              int `json:"averageAssists"`              // Dominion only
	AverageChampionsKilled      int `json:"averageChampionsKilled"`      // Dominion only
	AverageCombatPlayerScore    int `json:"averageCombatPlayerScore"`    // Dominion only
	AverageNodeCapture          int `json:"averageNodeCapture"`          // Dominion only
	AverageNodeCaptureAssist    int `json:"averageNodeCaptureAssist"`    // Dominion only
	AverageNodeNeutralize       int `json:"averageNodeNeutralize"`       // Dominion only
	AverageNodeNeutralizeAssist int `json:"averageNodeNeutralizeAssist"` // Dominion only
	AverageNumDeaths            int `json:"averageNumDeaths"`            // Dominion only
	AverageObjectivePlayerScore int `json:"averageObjectivePlayerScore"` // Dominion only
	AverageTeamObjective        int `json:"averageTeamObjective"`        // Dominion only
	AverageTotalPlayerScore     int `json:"averageTotalPlayerScore"`     // Dominion only
	BotGamesPlayed              int `json:"botGamesPlayed"`
	KillingSpree                int `json:"killingSpree"`
	MaxAssists                  int `json:"maxAssists"` // Dominion only
	MaxChampionsKilled          int `json:"maxChampionsKilled"`
	MaxCombatPlayerScore        int `json:"maxCombatPlayerScore"` // Dominion only
	MaxLargestCriticalStrike    int `json:"maxLargestCriticalStrike"`
	MaxLargestKillingSpree      int `json:"maxLargestKillingSpree"`
	MaxNodeCapture              int `json:"maxNodeCapture"`          // Dominion only
	MaxNodeCaptureAssist        int `json:"maxNodeCaptureAssist"`    // Dominion only
	MaxNodeNeutralize           int `json:"maxNodeNeutralize"`       // Dominion only
	MaxNodeNeutralizeAssist     int `json:"maxNodeNeutralizeAssist"` // Dominion only
	MaxNumDeaths                int `json:"maxNumDeaths"`            // Only returned for ranked statistics
	MaxObjectivePlayerScore     int `json:"maxObjectivePlayerScore"` // Dominion only
	MaxTeamObjective            int `json:"maxTeamObjective"`        // Dominion only
	MaxTimePlayed               int `json:"maxTimePlayed"`
	MaxTimeSpentLiving          int `json:"maxTimeSpentLiving"`
	MaxTotalPlayerScore         int `json:"maxTotalPlayerScore"` // Dominion only
	MostChampionKillsPerSession int `json:"mostChampionKillsPerSession"`
	MostSpellsCast              int `json:"mostSpellsCast"`
	NormalGamesPlayed           int `json:"normalGamesPlayed"`
	RankedPremadeGamesPlayed    int `json:"rankedPremadeGamesPlayed"`
	RankedSoloGamesPlayed       int `json:"rankedSoloGamesPlayed"`
	TotalAssists                int `json:"totalAssists"`
	TotalChampionKills          int `json:"totalChampionKills"`
	TotalDamageDealt            int `json:"totalDamageDealt"`
	TotalDamageTaken            int `json:"totalDamageTaken"`
	TotalDeathsPerSession       int `json:"totalDeathsPerSession"` // Only returned for ranked statistics
	TotalDoubleKills            int `json:"totalDoubleKills"`
	TotalFirstBlood             int `json:"totalFirstBlood"`
	TotalGoldEarned             int `json:"totalGoldEarned"`
	TotalHeal                   int `json:"totalHeal"`
	TotalMagicDamageDealt       int `json:"totalMagicDamageDealt"`
	TotalMinionKills            int `json:"totalMinionKills"`
	TotalNeutralMinionsKilled   int `json:"totalNeutralMinionsKilled"`
	TotalNodeCapture            int `json:"totalNodeCapture"`    // Dominion only
	TotalNodeNeutralize         int `json:"totalNodeNeutralize"` // Dominion only
	TotalPentaKills             int `json:"totalPentaKills"`
	TotalPhysicalDamageDealt    int `json:"totalPhysicalDamageDealt"`
	TotalQuadraKills            int `json:"totalQuadraKills"`
	TotalSessionsLost           int `json:"totalSessionsLost"`
	TotalSessionsPlayed         int `json:"totalSessionsPlayed"`
	TotalSessionsWon            int `json:"totalSessionsWon"`
	TotalTripleKills            int `json:"totalTripleKills"`
	TotalTurretsKilled          int `json:"totalTurretsKilled"`
	TotalUnrealKills            int `json:"totalUnrealKills"`
}

// GetRankedStats : Get ranked stats by summoner ID
func (c *RiotClient) GetRankedStats(region, summonerID, season string) (SRankedStatsDTO, error) {
	var stats SRankedStatsDTO
	// Sets params for the season
	params := &url.Values{"season": {season}}

	// Performs the http request on Riots API to retrieve the ranked stats data
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.3/stats/by-summoner/"+
		summonerID+"/ranked", region, params)
	if err != nil {
		return stats, err
	}

	// Unmarshals the response JSON into a SRankedStatsDTO struct
	if err := json.Unmarshal(resBody, &stats); err != nil {
		return stats, err
	}
	return stats, nil
}

// GetStatsSummaries : Get player stats summaries by summoner ID
func (c *RiotClient) GetStatsSummaries(region, summonerID, season string) (SPlayerStatsSummaryListDTO, error) {
	var stats SPlayerStatsSummaryListDTO
	// Sets params for the season
	params := &url.Values{"season": {season}}

	// Performs the http request on Riots API to retrieve the stats summary data
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.3/stats/by-summoner/"+
		summonerID+"/summary", region, params)
	if err != nil {
		return stats, err
	}

	// Unmarshals the response JSON into a SRankedStatsDTO struct
	if err := json.Unmarshal(resBody, &stats); err != nil {
		return stats, err
	}
	return stats, nil
}
