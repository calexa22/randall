package randall

import (
	"fmt"
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

type AssignedTeammatesPatchRequest struct {
	TeammateIds []uint `json:"teammate_ids"`
}

type CreateUserRequest struct {
	FirstName                   string   `json:"first_name"`
	LastName                    string   `json:"last_name"`
	Email                       string   `json:"email"`
	Timezone                    string   `json:"timezone,omitempty"`
	HasAccesToAllFutureProjects bool     `json:"has_access_to_all_future_projects,omitempty"`
	IsContractor                bool     `json:"is_contractor,omitempty"`
	IsActive                    *bool    `json:"is_active,omitempty"`
	WeeklyCapacity              uint     `json:"weekly_capacity,omitempty"`
	DefaultHourlyRate           float32  `json:"default_hourly_rate,omitempty"`
	CostRate                    float32  `json:"cost_rate,omitempty"`
	Roles                       []string `json:"roles,omitempty"`
	AccessRoles                 []string `json:"access_roles,omitempty"`
}

type UpdateUserPatchRequest struct {
	FirstName                   *string  `json:"first_name,omitempty"`
	LastName                    *string  `json:"last_name,omitempty"`
	Email                       *string  `json:"email,omitempty"`
	Timezone                    *string  `json:"timezone,omitempty"`
	HasAccesToAllFutureProjects *bool    `json:"has_access_to_all_future_projects,omitempty"`
	IsContractor                *bool    `json:"is_contractor,omitempty"`
	IsActive                    *bool    `json:"is_active,omitempty"`
	WeeklyCapacity              *uint    `json:"weekly_capacity,omitempty"`
	DefaultHourlyRate           *float32 `json:"default_hourly_rate,omitempty"`
	CostRate                    *float32 `json:"cost_rate,omitempty"`
	Roles                       []string `json:"roles,omitempty"`
	AccessRoles                 []string `json:"access_roles,omitempty"`
}

type CreateBillableRateRequest struct {
	Amount    float32     `json:"amount"`
	StartDate HarvestDate `json:"date,omitempty"`
}

type CreateCostRateRequest struct {
	Amount    float32     `json:"amount"`
	StartDate HarvestDate `json:"date,omitempty"`
}

func newUsersV2(client *internalClient) UsersApi {
	return UsersApi{
		baseUrl: baseUsersV2Url,
		client:  client,
	}
}

// Retrieves the currently authenticated user. Returns a user object and a 200 OK response code.
func (api UsersApi) MyUser() (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/me", api.baseUrl))
}

// Retrieves the currently authenticated user. Returns a user object and a 200 OK response code.
func (api UsersApi) AllUsers(params ...QueryStringProvider) (HarvestResponse, error) {
	query, err := getUrlValues(params...)

	if err != nil {
		return HarvestResponse{}, err
	}

	return api.client.DoGet(api.baseUrl, query)
}

// Retrieves the user with the give UserID. Returns a user object and a 200 OK response code if valid ID provided.
func (api UsersApi) GetUser(userId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.baseUrl, userId))
}

func (api UsersApi) CreateUser(req CreateUserRequest) (HarvestResponse, error) {
	return api.client.DoPost(api.baseUrl, req)
}

func (api UsersApi) UpdateUser(userId uint, req UpdateUserPatchRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.baseUrl, userId), req)
}

func (api UsersApi) ArchiveUser(userId uint) (HarvestResponse, error) {
	isActive := false
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.baseUrl, userId), UpdateUserPatchRequest{
		IsActive: &isActive,
	})
}

func (api UsersApi) DeleteUser(userId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d", api.baseUrl, userId))
}

func (api UsersApi) UnarchiveUser(userId uint) (HarvestResponse, error) {
	isActive := true
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.baseUrl, userId), UpdateUserPatchRequest{
		IsActive: &isActive,
	})
}

func (api UsersApi) GetAssignedTeammates(userId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d/teammates", api.baseUrl, userId))
}

func (api UsersApi) UpdateAssignedTeammates(userId uint, teammateIds AssignedTeammatesPatchRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d/teammates", api.baseUrl, userId), teammateIds)
}

func (api UsersApi) GetBillableRates(userId uint, params ...QueryStringProvider) (HarvestResponse, error) {
	query, err := getUrlValues(params...)

	if err != nil {
		return HarvestResponse{}, err
	}

	return api.client.DoGet(fmt.Sprintf("%s/%d/billable_rates", api.baseUrl, userId), query)
}

func (api UsersApi) GetBillableRate(userId, billableRateId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d/billable_rates/%d", api.baseUrl, userId, billableRateId))
}

func (api UsersApi) CreateBillableRate(userId uint, req CreateBillableRateRequest) (HarvestResponse, error) {
	return api.client.DoPost(fmt.Sprintf("%s/%d", api.baseUrl, userId), req)
}

func (api UsersApi) GetCostRates(userId uint, params ...QueryStringProvider) (HarvestResponse, error) {
	query, err := getUrlValues(params...)

	if err != nil {
		return HarvestResponse{}, err
	}

	return api.client.DoGet(fmt.Sprintf("%s/%d/cost_rates", api.baseUrl, userId), query)
}

func (api UsersApi) GetCostRate(userId, costRateId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d/cost_rates/%d", api.baseUrl, userId, costRateId))
}

func (api UsersApi) CreateCostRate(userId uint, req CreateCostRateRequest) (HarvestResponse, error) {
	return api.client.DoPost(fmt.Sprintf("%s/%d", api.baseUrl, userId), req)
}

func (api UsersApi) GetActiveProjectAssignments(userId uint, params ...QueryStringProvider) (HarvestResponse, error) {
	query, err := getUrlValues(params...)

	if err != nil {
		return HarvestResponse{}, err
	}

	return api.client.DoGet(fmt.Sprintf("%s/%d/project_assignments", api.baseUrl, userId), query)
}

func (api UsersApi) GetMyActiveProjectAssignments(params ...QueryStringProvider) (HarvestResponse, error) {
	query, err := getUrlValues(params...)

	if err != nil {
		return HarvestResponse{}, err
	}

	return api.client.DoGet(fmt.Sprintf("%s/me/project_assignments", api.baseUrl), query)
}
