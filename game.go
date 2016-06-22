// Package goriot : Currently up to date as of game-v1.3
package goriot

import (
	"encoding/json"
	"strconv"
)

// GetRecentGames : Gets a list of recent games for the summoner passed
func (c *RiotClient) GetRecentGames(region string, summonerID int64) (*RecentGamesDTO, error) {
	// Performs the http request on Riots API to retrieve the current games information
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.3/game/by-summoner/"+
		strconv.FormatInt(summonerID, 10)+"/recent", region, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshals the response JSON into our FeaturedGames struct
	recentGamesDTO := new(RecentGamesDTO)
	if err := json.Unmarshal(resBody, recentGamesDTO); err != nil {
		return nil, err
	}
	return recentGamesDTO, nil
}
