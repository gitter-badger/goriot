// Package goriot : Currently up to date as of lol-status-v1.0
package goriot

import "encoding/json"

// GetShards : Get shard list
func (c *RiotClient) GetShards() (*[]Shard, error) {
	// Performs the http request on Riots API to retrieve a list of shards
	resBody, err := c.riotRequest("/shards", "", nil)
	if err != nil {
		return nil, err
	}

	// Unmarshals the response JSON into our array of shards struct
	shards := new([]Shard)
	if err := json.Unmarshal(resBody, shards); err != nil {
		return nil, err
	}
	return shards, nil
}

// GetShardStatus : Get shard status. Returns the data available on the status.leagueoflegends.com website for the given region
func (c *RiotClient) GetShardStatus(region string) (*ShardStatus, error) {
	// Performs the http request on Riots API to retrieve a specific shard status
	resBody, err := c.riotRequest("/shards/"+region, region, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshals the response JSON into our array of shards struct
	shardStatus := new(ShardStatus)
	if err := json.Unmarshal(resBody, shardStatus); err != nil {
		return nil, err
	}
	return shardStatus, nil
}
