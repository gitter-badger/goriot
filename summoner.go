// Package goriot : Currently up to date as of summoner-v1.4
package goriot

import (
	"encoding/json"
	"strings"
)

// GetSummonersByName : Get summoner objects mapped by standardized summoner name for a given list of summoner names
func (c *RiotClient) GetSummonersByName(region string, summonerNames ...string) (map[string]SummonerDTO, error) {
	var summoners map[string]SummonerDTO
	// Performs the http request on Riots API to retrieve the summoners information
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.4/summoner/by-name/"+
		strings.Join(summonerNames, ","), region, nil)
	if err != nil {
		return summoners, err
	}

	// Unmarshals the response JSON into a SummonerDTO map
	if err := json.Unmarshal(resBody, &summoners); err != nil {
		return summoners, err
	}
	return summoners, nil
}

// GetSummonersByID : Get summoner objects mapped by summoner ID for a given list of summoner IDs
func (c *RiotClient) GetSummonersByID(region string, summonerIDs ...string) (map[string]SummonerDTO, error) {
	var summoners map[string]SummonerDTO
	// Performs the http request on Riots API to retrieve the summoners information
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.4/summoner/by-name/"+
		strings.Join(summonerIDs, ","), region, nil)
	if err != nil {
		return summoners, err
	}

	// Unmarshals the response JSON into a SummonerDTO map
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

	// Unmarshals the response JSON into a SummonerDTO map
	if err := json.Unmarshal(resBody, &names); err != nil {
		return names, err
	}
	return names, nil
}

// GetMasteryPagesBySummonerID : Get mastery pages mapped by summoner ID for a given list of summoner IDs
func (c *RiotClient) GetMasteryPagesBySummonerID(region string, summonerIDs ...string) (map[string]MasteryPagesDTO, error) {
	var masteries map[string]MasteryPagesDTO
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
func (c *RiotClient) GetRunePagesBySummonerID(region string, summonerIDs ...string) (map[string]RunePagesDTO, error) {
	var runes map[string]RunePagesDTO
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
