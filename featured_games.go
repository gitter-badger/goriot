// Package goriot : Currently up to date as of featured-games-v1.0
package goriot

import "encoding/json"

// FGFeaturedGames : A collection of featured in-progress games
type FGFeaturedGames struct {
	ClientRefreshInterval int64                `json:"clientRefreshInterval"` // The suggested interval to wait before requesting FeaturedGames again
	GameList              []FGFeaturedGameInfo `json:"gameList"`              // The list of featured games
}

// FGFeaturedGameInfo : A struct containing information on a particular featured in-progress game
type FGFeaturedGameInfo struct {
	BannedChampions   []FGBannedChampion `json:"bannedChampions"`   // Banned champion information
	GameID            int64              `json:"gameId"`            // The ID of the game
	GameLength        int64              `json:"gameLength"`        // The amount of time in seconds that has passed since the game started
	GameMode          string             `json:"gameMode"`          // The game mode (Legal values: CLASSIC, ODIN, ARAM, TUTORIAL, ONEFORALL, ASCENSION, FIRSTBLOOD, KINGPORO)
	GameQueueConfigID int64              `json:"gameQueueConfigId"` // The queue type (queue types are documented on the Game Constants page)
	GameStartTime     int64              `json:"gameStartTime"`     // The game start time represented in epoch milliseconds
	GameType          string             `json:"gameType"`          // The game type (Legal values: CUSTOM_GAME, MATCHED_GAME, TUTORIAL_GAME)
	MapID             int64              `json:"mapId"`             // The ID of the map
	Observers         FGObserver         `json:"observers"`         // The observer information
	Participants      []FGParticipant    `json:"participants"`      // The participant information
	PlatformID        string             `json:"platformId"`        // The ID of the platform on which the game is being played
}

// FGBannedChampion : A struct containing information on a banned champion
type FGBannedChampion struct {
	ChampionID int64 `json:"championId"` // The ID of the banned champion
	PickTurn   int   `json:"pickTurn"`   // The turn during which the champion was banned
	TeamID     int64 `json:"teamId"`     // The ID of the team that banned the champion
}

// FGObserver : A struct containing information on game observers
type FGObserver struct {
	EncryptionKey string `json:"encryptionKey"` // Key used to decrypt the spectator grid game data for playback
}

// FGParticipant : A struct containing information on a game participant
type FGParticipant struct {
	Bot           bool   `json:"bot"`           // Flag indicating whether or not this participant is a bot
	ChampionID    int64  `json:"championId"`    // The ID of the champion played by this participant
	ProfileIconID int64  `json:"profileIconId"` // The ID of the profile icon used by this participant
	Spell1ID      int64  `json:"spell1Id"`      // The ID of the first summoner spell used by this participant
	Spell2ID      int64  `json:"spell2Id"`      // The ID of the second summoner spell used by this participant
	SummonerName  string `json:"summonerName"`  // The summoner name of this participant
	TeamID        int64  `json:"teamId"`        // The team ID of this participant, indicating the participant's team
}

// GetFeaturedGames : Gets a list of featured in-progress games
func (c *RiotClient) GetFeaturedGames(region string) (FGFeaturedGames, error) {
	var featuredGames FGFeaturedGames
	// Performs the http request on Riots API to retrieve the current games information
	resBody, err := c.riotRequest("/observer-mode/rest/featured", region, nil)
	if err != nil {
		return featuredGames, err
	}

	// Unmarshals the response JSON into our FeaturedGames struct
	if err := json.Unmarshal(resBody, &featuredGames); err != nil {
		return featuredGames, err
	}
	return featuredGames, nil
}
