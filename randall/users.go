package randall

import (
	"encoding/json"
	"fmt"
)

// Encapsulates the Harvest API methods under /users
type UsersApi struct {
	baseUrl string
	hv      *HeaderValues
}

func newUserV2(domain string, hv *HeaderValues) UsersApi {
	return UsersApi{
		baseUrl: fmt.Sprintf("%s/v2/users", domain),
		hv:      hv,
	}
}

// Retrieves the currently authenticated user. Returns a user object and a 200 OK response code.
func (api *UsersApi) Me() (*JsonResponse, error) {
	req := newRequest("GET", fmt.Sprintf("%s/me", api.baseUrl), *api.hv)

	resp, err := client.Do(req)

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
