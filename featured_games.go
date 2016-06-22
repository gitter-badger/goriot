// Package goriot : Currently up to date as of featured-games-v1.0
package goriot

import "encoding/json"

// GetFeaturedGames : Gets a list of featured in-progress games
func (c *RiotClient) GetFeaturedGames(region string) (FeaturedGames, error) {
	var featuredGames FeaturedGames
	// Performs the http request on Riots API to retrieve the current games information
	resBody, err := c.riotRequest("/observer-mode/rest/featured", region, nil)
	if err != nil {
		return featuredGames, err
	}

	// Unmarshals the response JSON into our FeaturedGames struct
	if err := json.Unmarshal(resBody, &featuredGames); err != nil {
		return featuredGames, err
	}
	return featuredGames, nil
}
