// Package goriot : Currently up to date as of league-v2.5
package goriot

import (
	"encoding/json"
	"net/url"
	"strings"
)

const (
	RankedSolo5x5 = "RANKED_SOLO_5x5"
	RankedTeam3x3 = "RANKED_TEAM_3x3"
	RankedTeam5x5 = "RANKED_TEAM_5x5"
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

// GetLeaguesByTeamID : Get leagues mapped by team ID for a given list of team IDs
func (c *RiotClient) GetLeaguesByTeamID(region string, teamIDs ...string) (map[string][]LeagueDTO, error) {
	var leagueEntryMap map[string][]LeagueDTO
	// Performs the http request on Riots API to retrieve all leagues
	resBody, err := c.riotRequest("/api/lol/"+region+"/v2.5/league/by-team/"+
		strings.Join(teamIDs, ","), region, nil)
	if err != nil {
		return leagueEntryMap, err
	}

	// Unmarshals the response JSON into our map of leagues
	if err := json.Unmarshal(resBody, &leagueEntryMap); err != nil {
		return leagueEntryMap, err
	}
	return leagueEntryMap, nil
}

// GetLeagueEntriesByTeamID : Get league entries mapped by team ID for a given list of team IDs
func (c *RiotClient) GetLeagueEntriesByTeamID(region string, teamIDs ...string) (map[string][]LeagueDTO, error) {
	var leagueEntryMap map[string][]LeagueDTO
	// Performs the http request on Riots API to retrieve all leagues
	resBody, err := c.riotRequest("/api/lol/"+region+"/v2.5/league/by-team/"+
		strings.Join(teamIDs, ",")+"/entry", region, nil)
	if err != nil {
		return leagueEntryMap, err
	}

	// Unmarshals the response JSON into our map of leagues
	if err := json.Unmarshal(resBody, &leagueEntryMap); err != nil {
		return leagueEntryMap, err
	}
	return leagueEntryMap, nil
}

// GetChallengerLeagues : Get challenger tier leagues
func (c *RiotClient) GetChallengerLeagues(region string, teamType string) (LeagueDTO, error) {
	var leagueDTO LeagueDTO
	// Sets params for the type of team
	params := &url.Values{"type": {teamType}}

	// Performs the http request to Riots API to retrieve all challenger leagues
	resBody, err := c.riotRequest("/api/lol/"+region+
		"/v2.5/league/challenger", region, params)
	if err != nil {
		return leagueDTO, err
	}

	// Unmarshals the response JSON into a LeagueDTO
	if err := json.Unmarshal(resBody, leagueDTO); err != nil {
		return leagueDTO, err
	}
	return leagueDTO, nil
}

// GetMasterLeagues : Get master tier leagues
func (c *RiotClient) GetMasterLeagues(region string, teamType string) (LeagueDTO, error) {
	var leagueDTO LeagueDTO
	// Sets params for the type of team
	params := &url.Values{"type": {teamType}}

	// Performs the http request to Riots API to retrieve all master leagues
	resBody, err := c.riotRequest("/api/lol/"+region+
		"/v2.5/league/master", region, params)
	if err != nil {
		return leagueDTO, err
	}

	// Unmarshals the response JSON into a LeagueDTO
	if err := json.Unmarshal(resBody, leagueDTO); err != nil {
		return leagueDTO, err
	}
	return leagueDTO, nil
}
