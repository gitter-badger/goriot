// Package goriot : Currently up to date as of static-v1.2
package goriot

import "encoding/json"

// LSDRRealmDTO : A struct containing realm information
type LSDRRealmDTO struct {
	CDN            string            `json:"cdn"`            // The base CDN url
	CSS            string            `json:"css"`            // Latest changed version of Dragon Magic's css file
	DD             string            `json:"dd"`             // Latest changed version of Dragon Magic
	L              string            `json:"l"`              // Default language for this realm
	LG             string            `json:"lg"`             //	Legacy script mode for IE6 or older
	N              map[string]string `json:"n"`              // Latest changed version for each data type listed
	Profileiconmax int               `json:"profileiconmax"` // Special behavior number identifying the largest profileicon id that can be used under 500
	Store          string            `json:"store"`          // Additional api data drawn from other sources that may be related to data dragon functionality
	V              string            `json:"v"`              // Current version of this file for this realm
}

// GetStaticRealm : Retrieve realm data
func (c *RiotClient) GetStaticRealm(region string) (LSDRRealmDTO, error) {
	var realm LSDRRealmDTO

	// Performs the http request on Riots API to retrieve all language strings
	resBody, err := c.riotRequest("/api/lol/static-data/"+region+"/v1.2/realm", "global", nil)
	if err != nil {
		return realm, err
	}

	// Unmarshals the response JSON into a LSDRRealmDTO struct
	if err = json.Unmarshal(resBody, &realm); err != nil {
		return realm, err
	}
	return realm, nil
}
