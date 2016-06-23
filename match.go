// Package goriot : Currently up to date as of match-v2.2
package goriot

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// MatchOpts : A struct containing optional parameters for the match endpoint
type MatchOpts struct {
	IncludeTimeline bool
}

// MMatchDetail : A struct containing a match detail information
type MMatchDetail struct {
	MapID                 int                    `json:"mapId"`                 // Match map ID
	MatchCreation         int64                  `json:"matchCreation"`         // Match creation time
	MatchDuration         int64                  `json:"matchDuration"`         // Match duration
	MatchID               int64                  `json:"matchId"`               // ID of the match
	MatchMode             string                 `json:"matchMode"`             // Match mode (Legal values: CLASSIC, ODIN, ARAM, TUTORIAL, ONEFORALL, ASCENSION, FIRSTBLOOD, KINGPORO)
	MatchType             string                 `json:"matchType"`             // Match type (Legal values: CUSTOM_GAME, MATCHED_GAME, TUTORIAL_GAME)
	MatchVersion          string                 `json:"matchVersion"`          // Match version
	ParticipantIdentities []MParticipantIdentity `json:"participantIdentities"` // Participant identity information
	Participants          []MMatchParticipant    `json:"participants"`          // Participant information
	PlatformID            string                 `json:"platformId"`            // Platform ID of the match
	QueueType             string                 `json:"queueType"`             // Match queue type
	Region                string                 `json:"region"`                // Region where the match was played
	Season                string                 `json:"season"`                // Season match was played (Legal values: PRESEASON3, SEASON3, PRESEASON2014, SEASON2014, PRESEASON2015, SEASON2015, PRESEASON2016, SEASON2016)
	Teams                 []MTeam                `json:"teams"`                 // Team information
	Timeline              MTimeline              `json:"timeline"`              // Match timeline data (not included by default)
}

// MParticipant : A struct containing match participant information
type MParticipant struct {
	ChampionID                int                  `json:"championId"`                //	Champion ID
	HighestAchievedSeasonTier string               `json:"HighestAchievedSeasonTier"` // Highest ranked tier achieved for the previous season, if any, otherwise null
	Masteries                 []MMastery           `json:"masteries"`                 // List of mastery information
	ParticipantID             int                  `json:"participantId"`             // Participant ID
	Runes                     []MRune              `json:"runes"`                     // List of rune information
	Spell1ID                  int                  `json:"spell1Id"`                  //	First summoner spell ID
	Spell2ID                  int                  `json:"spell2Id"`                  // Second summoner spell ID
	Stats                     MParticipantStats    `json:"stats"`                     // Participant statistics
	TeamID                    int                  `json:"teamID"`                    // Team ID
	Timeline                  MParticipantTimeline `json:"timeline"`                  // Timeline data
}

// MParticipantIdentity : A struct containing participant identity information
type MParticipantIdentity struct {
	ParticipantID int     `json:"participantId"` // Participant ID
	Player        MPlayer `json:"player"`        // Player Information
}

// MTeam : A struct containing team information
type MTeam struct {
	Bans                 []MBannedChampion `json:"bans"`                 // If game was draft mode, contains banned champion data, otherwise null
	BaronKills           int               `json:"baronKills"`           // Number of times the team killed baron
	DominionVictoryScore int64             `json:"dominionVictoryScore"` // If game was a dominion game, specifies the points the team had at game end, otherwise null
	DragonKills          int               `json:"dragonKills"`          // Number of times the team killed dragon
	FirstBaron           bool              `json:"firstBaron"`           // Flag indicating whether or not the team got the first baron kill
	FirstBlood           bool              `json:"firstBlood"`           // Flag indicating whether or not the team got first blood
	FirstDragon          bool              `json:"firstDragon"`          // Flag indicating whether or not the team got the first dragon kill
	FirstInhibitor       bool              `json:"firstInhibitor"`       // Flag indicating whether or not the team destroyed the first inhibitor
	FirstRiftHerald      bool              `json:"firstRiftHerald"`      // 	Flag indicating whether or not the team got the first rift herald kill
	FirstTower           bool              `json:"firstTower"`           // Flag indicating whether or not the team destroyed the first tower
	InhibitorKills       int               `json:"inhibitorKills"`       // Number of inhibitors the team destroyed
	RiftHeraldKills      int               `json:"riftHeraldKills"`      // Number of times the team killed rift herald
	TeamID               int               `json:"teamId"`               // Team ID
	TowerKills           int               `json:"towerKills"`           // Number of towers the team destroyed
	VilemawKills         int               `json:"vilemawKills"`         // Number of times the team killed vilemaw
	Winner               bool              `json:"winner"`               // Flag indicating whether or not the team won
}

// MTimeline : A struct containing game timeline information
type MTimeline struct {
	FrameInterval int64    `json:"frameInterval"` // Time between each returned frame in milliseconds
	Frames        []MFrame `json:"frames"`        // List of timeline frames for the game
}

// MMastery : A struct containing mastery information
type MMastery struct {
	MasterID int64 `json:"masterId"` // Mastery ID
	Rank     int64 `json:"rank"`     // Mastery rank
}

// MParticipantStats : A struct containing participant statistics information
type MParticipantStats struct {
	Assists                         int64 `json:"assists"`                         // Number of assists
	ChampLevel                      int64 `json:"champLevel"`                      // Champion level achieved
	CombatPlayerScore               int64 `json:"combatPlayerScore"`               // If game was a dominion game, player's combat score, otherwise 0
	Deaths                          int64 `json:"deaths"`                          // Number of deaths
	DoubleKills                     int64 `json:"doubleKills"`                     // Number of double kills
	FirstBloodAssist                bool  `json:"firstBloodAssist"`                // Flag indicating if participant got an assist on first blood
	FirstBloodKill                  bool  `json:"firstBloodKill"`                  // Flag indicating if participant got first blood
	FirstInhibitorAssist            bool  `json:"firstInhibitorAssist"`            // Flag indicating if participant got an assist on the first inhibitor
	FirstInhibitorKill              bool  `json:"firstInhibitorKill"`              // Flag indicating if participant destroyed the first inhibitor
	FirstTowerAssist                bool  `json:"firstTowerAssist"`                // Flag indicating if participant got an assist on the first tower
	FirstTowerKill                  bool  `json:"firstTowerKill"`                  // Flag indicating if participant destroyed the first tower
	GoldEarned                      int64 `json:"goldEarned"`                      // Gold earned
	GoldSpent                       int64 `json:"goldSpent"`                       // Gold spent
	InhibitorKills                  int64 `json:"inhibitorKills"`                  // Number of inhibitor kills
	Item0                           int64 `json:"item0"`                           // First item ID
	Item1                           int64 `json:"item1"`                           // Second item ID
	Item2                           int64 `json:"item2"`                           // Third item ID
	Item3                           int64 `json:"item3"`                           // Fourth item ID
	Item4                           int64 `json:"item4"`                           // Fifth item ID
	Item5                           int64 `json:"item5"`                           // Sixth item ID
	Item6                           int64 `json:"item6"`                           // Seventh item ID
	KillingSprees                   int64 `json:"killingSprees"`                   // Number of killing sprees
	Kills                           int64 `json:"kills"`                           // Number of kills
	LargestCriticalStrike           int64 `json:"largestCriticalStrike"`           // Largest critical strike
	LargestKillingSpree             int64 `json:"largestKillingSpree"`             // Largest killing spree
	LargestMultiKill                int64 `json:"largestMultiKill"`                // Largest multi kill
	MagicDamageDealt                int64 `json:"magicDamageDealt"`                // Magical damage dealt
	MagicDamageDealtToChampions     int64 `json:"magicDamageDealtToChampions"`     // Magical damage dealt to champions
	MagicDamageTaken                int64 `json:"magicDamageTaken"`                // Magic damage taken
	MinionsKilled                   int64 `json:"minionsKilled"`                   // Minions killed
	NeutralMinionsKilled            int64 `json:"neutralMinionsKilled"`            // Neutral minions killed
	NeutralMinionsKilledEnemyJungle int64 `json:"neutralMinionsKilledEnemyJungle"` // Neutral jungle minions killed in the enemy team's jungle
	NeutralMinionsKilledTeamJungle  int64 `json:"neutralMinionsKilledTeamJungle"`  // Neutral jungle minions killed in your team's jungle
	NodeCapture                     int64 `json:"nodeCapture"`                     // If game was a dominion game, number of node captures
	NodeCaptureAssist               int64 `json:"nodeCaptureAssist"`               // If game was a dominion game, number of node capture assists
	NodeNeutralize                  int64 `json:"nodeNeutralize"`                  // If game was a dominion game, number of node neutralizations
	NodeNeutralizeAssist            int64 `json:"nodeNeutralizeAssist"`            // If game was a dominion game, number of node neutralization assists
	ObjectivePlayerScore            int64 `json:"objectivePlayerScore"`            // If game was a dominion game, player's objectives score, otherwise 0
	PentaKills                      int64 `json:"pentaKills"`                      // Number of penta kills
	PhysicalDamageDealt             int64 `json:"physicalDamageDealt"`             // Physical damage dealt
	PhysicalDamageDealtToChampions  int64 `json:"physicalDamageDealtToChampions"`  // Physical damage dealt to champions
	PhysicalDamageTaken             int64 `json:"physicalDamageTaken"`             // Physical damage taken
	QuadraKills                     int64 `json:"quadraKills"`                     // Number of quadra kills
	SightWardsBoughtInGame          int64 `json:"sightWardsBoughtInGame"`          // Sight wards purchased
	TeamObjective                   int64 `json:"teamObjective"`                   // If game was a dominion game, number of completed team objectives (i.e., quests)
	TotalDamageDealt                int64 `json:"totalDamageDealt"`                // Total damage dealt
	TotalDamageDealtToChampions     int64 `json:"totalDamageDealtToChampions"`     // Total damage dealt to champions
	TotalDamageTaken                int64 `json:"totalDamageTaken"`                //	Total damage taken
	TotalHeal                       int64 `json:"totalHeal"`                       // Total heal amount
	TotalPlayerScore                int64 `json:"totalPlayerScore"`                // If game was a dominion game, player's total score, otherwise 0
	TotalScoreRank                  int64 `json:"totalScoreRank"`                  // If game was a dominion game, team rank of the player's total score (e.g., 1-5)
	TotalTimeCrowdControlDealt      int64 `json:"totalTimeCrowdControlDealt"`      // totalTimeCrowdControlDealt
	TotalUnitsHealed                int64 `json:"totalUnitsHealed"`                // Total units healed
	TowerKills                      int64 `json:"towerKills"`                      // Number of tower kills
	TripleKills                     int64 `json:"tripleKills"`                     // Number of triple kills
	TrueDamageDealt                 int64 `json:"trueDamageDealt"`                 // True damage dealt
	TrueDamageDealtToChampions      int64 `json:"trueDamageDealtToChampions"`      // True damage dealt to champions
	TrueDamageTaken                 int64 `json:"trueDamageTaken"`                 // True damage taken
	UnrealKills                     int64 `json:"unrealKills"`                     // Number of unreal kills
	VisionWardsBoughtInGame         int64 `json:"visionWardsBoughtInGame"`         // Vision wards purchased
	WardsKilled                     int64 `json:"wardsKilled"`                     // Number of wards killed
	WardsPlaced                     int64 `json:"wardsPlaced"`                     // Number of wards placed
	Winner                          bool  `json:"winner"`                          // Flag indicating whether or not the participant won
}

// MParticipantTimeline : A struct containing all timeline information
type MParticipantTimeline struct {
	AncientGolemAssistsPerMinCounts MParticipantTimelineData `json:"ancientGolemAssistsPerMinCounts"` // Ancient golem assists per minute timeline counts
	AncientGolemKillsPerMinCounts   MParticipantTimelineData `json:"AncientGolemKillsPerMinCounts"`   // Ancient golem kills per minute timeline counts
	AssistedLaneDeathsPerMinDeltas  MParticipantTimelineData `json:"assistedLaneDeathsPerMinDeltas"`  // Assisted lane deaths per minute timeline data
	AssistedLaneKillsPerMinDeltas   MParticipantTimelineData `json:"assistedLaneKillsPerMinDeltas"`   // Assisted lane kills per minute timeline data
	BaronAssistsPerMinCounts        MParticipantTimelineData `json:"baronAssistsPerMinCounts"`        // Baron assists per minute timeline counts
	BaronKillsPerMinCounts          MParticipantTimelineData `json:"baronKillsPerMinCounts"`          // Baron kills per minute timeline counts
	CreepsPerMinDeltas              MParticipantTimelineData `json:"creepsPerMinDeltas"`              // Creeps per minute timeline data
	CSDiffPerMinDeltas              MParticipantTimelineData `json:"csDiffPerMinDeltas"`              // Creep score difference per minute timeline data
	DamageTakenDiffPerMinDeltas     MParticipantTimelineData `json:"damageTakenDiffPerMinDeltas"`     // Damage taken difference per minute timeline data
	DamageTakenPerMinDeltas         MParticipantTimelineData `json:"damageTakenPerMinDeltas"`         // Damage taken per minute timeline data
	DragonAssistsPerMinCounts       MParticipantTimelineData `json:"dragonAssistsPerMinCounts"`       // Dragon assists per minute timeline counts
	DragonKillsPerMinCounts         MParticipantTimelineData `json:"dragonKillsPerMinCounts"`         // Dragon kills per minute timeline counts
	ElderLizardAssistsPerMinCounts  MParticipantTimelineData `json:"elderLizardAssistsPerMinCounts"`  // Elder lizard assists per minute timeline counts
	ElderLizardKillsPerMinCounts    MParticipantTimelineData `json:"elderLizardKillsPerMinCounts"`    // Elder lizard kills per minute timeline counts
	GoldPerMinDeltas                MParticipantTimelineData `json:"goldPerMinDeltas"`                // Gold per minute timeline data
	InhibitorAssistsPerMinCounts    MParticipantTimelineData `json:"inhibitorAssistsPerMinCounts"`    // Inhibitor assists per minute timeline counts
	InhibitorKillsPerMinCounts      MParticipantTimelineData `json:"inhibitorKillsPerMinCounts"`      // Inhibitor kills per minute timeline counts
	Lane                            string                   `json:"lane"`                            // Participant's lane (Legal values: MID, MIDDLE, TOP, JUNGLE, BOT, BOTTOM)
	Role                            string                   `json:"role"`                            // Participant's role (Legal values: DUO, NONE, SOLO, DUO_CARRY, DUO_SUPPORT)
	TowerAssistsPerMinCounts        MParticipantTimelineData `json:"towerAssistsPerMinCounts"`        // Tower assists per minute timeline counts
	TowerKillsPerMinCounts          MParticipantTimelineData `json:"towerKillsPerMinCounts"`          // Tower kills per minute timeline counts
	TowerKillsPerMinDeltas          MParticipantTimelineData `json:"towerKillsPerMinDeltas"`          // Tower kills per minute timeline data
	VilemawAssistsPerMinCounts      MParticipantTimelineData `json:"vilemawAssistsPerMinCounts"`      // Vilemaw assists per minute timeline counts
	VilemawKillsPerMinCounts        MParticipantTimelineData `json:"vilemawKillsPerMinCounts"`        // Vilemaw kills per minute timeline counts
	WardsPerMinDeltas               MParticipantTimelineData `json:"wardsPerMinDeltas"`               // Wards placed per minute timeline data
	XPDiffPerMinDeltas              MParticipantTimelineData `json:"xpDiffPerMinDeltas"`              // Experience difference per minute timeline data
	XPPerMinDeltas                  MParticipantTimelineData `json:"xpPerMinDeltas"`                  // Experience per minute timeline data
}

// MRune : A struct containing rune information on a match participant
type MRune struct {
	Rank   int64 `json:"rank"`   // Rune rank
	RuneID int64 `json:"runeId"` // Rune ID
}

// MPlayer : A struct containing match player information
type MPlayer struct {
	MatchHistoryURI string `json:"matchHistoryUri"` // Match history URI
	ProfileIcon     int    `json:"profileIcon"`     // Profile icon ID
	SummonerID      int64  `json:"summonerId"`      // Summoner ID
	SummonerName    string `json:"summonerName"`    // Summoner name
}

// MBannedChampion : A struct containing information on a banned champion
type MBannedChampion struct {
	ChampionID int64 `json:"championId"` // Banned champion ID
	PickTurn   int   `json:"pickTurn"`   // Turn during which the champion was banned
}

// MFrame : A struct containing game frame information
type MFrame struct {
	Events            []MEvent                     `json:"events"`            // List of events for this frame
	ParticipantFrames map[string]MParticipantFrame `json:"participantFrames"` // Map of each participant ID to the participant's information for the frame
	Timestamp         int64                        `json:"timestamp"`         // Map of each participant ID to the participant's information for the frame
}

// MParticipantTimelineData : A struct containing timeline data
type MParticipantTimelineData struct {
	TenToTwenty    float64 `json:"tenToTwenty"`    // Value per minute from 10 min to 20 min
	ThirtyToEnd    float64 `json:"thirtyToEnd"`    // Value per minute from 30 min to the end of the game
	TwentyToThirty float64 `json:"twentyToThirty"` // Value per minute from 20 min to 30 min
	ZeroToTen      float64 `json:"zeroToTen"`      // Value per minute from the beginning of the game to 10 min
}

// MEvent : A struct containing game event information
type MEvent struct {
	AscendedType            string    `json:"ascendedType"`            // The ascended type of the event
	AssistingParticipantIds []int     `json:"assistingParticipantIds"` // The assisting participant IDs of the event
	BuildingType            string    `json:"buildingType"`            // The building type of the event
	CreatorID               int       `json:"creatorId"`               // The creator ID of the event
	EventType               string    `json:"eventType"`               // Event type
	ItemAfter               int       `json:"itemAfter"`               // The ending item ID of the event
	ItemBefore              int       `json:"itemBefore"`              // The starting item ID of the event
	ItemID                  int       `json:"ItemID"`                  // The item ID of the event
	KillerID                int       `json:"killerId"`                // The killer ID of the event
	LaneType                string    `json:"laneType"`                // The lane type of the event
	LevelUpType             string    `json:"levelUpType"`             // The level up type of the event
	MonsterType             string    `json:"monsterType"`             // The monster type of the event
	ParticipantID           int       `json:"participantId"`           // The participant ID of the event
	PointCaptured           string    `json:"pointCaptured"`           // The point captured in the event
	Position                MPosition `json:"position"`                // The position of the event
	SkillSlot               int       `json:"skillSlot"`               // The skill slot of the event
	TeamID                  int       `json:"teamId"`                  // The team ID of the event
	Timestamp               int64     `json:"timestamp"`               // Represents how many milliseconds into the game the event occurred
	TowerType               string    `json:"towerType"`               // The tower type of the event
	VictimID                int       `json:"victimId"`                // The victim ID of the event
	WardType                string    `json:"wardType"`                // 	The ward type of the event
}

// MParticipantFrame : A struct containing participant fame information
type MParticipantFrame struct {
	CurrentGold         int       `json:"currentGold"`         // Participant's current gold
	DominionScore       int       `json:"dominionScore"`       // Dominion score of the participant
	JungleMinionsKilled int       `json:"jungleMinionsKilled"` // Number of jungle minions killed by participant
	Level               int       `json:"level"`               // Participant's current level
	MinionsKilled       int       `json:"minionsKilled"`       // Number of minions killed by participant
	ParticipantID       int       `json:"participantId"`       // Participant ID
	Position            MPosition `json:"position"`            // Participant's position
	TeamScore           int       `json:"teamScore"`           // Team score of the participant
	TotalGold           int       `json:"totalGold"`           // Participant's total gold
	XP                  int       `json:"xp"`                  // Experience earned by participant
}

// MPosition : A struct containing participant frame position information
type MPosition struct {
	X int `json:"x"` // X coordinate of the position
	Y int `json:"y"` // Y coordinate of the position
}

// GetMatchByID : Retrieve match by match ID
func (c *RiotClient) GetMatchByID(region string, matchID int64, opts *MatchOpts) (MMatchDetail, error) {
	var match MMatchDetail
	// Builds out query params based on options passed
	params := &url.Values{}
	if opts != nil {
		if opts.IncludeTimeline == true {
			params.Add("includeTimeline", "true")
		}
	}

	// Performs the http request on Riots API to retrieve the match information
	resBody, err := c.riotRequest("/api/lol/"+region+"/v2.2/match/"+
		strconv.FormatInt(matchID, 10), region, params)
	if err != nil {
		return match, err
	}

	// Unmarshals the response JSON into our MMatchDetail struct
	if err := json.Unmarshal(resBody, &match); err != nil {
		return match, err
	}
	return match, nil
}
