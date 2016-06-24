// Package goriot : Currently up to date as of static-v1.2
package goriot

import (
	"encoding/json"
	"net/url"
)

// StaticLangugesOpts : A struct containing optional paramenters used with static language endpoints
type StaticLangugesOpts struct {
	Locale  string
	Version string
}

// LSDLLanguageStringsDTO : A struct containing languages strings data
type LSDLLanguageStringsDTO struct {
	Data    map[string]string `json:"data"`
	Type    string            `json:"type"`
	Version string            `json:"version"`
}

// GetLanguageStrings : Retrieve language strings data
func (c *RiotClient) GetLanguageStrings(region string, opts *StaticLangugesOpts) (LSDLLanguageStringsDTO, error) {
	var languageStrings LSDLLanguageStringsDTO
	params := &url.Values{}
	// Builds out query params based on options passed
	if opts != nil {
		if opts.Locale != "" {
			params.Add("locale", opts.Locale)
		}
		if opts.Version != "" {
			params.Add("version", opts.Version)
		}
	}

	// Performs the http request on Riots API to retrieve all language strings
	resBody, err := c.riotRequest("/api/lol/static-data/"+region+"/v1.2/language-strings", "global", params)
	if err != nil {
		return languageStrings, err
	}

	// Unmarshals the response JSON into a LSDCChampionListDTO struct
	if err = json.Unmarshal(resBody, &languageStrings); err != nil {
		return languageStrings, err
	}
	return languageStrings, nil
}

// GetLanguages : Retrieve supported languages data
func (c *RiotClient) GetLanguages(region string) ([]string, error) {
	var languageStrings []string

	// Performs the http request on Riots API to retrieve all language strings
	resBody, err := c.riotRequest("/api/lol/static-data/"+region+"/v1.2/languages", "global", nil)
	if err != nil {
		return languageStrings, err
	}

	// Unmarshals the response JSON into a LSDCChampionListDTO struct
	if err = json.Unmarshal(resBody, &languageStrings); err != nil {
		return languageStrings, err
	}
	return languageStrings, nil
}
