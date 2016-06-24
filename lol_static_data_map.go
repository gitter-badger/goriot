// Package goriot : Currently up to date as of static-v1.2
package goriot

import (
	"encoding/json"
	"net/url"
)

// StaticMapOpts : A struct containing optional paramenters used with static map endpoints
type StaticMapOpts struct {
	Locale  string
	Version string
}

// LSDMMapDataDTO : A struct containing map data
type LSDMMapDataDTO struct {
	Data    map[string]LSDMMapDetailsDTO `json:"data"`
	Type    string                       `json:"type"`
	Version string                       `json:"version"`
}

// LSDMMapDetailsDTO : A struct containing map detail data
type LSDMMapDetailsDTO struct {
	Image                 LSDMImageDTO `json:"image"`
	MapID                 int64        `json:"mapId"`
	MapName               string       `json:"mapName"`
	UnpurchasableItemList []int64      `json:"unpurchasableItemList"`
}

// LSDMImageDTO : A struct containing image data
type LSDMImageDTO struct {
	Full   string `json:"full"`
	Group  string `json:"group"`
	H      int    `json:"h"`
	Sprite string `json:"sprite"`
	W      int    `json:"w"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

// GetStaticMaps : Retrieve map data
func (c *RiotClient) GetStaticMaps(region string, opts *StaticMapOpts) (LSDMMapDataDTO, error) {
	var maps LSDMMapDataDTO
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
	resBody, err := c.riotRequest("/api/lol/static-data/"+region+"/v1.2/map", "global", params)
	if err != nil {
		return maps, err
	}

	// Unmarshals the response JSON into a LSDCChampionListDTO struct
	if err = json.Unmarshal(resBody, &maps); err != nil {
		return maps, err
	}
	return maps, nil
}
