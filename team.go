// Package goriot : Currently up to date as of team-v2.4
package goriot

import (
	"encoding/json"
	"strings"
)

// TTeamDTO : A struct containing team information
type TTeamDTO struct {
	CreateDate                    int64                     `json:"createDate"` // Date that team was created specified as epoch milliseconds
	FullID                        string                    `json:"fullId"`
	LastGameDate                  int64                     `json:"lastGameDate"`                  // Date that last game played by team ended specified as epoch milliseconds
	LastJoinDate                  int64                     `json:"lastJoinDate"`                  // Date that last member joined specified as epoch milliseconds
	LastJoinedRankedTeamQueueDate int64                     `json:"lastJoinedRankedTeamQueueDate"` // Date that team last joined the ranked team queue specified as epoch milliseconds
	MatchHistory                  []TMatchHistorySummaryDTO `json:"matchHistory"`
	ModifyDate                    int64                     `json:"modifyDate"` // Date that team was last modified specified as epoch milliseconds
	Name                          string                    `json:"name"`
	Roster                        TRosterDTO                `json:"roster"`
	SecondLastJoinDate            int64                     `json:"secondLastJoinDate"` // Date that second to last member joined specified as epoch milliseconds
	Status                        string                    `json:"status"`
	Tag                           string                    `json:"tag"`
	TeamStatDetails               []TTeamStatDetailDTO      `json:"teamStatDetails"`
	ThirdLastJoinDate             int64                     `json:"thirdLastJoinDate"` // Date that third to last member joined specified as epoch milliseconds
}

// TMatchHistorySummaryDTO : A struct containing a summary of a summoners match history
type TMatchHistorySummaryDTO struct {
	Assists           int    `json:"assists"`
	Date              int64  `json:"date"` // Date that match was completed specified as epoch milliseconds
	Deaths            int    `json:"deaths"`
	GameID            int64  `json:"gameId"`
	GameMode          string `json:"gameMode"`
	Invalid           bool   `json:"invalid"`
	Kills             int    `json:"kills"`
	MapID             int    `json:"mapId"`
	OpposingTeamKills int    `json:"opposingTeamKills"`
	OpposingTeamName  string `json:"opposingTeamName"`
	Win               bool   `json:"win"`
}

// TRosterDTO : A struct containing a games/teams members
type TRosterDTO struct {
	MemberList []TTeamMemberInfoDTO `json:"memberList"`
	OwnerID    int64                `json:"ownerId"`
}

// TTeamStatDetailDTO : A struct containing team statistics detail information
type TTeamStatDetailDTO struct {
	AverageGamesPlayed int    `json:"averageGamesPlayed"`
	Losses             int    `json:"losses"`
	TeamStatType       string `json:"teamStatType"`
	Wins               int    `json:"wins"`
}

// TTeamMemberInfoDTO : A struct containing team member information
type TTeamMemberInfoDTO struct {
	InviteDate int64  `json:"inviteDate"` // Date that team member was invited to team specified as epoch milliseconds
	JoinDate   int64  `json:"joinDate"`   // Date that team member joined team specified as epoch milliseconds
	PlayerID   int64  `json:"playerId"`
	Status     string `json:"status"`
}

// GetTeamsBySummonerIDs : Get teams mapped by summoner ID for a given list of summoner IDs
func (c *RiotClient) GetTeamsBySummonerIDs(region string, summonerIDs ...string) (map[string][]TTeamDTO, error) {
	var teams map[string][]TTeamDTO
	// Performs the http request on Riots API to retrieve team data
	resBody, err := c.riotRequest("/api/lol/"+region+"/v2.4/team/by-summoner/"+
		strings.Join(summonerIDs, ","), region, nil)
	if err != nil {
		return teams, err
	}

	// Unmarshals the response JSON into a map of TTeamDTO arrays
	if err = json.Unmarshal(resBody, &teams); err != nil {
		return teams, err
	}
	return teams, nil
}

// GetTeamsByTeamIDs : Get teams mapped by team ID for a given list of team IDs
func (c *RiotClient) GetTeamsByTeamIDs(region string, teamIDs ...string) (map[string]TTeamDTO, error) {
	var teams map[string]TTeamDTO
	// Performs the http request on Riots API to retrieve team data
	resBody, err := c.riotRequest("/api/lol/"+region+"/v2.4/team/"+
		strings.Join(teamIDs, ","), region, nil)
	if err != nil {
		return teams, err
	}

	// Unmarshals the response JSON into a map of TTeamDTOs
	if err = json.Unmarshal(resBody, &teams); err != nil {
		return teams, err
	}
	return teams, nil
}
