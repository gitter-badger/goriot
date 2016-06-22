// Package goriot : Currently up to date as of league-v2.5
package goriot

import (
	"encoding/json"
	"strings"
)

// GetLeaguesBySummonerID : Get leagues mapped by summoner ID for a given list of summoner IDs
func (c *RiotClient) GetLeaguesBySummonerID(region string, summonerIDs ...string) (map[string][]LeagueDTO, error) {
	var leaguesMap map[string][]LeagueDTO
	// Performs the http request to Riots API to retrieve all leagues
	resBody, err := c.riotRequest("/api/lol/"+region+"/v2.5/league/by-summoner/"+
		strings.Join(summonerIDs, ","), region, nil)
	if err != nil {
		return leaguesMap, err
	}

	// Unmarshals the response JSON into our map of leagues
	if err := json.Unmarshal(resBody, &leaguesMap); err != nil {
		return leaguesMap, err
	}
	return leaguesMap, nil
}

// GetLeagueEntriesBySummonerID : Get league entries mapped by summoner ID for a given list of summoner IDs
func (c *RiotClient) GetLeagueEntriesBySummonerID(region string, summonerIDs ...string) (map[string][]LeagueDTO, error) {
	var leagueEntryMap map[string][]LeagueDTO
	// Performs the http request on Riots API to retrieve all leagues
	resBody, err := c.riotRequest("/api/lol/"+region+"/v2.5/league/by-summoner/"+
		strings.Join(summonerIDs, ",")+"/entry", region, nil)
	if err != nil {
		return leagueEntryMap, err
	}

	// Unmarshals the response JSON into our map of leagues
	if err := json.Unmarshal(resBody, &leagueEntryMap); err != nil {
		return leagueEntryMap, err
	}
	return leagueEntryMap, nil
}

// TODO SEEMS ITS BROKEN ON RIOTS END RIGHT NOW - FIX LATER
// GetLeaguesByTeamID : Get leagues mapped by team ID for a given list of team IDs
// func (c *RiotClient) GetLeaguesByTeamID(region string, teamIDs ...string) (*map[string][]LeagueDTO, error) {
// 	// Performs the http request on Riots API to retrieve all leagues
// 	resBody, err := c.riotRequest("/api/lol/"+region+"/v2.5/league/by-team/"+strings.Join(teamIDs, ","), region, nil)
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
// 	// Performs the http request on Riots API to retrieve all leagues
// 	resBody, err := c.riotRequest("/api/lol/"+region+"/v2.5/league/by-team/"+strings.Join(teamIDs, ",")+"/entry", region, nil)
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
// 	resBody, err := c.riotRequest("/api/lol/"+region+"/v2.5/league/challenger", region, nil)
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
