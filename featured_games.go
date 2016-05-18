// Currently up to date as of featured-games-v1.0
package main

import "encoding/json"

// GetFeaturedGames : Gets a list of featured in-progress games
func (c *RiotClient) GetFeaturedGames() (*FeaturedGames, error) {
	// Performs the http request to Riots API to retrieve the current games information
	resBody, err := c.riotRequest("/observer-mode/rest/featured", nil)
	if err != nil {
		return nil, err
	}

	// Unmarshals the response JSON into our FeaturedGames struct
	featuredGames := new(FeaturedGames)
	if err := json.Unmarshal(resBody, featuredGames); err != nil {
		return nil, err
	}
	return featuredGames, nil
}
