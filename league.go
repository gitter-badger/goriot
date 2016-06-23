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

// LLeagueDTO : A struct containing league informations
type LLeagueDTO struct {
	Entries       []LLeagueEntryDTO `json:"entries"`       // The requested league entries.
	Name          string            `json:"name"`          // This name is an internal place-holder name only. Display and localization of names in the game client are handled client-side.
	ParticipantID string            `json:"participantId"` // Specifies the relevant participant that is a member of this league (i.e., a requested summoner ID, a requested team ID, or the ID of a team to which one of the requested summoners belongs). Only present when full league is requested so that participant's entry can be identified. Not present when individual entry is requested.
	Queue         string            `json:"queue"`         // The league's queue type. (Legal values: RANKED_SOLO_5x5, RANKED_TEAM_3x3, RANKED_TEAM_5x5)
	Tier          string            `json:"tier"`          // The league's tier. (Legal values: CHALLENGER, MASTER, DIAMOND, PLATINUM, GOLD, SILVER, BRONZE)
}

// LLeagueEntryDTO : A struct containing league participant information representing a summoner or team
type LLeagueEntryDTO struct {
	Division         string         `json:"division"`         // The league division of the participant.
	IsFreshBlood     bool           `json:"isFreshBlood"`     // Specifies if the participant is fresh blood.
	IsHotStreak      bool           `json:"isHotStreak"`      // Specifies if the participant is on a hot streak.
	IsInactive       bool           `json:"isInactive"`       // Specifies if the participant is inactive.
	IsVeteran        bool           `json:"isVeteran"`        // Specifies if the participant is a veteran.
	LeaguePoints     int            `json:"leaguePoints"`     // The league points of the participant.
	Losses           int            `json:"losses"`           // The number of losses for the participant.
	MiniSeries       LMiniSeriesDTO `json:"miniSeries"`       // Mini series data for the participant. Only present if the participant is currently in a mini series.
	PlayerOrTeamID   string         `json:"playerOrTeamId"`   // The ID of the participant (i.e., summoner or team) represented by this entry.
	PlayerOrTeamName string         `json:"playerOrTeamName"` // The name of the the participant (i.e., summoner or team) represented by this entry.
	Wins             int            `json:"wins"`             // The number of wins for the participant.
}

// LMiniSeriesDTO : A struct containing mini series information
type LMiniSeriesDTO struct {
	Losses   int    `json:"losses"`   // Number of current losses in the mini series.
	Progress string `json:"progress"` // String showing the current, sequential mini series progress where 'W' represents a win, 'L' represents a loss, and 'N' represents a game that hasn't been played yet.
	Target   int    `json:"target"`   // Number of wins required for promotion.
	Wins     int    `json:"wins"`     // 	Number of current wins in the mini series.
}

// GetLeaguesBySummonerID : Get leagues mapped by summoner ID for a given list of summoner IDs
func (c *RiotClient) GetLeaguesBySummonerID(region string, summonerIDs ...string) (map[string][]LLeagueDTO, error) {
	var leaguesMap map[string][]LLeagueDTO
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
func (c *RiotClient) GetLeagueEntriesBySummonerID(region string, summonerIDs ...string) (map[string][]LLeagueDTO, error) {
	var leagueEntryMap map[string][]LLeagueDTO
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
// NOTE : This func is most likely returning no data as 5's teams are no longer
// a think and can therefore not be in a league
func (c *RiotClient) GetLeaguesByTeamID(region string, teamIDs ...string) (map[string][]LLeagueDTO, error) {
	var leagueEntryMap map[string][]LLeagueDTO
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
// NOTE : This func is most likely returning no data as 5's teams are no longer
// a think and can therefore not be in a league
func (c *RiotClient) GetLeagueEntriesByTeamID(region string, teamIDs ...string) (map[string][]LLeagueDTO, error) {
	var leagueEntryMap map[string][]LLeagueDTO
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
func (c *RiotClient) GetChallengerLeagues(region string, teamType string) (LLeagueDTO, error) {
	var leagueDTO LLeagueDTO
	// Sets params for the type of team
	params := &url.Values{"type": {teamType}}

	// Performs the http request to Riots API to retrieve all challenger leagues
	resBody, err := c.riotRequest("/api/lol/"+region+
		"/v2.5/league/challenger", region, params)
	if err != nil {
		return leagueDTO, err
	}

	// Unmarshals the response JSON into a LLeagueDTO
	if err := json.Unmarshal(resBody, leagueDTO); err != nil {
		return leagueDTO, err
	}
	return leagueDTO, nil
}

// GetMasterLeagues : Get master tier leagues
func (c *RiotClient) GetMasterLeagues(region string, teamType string) (LLeagueDTO, error) {
	var leagueDTO LLeagueDTO
	// Sets params for the type of team
	params := &url.Values{"type": {teamType}}

	// Performs the http request to Riots API to retrieve all master leagues
	resBody, err := c.riotRequest("/api/lol/"+region+
		"/v2.5/league/master", region, params)
	if err != nil {
		return leagueDTO, err
	}

	// Unmarshals the response JSON into a LLeagueDTO
	if err := json.Unmarshal(resBody, leagueDTO); err != nil {
		return leagueDTO, err
	}
	return leagueDTO, nil
}
