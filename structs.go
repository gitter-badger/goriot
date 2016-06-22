package goriot

// AggregatedStatsDTO :
type AggregatedStatsDTO struct {
	AverageAssists              int `json:"averageAssists"`              // Dominion only
	AverageChampionsKilled      int `json:"averageChampionsKilled"`      // Dominion only
	AverageCombatPlayerScore    int `json:"averageCombatPlayerScore"`    // Dominion only
	AverageNodeCapture          int `json:"averageNodeCapture"`          // Dominion only
	AverageNodeCaptureAssist    int `json:"averageNodeCaptureAssist"`    // Dominion only
	AverageNodeNeutralize       int `json:"averageNodeNeutralize"`       // Dominion only
	AverageNodeNeutralizeAssist int `json:"averageNodeNeutralizeAssist"` // Dominion only
	AverageNumDeaths            int `json:"averageNumDeaths"`            // Dominion only
	AverageObjectivePlayerScore int `json:"averageObjectivePlayerScore"` // Dominion only
	AverageTeamObjective        int `json:"averageTeamObjective"`        // Dominion only
	AverageTotalPlayerScore     int `json:"averageTotalPlayerScore"`     // Dominion only
	BotGamesPlayed              int `json:"botGamesPlayed"`
	KillingSpree                int `json:"killingSpree"`
	MaxAssists                  int `json:"maxAssists"` // Dominion only
	MaxChampionsKilled          int `json:"maxChampionsKilled"`
	MaxCombatPlayerScore        int `json:"maxCombatPlayerScore"` // Dominion only
	MaxLargestCriticalStrike    int `json:"maxLargestCriticalStrike"`
	MaxLargestKillingSpree      int `json:"maxLargestKillingSpree"`
	MaxNodeCapture              int `json:"maxNodeCapture"`          // Dominion only
	MaxNodeCaptureAssist        int `json:"maxNodeCaptureAssist"`    // Dominion only
	MaxNodeNeutralize           int `json:"maxNodeNeutralize"`       // Dominion only
	MaxNodeNeutralizeAssist     int `json:"maxNodeNeutralizeAssist"` // Dominion only
	MaxNumDeaths                int `json:"maxNumDeaths"`            // Only returned for ranked statistics
	MaxObjectivePlayerScore     int `json:"maxObjectivePlayerScore"` // Dominion only
	MaxTeamObjective            int `json:"maxTeamObjective"`        // Dominion only
	MaxTimePlayed               int `json:"maxTimePlayed"`
	MaxTimeSpentLiving          int `json:"maxTimeSpentLiving"`
	MaxTotalPlayerScore         int `json:"maxTotalPlayerScore"` // Dominion only
	MostChampionKillsPerSession int `json:"mostChampionKillsPerSession"`
	MostSpellsCast              int `json:"mostSpellsCast"`
	NormalGamesPlayed           int `json:"normalGamesPlayed"`
	RankedPremadeGamesPlayed    int `json:"rankedPremadeGamesPlayed"`
	RankedSoloGamesPlayed       int `json:"rankedSoloGamesPlayed"`
	TotalAssists                int `json:"totalAssists"`
	TotalChampionKills          int `json:"totalChampionKills"`
	TotalDamageDealt            int `json:"totalDamageDealt"`
	TotalDamageTaken            int `json:"totalDamageTaken"`
	TotalDeathsPerSession       int `json:"totalDeathsPerSession"` // Only returned for ranked statistics
	TotalDoubleKills            int `json:"totalDoubleKills"`
	TotalFirstBlood             int `json:"totalFirstBlood"`
	TotalGoldEarned             int `json:"totalGoldEarned"`
	TotalHeal                   int `json:"totalHeal"`
	TotalMagicDamageDealt       int `json:"totalMagicDamageDealt"`
	TotalMinionKills            int `json:"totalMinionKills"`
	TotalNeutralMinionsKilled   int `json:"totalNeutralMinionsKilled"`
	TotalNodeCapture            int `json:"totalNodeCapture"`    // Dominion only
	TotalNodeNeutralize         int `json:"totalNodeNeutralize"` // Dominion only
	TotalPentaKills             int `json:"totalPentaKills"`
	TotalPhysicalDamageDealt    int `json:"totalPhysicalDamageDealt"`
	TotalQuadraKills            int `json:"totalQuadraKills"`
	TotalSessionsLost           int `json:"totalSessionsLost"`
	TotalSessionsPlayed         int `json:"totalSessionsPlayed"`
	TotalSessionsWon            int `json:"totalSessionsWon"`
	TotalTripleKills            int `json:"totalTripleKills"`
	TotalTurretsKilled          int `json:"totalTurretsKilled"`
	TotalUnrealKills            int `json:"totalUnrealKills"`
}

// BannedChampion : A struct containing information on a banned champion
type BannedChampion struct {
	ChampionID int64 `json:"championId"` // The ID of the banned champion
	PickTurn   int   `json:"pickTurn"`   // The turn during which the champion was banned
	TeamID     int64 `json:"teamId"`     // The ID of the team that banned the champion
}

// ChampionListDTO : A struct containing a collection of champion information
type ChampionListDTO struct {
	Chamions []ChampionDTO `json:"champions"` // 	The collection of champion information.
}

// ChampionDTO : A struct containing champion information
type ChampionDTO struct {
	Active            bool `json:"active"`            // Indicates if the champion is active.
	BotEnabled        bool `json:"botEnabled"`        // Bot enabled flag (for custom games).
	BotMmEnabled      bool `json:"botMmEnabled"`      // Bot Match Made enabled flag (for Co-op vs. AI games).
	FreeToPlay        bool `json:"freeToPlay"`        // Indicates if the champion is free to play. Free to play champions are rotated periodically.
	ID                int  `json:"id"`                // Champion ID. For static information correlating to champion IDs, please refer to the LoL Static Data API.
	RankedPlayEnabled bool `json:"rankedPlayEnabled"` // Ranked play enabled flag.
}

// ChampionMasteryDTO : A struct containing single Champion Mastery information for player and champion combination
type ChampionMasteryDTO struct {
	ChampionID                   int64  `json:"championId"`                   // 	Champion ID for this entry
	ChampionLevel                int    `json:"championLevel"`                // Champion level for specified player and champion combination
	ChampionPoints               int    `json:"championPoints"`               // Total number of champion points for this player and champion combination - they are used to determine championLevel
	ChampionPointsSinceLastLevel int64  `json:"championPointsSinceLastLevel"` // Number of points earned since current level has been achieved. Zero if player reached maximum champion level for this champion
	ChampionPointsUntilNextLevel int64  `json:"championPointsUntilNextLevel"` // Number of points needed to achieve next level. Zero if player reached maximum champion level for this champion
	ChestGranted                 bool   `json:"chestGranted"`                 // Is chest granted for this champion or not in current season
	HighestGrade                 string `json:"highestGrade"`                 // The highest grade of this champion of current season
	LastPlayTime                 int64  `json:"lastPlayTime"`                 // Last time this champion was played by this player - in Unix milliseconds time format
	PlayerID                     int64  `json:"playerId"`                     // Player ID for this entry
}

// ChampionStatsDTO : A struct containing a collection of champion stats information
type ChampionStatsDTO struct {
	ID    int                // Champion ID
	Stats AggregatedStatsDTO // Aggregated stats associated with the champion
}

// CurrentGameInfo : A struct containing information on a in-progress game
type CurrentGameInfo struct {
	BannedChampions   []BannedChampion         `json:"bannedChampions"`   // Banned champion information
	GameID            int64                    `json:"gameId"`            // The ID of the game
	GameLength        int64                    `json:"gameLength"`        // The amount of time in seconds that has passed since the game started
	GameMode          string                   `json:"gameMode"`          // The game mode (Legal values: CLASSIC, ODIN, ARAM, TUTORIAL, ONEFORALL, ASCENSION, FIRSTBLOOD, KINGPORO)
	GameQueueConfigID int64                    `json:"gameQueueConfigId"` // The queue type (queue types are documented on the Game Constants page)
	GameStartTime     int64                    `json:"gameStartTime"`     // The game start time represented in epoch milliseconds
	GameType          string                   `json:"gameType"`          // The game type (Legal values: CUSTOM_GAME, MATCHED_GAME, TUTORIAL_GAME)
	MapID             int64                    `json:"mapId"`             // The ID of the map
	Observers         Observer                 `json:"observers"`         // The observer information
	Participants      []CurrentGameParticipant `json:"participants"`      // The participant information
	PlatformID        string                   `json:"platformId"`        // The ID of the platform on which the game is being played
}

// CurrentGameParticipant : A struct containing information on a player
type CurrentGameParticipant struct {
	Bot           bool      `json:"bot"`           // Flag indicating whether or not this participant is a bot
	ChampionID    int64     `json:"championId"`    // The ID of the champion played by this participant
	Masteries     []Mastery `json:"masteries"`     // The masteries used by this participant
	ProfileIconID int64     `json:"profileIconId"` // The ID of the profile icon used by this participant
	Runes         []Rune    `json:"runes"`         // The runes used by this participant
	Spell1ID      int64     `json:"spell1Id"`      // The ID of the first summoner spell used by this participant
	Spell2ID      int64     `json:"spell2Id"`      // The ID of the second summoner spell used by this participant
	SummonerID    int64     `json:"summonerId"`    // The summoner ID of this participant
	SummonerName  string    `json:"summonerName"`  // The summoner name of this participant
	TeamID        int64     `json:"teamId"`        // The team ID of this participant, indicating the participant's team
}

// FeaturedGames : A collection of featured in-progress games
type FeaturedGames struct {
	ClientRefreshInterval int64              `json:"clientRefreshInterval"` // The suggested interval to wait before requesting FeaturedGames again
	GameList              []FeaturedGameInfo `json:"gameList"`              // The list of featured games
}

// FeaturedGameInfo : A struct containing information on a particular featured in-progress game
type FeaturedGameInfo struct {
	BannedChampions   []BannedChampion `json:"bannedChampions"`   // Banned champion information
	GameID            int64            `json:"gameId"`            // The ID of the game
	GameLength        int64            `json:"gameLength"`        // The amount of time in seconds that has passed since the game started
	GameMode          string           `json:"gameMode"`          // The game mode (Legal values: CLASSIC, ODIN, ARAM, TUTORIAL, ONEFORALL, ASCENSION, FIRSTBLOOD, KINGPORO)
	GameQueueConfigID int64            `json:"gameQueueConfigId"` // The queue type (queue types are documented on the Game Constants page)
	GameStartTime     int64            `json:"gameStartTime"`     // The game start time represented in epoch milliseconds
	GameType          string           `json:"gameType"`          // The game type (Legal values: CUSTOM_GAME, MATCHED_GAME, TUTORIAL_GAME)
	MapID             int64            `json:"mapId"`             // The ID of the map
	Observers         Observer         `json:"observers"`         // The observer information
	Participants      []Participant    `json:"participants"`      // The participant information
	PlatformID        string           `json:"platformId"`        // The ID of the platform on which the game is being played
}

// GameDTO : A struct containing game information
type GameDTO struct {
	ChampionID    int         `json:"championId"`    // Champion ID associated with game.
	CreateDate    int64       `json:"createDate"`    // Date that end game data was recorded, specified as epoch milliseconds.
	FellowPlayers []PlayerDTO `json:"fellowPlayers"` // Other players associated with the game.
	GameID        int64       `json:"gameId"`        // Game ID.
	GameMode      string      `json:"gameMode"`      // Game mode. (Legal values: CLASSIC, ODIN, ARAM, TUTORIAL, ONEFORALL, ASCENSION, FIRSTBLOOD, KINGPORO)
	GameType      string      `json:"gameType"`      // Game type. (Legal values: CUSTOM_GAME, MATCHED_GAME, TUTORIAL_GAME)
	Invalid       bool        `json:"invalid"`       // Invalid flag.
	IPEarned      int         `json:"ipEarned"`      // IP Earned.
	Level         int         `json:"level"`         // Level.
	MapID         int         `json:"mapId"`         // Map ID.
	Spell1        int         `json:"spell1"`        // ID of first summoner spell.
	Spell2        int         `json:"spell2"`        // ID of second summoner spell.
	Stats         RawStatsDTO `json:"stats"`         // Statistics associated with the game for this summoner.
	SubType       string      `json:"subType"`       // Game sub-type. (Legal values: NONE, NORMAL, BOT, RANKED_SOLO_5x5, RANKED_PREMADE_3x3, RANKED_PREMADE_5x5, ODIN_UNRANKED, RANKED_TEAM_3x3, RANKED_TEAM_5x5, NORMAL_3x3, BOT_3x3, CAP_5x5, ARAM_UNRANKED_5x5, ONEFORALL_5x5, FIRSTBLOOD_1x1, FIRSTBLOOD_2x2, SR_6x6, URF, URF_BOT, NIGHTMARE_BOT, ASCENSION, HEXAKILL, KING_PORO, COUNTER_PICK, BILGEWATER)
	TeamID        int         `json:"teamId"`        // Team ID associated with game. Team ID 100 is blue team. Team ID 200 is purple team.
}

// Incident : A struct containing information on a service incident
type Incident struct {
	Active    bool      `json:"active"`
	CreatedAt string    `json:"created_at"`
	ID        int64     `json:"id"`
	Updates   []Message `json:"updates"`
}

// LeagueDTO : A struct containing league informations
type LeagueDTO struct {
	Entries       []LeagueEntryDTO `json:"entries"`       // The requested league entries.
	Name          string           `json:"name"`          // This name is an internal place-holder name only. Display and localization of names in the game client are handled client-side.
	ParticipantID string           `json:"participantId"` // Specifies the relevant participant that is a member of this league (i.e., a requested summoner ID, a requested team ID, or the ID of a team to which one of the requested summoners belongs). Only present when full league is requested so that participant's entry can be identified. Not present when individual entry is requested.
	Queue         string           `json:"queue"`         // The league's queue type. (Legal values: RANKED_SOLO_5x5, RANKED_TEAM_3x3, RANKED_TEAM_5x5)
	Tier          string           `json:"tier"`          // The league's tier. (Legal values: CHALLENGER, MASTER, DIAMOND, PLATINUM, GOLD, SILVER, BRONZE)
}

// LeagueEntryDTO : A struct containing league participant information representing a summoner or team
type LeagueEntryDTO struct {
	Division         string        `json:"division"`         // The league division of the participant.
	IsFreshBlood     bool          `json:"isFreshBlood"`     // Specifies if the participant is fresh blood.
	IsHotStreak      bool          `json:"isHotStreak"`      // Specifies if the participant is on a hot streak.
	IsInactive       bool          `json:"isInactive"`       // Specifies if the participant is inactive.
	IsVeteran        bool          `json:"isVeteran"`        // Specifies if the participant is a veteran.
	LeaguePoints     int           `json:"leaguePoints"`     // The league points of the participant.
	Losses           int           `json:"losses"`           // The number of losses for the participant.
	MiniSeries       MiniSeriesDTO `json:"miniSeries"`       // Mini series data for the participant. Only present if the participant is currently in a mini series.
	PlayerOrTeamID   string        `json:"playerOrTeamId"`   // The ID of the participant (i.e., summoner or team) represented by this entry.
	PlayerOrTeamName string        `json:"playerOrTeamName"` // The name of the the participant (i.e., summoner or team) represented by this entry.
	Wins             int           `json:"wins"`             // The number of wins for the participant.
}

// Mastery : A struct containing information on a players mastery
type Mastery struct {
	MasteryID int64 `json:"masteryId"` // The ID of the mastery
	Rank      int   `json:"rank"`      // The number of points put into this mastery by the user
}

// MasteryDTO : A struct containing a users mastery information
type MasteryDTO struct {
	ID   int `json:"id"`   // Mastery ID. For static information correlating to masteries, please refer to the LoL Static Data API
	Rank int `json:"rank"` // Mastery rank (i.e., the number of points put into this mastery)
}

// MasteryPagesDTO : A struct containing a users mastery pages
type MasteryPagesDTO struct {
	Page       []MasteryPageDTO `json:"pages"`      // Collection of mastery pages associated with the summoner
	SummonerID int64            `json:"summonerId"` // Summoner ID
}

// MasteryPageDTO : A struct containing information on a mastery page
type MasteryPageDTO struct {
	Current   bool         `json:"current"`   // Indicates if the mastery page is the current mastery page
	ID        int64        `json:"id"`        //	Mastery page ID
	Masteries []MasteryDTO `json:"masteries"` // Collection of masteries associated with the mastery page
	Name      string       `json:"name"`      // Mastery page name
}

// Message : A struct containing message information of an incident
type Message struct {
	Author       string        `json:"author"`
	Content      string        `json:"content"`
	CreatedAt    string        `json:"created_at"`
	ID           string        `json:"id"`
	Severity     string        `json:"serverity"`
	Translations []Translation `json:"translations"`
	UpdatedAt    string        `json:"updated_at"`
}

// MatchHistorySummaryDTO : A struct containing a summary of a summoners match history
type MatchHistorySummaryDTO struct {
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

// MatchList : A struct containing match list information
type MatchList struct {
	EndIndex   int              `json:"endIndex"`
	Matches    []MatchReference `json:"matches"`
	StartIndex int              `json:"startIndex"`
	TotalGames int              `json:"totalGames"`
}

// MatchReference : A struct containing match refernce information
type MatchReference struct {
	Champion   int64  `json:"champion"`
	Lane       string `json:"lane"` // Legal values: MID, MIDDLE, TOP, JUNGLE, BOT, BOTTOM
	MatchID    int64  `json:"matchId"`
	PlatformID string `json:"platformId"`
	Queue      string `json:"queue"` // Legal values: TEAM_BUILDER_DRAFT_RANKED_5x5, RANKED_SOLO_5x5, RANKED_TEAM_3x3, RANKED_TEAM_5x5
	Region     string `json:"region"`
	Role       string `json:"role"`   // Legal values: DUO, NONE, SOLO, DUO_CARRY, DUO_SUPPORT
	Season     string `json:"season"` // Legal values: PRESEASON3, SEASON3, PRESEASON2014, SEASON2014, PRESEASON2015, SEASON2015, PRESEASON2016, SEASON2016
	Timestamp  int64  `json:"timestamp"`
}

// MiniSeriesDTO : A struct containing mini series information
type MiniSeriesDTO struct {
	Losses   int    `json:"losses"`   // Number of current losses in the mini series.
	Progress string `json:"progress"` // String showing the current, sequential mini series progress where 'W' represents a win, 'L' represents a loss, and 'N' represents a game that hasn't been played yet.
	Target   int    `json:"target"`   // Number of wins required for promotion.
	Wins     int    `json:"wins"`     // 	Number of current wins in the mini series.
}

// Observer : A struct containing information on game observers
type Observer struct {
	EncryptionKey string `json:"encryptionKey"` // Key used to decrypt the spectator grid game data for playback
}

// Participant : A struct containing information on a game participant
type Participant struct {
	Bot           bool   `json:"bot"`           // Flag indicating whether or not this participant is a bot
	ChampionID    int64  `json:"championId"`    // The ID of the champion played by this participant
	ProfileIconID int64  `json:"profileIconId"` // The ID of the profile icon used by this participant
	Spell1ID      int64  `json:"spell1Id"`      // The ID of the first summoner spell used by this participant
	Spell2ID      int64  `json:"spell2Id"`      // The ID of the second summoner spell used by this participant
	SummonerName  string `json:"summonerName"`  // The summoner name of this participant
	TeamID        int64  `json:"teamId"`        // The team ID of this participant, indicating the participant's team
}

// PlayerDTO : A struct containing player information
type PlayerDTO struct {
	ChampionID int   `json:"championId"` // Champion id associated with player.
	SummonerID int64 `json:"summonerId"` // Summoner id associated with player.
	TeamID     int   `json:"teamId"`     // Team id associated with player.
}

// PlayerStatsSummaryDTO : A struct containing player stats summary information
type PlayerStatsSummaryDTO struct {
	AggregatedStats       AggregatedStatsDTO `json:"aggregatedStats"`       // Aggregated stats
	Losses                int                `json:"losses"`                // Number of losses for this queue type. Returned for ranked queue types only
	ModifyDate            int64              `json:"modifyDate"`            // Date stats were last modified specified as epoch milliseconds
	PlayerStatSummaryType string             `json:"playerStatSummaryType"` // Player stats summary type
	Wins                  int                `json:"wins"`                  //	Number of wins for this queue type
}

// PlayerStatsSummaryListDTO : A struct containing a collection of player stats summary information
type PlayerStatsSummaryListDTO struct {
	PlayerStatSummaries []PlayerStatsSummaryDTO `json:"playerStatSummaries"` // Collection of player stats summaries associated with the summoner
	SummonerID          int64                   `json:"summonerId"`          // Summoner ID
}

// RankedStatsDTO : A struct containing ranked stats information
type RankedStatsDTO struct {
	Champions  []ChampionStatsDTO `json:"champions"`  // Collection of aggregated stats summarized by champion
	ModifyDate int64              `json:"modifyDate"` // Date stats were last modified specified as epoch milliseconds
	SummonerID int64              `json:"summonerId"` // Summoner ID
}

// RawStatsDTO : A struct containing raw game stat information
type RawStatsDTO struct {
	Assists                         int  `json:"assists"`
	BarracksKilled                  int  `json:"barracksKilled"` // Number of enemy inhibitors killed.
	BountyLevel                     int  `json:"bountyLevel"`
	ChampionsKilled                 int  `json:"championsKilled"`
	CombatPlayerScore               int  `json:"combatPlayerScore"`
	ConsumablesPurchased            int  `json:"consumablesPurchased"`
	DamageDealtPlayer               int  `json:"damageDealtPlayer"`
	DoubleKills                     int  `json:"doubleKills"`
	FirstBlood                      int  `json:"firstBlood"`
	Gold                            int  `json:"gold"`
	GoldEarned                      int  `json:"goldEarned"`
	GoldSpent                       int  `json:"goldSpent"`
	Item0                           int  `json:"item0"`
	Item1                           int  `json:"item1"`
	Item2                           int  `json:"item2"`
	Item3                           int  `json:"item3"`
	Item4                           int  `json:"item4"`
	Item5                           int  `json:"item5"`
	Item6                           int  `json:"item6"`
	ItemsPurchased                  int  `json:"itemsPurchased"`
	KillingSprees                   int  `json:"killingSprees"`
	LargestCriticalStrike           int  `json:"largestCriticalStrike"`
	LargestKillingSpree             int  `json:"largestKillingSpree"`
	LargestMultiKill                int  `json:"largestMultiKill"`
	LegendaryItemsCreated           int  `json:"legendaryItemsCreated"` // Number of tier 3 items built.
	Level                           int  `json:"level"`
	MagicDamageDealtPlayer          int  `json:"magicDamageDealtPlayer"`
	MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
	MagicDamageTaken                int  `json:"magicDamageTaken"`
	MinionsDenied                   int  `json:"minionsDenied"`
	MinionsKilled                   int  `json:"minionsKilled"`
	NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
	NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
	NeutralMinionsKilledYourJungle  int  `json:"neutralMinionsKilledYourJungle"`
	NexusKilled                     bool `json:"nexusKilled"` // Flag specifying if the summoner got the killing blow on the nexus.
	NodeCapture                     int  `json:"nodeCapture"`
	NodeCaptureAssist               int  `json:"nodeCaptureAssist"`
	NodeNeutralize                  int  `json:"nodeNeutralize"`
	NodeNeutralizeAssist            int  `json:"nodeNeutralizeAssist"`
	NumDeaths                       int  `json:"numDeaths"`
	NumItemsBought                  int  `json:"numItemsBought"`
	ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
	PentaKills                      int  `json:"pentaKills"`
	PhysicalDamageDealtPlayer       int  `json:"physicalDamageDealtPlayer"`
	PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
	PlayerPosition                  int  `json:"playerPosition"` // Player position (Legal values: TOP(1), MIDDLE(2), JUNGLE(3), BOT(4))
	PlayerRole                      int  `json:"playerRole"`     // Player role (Legal values: DUO(1), SUPPORT(2), CARRY(3), SOLO(4))
	PlayerScore0                    int  `json:"playerScore0"`
	PlayerScore1                    int  `json:"playerScore1"`
	PlayerScore2                    int  `json:"playerScore2"`
	PlayerScore3                    int  `json:"playerScore3"`
	PlayerScore4                    int  `json:"playerScore4"`
	PlayerScore5                    int  `json:"playerScore5"`
	PlayerScore6                    int  `json:"playerScore6"`
	PlayerScore7                    int  `json:"playerScore7"`
	PlayerScore8                    int  `json:"playerScore8"`
	PlayerScore9                    int  `json:"playerScore9"`
	QuadraKills                     int  `json:"quadraKills"`
	SightWardsBought                int  `json:"sightWardsBought"`
	Spell1Cast                      int  `json:"spell1Cast"` // Number of times first champion spell was cast.
	Spell2Cast                      int  `json:"spell2Cast"` // Number of times second champion spell was cast.
	Spell3Cast                      int  `json:"spell3Cast"` // Number of times third champion spell was cast.
	Spell4Cast                      int  `json:"spell4Cast"` // Number of times fourth champion spell was cast.
	SummonSpell1Cast                int  `json:"summonSpell1Cast"`
	SummonSpell2Cast                int  `json:"summonSpell2Cast"`
	SuperMonsterKilled              int  `json:"superMonsterKilled"`
	Team                            int  `json:"team"`
	TeamObjective                   int  `json:"teamObjective"`
	TimePlayed                      int  `json:"timePlayed"`
	TotalDamageDealt                int  `json:"totalDamageDealt"`
	TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
	TotalDamageTaken                int  `json:"totalDamageTaken"`
	TotalHeal                       int  `json:"totalHeal"`
	TotalPlayerScore                int  `json:"totalPlayerScore"`
	TotalScoreRank                  int  `json:"totalScoreRank"`
	TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
	TotalUnitsHealed                int  `json:"totalUnitsHealed"`
	TripleKills                     int  `json:"tripleKills"`
	TrueDamageDealtPlayer           int  `json:"trueDamageDealtPlayer"`
	TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                 int  `json:"trueDamageTaken"`
	TurretsKilled                   int  `json:"turretsKilled"`
	UnrealKills                     int  `json:"unrealKills"`
	VictoryPointTotal               int  `json:"victoryPointTotal"`
	VisionWardsBought               int  `json:"visionWardsBought"`
	WardKilled                      int  `json:"wardKilled"`
	WardPlaced                      int  `json:"wardPlaced"`
	Win                             bool `json:"win"` // Flag specifying whether or not this game was won.
}

// RecentGamesDTO : A struct containing recent games information
type RecentGamesDTO struct {
	Games      []GameDTO `json:"games"`      // Collection of recent games played (max 10).
	SummonerID int64     `json:"summonerId"` // Summoner ID.
}

// RosterDTO : A struct containing a games/teams members
type RosterDTO struct {
	MemberList []TeamMemberInfoDTO `json:"memberList"`
	OwnerID    int64               `json:"ownerId"`
}

// Rune : A struct containing information on a players rune
type Rune struct {
	Count  int   `json:"count"`  // The count of this rune used by the participant
	RuneID int64 `json:"runeId"` // The ID of the rune
}

// RunePageDTO : A struct containing a players rune slots
type RunePageDTO struct {
	Current bool          `json:"current"` // Indicates if the page is the current page
	ID      int64         `json:"id"`      // Rune page ID
	Name    string        `json:"name"`    // Rune page name
	Slots   []RuneSlotDTO `json:"slots"`   // Collection of rune slots associated with the rune page
}

// RunePagesDTO : A struct containing a players rune pages
type RunePagesDTO struct {
	Pages      []RunePageDTO `json:"pages"`      // Collection of rune pages associated with the summoner
	SummonerID int64         `json:"summonerId"` // Summoner ID
}

// RuneSlotDTO : A struct containing information on a players rune page slot
type RuneSlotDTO struct {
	RuneID     int `json:"runeId"`     // This object contains rune slot information
	RuneSlotID int `json:"runeSlotId"` // Rune slot ID
}

// Service : Information on a service that is running on a shard
type Service struct {
	Incidents []Incident `json:"incidents"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
	Status    string     `json:"status"`
}

// Shard : A struct containing status information
type Shard struct {
	Hostname  string   `json:"hostname"`
	Locales   []string `json:"locales"`
	Name      string   `json:"name"`
	RegionTag string   `json:"region_tag"`
	Slug      string   `json:"slug"`
}

// ShardStatus : The server status information pertaining to a particular region
type ShardStatus struct {
	Hostname  string    `json:"hostname"`
	Locales   []string  `json:"locales"`
	Name      string    `json:"name"`
	RegionTag string    `json:"region_tag"`
	Services  []Service `json:"services"`
	Slug      string    `json:"slug"`
}

// SummonerDTO : A struct containing summoner information
type SummonerDTO struct {
	ID            int64  `json:"id"`            // Summoner ID
	Name          string `json:"name"`          // Summoner name
	ProfileIconID int    `json:"profileIconId"` // ID of the summoner icon associated with the summoner
	RevisionDate  int64  `json:"revisionDate"`  // Date summoner was last modified specified as epoch milliseconds
	SummonerLevel int64  `json:"summonerLevel"` // 	Summoner level associated with the summoner
}

// TeamDTO : A struct containing team information
type TeamDTO struct {
	CreateDate                    int64                    `json:"createDate"` // Date that team was created specified as epoch milliseconds
	FullID                        string                   `json:"fullId"`
	LastGameDate                  int64                    `json:"lastGameDate"`                  // Date that last game played by team ended specified as epoch milliseconds
	LastJoinDate                  int64                    `json:"lastJoinDate"`                  // Date that last member joined specified as epoch milliseconds
	LastJoinedRankedTeamQueueDate int64                    `json:"lastJoinedRankedTeamQueueDate"` // Date that team last joined the ranked team queue specified as epoch milliseconds
	MatchHistory                  []MatchHistorySummaryDTO `json:"matchHistory"`
	ModifyDate                    int64                    `json:"modifyDate"` // Date that team was last modified specified as epoch milliseconds
	Name                          string                   `json:"name"`
	Roster                        RosterDTO                `json:"roster"`
	SecondLastJoinDate            int64                    `json:"secondLastJoinDate"` // Date that second to last member joined specified as epoch milliseconds
	Status                        string                   `json:"status"`
	Tag                           string                   `json:"tag"`
	TeamStatDetails               []TeamStatDetailDTO      `json:"teamStatDetails"`
	ThirdLastJoinDate             int64                    `json:"thirdLastJoinDate"` // Date that third to last member joined specified as epoch milliseconds
}

// TeamStatDetailDTO : A struct containing team statistics detail information
type TeamStatDetailDTO struct {
	AverageGamesPlayed int    `json:"averageGamesPlayed"`
	Losses             int    `json:"losses"`
	TeamStatType       string `json:"teamStatType"`
	Wins               int    `json:"wins"`
}

// TeamMemberInfoDTO : A struct containing team member information
type TeamMemberInfoDTO struct {
	InviteDate int64  `json:"inviteDate"` // Date that team member was invited to team specified as epoch milliseconds
	JoinDate   int64  `json:"joinDate"`   // Date that team member joined team specified as epoch milliseconds
	PlayerID   int64  `json:"playerId"`
	Status     string `json:"status"`
}

// Translation : A struct containing translation information
type Translation struct {
	Content   string `json:"content"`
	Locale    string `json:"locale"`
	UpdatedAt string `json:"updated_at"`
}
