// Package goriot : Currently up to date as of championmastery-v???
package goriot

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// ChampionMasteryOpts : A struct containing optional parameters for the championmastery endpoint
type ChampionMasteryOpts struct {
	Count int
}

// CMChampionMasteryDTO : A struct containing single Champion Mastery information for player and champion combination
type CMChampionMasteryDTO struct {
	ChampionID                   int64  `json:"championId"`                   // 	Champion ID for this entry
	ChampionLevel                int    `json:"championLevel"`                // Champion level for specified player and champion combination
	ChampionPoints               int    `json:"championPoints"`               // Total number of champion points for this player and champion combination - they are used to determine championLevel
	ChampionPointsSinceLastLevel int64  `json:"championPointsSinceLastLevel"` // Number of points earned since current level has been achieved. Zero if player reached maximum champion level for this champion
	ChampionPointsUntilNextLevel int64  `json:"championPointsUntilNextLevel"` // Number of points needed to achieve next level. Zero if player reached maximum champion level for this champion
	ChestGranted                 bool   `json:"chestGranted"`                 // Is chest granted for this champion or not in current season
	HighestGrade                 string `json:"highestGrade"`                 // The highest grade of this champion of current season
	LastPlayTime                 int64  `json:"lastPlayTime"`                 // Last time this champion was played by this player - in Unix milliseconds time format
	PlayerID                     int64  `json:"playerId"`                     // Player ID for this entry
}

// GetChampionMastery : Get a champion mastery by player id and champion id
func (c *RiotClient) GetChampionMastery(region string, playerID, championID int64) (CMChampionMasteryDTO, error) {
	var championMasteryDTO CMChampionMasteryDTO
	// Performs the http request on Riots API to retrieve a players champion mastery
	resBody, err := c.riotRequest("/championmastery/location/"+
		region+"1/player/"+strconv.FormatInt(playerID, 10)+
		"/champion/"+strconv.FormatInt(championID, 10), region, nil)
	if err != nil {
		return championMasteryDTO, err
	}

	// Unmarshals the response JSON into our CMChampionMasteryDTO struct
	if err := json.Unmarshal(resBody, &championMasteryDTO); err != nil {
		return championMasteryDTO, err
	}
	return championMasteryDTO, nil
}

// GetAllChampionMasteries : Get all champion mastery entries sorted by number of champion points descending
func (c *RiotClient) GetAllChampionMasteries(region string, playerID int64) ([]CMChampionMasteryDTO, error) {
	var championMasteryDTOs []CMChampionMasteryDTO
	// Performs the http request on Riots API to retrieve a players champion masteries
	resBody, err := c.riotRequest("/championmastery/location/"+
		region+"1/player/"+strconv.FormatInt(playerID, 10)+
		"/champions", region, nil)
	if err != nil {
		return championMasteryDTOs, err
	}

	// Unmarshals the response JSON into our array of CMChampionMasteryDTO structs
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
func (c *RiotClient) GetTopChampionMasteries(region string, playerID int64, opts *ChampionMasteryOpts) ([]CMChampionMasteryDTO, error) {
	var championMasteryDTOs []CMChampionMasteryDTO
	// Builds out query params based on options passed
	params := &url.Values{}
	if opts != nil {
		if opts.Count != 0 {
			params.Add("count", strconv.Itoa(opts.Count))
		}
	}

	// Performs the http request on Riots API to retrieve the players top masteries
	resBody, err := c.riotRequest("/championmastery/location/"+
		region+"1/player/"+strconv.FormatInt(playerID, 10)+
		"/topchampions", region, params)
	if err != nil {
		return championMasteryDTOs, err
	}

	// Unmarshals the response JSON into our array of CMChampionMasteryDTO structs
	if err := json.Unmarshal(resBody, &championMasteryDTOs); err != nil {
		return championMasteryDTOs, err
	}
	return championMasteryDTOs, nil
}
