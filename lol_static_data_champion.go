// Package goriot : Currently up to date as of static-v1.2
package goriot

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// StaticChampionOpts : A struct containing optional paramenters used with static champion endpoint
type StaticChampionOpts struct {
	Locale    string
	Version   string
	DataByID  bool
	ChampData string
}

// LSDCChampionListDTO : A struct containing champion list data
type LSDCChampionListDTO struct {
	Data    map[string]LSDCChampionDTO `json:"data"`
	Format  string                     `json:"format"`
	Keys    map[string]string          `json:"keys"`
	Type    string                     `json:"type"`
	Version string                     `json:"version"`
}

// LSDCChampionDTO : A struct containing champion data
type LSDCChampionDTO struct {
	AllyTips    []string               `json:"allytips"`
	Blurb       string                 `json:"blurb"`
	Enemytips   []string               `json:"enemytips"`
	ID          int                    `json:"id"`
	Image       LSDCImageDTO           `json:"image"`
	Info        LSDCInfoDTO            `json:"info"`
	Key         string                 `json:"key"`
	Lore        string                 `json:"lore"`
	Name        string                 `json:"name"`
	Partype     string                 `json:"string"`
	Passive     LSDCPassiveDTO         `json:"passive"`
	Recommended []LSDCRecommendedDTO   `json:"recommended"`
	Skins       []LSDCSkinDTO          `json:"skins"`
	Spells      []LSDCChampionSpellDTO `json:"spells"`
	Stats       LSDCStatsDTO           `json:"stats"`
	Tags        []string               `json:"tags"`
	Title       string                 `json:"title"`
}

// LSDCChampionSpellDTO : A struct containing champion spell data
type LSDCChampionSpellDTO struct {
	AltImages            []LSDCImageDTO     `json:"altimages"`
	Cooldown             []float64          `json:"cooldown"`
	CooldownBurn         string             `json:"cooldownBurn"`
	Cost                 []int              `json:"cost"`
	CostBurn             string             `json:"costBurn"`
	CostType             string             `json:"costType"`
	Description          string             `json:"description"`
	Effect               [][]float64        `json:"effect"`
	EffectBurn           []string           `json:"effectBurn"`
	Image                LSDCImageDTO       `json:"image"`
	Key                  string             `json:"key"`
	Leveltip             LSDCLevelTipDTO    `json:"leveltip"`
	MaxRank              int                `json:"maxrank"`
	Name                 string             `json:"name"`
	Range                string             `json:"range"` // This field is either a List of Integer or the String 'self' for spells that target one's own champion
	RangeBurn            string             `json:"rangeBurn"`
	Resource             string             `json:"resource"`
	SanitizedDescription string             `json:"sanitizedDescription"`
	SanitizedTooltip     string             `json:"sanitizedTooltip"`
	Tooltip              string             `json:"tooltip"`
	Vars                 []LSDCSpellVarsDTO `json:"vars"`
}

// LSDCImageDTO : A struct containing image data
type LSDCImageDTO struct {
	Full   string `json:"full"`
	Group  string `json:"group"`
	H      int    `json:"h"`
	Sprite string `json:"sprite"`
	W      int    `json:"w"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

// LSDCInfoDTO : A struct containing champion information
type LSDCInfoDTO struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Difficulty int `json:"difficulty"`
	Magic      int `json:"magic"`
}

// LSDCPassiveDTO : A struct containing champion passive data
type LSDCPassiveDTO struct {
	Description          string       `json:"description"`
	Image                LSDCImageDTO `json:"image"`
	Name                 string       `json:"name"`
	SanitizedDescription string       `json:"sanitizedDescription"`
}

// LSDCRecommendedDTO : A struct containing champion recommended data
type LSDCRecommendedDTO struct {
	Blocks   []LSDCBlockDTO `json:"blocks"`
	Champion string         `json:"champion"`
	Map      string         `json:"map"`
	Mode     string         `json:"mode"`
	Priority bool           `json:"priority"`
	Title    string         `json:"title"`
	Type     string         `json:"type"`
}

// LSDCSkinDTO : A struct containing champion skin data
type LSDCSkinDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Num  string `json:"num"`
}

// LSDCStatsDTO : A struct containing champion stats data
type LSDCStatsDTO struct {
	Armor                float64 `json:"armor"`
	ArmorPerLevel        float64 `json:"armorperlevel"`
	AttackDamage         float64 `json:"attackdamage"`
	AttackDamagePerLevel float64 `json:"attackdamageperlevel"`
	AttackRange          float64 `json:"attackrange"`
	AttackSpeedOffset    float64 `json:"attackspeedoffset"`
	AttackSpeedPerLevel  float64 `json:"attackspeedperlevel"`
	Crit                 float64 `json:"crit"`
	CritPerLevel         float64 `json:"critperlevel"`
	HP                   float64 `json:"hp"`
	HPPerLevel           float64 `json:"hpperlevel"`
	HPRegen              float64 `json:"hpregen"`
	HPRegenPerLevel      float64 `json:"hpregenperlevel"`
	MoveSpeed            float64 `json:"movespeed"`
	MP                   float64 `json:"mp"`
	MPPerLevel           float64 `json:"mpperlevel"`
	MPRegen              float64 `json:"mpregen"`
	MPRegenPerLevel      float64 `json:"mpregenperlevel"`
	SpellBlock           float64 `json:"spellblock"`
	SpellBlockPerLevel   float64 `json:"spellblockperlevel"`
}

// LSDCLevelTipDTO : A struct containing champion level tip data
type LSDCLevelTipDTO struct {
	Effect []string `json:"effect"`
	Label  []string `json:"label"`
}

// LSDCSpellVarsDTO : A struct containing spell vars data
type LSDCSpellVarsDTO struct {
	Coeff     []float64 `json:"coeff"`
	Dyn       string    `json:"dyn"`
	Key       string    `json:"key"`
	Link      string    `json:"link"`
	RanksWith string    `json:"ranksWith"`
}

// LSDCBlockDTO : A struct containing champion recommended block data
type LSDCBlockDTO struct {
	Items   []LSDCBlockItemDTO `json:"items"`
	RecMath bool               `json:"recMath"`
	Type    string             `json:"type"`
}

// LSDCBlockItemDTO : A struct containing champion recommended block item data
type LSDCBlockItemDTO struct {
	Count int `json:"count"`
	ID    int `json:"id"`
}

// GetStaticChampions : Retrieves champion list
func (c *RiotClient) GetStaticChampions(region string, opts *StaticChampionOpts) (LSDCChampionListDTO, error) {
	var champions LSDCChampionListDTO
	params := &url.Values{}
	// Builds out query params based on options passed
	if opts != nil {
		if opts.Locale != "" {
			params.Add("locale", opts.Locale)
		}
		if opts.Version != "" {
			params.Add("version", opts.Version)
		}
		if opts.DataByID == true {
			params.Add("dataById", "true")
		}
		if opts.ChampData != "" {
			params.Add("champData", opts.ChampData)
		}
	}

	// Performs the http request on Riots API to retrieve version data
	resBody, err := c.riotRequest("/api/lol/static-data/"+region+"/v1.2/champion", "global", params)
	if err != nil {
		return champions, err
	}

	// Unmarshals the response JSON into a LSDCChampionListDTO struct
	if err = json.Unmarshal(resBody, &champions); err != nil {
		return champions, err
	}
	return champions, nil
}

// GetStaticChampionByID : Retrieves a champion by its id
func (c *RiotClient) GetStaticChampionByID(region string, championID int, opts *StaticChampionOpts) (LSDCChampionDTO, error) {
	var champion LSDCChampionDTO
	params := &url.Values{}
	// Builds out query params based on options passed
	if opts != nil {
		if opts.Locale != "" {
			params.Add("locale", opts.Locale)
		}
		if opts.Version != "" {
			params.Add("version", opts.Version)
		}
		if opts.ChampData != "" {
			params.Add("champData", opts.ChampData)
		}
	}

	// Performs the http request on Riots API to retrieve version data
	resBody, err := c.riotRequest("/api/lol/static-data/"+region+"/v1.2/champion"+strconv.Itoa(championID), "global", params)
	if err != nil {
		return champion, err
	}

	// Unmarshals the response JSON into a LSDCChampionDTO struct
	if err = json.Unmarshal(resBody, &champion); err != nil {
		return champion, err
	}
	return champion, nil
}
