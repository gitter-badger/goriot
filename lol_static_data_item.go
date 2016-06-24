// Package goriot : Currently up to date as of static-v1.2
package goriot

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// StaticItemOpts : A struct containing optional paramenters used with static item endpoint
type StaticItemOpts struct {
	Locale       string
	Version      string
	ItemListData string
	ItemData     string
}

// LSDIItemListDTO : A struct containing item list data
type LSDIItemListDTO struct {
	Basic   LSDIBasicDataDTO       `json:"basic"`
	Data    map[string]LSDIItemDTO `json:"data"`
	Groups  []LSDIGroupDTO         `json:"groups"`
	Tree    []LSDIItemTreeDTO      `json:"type"`
	Version string                 `json:"version"`
}

// LSDIBasicDataDTO : A struct containing basic data
type LSDIBasicDataDTO struct {
	Colloq               string                `json:"colloq"`
	ConsumeOnFull        bool                  `json:"consumeOnFull"`
	Consumed             bool                  `json:"consumed"`
	Depth                int                   `json:"depth"`
	Description          string                `json:"description"`
	From                 []string              `json:"from"`
	Gold                 LSDIGoldDTO           `json:"gold"` // Data Dragon includes the gold field for basic data, which is shared by both rune and item. However, only items have a gold field on them, representing their gold cost in the store. Since runes are not sold in the store, they have no gold cost.
	Group                string                `json:"group"`
	HideFromAll          bool                  `json:"hideFromAll"`
	ID                   int                   `json:"id"`
	Image                LSDIImageDTO          `json:"image"`
	InStore              bool                  `json:"inStore"`
	Into                 []string              `json:"into"`
	Maps                 map[string]bool       `json:"maps"`
	Name                 string                `json:"name"`
	Plaintext            string                `json:"plaintext"`
	RequiredChampion     string                `json:"requiredChampion"`
	Rune                 LSDIMetaDataDTO       `json:"rune"`
	SanitizedDescription string                `json:"sanitizedDescription"`
	SpecialRecipe        int                   `json:"specialRecipe"`
	Stacks               int                   `json:"stacks"`
	Stats                LSDIBasicDataStatsDto `json:"stats"`
	Tags                 []string              `json:"tags"`
}

// LSDIGroupDTO : A struct containing group data
type LSDIGroupDTO struct {
	MaxGroupOwnable string `json:"maxGroupOwnable"`
	Key             string `json:"key"`
}

// LSDIItemDTO : A struct containing item data
type LSDIItemDTO struct {
	Colloq               string                `json:"colloq"`
	ConsumeOnFull        bool                  `json:"consumeOnFull"`
	Consumed             bool                  `json:"consumed"`
	Depth                int                   `json:"depth"`
	Description          string                `json:"description"`
	Effect               map[string]string     `json:"effect"`
	From                 []string              `json:"from"`
	Gold                 LSDIGoldDTO           `json:"gold"` // Data Dragon includes the gold field for basic data, which is shared by both rune and item. However, only items have a gold field on them, representing their gold cost in the store. Since runes are not sold in the store, they have no gold cost.
	Group                string                `json:"group"`
	HideFromAll          bool                  `json:"hideFromAll"`
	ID                   int                   `json:"id"`
	Image                LSDIImageDTO          `json:"image"`
	InStore              bool                  `json:"inStore"`
	Into                 []string              `json:"into"`
	Maps                 map[string]bool       `json:"maps"`
	Name                 string                `json:"name"`
	Plaintext            string                `json:"plaintext"`
	RequiredChampion     string                `json:"requiredChampion"`
	Rune                 LSDIMetaDataDTO       `json:"rune"`
	SanitizedDescription string                `json:"sanitizedDescription"`
	SpecialRecipe        int                   `json:"specialRecipe"`
	Stacks               int                   `json:"stacks"`
	Stats                LSDIBasicDataStatsDto `json:"stats"`
	Tags                 []string              `json:"tags"`
}

// LSDIItemTreeDTO : A struct containing item tree data
type LSDIItemTreeDTO struct {
	Header string   `json:"header"`
	Tags   []string `json:"tags"`
}

// LSDIBasicDataStatsDto : A struct containing basic data stats information
type LSDIBasicDataStatsDto struct {
	FlatArmorMod                        float64 `json:"FlatArmorMod"`
	FlatAttackSpeedMod                  float64 `json:"FlatAttackSpeedMod"`
	FlatBlockMod                        float64 `json:"FlatBlockMod"`
	FlatCritChanceMod                   float64 `json:"FlatCritChanceMod"`
	FlatCritDamageMod                   float64 `json:"FlatCritDamageMod"`
	FlatEXPBonus                        float64 `json:"FlatEXPBonus"`
	FlatEnergyPoolMod                   float64 `json:"FlatEnergyPoolMod"`
	FlatEnergyRegenMod                  float64 `json:"FlatEnergyRegenMod"`
	FlatHPPoolMod                       float64 `json:"FlatHPPoolMod"`
	FlatHPRegenMod                      float64 `json:"FlatHPRegenMod"`
	FlatMPPoolMod                       float64 `json:"FlatMPPoolMod"`
	FlatMPRegenMod                      float64 `json:"FlatMPRegenMod"`
	FlatMagicDamageMod                  float64 `json:"FlatMagicDamageMod"`
	FlatMovementSpeedMod                float64 `json:"FlatMovementSpeedMod"`
	FlatPhysicalDamageMod               float64 `json:"FlatPhysicalDamageMod"`
	FlatSpellBlockMod                   float64 `json:"FlatSpellBlockMod"`
	PercentArmorMod                     float64 `json:"PercentArmorMod"`
	PercentAttackSpeedMod               float64 `json:"PercentAttackSpeedMod"`
	PercentBlockMod                     float64 `json:"PercentBlockMod"`
	PercentCritChanceMod                float64 `json:"PercentCritChanceMod"`
	PercentCritDamageMod                float64 `json:"PercentCritDamageMod"`
	PercentDodgeMod                     float64 `json:"PercentDodgeMod"`
	PercentEXPBonus                     float64 `json:"PercentEXPBonus"`
	PercentHPPoolMod                    float64 `json:"PercentHPPoolMod"`
	PercentHPRegenMod                   float64 `json:"PercentHPRegenMod"`
	PercentLifeStealMod                 float64 `json:"PercentLifeStealMod"`
	PercentMPPoolMod                    float64 `json:"PercentMPPoolMod"`
	PercentMPRegenMod                   float64 `json:"PercentMPRegenMod"`
	PercentMagicDamageMod               float64 `json:"PercentMagicDamageMod"`
	PercentMovementSpeedMod             float64 `json:"PercentMovementSpeedMod"`
	PercentPhysicalDamageMod            float64 `json:"PercentPhysicalDamageMod"`
	PercentSpellBlockMod                float64 `json:"PercentSpellBlockMod"`
	PercentSpellVampMod                 float64 `json:"PercentSpellVampMod"`
	RFlatArmorModPerLevel               float64 `json:"rFlatArmorModPerLevel"`
	RFlatArmorPenetrationMod            float64 `json:"rFlatArmorPenetrationMod"`
	RFlatArmorPenetrationModPerLevel    float64 `json:"rFlatArmorPenetrationModPerLevel"`
	RFlatCritChanceModPerLevel          float64 `json:"rFlatCritChanceModPerLevel"`
	RFlatCritDamageModPerLevel          float64 `json:"rFlatCritDamageModPerLevel"`
	RFlatDodgeMod                       float64 `json:"rFlatDodgeMod"`
	RFlatDodgeModPerLevel               float64 `json:"rFlatDodgeModPerLevel"`
	RFlatEnergyModPerLevel              float64 `json:"rFlatEnergyModPerLevel"`
	RFlatEnergyRegenModPerLevel         float64 `json:"rFlatEnergyRegenModPerLevel"`
	RFlatGoldPer10Mod                   float64 `json:"rFlatGoldPer10Mod"`
	RFlatHPModPerLevel                  float64 `json:"rFlatHPModPerLevel"`
	RFlatHPRegenModPerLevel             float64 `json:"rFlatHPRegenModPerLevel"`
	RFlatMPModPerLevel                  float64 `json:"rFlatMPModPerLevel"`
	RFlatMPRegenModPerLevel             float64 `json:"rFlatMPRegenModPerLevel"`
	RFlatMagicDamageModPerLevel         float64 `json:"rFlatMagicDamageModPerLevel"`
	RFlatMagicPenetrationMod            float64 `json:"rFlatMagicPenetrationMod"`
	RFlatMagicPenetrationModPerLevel    float64 `json:"rFlatMagicPenetrationModPerLevel"`
	RFlatMovementSpeedModPerLevel       float64 `json:"rFlatMovementSpeedModPerLevel"`
	RFlatPhysicalDamageModPerLevel      float64 `json:"rFlatPhysicalDamageModPerLevel"`
	RFlatSpellBlockModPerLevel          float64 `json:"rFlatSpellBlockModPerLevel"`
	RFlatTimeDeadMod                    float64 `json:"rFlatTimeDeadMod"`
	RFlatTimeDeadModPerLevel            float64 `json:"rFlatTimeDeadModPerLevel"`
	RPercentArmorPenetrationMod         float64 `json:"rPercentArmorPenetrationMod"`
	RPercentArmorPenetrationModPerLevel float64 `json:"rPercentArmorPenetrationModPerLevel"`
	RPercentAttackSpeedModPerLevel      float64 `json:"rPercentAttackSpeedModPerLevel"`
	RPercentCooldownMod                 float64 `json:"rPercentCooldownMod"`
	RPercentCooldownModPerLevel         float64 `json:"rPercentCooldownModPerLevel"`
	RPercentMagicPenetrationMod         float64 `json:"rPercentMagicPenetrationMod"`
	RPercentMagicPenetrationModPerLevel float64 `json:"rPercentMagicPenetrationModPerLevel"`
	RPercentMovementSpeedModPerLevel    float64 `json:"rPercentMovementSpeedModPerLevel"`
	RPercentTimeDeadMod                 float64 `json:"rPercentTimeDeadMod"`
	RPercentTimeDeadModPerLevel         float64 `json:"rPercentTimeDeadModPerLevel"`
}

// LSDIGoldDTO : A struct containing group data
type LSDIGoldDTO struct {
	Base        int  `json:"base"`
	Purchasable bool `json:"purchasable"`
	Sell        int  `json:"sell"`
	Total       int  `json:"total"`
}

// LSDIImageDTO : A struct containing image data
type LSDIImageDTO struct {
	Full   string `json:"full"`
	Group  string `json:"group"`
	H      int    `json:"h"`
	Sprite string `json:"sprite"`
	W      int    `json:"w"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

// LSDIMetaDataDTO : A struct containing meta data
type LSDIMetaDataDTO struct {
	IsRune bool   `json:"isRune"`
	Tier   string `json:"tier"`
	Type   string `json:"type"`
}

// GetStaticItems : Retrieves item list
func (c *RiotClient) GetStaticItems(region string, opts *StaticItemOpts) (LSDIItemListDTO, error) {
	var items LSDIItemListDTO
	params := &url.Values{}
	// Builds out query params based on options passed
	if opts != nil {
		if opts.Locale != "" {
			params.Add("locale", opts.Locale)
		}
		if opts.Version != "" {
			params.Add("version", opts.Version)
		}
		if opts.ItemListData != "" {
			params.Add("itemListData", opts.ItemListData)
		}
	}

	// Performs the http request on Riots API to retrieve all items
	resBody, err := c.riotRequest("/api/lol/static-data/"+region+"/v1.2/item", region, params)
	if err != nil {
		return items, err
	}

	// Unmarshals the response JSON into a LSDCChampionListDTO struct
	if err = json.Unmarshal(resBody, &items); err != nil {
		return items, err
	}
	return items, nil
}

// GetStaticItemByID : Retrieves item by its unique id
func (c *RiotClient) GetStaticItemByID(region string, id int, opts *StaticItemOpts) (LSDIItemDTO, error) {
	var item LSDIItemDTO
	params := &url.Values{}
	// Builds out query params based on options passed
	if opts != nil {
		if opts.Locale != "" {
			params.Add("locale", opts.Locale)
		}
		if opts.Version != "" {
			params.Add("version", opts.Version)
		}
		if opts.ItemData != "" {
			params.Add("itemData", opts.ItemListData)
		}
	}

	// Performs the http request on Riots API to retrieve the specified item
	resBody, err := c.riotRequest("/api/lol/static-data/"+region+"/v1.2/item/"+strconv.Itoa(id), region, params)
	if err != nil {
		return item, err
	}

	// Unmarshals the response JSON into a LSDCChampionListDTO struct
	if err = json.Unmarshal(resBody, &item); err != nil {
		return item, err
	}
	return item, nil
}
