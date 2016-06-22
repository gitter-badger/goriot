package goriot

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// RiotClient : A globally used struct used to store
// information needed to make requests to the Riot API
type RiotClient struct {
	APIKey string
}

// NewRiotClient : A simple function used to create a client that will bind to Riots API
func NewRiotClient(apiKey string) (*RiotClient, error) {
	// Performs a test request in order to validate the api key
	rc := new(RiotClient)
	rc.APIKey = apiKey
	_, err := rc.GetVersions("na")
	if err != nil {
		return nil, err
	}
	return rc, nil
}

// riotRequest : Func responsible for all http requests made to the Riot API
func (c *RiotClient) riotRequest(uri, region string, params *url.Values) ([]byte, error) {
	// Builds out the request based on passed parameters
	if params == nil {
		params = &url.Values{}
	}
	params.Add("api_key", c.APIKey)

	// Builds out the correct url
	url := ""
	if strings.Contains(uri, "shards") {
		url = "http://status.leagueoflegends.com" + uri
	} else {
		url = "https://" + strings.ToLower(region) + ".api.pvp.net" + uri + "?" + params.Encode()
	}

	// Builds out a request struct to be executed on the Riot API
	req, err := http.NewRequest("GET", url, nil)
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

	// Checks response code and returns err if not 200
	if res.StatusCode != 200 {
		// TODO handle error codes
		return nil, nil
	}

	// Reads and returns the body of the http.Response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
