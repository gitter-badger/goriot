// Package goriot : Currently up to date as of static-v1.2
package goriot

import "encoding/json"

// GetVersions : Retrieve version data
func (c *RiotClient) GetVersions(region string) (*[]string, error) {
	// Performs the http request to Riots API to retrieve version data
	resBody, err := c.riotRequest("/api/lol/static-data/"+region+"/v1.2/versions", "global", nil)
	if err != nil {
		return nil, err
	}

	// Unmarshals the response JSON into an array of version strings
	versions := new([]string)
	if err = json.Unmarshal(resBody, versions); err != nil {
		return nil, err
	}
	return versions, nil
}
