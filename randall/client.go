package randall

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const domain = "https://api.harvestapp.com"

var client *http.Client

// The interface used to interact with the entire Harvest API
type Client struct {
	HeaderValues *HeaderValues
	Users        UsersApi
}

func New(hv HeaderValues) *Client {
	client = &http.Client{}
	return &Client{
		HeaderValues: &hv,
		Users:        newUserV2(domain, &hv),
	}
}

func newRequest(method, url string, hv HeaderValues, queryStr ...map[string]string) *http.Request {
	r, _ := http.NewRequest("GET", url, nil)

	if len(queryStr) > 0 {
		query := r.URL.Query()
		for key, value := range queryStr[0] {
			query.Set(key, value)
		}

		r.URL.RawQuery = query.Encode()
	}

	r.Header.Set("User-Agent", fmt.Sprintf("%s (%s)", hv.UserAgentApp, hv.UserAgentEmail))
	r.Header.Set("Harvest-Account-ID", hv.AccountId)
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", hv.AccessToken))

	return r
}

func readJsonResponse(r *http.Request) (*JsonResponse, error) {
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

	jsonResp := JsonResponse{
		StatusCode: resp.StatusCode,
		Data:       data,
	}

	return &jsonResp, nil
}
