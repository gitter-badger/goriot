// Package goriot : Currently up to date as of static-v1.2
package goriot

import "encoding/json"

// GetVersions : Retrieve version data
func (c *RiotClient) GetVersions(region string) ([]string, error) {
	var versions []string
	// Performs the http request on Riots API to retrieve version data
	resBody, err := c.riotRequest("/api/lol/static-data/"+region+"/v1.2/versions", "global", nil)
	if err != nil {
		return versions, err
	}

	// Unmarshals the response JSON into an array of version strings
	if err = json.Unmarshal(resBody, &versions); err != nil {
		return versions, err
	}
	return versions, nil
}
