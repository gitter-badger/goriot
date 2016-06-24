// Package goriot : Currently up to date as of static-v1.2
package goriot

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// StaticMasteryOpts : A struct containing optional paramenters used with static mastery endpoints
type StaticMasteryOpts struct {
	Locale          string
	Version         string
	MasteryData     string
	MasteryListData string
}

// LSDMMasteryListDTO : A struct containing mastery list data
type LSDMMasteryListDTO struct {
	Data    map[string]LSDMMasteryDTO `json:"data"`
	Tree    LSDMMasteryTreeDTO        `json:"tree"`
	Type    string                    `json:"type"`
	Version string                    `json:"version"`
}

// LSDMMasteryDTO : A struct containing mastery data
type LSDMMasteryDTO struct {
	Description          []string     `json:"description"`
	ID                   int          `json:"id"`
	Image                LSDMImageDTO `json:"image"`
	MasteryTree          string       `json:"masteryTree"`
	Name                 string       `json:"name"`
	Prereq               string       `json:"prereq"`
	Ranks                int          `json:"ranks"`
	SanitizedDescription []string     `json:"sanitizedDescription"`
}

// LSDMMasteryTreeDTO : A struct containing mastery tree data
type LSDMMasteryTreeDTO struct {
	Cunning  []LSDMMasteryTreeListDTO `json:"Cunning"`
	Ferocity []LSDMMasteryTreeListDTO `json:"Ferocity"`
	Resolve  []LSDMMasteryTreeListDTO `json:"Resolve"`
}

// LSDMMasteryTreeListDTO : A struct containing mastery tree list data
type LSDMMasteryTreeListDTO struct {
	MasteryTreeItems []LSDMMasteryTreeItemDTO `json:"masteryTreeItems"`
}

// LSDMMasteryTreeItemDTO : A struct containing mastery tree item data
type LSDMMasteryTreeItemDTO struct {
	MasteryID int    `json:"masteryId"`
	Prereq    string `json:"prereq"`
}

// GetStaticMasteries : Retrieves mastery list
func (c *RiotClient) GetStaticMasteries(region string, opts *StaticMasteryOpts) (LSDMMasteryListDTO, error) {
	var masteries LSDMMasteryListDTO
	params := &url.Values{}
	// Builds out query params based on options passed
	if opts != nil {
		if opts.Locale != "" {
			params.Add("locale", opts.Locale)
		}
		if opts.Version != "" {
			params.Add("version", opts.Version)
		}
		if opts.MasteryListData != "" {
			params.Add("masteryListData", opts.MasteryListData)
		}
	}

	// Performs the http request on Riots API to retrieve all language strings
	resBody, err := c.riotRequest("/api/lol/static-data/"+region+"/v1.2/mastery", "global", params)
	if err != nil {
		return masteries, err
	}

	// Unmarshals the response JSON into a LSDCChampionListDTO struct
	if err = json.Unmarshal(resBody, &masteries); err != nil {
		return masteries, err
	}
	return masteries, nil
}

// GetStaticMasteryByID : Retrieves mastery item by its unique id
func (c *RiotClient) GetStaticMasteryByID(region string, id int, opts *StaticMasteryOpts) (LSDMMasteryDTO, error) {
	var mastery LSDMMasteryDTO
	params := &url.Values{}
	// Builds out query params based on options passed
	if opts != nil {
		if opts.Locale != "" {
			params.Add("locale", opts.Locale)
		}
		if opts.Version != "" {
			params.Add("version", opts.Version)
		}
		if opts.MasteryData != "" {
			params.Add("masteryData", opts.MasteryData)
		}
	}

	// Performs the http request on Riots API to retrieve all language strings
	resBody, err := c.riotRequest("/api/lol/static-data/"+region+"/v1.2/mastery/"+strconv.Itoa(id), "global", params)
	if err != nil {
		return mastery, err
	}

	// Unmarshals the response JSON into a LSDCChampionListDTO struct
	if err = json.Unmarshal(resBody, &mastery); err != nil {
		return mastery, err
	}
	return mastery, nil
}
