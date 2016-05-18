// Currently up to date as of champion-v1.2
package main

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// GetChampion : Retrieve all Champions
func (c *RiotClient) GetChampions(freeToPlay bool) (*ChampionListDTO, error) {
	// Uses the free to play optional value from func params
	params := &url.Values{
		"freeToPlay": {strconv.FormatBool(freeToPlay)},
	}

	// Performs the http request to Riots API to retrieve all the champions
	resBody, err := c.riotRequest("/api/lol/"+c.Region+"/v1.2/champion", params)
	if err != nil {
		return nil, err
	}

	// Unmarshals the response JSON into our array of champions
	championListDTO := new(ChampionListDTO)
	if err := json.Unmarshal(resBody, championListDTO); err != nil {
		return nil, err
	}
	return championListDTO, nil
}

// GetChampionByID : Gets data on a particular champion
func (c *RiotClient) GetChampionByID(id int) (*ChampionDTO, error) {
	// Performs the http request to Riots API to retrieve the champion data
	resBody, err := c.riotRequest("/api/lol/"+c.Region+"/v1.2/champion/"+strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}

	// Unmarshals the response JSON into our champion object
	championDTO := new(ChampionDTO)
	if err := json.Unmarshal(resBody, championDTO); err != nil {
		return nil, err
	}
	return championDTO, nil
}
