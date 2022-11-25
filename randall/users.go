package randall

import (
	"fmt"
	"strconv"
)

const (
	baseUsersV2Url          = "v2/users"
	AccessRoleMember        = "member"
	AccessRoleManager       = "manager"
	AccessRoleAdministrator = "administrator"
)

// Encapsulates the Harvest API methods under /users
type UsersApi struct {
	baseUrl string
	client  *internalClient
}

func newUsersV2(client *internalClient) UsersApi {
	return UsersApi{
		baseUrl: baseUsersV2Url,
		client:  client,
	}
}

// Retrieves the currently authenticated user. Returns a user object and a 200 OK response code.
func (api UsersApi) Me() (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/me", api.baseUrl))
}

// Retrieves the currently authenticated user. Returns a user object and a 200 OK response code.
func (api UsersApi) All(activeOnly bool, page int, perPage int) (HarvestResponse, error) {
	query := map[string]string{
		"is_active": strconv.FormatBool(activeOnly),
		"page":      strconv.Itoa(max(page, 1)),
		"per_page":  strconv.Itoa(min(max(page, 1), 2000)),
	}

	return api.client.DoGet(api.baseUrl, query)
}

// Retrieves the user with the give UserID. Returns a user object and a 200 OK response code if valid ID provided.
func (api UsersApi) Get(userId int) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.baseUrl, userId))
}
