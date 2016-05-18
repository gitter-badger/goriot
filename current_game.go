// Currently up to date as of current-game-v1.0
package main

import (
	"encoding/json"
	"strconv"
	"strings"
)

// GetCurrentGameInfo : Gets information on the current game by the summoner ID
func (c *RiotClient) GetCurrentGameInfo(summonerID int64) (*CurrentGameInfo, error) {
	// Performs the http request to Riots API to retrieve the current games information
	resBody, err := c.riotRequest("/observer-mode/rest/consumer/getSpectatorGameInfo/"+
		strings.ToUpper(c.Region)+"1/"+strconv.FormatInt(summonerID, 10), nil)
	if err != nil {
		return nil, err
	}

	// Unmarshals the response JSON into our CurrentGameInfo struct
	currentGameInfo := new(CurrentGameInfo)
	if err := json.Unmarshal(resBody, currentGameInfo); err != nil {
		return nil, err
	}
	return currentGameInfo, nil
}
