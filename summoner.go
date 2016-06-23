// Package goriot : Currently up to date as of summoner-v1.4
package goriot

import (
	"encoding/json"
	"strings"
)

// SSummonerDTO : A struct containing summoner information
type SSummonerDTO struct {
	ID            int64  `json:"id"`            // Summoner ID
	Name          string `json:"name"`          // Summoner name
	ProfileIconID int    `json:"profileIconId"` // ID of the summoner icon associated with the summoner
	RevisionDate  int64  `json:"revisionDate"`  // Date summoner was last modified specified as epoch milliseconds
	SummonerLevel int64  `json:"summonerLevel"` // 	Summoner level associated with the summoner
}

// SMasteryPagesDTO : A struct containing a users mastery pages
type SMasteryPagesDTO struct {
	Page       []SMasteryPageDTO `json:"pages"`      // Collection of mastery pages associated with the summoner
	SummonerID int64             `json:"summonerId"` // Summoner ID
}

// SMasteryPageDTO : A struct containing information on a mastery page
type SMasteryPageDTO struct {
	Current   bool          `json:"current"`   // Indicates if the mastery page is the current mastery page
	ID        int64         `json:"id"`        //	Mastery page ID
	Masteries []SMasteryDTO `json:"masteries"` // Collection of masteries associated with the mastery page
	Name      string        `json:"name"`      // Mastery page name
}

// SMasteryDTO : A struct containing a users mastery information
type SMasteryDTO struct {
	ID   int `json:"id"`   // Mastery ID. For static information correlating to masteries, please refer to the LoL Static Data API
	Rank int `json:"rank"` // Mastery rank (i.e., the number of points put into this mastery)
}

// SRunePagesDTO : A struct containing a players rune pages
type SRunePagesDTO struct {
	Pages      []SRunePageDTO `json:"pages"`      // Collection of rune pages associated with the summoner
	SummonerID int64          `json:"summonerId"` // Summoner ID
}

// SRunePageDTO : A struct containing a players rune slots
type SRunePageDTO struct {
	Current bool           `json:"current"` // Indicates if the page is the current page
	ID      int64          `json:"id"`      // Rune page ID
	Name    string         `json:"name"`    // Rune page name
	Slots   []SRuneSlotDTO `json:"slots"`   // Collection of rune slots associated with the rune page
}

// SRuneSlotDTO : A struct containing information on a players rune page slot
type SRuneSlotDTO struct {
	RuneID     int `json:"runeId"`     // This object contains rune slot information
	RuneSlotID int `json:"runeSlotId"` // Rune slot ID
}

// GetSummonersByName : Get summoner objects mapped by standardized summoner name for a given list of summoner names
func (c *RiotClient) GetSummonersByName(region string, summonerNames ...string) (map[string]SSummonerDTO, error) {
	var summoners map[string]SSummonerDTO
	// Performs the http request on Riots API to retrieve the summoners information
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.4/summoner/by-name/"+
		strings.Join(summonerNames, ","), region, nil)
	if err != nil {
		return summoners, err
	}

	// Unmarshals the response JSON into a SSummonerDTO map
	if err := json.Unmarshal(resBody, &summoners); err != nil {
		return summoners, err
	}
	return summoners, nil
}

// GetSummonersByID : Get summoner objects mapped by summoner ID for a given list of summoner IDs
func (c *RiotClient) GetSummonersByID(region string, summonerIDs ...string) (map[string]SSummonerDTO, error) {
	var summoners map[string]SSummonerDTO
	// Performs the http request on Riots API to retrieve the summoners information
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.4/summoner/by-name/"+
		strings.Join(summonerIDs, ","), region, nil)
	if err != nil {
		return summoners, err
	}

	// Unmarshals the response JSON into a SSummonerDTO map
	if err := json.Unmarshal(resBody, &summoners); err != nil {
		return summoners, err
	}
	return summoners, nil
}

// GetSummonerNamesByID : Get summoner names mapped by summoner ID for a given list of summoner IDs
func (c *RiotClient) GetSummonerNamesByID(region string, summonerIDs ...string) (map[string]string, error) {
	var names map[string]string
	// Performs the http request on Riots API to retrieve the summoners information
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.4/summoner/"+
		strings.Join(summonerIDs, ",")+"/name", region, nil)
	if err != nil {
		return names, err
	}

	// Unmarshals the response JSON into a map of summoner names
	if err := json.Unmarshal(resBody, &names); err != nil {
		return names, err
	}
	return names, nil
}

// GetMasteryPagesBySummonerID : Get mastery pages mapped by summoner ID for a given list of summoner IDs
func (c *RiotClient) GetMasteryPagesBySummonerID(region string, summonerIDs ...string) (map[string]SMasteryPagesDTO, error) {
	var masteries map[string]SMasteryPagesDTO
	// Performs the http request on Riots API to retrieve the mastery page information
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.4/summoner/"+
		strings.Join(summonerIDs, ",")+"/masteries", region, nil)
	if err != nil {
		return masteries, err
	}

	// Unmarshals the response JSON into a MasteryPagesDTO map
	if err := json.Unmarshal(resBody, &masteries); err != nil {
		return masteries, err
	}
	return masteries, nil
}

// GetRunePagesBySummonerID : Get rune pages mapped by summoner ID for a given list of summoner IDs
func (c *RiotClient) GetRunePagesBySummonerID(region string, summonerIDs ...string) (map[string]SRunePagesDTO, error) {
	var runes map[string]SRunePagesDTO
	// Performs the http request on Riots API to retrieve the summoners information
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.4/summoner/"+
		strings.Join(summonerIDs, ",")+"/runes", region, nil)
	if err != nil {
		return runes, err
	}

	// Unmarshals the response JSON into a RunePagesDTO map
	if err := json.Unmarshal(resBody, &runes); err != nil {
		return runes, err
	}
	return runes, nil
}
