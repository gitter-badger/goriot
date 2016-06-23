// Package goriot : Currently up to date as of current-game-v1.0
package goriot

import (
	"encoding/json"
	"strconv"
)

// CGCurrentGameInfo : A struct containing information on a in-progress game
type CGCurrentGameInfo struct {
	BannedChampions   []CGBannedChampion `json:"bannedChampions"`   // Banned champion information
	GameID            int64              `json:"gameId"`            // The ID of the game
	GameLength        int64              `json:"gameLength"`        // The amount of time in seconds that has passed since the game started
	GameMode          string             `json:"gameMode"`          // The game mode (Legal values: CLASSIC, ODIN, ARAM, TUTORIAL, ONEFORALL, ASCENSION, FIRSTBLOOD, KINGPORO)
	GameQueueConfigID int64              `json:"gameQueueConfigId"` // The queue type (queue types are documented on the Game Constants page)
	GameStartTime     int64              `json:"gameStartTime"`     // The game start time represented in epoch milliseconds
	GameType          string             `json:"gameType"`          // The game type (Legal values: CUSTOM_GAME, MATCHED_GAME, TUTORIAL_GAME)
	MapID             int64              `json:"mapId"`             // The ID of the map
	Observers         CGObserver         `json:"observers"`         // The observer information
	Participants      []CGParticipant    `json:"participants"`      // The participant information
	PlatformID        string             `json:"platformId"`        // The ID of the platform on which the game is being played
}

// CGBannedChampion : A struct containing information on a banned champion
type CGBannedChampion struct {
	ChampionID int64 `json:"championId"` // The ID of the banned champion
	PickTurn   int   `json:"pickTurn"`   // The turn during which the champion was banned
	TeamID     int64 `json:"teamId"`     // The ID of the team that banned the champion
}

// CGParticipant : A struct containing information on a player
type CGParticipant struct {
	Bot           bool        `json:"bot"`           // Flag indicating whether or not this participant is a bot
	ChampionID    int64       `json:"championId"`    // The ID of the champion played by this participant
	Masteries     []CGMastery `json:"masteries"`     // The masteries used by this participant
	ProfileIconID int64       `json:"profileIconId"` // The ID of the profile icon used by this participant
	Runes         []CGRune    `json:"runes"`         // The runes used by this participant
	Spell1ID      int64       `json:"spell1Id"`      // The ID of the first summoner spell used by this participant
	Spell2ID      int64       `json:"spell2Id"`      // The ID of the second summoner spell used by this participant
	SummonerID    int64       `json:"summonerId"`    // The summoner ID of this participant
	SummonerName  string      `json:"summonerName"`  // The summoner name of this participant
	TeamID        int64       `json:"teamId"`        // The team ID of this participant, indicating the participant's team
}

// CGObserver : A struct containing information on game observers
type CGObserver struct {
	EncryptionKey string `json:"encryptionKey"` // Key used to decrypt the spectator grid game data for playback
}

// CGMastery : A struct containing information on a players mastery
type CGMastery struct {
	MasteryID int64 `json:"masteryId"` // The ID of the mastery
	Rank      int   `json:"rank"`      // The number of points put into this mastery by the user
}

// CGRune : A struct containing information on a players rune
type CGRune struct {
	Count  int   `json:"count"`  // The count of this rune used by the participant
	RuneID int64 `json:"runeId"` // The ID of the rune
}

// GetCurrentGameInfo : Gets information on the current game by the summoner ID
func (c *RiotClient) GetCurrentGameInfo(region string, summonerID int64) (CGCurrentGameInfo, error) {
	var currentGameInfo CGCurrentGameInfo
	// Performs the http request on Riots API to retrieve the current games information
	resBody, err := c.riotRequest("/observer-mode/rest/consumer/getSpectatorGameInfo/"+
		region+"1/"+strconv.FormatInt(summonerID, 10), region, nil)
	if err != nil {
		return currentGameInfo, err
	}

	// Unmarshals the response JSON into our CGCurrentGameInfo struct
	if err := json.Unmarshal(resBody, &currentGameInfo); err != nil {
		return currentGameInfo, err
	}
	return currentGameInfo, nil
}
