// Package goriot : Currently up to date as of lol-status-v1.0
package goriot

import "encoding/json"

// LSShard : A struct containing status information
type LSShard struct {
	Hostname  string   `json:"hostname"`
	Locales   []string `json:"locales"`
	Name      string   `json:"name"`
	RegionTag string   `json:"region_tag"`
	Slug      string   `json:"slug"`
}

// LSShardStatus : The server status information pertaining to a particular region
type LSShardStatus struct {
	Hostname  string      `json:"hostname"`
	Locales   []string    `json:"locales"`
	Name      string      `json:"name"`
	RegionTag string      `json:"region_tag"`
	Services  []LSService `json:"services"`
	Slug      string      `json:"slug"`
}

// LSService : Information on a service that is running on a shard
type LSService struct {
	Incidents []LSIncident `json:"incidents"`
	Name      string       `json:"name"`
	Slug      string       `json:"slug"`
	Status    string       `json:"status"`
}

// LSIncident : A struct containing information on a service incident
type LSIncident struct {
	Active    bool        `json:"active"`
	CreatedAt string      `json:"created_at"`
	ID        int64       `json:"id"`
	Updates   []LSMessage `json:"updates"`
}

// LSMessage : A struct containing message information of an incident
type LSMessage struct {
	Author       string          `json:"author"`
	Content      string          `json:"content"`
	CreatedAt    string          `json:"created_at"`
	ID           string          `json:"id"`
	Severity     string          `json:"serverity"`
	Translations []LSTranslation `json:"translations"`
	UpdatedAt    string          `json:"updated_at"`
}

// LSTranslation : A struct containing translation information
type LSTranslation struct {
	Content   string `json:"content"`
	Locale    string `json:"locale"`
	UpdatedAt string `json:"updated_at"`
}

// GetShards : Get shard list
func (c *RiotClient) GetShards() ([]LSShard, error) {
	var shards []LSShard
	// Performs the http request on Riots API to retrieve a list of shards
	resBody, err := c.riotRequest("/shards", "", nil)
	if err != nil {
		return shards, err
	}

	// Unmarshals the response JSON into our array of shards struct
	if err := json.Unmarshal(resBody, &shards); err != nil {
		return shards, err
	}
	return shards, nil
}

// GetShardStatus : Get shard status. Returns the data available on the status.leagueoflegends.com website for the given region
func (c *RiotClient) GetShardStatus(region string) (LSShardStatus, error) {
	var shardStatus LSShardStatus
	// Performs the http request on Riots API to retrieve a specific shard status
	resBody, err := c.riotRequest("/shards/"+region, region, nil)
	if err != nil {
		return shardStatus, err
	}

	// Unmarshals the response JSON into our array of shards struct
	if err := json.Unmarshal(resBody, &shardStatus); err != nil {
		return shardStatus, err
	}
	return shardStatus, nil
}
