// Package goriot : Currently up to date as of league-v2.5
package goriot

import (
	"encoding/json"
	"strconv"
	"strings"
)

// GetLeaguesBySummonerID : Get leagues mapped by summoner ID for a given list of summoner IDs
func (c *RiotClient) GetLeaguesBySummonerID(region string, summonerIDs ...int64) (*map[string][]LeagueDTO, error) {
	// Builds the comma delimited string of summoner ids used for query
	sums := ""
	for k, sumID := range summonerIDs {
		if k == 0 {
			sums = strconv.FormatInt(sumID, 10)
		} else {
			sums = sums + "," + strconv.FormatInt(sumID, 10)
		}
	}

	// Performs the http request to Riots API to retrieve all leagues
	resBody, err := c.riotRequest("/api/lol/"+strings.ToLower(region)+
		"/v2.5/league/by-summoner/"+sums, region, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshals the response JSON into our map of leagues
	leaguesMap := new(map[string][]LeagueDTO)
	if err := json.Unmarshal(resBody, leaguesMap); err != nil {
		return nil, err
	}
	return leaguesMap, nil
}

// GetLeagueEntriesBySummonerID : Get league entries mapped by summoner ID for a given list of summoner IDs
func (c *RiotClient) GetLeagueEntriesBySummonerID(region string, summonerIDs ...int64) (*map[string][]LeagueDTO, error) {
	// Builds the comma delimited string of summoner ids used for query
	sums := ""
	for k, sumID := range summonerIDs {
		if k == 0 {
			sums = strconv.FormatInt(sumID, 10)
		} else {
			sums = sums + "," + strconv.FormatInt(sumID, 10)
		}
	}

	// Performs the http request to Riots API to retrieve all leagues
	resBody, err := c.riotRequest("/api/lol/"+strings.ToLower(region)+
		"/v2.5/league/by-summoner/"+sums+"/entry", region, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshals the response JSON into our map of leagues
	leagueEntryMap := new(map[string][]LeagueDTO)
	if err := json.Unmarshal(resBody, leagueEntryMap); err != nil {
		return nil, err
	}
	return leagueEntryMap, nil
}

// TODO SEEMS ITS BROKEN ON RIOTS END RIGHT NOW - FIX LATER
// GetLeaguesByTeamID : Get leagues mapped by team ID for a given list of team IDs
// func (c *RiotClient) GetLeaguesByTeamID(region string, teamIDs ...string) (*map[string][]LeagueDTO, error) {
// 	// Builds the comma delimited string of team ids used for query
// 	teams := ""
// 	for k, teamID := range teamIDs {
// 		if k == 0 {
// 			teams = teamID
// 		} else {
// 			teams = teams + "," + teamID
// 		}
// 	}
//
// 	// Performs the http request to Riots API to retrieve all leagues
// 	resBody, err := c.riotRequest("/api/lol/"+strings.ToLower(region)+"/v2.5/league/by-team/"+teams, region, nil)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Unmarshals the response JSON into our map of leagues
// 	leagueEntryMap := new(map[string][]LeagueDTO)
// 	if err := json.Unmarshal(resBody, leagueEntryMap); err != nil {
// 		return nil, err
// 	}
// 	return leagueEntryMap, nil
// }

// TODO SEEMS ITS BROKEN ON RIOTS END RIGHT NOW - FIX LATER
// GetLeagueEntriesByTeamID : Get league entries mapped by team ID for a given list of team IDs
// func (c *RiotClient) GetLeagueEntriesByTeamID(region string, teamIDs ...string) (*map[string][]LeagueDTO, error) {
// 	// Builds the comma delimited string of team ids used for query
// 	teams := ""
// 	for k, teamID := range teamIDs {
// 		if k == 0 {
// 			teams = teamID
// 		} else {
// 			teams = teams + "," + teamID
// 		}
// 	}
//
// 	// Performs the http request to Riots API to retrieve all leagues
// 	resBody, err := c.riotRequest("/api/lol/"+strings.ToLower(region)+"/v2.5/league/by-team/"+teams+"/entry", region, nil)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Unmarshals the response JSON into our map of leagues
// 	leagueEntryMap := new(map[string][]LeagueDTO)
// 	if err := json.Unmarshal(resBody, leagueEntryMap); err != nil {
// 		return nil, err
// 	}
// 	return leagueEntryMap, nil
// }

// GetChallengerLeague : Get challenger tier league
// func (c *RiotClient) GetChallengerLeague(region string) (*map[string][]LeagueDTO, error) {
// 	// Performs the http request to Riots API to retrieve all challenger league
// 	resBody, err := c.riotRequest("/api/lol/"+strings.ToLower(region)+"/v2.5/league/challenger", region, nil)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Unmarshals the response JSON into our map of leagues
// 	leagueEntryMap := new(map[string][]LeagueDTO)
// 	if err := json.Unmarshal(resBody, leagueEntryMap); err != nil {
// 		return nil, err
// 	}
// 	return leagueEntryMap, nil
// }
