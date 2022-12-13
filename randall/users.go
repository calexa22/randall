package randall

import (
	"fmt"
	"time"
)

const (
	AccessRoleMember        = "member"
	AccessRoleManager       = "manager"
	AccessRoleAdministrator = "administrator"
)

// Encapsulates the Harvest API methods under /users
type UsersApi struct {
	baseUrl string
	client  *internalClient
}

type CreateUserRequest struct {
	FirstName                   string   `json:"first_name"`
	LastName                    string   `json:"last_name"`
	Email                       string   `json:"email"`
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

type UpdateUserRequest struct {
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

type UpdateAssignedTeammatesRequest struct {
	TeammateIds []uint `json:"teammate_ids"`
}

type CreateBillableRateRequest struct {
	Amount    float32   `json:"amount"`
	StartDate time.Time `json:"date,omitempty" layout:"2006-01-02"`
}

type CreateCostRateRequest struct {
	Amount    float32   `json:"amount"`
	StartDate time.Time `json:"date,omitempty" layout:"2006-01-02"`
}

func newUsersV2(client *internalClient) UsersApi {
	return UsersApi{
		baseUrl: "v2/users",
		client:  client,
	}
}

// Retrieves the currently authenticated user. Returns a user object and a 200 OK response code.
func (api UsersApi) MyUser() (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/me", api.baseUrl))
}

// Retrieves the currently authenticated user. Returns a user object and a 200 OK response code.
func (api UsersApi) AllUsers(params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(api.baseUrl, getOptionalCollectionParams(params))
}

// Retrieves the user with the give UserID. Returns a user object and a 200 OK response code if valid ID provided.
func (api UsersApi) GetUser(userId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d", api.baseUrl, userId))
}

func (api UsersApi) CreateUser(req CreateUserRequest) (HarvestResponse, error) {
	return api.client.doPost(api.baseUrl, req)
}

func (api UsersApi) UpdateUser(userId uint, req UpdateUserRequest) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d", api.baseUrl, userId), req)
}

func (api UsersApi) ArchiveUser(userId uint) (HarvestResponse, error) {
	isActive := false
	return api.client.doPatch(fmt.Sprintf("%s/%d", api.baseUrl, userId), UpdateUserRequest{
		IsActive: &isActive,
	})
}

func (api UsersApi) DeleteUser(userId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d", api.baseUrl, userId))
}

func (api UsersApi) UnarchiveUser(userId uint) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d", api.baseUrl, userId), UpdateUserRequest{
		IsActive: OptionalBool(true),
	})
}

func (api UsersApi) GetAssignedTeammates(userId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d/teammates", api.baseUrl, userId))
}

func (api UsersApi) UpdateAssignedTeammates(userId uint, teammateIds UpdateAssignedTeammatesRequest) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d/teammates", api.baseUrl, userId), teammateIds)
}

func (api UsersApi) GetBillableRates(userId uint, params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(
		fmt.Sprintf("%s/%d/billable_rates", api.baseUrl, userId),
		getOptionalCollectionParams(params))
}

func (api UsersApi) GetBillableRate(userId, billableRateId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d/billable_rates/%d", api.baseUrl, userId, billableRateId))
}

func (api UsersApi) CreateBillableRate(userId uint, req CreateBillableRateRequest) (HarvestResponse, error) {
	return api.client.doPost(fmt.Sprintf("%s/%d", api.baseUrl, userId), req)
}

func (api UsersApi) GetCostRates(userId uint, params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(
		fmt.Sprintf("%s/%d/cost_rates", api.baseUrl, userId),
		getOptionalCollectionParams(params),
	)
}

func (api UsersApi) GetCostRate(userId, costRateId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d/cost_rates/%d", api.baseUrl, userId, costRateId))
}

func (api UsersApi) CreateCostRate(userId uint, req CreateCostRateRequest) (HarvestResponse, error) {
	return api.client.doPost(fmt.Sprintf("%s/%d", api.baseUrl, userId), req)
}

func (api UsersApi) GetActiveProjectAssignments(userId uint, params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(
		fmt.Sprintf("%s/%d/project_assignments", api.baseUrl, userId),
		getOptionalCollectionParams(params),
	)
}

func (api UsersApi) GetMyActiveProjectAssignments(params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(
		fmt.Sprintf("%s/me/project_assignments", api.baseUrl),
		getOptionalCollectionParams(params),
	)
}
