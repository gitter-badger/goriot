// Package goriot : Currently up to date as of team-v2.4
package goriot

import (
	"encoding/json"
	"strings"
)

// GetTeamsBySummonerIDs : Get teams mapped by summoner ID for a given list of summoner IDs
func (c *RiotClient) GetTeamsBySummonerIDs(region string, summonerIDs ...string) (map[string][]TeamDTO, error) {
	var teams map[string][]TeamDTO
	// Performs the http request on Riots API to retrieve team data
	resBody, err := c.riotRequest("/api/lol/"+region+"/v2.4/team/by-summoner/"+
		strings.Join(summonerIDs, ","), region, nil)
	if err != nil {
		return teams, err
	}

	// Unmarshals the response JSON into a map of TeamDTO arrays
	if err = json.Unmarshal(resBody, &teams); err != nil {
		return teams, err
	}
	return teams, nil
}

// GetTeamsByTeamIDs : Get teams mapped by team ID for a given list of team IDs
func (c *RiotClient) GetTeamsByTeamIDs(region string, teamIDs ...string) (map[string]TeamDTO, error) {
	var teams map[string]TeamDTO
	// Performs the http request on Riots API to retrieve team data
	resBody, err := c.riotRequest("/api/lol/"+region+"/v2.4/team/"+
		strings.Join(teamIDs, ","), region, nil)
	if err != nil {
		return teams, err
	}

	// Unmarshals the response JSON into a map of TeamDTOs
	if err = json.Unmarshal(resBody, &teams); err != nil {
		return teams, err
	}
	return teams, nil
}
