// Package goriot : Currently up to date as of stats-v1.3
// TODO OPTIONAL PARAMS
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

// MMatchList : A struct containing match list information
type MMatchList struct {
	EndIndex   int               `json:"endIndex"`
	Matches    []MMatchReference `json:"matches"`
	StartIndex int               `json:"startIndex"`
	TotalGames int               `json:"totalGames"`
}

// MMatchReference : A struct containing match refernce information
type MMatchReference struct {
	Champion   int64  `json:"champion"`
	Lane       string `json:"lane"` // Legal values: MID, MIDDLE, TOP, JUNGLE, BOT, BOTTOM
	MatchID    int64  `json:"matchId"`
	PlatformID string `json:"platformId"`
	Queue      string `json:"queue"` // Legal values: TEAM_BUILDER_DRAFT_RANKED_5x5, RANKED_SOLO_5x5, RANKED_TEAM_3x3, RANKED_TEAM_5x5
	Region     string `json:"region"`
	Role       string `json:"role"`   // Legal values: DUO, NONE, SOLO, DUO_CARRY, DUO_SUPPORT
	Season     string `json:"season"` // Legal values: PRESEASON3, SEASON3, PRESEASON2014, SEASON2014, PRESEASON2015, SEASON2015, PRESEASON2016, SEASON2016
	Timestamp  int64  `json:"timestamp"`
}

// GetMatchList : Retrieve match list by summoner ID
func (c *RiotClient) GetMatchList(region, summonerID, season string) (MMatchList, error) {
	var matches MMatchList
	// Sets params for the season
	params := &url.Values{"season": {season}}

	// Performs the http request on Riots API to retrieve the ranked stats data
	resBody, err := c.riotRequest("/api/lol/"+region+
		"/v2.2/matchlist/by-summoner/"+summonerID, region, params)
	if err != nil {
		return matches, err
	}

	// Unmarshals the response JSON into a MMatchList struct
	if err := json.Unmarshal(resBody, &matches); err != nil {
		return matches, err
	}
	return matches, nil
}
