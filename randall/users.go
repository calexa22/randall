package randall

import (
	"fmt"
	"strconv"
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
	r := newRequest("GET", fmt.Sprintf("%s/me", api.baseUrl), *api.hv)
	return readJsonResponse(r)
}

// Retrieves the currently authenticated user. Returns a user object and a 200 OK response code.
func (api *UsersApi) Users(activeOnly bool, page int, perPage int) (*JsonResponse, error) {
	query := map[string]string{
		"is_active": strconv.FormatBool(activeOnly),
		"page":      strconv.Itoa(max(page, 1)),
		"per_page":  strconv.Itoa(min(max(page, 1), 2000)),
	}

	r := newRequest("GET", api.baseUrl, *api.hv, query)
	return readJsonResponse(r)
}

func (api *UsersApi) User(userId int) (*JsonResponse, error) {

	r := newRequest("GET", fmt.Sprintf("%s/users/%d", api.baseUrl, userId), *api.hv)
	return readJsonResponse(r)
}
