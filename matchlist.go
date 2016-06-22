// Package goriot : Currently up to date as of stats-v1.3
package goriot

import (
	"encoding/json"
	"net/url"
)

// const (
// 	Season3    = "SEASON3"
// 	Season2014 = "SEASON2014"
// 	Season2015 = "SEASON2015"
// 	Season2016 = "SEASON2016"
// )

// GetMatchList : Retrieve match list by summoner ID
func (c *RiotClient) GetMatchList(region, summonerID, season string) (MatchList, error) {
	var matches MatchList
	// Sets params for the season
	params := &url.Values{"season": {season}}

	// Performs the http request on Riots API to retrieve the ranked stats data
	resBody, err := c.riotRequest("/api/lol/"+region+
		"/v2.2/matchlist/by-summoner/"+summonerID, region, params)
	if err != nil {
		return matches, err
	}

	// Unmarshals the response JSON into a RankedStatsDTO struct
	if err := json.Unmarshal(resBody, &matches); err != nil {
		return matches, err
	}
	return matches, nil
}
