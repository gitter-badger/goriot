// Package goriot : Currently up to date as of champion-v1.2
package goriot

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// GetChampions : Retrieve all Champions
func (c *RiotClient) GetChampions(region string, freeToPlay bool) (ChampionListDTO, error) {
	var championListDTO ChampionListDTO
	// Uses the free to play optional value from func params
	params := &url.Values{"freeToPlay": {strconv.FormatBool(freeToPlay)}}

	// Performs the http request on Riots API to retrieve all the champions
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.2/champion", region, params)
	if err != nil {
		return championListDTO, err
	}

	// Unmarshals the response JSON into our array of champions
	if err := json.Unmarshal(resBody, &championListDTO); err != nil {
		return championListDTO, err
	}
	return championListDTO, nil
}

// GetChampionByID : Gets data on a particular champion
func (c *RiotClient) GetChampionByID(region string, id int) (ChampionDTO, error) {
	var championDTO ChampionDTO
	// Performs the http request on Riots API to retrieve the champion data
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.2/champion/"+strconv.Itoa(id), region, nil)
	if err != nil {
		return championDTO, err
	}

	// Unmarshals the response JSON into our champion object
	if err := json.Unmarshal(resBody, &championDTO); err != nil {
		return championDTO, err
	}
	return championDTO, nil
}
