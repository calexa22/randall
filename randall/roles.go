package randall

import (
	"fmt"
)

const baseRolesV2Url = "v2/roles"

// Encapsulates the Harvest API methods under /projects
type RolesApi struct {
	baseUrl string
	client  *internalClient
}

type CreateRoleRequest struct {
	Name    string `json:"name"`
	UserIds []uint `json:"user_ids,omitempty"`
}

type UpdateRolePatchRequest struct {
	Name    string `json:"name,omitempty"`
	UserIds []uint `json:"user_ids,omitempty"`
}

func newRolesV2(client *internalClient) RolesApi {
	return RolesApi{
		baseUrl: baseRolesV2Url,
		client:  client,
	}
}

// Retrieves a list of Roles.
func (api *RolesApi) GetAllRoles(params HarvestCollectionParams) (HarvestResponse, error) {
	query, err := getUrlValues(params)

	if err != nil {
		return HarvestResponse{}, err
	}

	return api.client.DoGet(api.baseUrl, query)
}

// Retrieves a Role with the given RoleID.
func (api *RolesApi) GetRole(roleId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.baseUrl, roleId))
}

func (api *RolesApi) CreateRole(req CreateRoleRequest) (HarvestResponse, error) {
	return api.client.DoPost(api.baseUrl, req)
}

func (api *RolesApi) UpdateRole(roleId uint, req UpdateRolePatchRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.baseUrl, roleId), req)
}

func (api *RolesApi) DeleteRole(roleId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d", api.baseUrl, roleId))
}
