// Package goriot : Currently up to date as of static-v1.2
package goriot

// // StaticOpts : A struct containing optional paramenters used with static endpoint
// type StaticOpts struct {
// 	Locale    string
// 	Version   string
// 	DataByID  bool
// 	ChampData string
// }
//
// // GetStaticChampions : Retrieves champion list
// func (c *RiotClient) GetStaticChampions(region string, opts *StaticOpts) (ChampionListDTO, error) {
// 	var champions ChampionListDTO
// 	params := &url.Values{}
// 	// Builds out query params based on options passed
// 	if opts != nil {
// 		if opts.Locale != "" {
// 			params.Add("locale", opts.Locale)
// 		}
// 		if opts.Version != "" {
// 			params.Add("version", opts.Version)
// 		}
// 		if opts.DataByID == true {
// 			params.Add("dataById", "true")
// 		}
// 		if opts.ChampData != "" {
// 			params.Add("champData", opts.ChampData)
// 		}
// 	}
//
// 	// Performs the http request on Riots API to retrieve version data
// 	resBody, err := c.riotRequest("/api/lol/static-data/"+region+"/v1.2/champion", region, params)
// 	if err != nil {
// 		return champions, err
// 	}
//
// 	// Unmarshals the response JSON into a ChampionListDTO struct
// 	if err = json.Unmarshal(resBody, &champions); err != nil {
// 		return champions, err
// 	}
// 	return champions, nil
// }
//
// // GetStaticChampionByID : Retrieves a champion by its id
// func (c *RiotClient) GetStaticChampionByID(region string, championID int, opts *StaticOpts) (ChampionDTO, error) {
// 	var champion ChampionDTO
// 	params := &url.Values{}
// 	// Builds out query params based on options passed
// 	if opts != nil {
// 		if opts.Locale != "" {
// 			params.Add("locale", opts.Locale)
// 		}
// 		if opts.Version != "" {
// 			params.Add("version", opts.Version)
// 		}
// 		if opts.ChampData != "" {
// 			params.Add("champData", opts.ChampData)
// 		}
// 	}
//
// 	// Performs the http request on Riots API to retrieve version data
// 	resBody, err := c.riotRequest("/api/lol/static-data/"+region+"/v1.2/champion"+strconv.Itoa(championID), region, params)
// 	if err != nil {
// 		return champion, err
// 	}
//
// 	// Unmarshals the response JSON into a ChampionListDTO struct
// 	if err = json.Unmarshal(resBody, &champion); err != nil {
// 		return champion, err
// 	}
// 	return champion, nil
// }
//
// // GetVersions : Retrieve version data
// func (c *RiotClient) GetVersions(region string) ([]string, error) {
// 	var versions []string
// 	// Performs the http request on Riots API to retrieve version data
// 	resBody, err := c.riotRequest("/api/lol/static-data/"+region+"/v1.2/versions", "global", nil)
// 	if err != nil {
// 		return versions, err
// 	}
//
// 	// Unmarshals the response JSON into an array of version strings
// 	if err = json.Unmarshal(resBody, &versions); err != nil {
// 		return versions, err
// 	}
// 	return versions, nil
// }
