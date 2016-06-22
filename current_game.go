// Package goriot : Currently up to date as of current-game-v1.0
package goriot

import (
	"encoding/json"
	"strconv"
)

// GetCurrentGameInfo : Gets information on the current game by the summoner ID
func (c *RiotClient) GetCurrentGameInfo(region string, summonerID int64) (CurrentGameInfo, error) {
	var currentGameInfo CurrentGameInfo
	// Performs the http request on Riots API to retrieve the current games information
	resBody, err := c.riotRequest("/observer-mode/rest/consumer/getSpectatorGameInfo/"+
		region+"1/"+strconv.FormatInt(summonerID, 10), region, nil)
	if err != nil {
		return currentGameInfo, err
	}

	// Unmarshals the response JSON into our CurrentGameInfo struct
	if err := json.Unmarshal(resBody, &currentGameInfo); err != nil {
		return currentGameInfo, err
	}
	return currentGameInfo, nil
}
