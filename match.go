// Package goriot : Currently up to date as of match-v2.2
package goriot

import (
	"encoding/json"
	"strconv"
)

// GetMatchByID : Retrieve match by match ID
func (c *RiotClient) GetMatchByID(region string, matchID int64, includeTimeLine bool) (MatchDetail, error) {
	var match MatchDetail
	// Performs the http request on Riots API to retrieve the match information
	resBody, err := c.riotRequest("/api/lol/"+region+"/v2.2/match/"+
		strconv.FormatInt(matchID, 10), region, nil)
	if err != nil {
		return match, err
	}

	// Unmarshals the response JSON into our MatchDetail struct
	if err := json.Unmarshal(resBody, &match); err != nil {
		return match, err
	}
	return match, nil
}
