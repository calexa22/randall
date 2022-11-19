package randall

import (
	"fmt"
	"strconv"
)

// Encapsulates the Harvest API methods under /users
type UsersApi struct {
	baseUrl string
	headers *HarvestHeaders
}

func newUserV2(domain string, headers *HarvestHeaders) UsersApi {
	return UsersApi{
		baseUrl: fmt.Sprintf("%s/v2/users", domain),
		headers: headers,
	}
}

// Retrieves the currently authenticated user. Returns a user object and a 200 OK response code.
func (api *UsersApi) Me() (*HarvestResponse, error) {
	r := newRequest("GET", fmt.Sprintf("%s/me", api.baseUrl), *api.headers)
	return readResponse(r)
}

// Retrieves the currently authenticated user. Returns a user object and a 200 OK response code.
func (api *UsersApi) Users(activeOnly bool, page int, perPage int) (*HarvestResponse, error) {
	query := map[string]string{
		"is_active": strconv.FormatBool(activeOnly),
		"page":      strconv.Itoa(max(page, 1)),
		"per_page":  strconv.Itoa(min(max(page, 1), 2000)),
	}

	r := newRequest("GET", api.baseUrl, *api.headers, query)
	return readResponse(r)
}

// Retrieves the user with the give UserID. Returns a user object and a 200 OK response code if valid ID provided.
func (api *UsersApi) User(userId int) (*HarvestResponse, error) {
	r := newRequest("GET", fmt.Sprintf("%s/users/%d", api.baseUrl, userId), *api.headers)
	return readResponse(r)
}
