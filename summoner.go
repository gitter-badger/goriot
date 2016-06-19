// Currently up to date as of summoner-v1.4
package goriot

import "encoding/json"

// GetSummonerByName : Get summoner objects mapped by standardized summoner name for a given list of summoner names
func (c *RiotClient) GetSummonerByName(region string, summonerName ...string) (*map[string]SummonerDTO, error) {
	// Builds the comma delimited string of summoner names used for query
	sums := ""
	for k, name := range summonerName {
		if k == 0 {
			sums = name
		} else {
			sums = sums + "," + name
		}
	}

	// Performs the http request to Riots API to retrieve the summoners information
	resBody, err := c.riotRequest("/api/lol/"+region+"/v1.4/summoner/by-name/"+sums, region, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshals the response JSON into a SummonerDTO map
	summoners := new(map[string]SummonerDTO)
	if err := json.Unmarshal(resBody, summoners); err != nil {
		return nil, err
	}

	return summoners, nil
}
