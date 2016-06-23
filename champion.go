// Package goriot : Currently up to date as of champion-v1.2
package goriot

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// ChampionOpts : A struct containing optional parameters for the champion endpoint
type ChampionOpts struct {
	FreeToPlay bool
}

// CChampionListDTO : A struct containing a collection of champion information
type CChampionListDTO struct {
	Champions []CChampionDTO `json:"champions"` // 	The collection of champion information.
}

// CChampionDTO : A struct containing champion information
type CChampionDTO struct {
	Active            bool `json:"active"`            // Indicates if the champion is active.
	BotEnabled        bool `json:"botEnabled"`        // Bot enabled flag (for custom games).
	BotMmEnabled      bool `json:"botMmEnabled"`      // Bot Match Made enabled flag (for Co-op vs. AI games).
	FreeToPlay        bool `json:"freeToPlay"`        // Indicates if the champion is free to play. Free to play champions are rotated periodically.
	ID                int  `json:"id"`                // Champion ID. For static information correlating to champion IDs, please refer to the LoL Static Data API.
	RankedPlayEnabled bool `json:"rankedPlayEnabled"` // Ranked play enabled flag.
}

// GetChampions : Retrieve all Champions
func (c *RiotClient) GetChampions(region string, opts *ChampionOpts) (CChampionListDTO, error) {
	var championListDTO CChampionListDTO
	// Builds out query params based on options passed
	params := &url.Values{}
	if opts != nil {
		if opts.FreeToPlay == true {
			params.Add("freeToPlay", "true")
		}
	}

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
func (c *RiotClient) GetChampionByID(region string, id int) (CChampionDTO, error) {
	var championDTO CChampionDTO
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
