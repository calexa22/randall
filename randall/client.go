package randall

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const domain = "https://api.harvestapp.com"

var client *http.Client

// The interface used to interact with the entire Harvest API.
type Client struct {
	// A collection of values to be assigned to the request headers required by Harvest (Auth, User-Agent, Harvest-Account-ID)
	Headers *HarvestHeaders
	// The interface used to make calls to /users endpoints under the given Client.
	Users UsersApi
}

// Initializes a new instance of Client. Requests through the Client will have the headers
// required by the Harvest API with the passed in values
func New(headers HarvestHeaders) *Client {
	client = &http.Client{}
	return &Client{
		Headers: &headers,
		Users:   newUserV2(domain, &headers),
	}
}

func newRequest(method, url string, headers HarvestHeaders, queryStr ...map[string]string) *http.Request {
	r, _ := http.NewRequest("GET", url, nil)

	if len(queryStr) > 0 {
		query := r.URL.Query()
		for key, value := range queryStr[0] {
			query.Set(key, value)
		}

		r.URL.RawQuery = query.Encode()
	}

	r.Header.Set("User-Agent", fmt.Sprintf("%s (%s)", headers.UserAgentApp, headers.UserAgentEmail))
	r.Header.Set("Harvest-Account-ID", headers.AccountId)
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", headers.AccessToken))

	return r
}

func readResponse(r *http.Request) (*HarvestResponse, error) {
	resp, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data map[string]any

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		return nil, err
	}

	harvestResp := HarvestResponse{
		StatusCode: resp.StatusCode,
		Data:       data,
	}

	return &harvestResp, nil
}
