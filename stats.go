// Package goriot : Currently up to date as of stats-v1.3
package goriot

import (
	"encoding/json"
	"net/url"
)

const (
	Season3    = "SEASON3"
	Season2014 = "SEASON2014"
	Season2015 = "SEASON2015"
	Season2016 = "SEASON2016"
)

// GetRankedStats : Get ranked stats by summoner ID
func (c *RiotClient) GetRankedStats(region, summonerID, season string) (RankedStatsDTO, error) {
	var stats RankedStatsDTO
	// Sets params for the season
	params := &url.Values{"season": {season}}

	// Performs the http request on Riots API to retrieve the ranked stats data
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.3/stats/by-summoner/"+
		summonerID+"/ranked", region, params)
	if err != nil {
		return stats, err
	}

	// Unmarshals the response JSON into a RankedStatsDTO struct
	if err := json.Unmarshal(resBody, &stats); err != nil {
		return stats, err
	}
	return stats, nil
}

// GetStatsSummaries : Get player stats summaries by summoner ID
func (c *RiotClient) GetStatsSummaries(region, summonerID, season string) (PlayerStatsSummaryListDTO, error) {
	var stats PlayerStatsSummaryListDTO
	// Sets params for the season
	params := &url.Values{"season": {season}}

	// Performs the http request on Riots API to retrieve the stats summary data
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.3/stats/by-summoner/"+
		summonerID+"/summary", region, params)
	if err != nil {
		return stats, err
	}

	// Unmarshals the response JSON into a RankedStatsDTO struct
	if err := json.Unmarshal(resBody, &stats); err != nil {
		return stats, err
	}
	return stats, nil
}
