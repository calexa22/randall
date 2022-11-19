package randall

import (
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

func newRequest(method, url string, hv HeaderValues) *http.Request {
	r, _ := http.NewRequest("GET", url, nil)

	r.Header.Set("User-Agent", fmt.Sprintf("%s (%s)", hv.UserAgentApp, hv.UserAgentEmail))
	r.Header.Set("Harvest-Account-ID", hv.AccountId)
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", hv.AccessToken))

	return r
}
