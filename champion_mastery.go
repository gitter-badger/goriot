// Package goriot : Currently up to date as of championmastery-v???
package goriot

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// GetChampionMastery : Get a champion mastery by player id and champion id
func (c *RiotClient) GetChampionMastery(region string, playerID, championID int64) (ChampionMasteryDTO, error) {
	var championMasteryDTO ChampionMasteryDTO
	// Performs the http request on Riots API to retrieve a players champion mastery
	resBody, err := c.riotRequest("/championmastery/location/"+
		region+"1/player/"+strconv.FormatInt(playerID, 10)+
		"/champion/"+strconv.FormatInt(championID, 10), region, nil)
	if err != nil {
		return championMasteryDTO, err
	}

	// Unmarshals the response JSON into our ChampionMasteryDTO struct
	if err := json.Unmarshal(resBody, &championMasteryDTO); err != nil {
		return championMasteryDTO, err
	}
	return championMasteryDTO, nil
}

// GetAllChampionMasteries : Get all champion mastery entries sorted by number of champion points descending
func (c *RiotClient) GetAllChampionMasteries(region string, playerID int64) ([]ChampionMasteryDTO, error) {
	var championMasteryDTOs []ChampionMasteryDTO
	// Performs the http request on Riots API to retrieve a players champion masteries
	resBody, err := c.riotRequest("/championmastery/location/"+
		region+"1/player/"+strconv.FormatInt(playerID, 10)+
		"/champions", region, nil)
	if err != nil {
		return championMasteryDTOs, err
	}

	// Unmarshals the response JSON into our array of ChampionMasteryDTO structs
	if err := json.Unmarshal(resBody, &championMasteryDTOs); err != nil {
		return championMasteryDTOs, err
	}
	return championMasteryDTOs, nil
}

// GetChampionMasteryScore : Get a player's total champion mastery score, which is sum of individual champion mastery levels
func (c *RiotClient) GetChampionMasteryScore(region string, playerID int64) (int, error) {
	var championScore int
	// Performs the http request on Riots API to retrieve a players total mastery score
	resBody, err := c.riotRequest("/championmastery/location/"+
		region+"1/player/"+strconv.FormatInt(playerID, 10)+
		"/score", region, nil)
	if err != nil {
		return championScore, err
	}

	// Unmarshals the response into our champion score integer
	if err := json.Unmarshal(resBody, &championScore); err != nil {
		return championScore, err
	}
	return championScore, nil
}

// GetTopChampionMasteries : Get specified number of top champion mastery entries sorted by number of champion points descending
func (c *RiotClient) GetTopChampionMasteries(region string, playerID int64, count int) ([]ChampionMasteryDTO, error) {
	var championMasteryDTOs []ChampionMasteryDTO
	// Specifies the optional count query parameter
	params := &url.Values{
		"count": {strconv.Itoa(count)},
	}

	// Performs the http request on Riots API to retrieve the players top masteries
	resBody, err := c.riotRequest("/championmastery/location/"+
		region+"1/player/"+strconv.FormatInt(playerID, 10)+
		"/topchampions", region, params)
	if err != nil {
		return championMasteryDTOs, err
	}

	// Unmarshals the response JSON into our array of ChampionMasteryDTO structs
	if err := json.Unmarshal(resBody, &championMasteryDTOs); err != nil {
		return championMasteryDTOs, err
	}
	return championMasteryDTOs, nil
}
