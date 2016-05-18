package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// RiotClient : A globally used struct used to store
// information needed to make requests to the Riot API
type RiotClient struct {
	APIKey string
	Region string
}

// NewRiotClient : A simple function used to create a client that will bind to Riots API
func NewRiotClient(apiKey string, region string) *RiotClient {
	return &RiotClient{
		APIKey: apiKey,
		Region: strings.ToLower(region),
	}
}

// riotRequest : Func responsible for all http requests made to the Riot API
func (c *RiotClient) riotRequest(uri string, params *url.Values) ([]byte, error) {
	// Builds out the request based on passed parameters
	if params == nil {
		params = &url.Values{}
	}
	params.Add("api_key", c.APIKey)
	log.Println(params)
	req, err := http.NewRequest("GET", "https://"+c.Region+".api.pvp.net"+uri+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	// Sets the headers used for all requests
	req.Header.Add("User-Agent", "Sdwolfe32-GoRiot")
	req.Header.Add("Accept-Language", "en-US")
	req.Header.Add("Accept-Charset", "ISO-8859-1,utf-8")

	// Executes the request
	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// Reads and returns the body of the http.Response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(body))
	return body, nil
}