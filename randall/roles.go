package randall

import (
	"fmt"
)

// Encapsulates the Harvest API methods under /roles.
type RolesApi struct {
	baseUrl string
	client  *internalClient
}

type CreateRoleRequest struct {
	// The name of the role.
	Name string `json:"name"`
	// An optional list of users to assign this new Role to.
	UserIds []uint `json:"user_ids,omitempty"`
}

type UpdateRoleRequest struct {
	// The name of the role.
	Name *string `json:"name,omitempty"`
	// An optional list of users to assign this new Role to.
	UserIds []uint `json:"user_ids,omitempty"`
}

func newRolesV2(client *internalClient) RolesApi {
	return RolesApi{
		baseUrl: "v2/roles",
		client:  client,
	}
}

// Retrieves a list of all Roles, with an optional query string.
func (api RolesApi) GetAllRoles(params ...HarvestCollectionParams) (HarvestResponse, error) {
	var param *HarvestCollectionParams

	if len(params) > 0 {
		param = &params[0]
	}
	return api.client.DoGet(api.baseUrl, param)
}

// Retrieves a Role with the given RoleID.
func (api RolesApi) GetRole(roleId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.baseUrl, roleId))
}

// Creates a new Role.
func (api RolesApi) CreateRole(req CreateRoleRequest) (HarvestResponse, error) {
	return api.client.DoPost(api.baseUrl, req)
}

// Updates a Role with the given RoleID.
func (api RolesApi) UpdateRole(roleId uint, req UpdateRoleRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.baseUrl, roleId), req)
}

// Deletes a Role with the given RoleID.
func (api RolesApi) DeleteRole(roleId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d", api.baseUrl, roleId))
}
