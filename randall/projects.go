package randall

import (
	"fmt"
)

const (
	baseProjectsV2Url = "v2/projects"

	ProjectBilledByProject = "Project"
	ProjectBilledByTask    = "Tasks"
	ProjectBilledByPeople  = "People"
	ProjectBilledByNone    = "none"

	ProjectBudgetByHoursPerProject  = "project"
	ProjectBudgetByTotalProjectFees = "project_cost"
	ProjectBudgetByHrsPerTask       = "task"
	ProjectBudgetByFeesPerTask      = "task_fees"
	ProjectBudgetByHrsPerPerson     = "person"
	ProjectBudgetByNone             = "none"
)

type CreateProjectRequest struct {
	ClientId                         uint         `json:"client_id"`
	Name                             string       `json:"name"`
	IsBillable                       bool         `json:"is_billable"`
	BillBy                           string       `json:"bill_by"`
	BudgetBy                         string       `json:"budget_by"`
	Code                             *string      `json:"code,omitempty"`
	IsActive                         *bool        `json:"is_active,omitempty"`
	IsFixedFee                       *bool        `json:"is_fixed_fee,omitempty"`
	HourlyRate                       *float32     `json:"hourly_rate,omitempty"`
	Budget                           *float32     `json:"budget,omitempty"`
	BudgetIsMonthly                  *bool        `json:"budget_is_monthly,omitempty"`
	NotifyWhenOverBudget             *bool        `json:"notify_when_over_budget,omitempty"`
	OverBudgetNotificationPercentage *float32     `json:"over_budget_notification_percentage,omitempty"`
	ShowBudgetToAll                  *bool        `json:"show_budget_to_all,omitempty"`
	CostBudget                       *float32     `json:"cost_budget,omitempty"`
	CostBudgetIncludeExpenses        *bool        `json:"cost_budget_include_expenses,omitempty"`
	Fee                              *float32     `json:"fee,omitempty"`
	Notes                            *string      `json:"notes,omitempty"`
	StartsOn                         *HarvestDate `json:"starts_on,omitempty"`
	EndsOn                           *HarvestDate `json:"ends_on,omitempty"`
}

type ProjectPatchRequest struct {
	ClientId                         *uint        `json:"client_id,omitempty"`
	Name                             *string      `json:"name,omitempty"`
	IsBillable                       *bool        `json:"is_billable,omitempty"`
	BillBy                           *string      `json:"bill_by,omitempty"`
	BudgetBy                         *string      `json:"budget_by,omitempty"`
	Code                             *string      `json:"code,omitempty"`
	IsActive                         *bool        `json:"is_active,omitempty"`
	IsFixedFee                       *bool        `json:"is_fixed_fee,omitempty"`
	HourlyRate                       *float32     `json:"hourly_rate,omitempty"`
	Budget                           *float32     `json:"budget,omitempty"`
	BudgetIsMonthly                  *bool        `json:"budget_is_monthly,omitempty"`
	NotifyWhenOverBudget             *bool        `json:"notify_when_over_budget,omitempty"`
	OverBudgetNotificationPercentage *float32     `json:"over_budget_notification_percentage,omitempty"`
	ShowBudgetToAll                  *bool        `json:"show_budget_to_all,omitempty"`
	CostBudget                       *float32     `json:"cost_budget,omitempty"`
	CostBudgetIncludeExpenses        *bool        `json:"cost_budget_include_expenses,omitempty"`
	Fee                              *float32     `json:"fee,omitempty"`
	Notes                            *string      `json:"notes,omitempty"`
	StartsOn                         *HarvestDate `json:"starts_on,omitempty"`
	EndsOn                           *HarvestDate `json:"ends_on,omitempty"`
}

type CreateUserAssignmentRequest struct {
	UserId           uint     `json:"user_id"`
	Name             string   `json:"name"`
	IsActive         *bool    `json:"is_active,omitempty"`
	IsProjectManager *bool    `json:"is_project_manager,omitempty"`
	UseDefaultRates  *bool    `json:"use_default_rates,omitempty"`
	HourlyRate       *float32 `json:"hourly_rate,omitempty"`
	Budget           *float32 `json:"budget,omitempty"`
}

type PatchUserAssignmentRequest struct {
	IsActive         *bool    `json:"is_active,omitempty"`
	IsProjectManager *bool    `json:"is_project_manager,omitempty"`
	UseDefaultRates  *bool    `json:"use_default_rates,omitempty"`
	HourlyRate       *float32 `json:"hourly_rate,omitempty"`
	Budget           *float32 `json:"budget,omitempty"`
}

type CreateTaskAssignmentRequest struct {
	TaskId     uint     `json:"task_id"`
	IsActive   *bool    `json:"is_active,omitempty"`
	Billable   *bool    `json:"billable,omitempty"`
	HourlyRate *float32 `json:"hourly_rate,omitempty"`
	Budget     *float32 `json:"budget,omitempty"`
}

type PatchTaskAssignmentRequest struct {
	IsActive   *bool    `json:"is_active,omitempty"`
	Billable   *bool    `json:"billable,omitempty"`
	HourlyRate *float32 `json:"hourly_rate,omitempty"`
	Budget     *float32 `json:"budget,omitempty"`
}

// Encapsulates the Harvest API methods under /projects
type ProjectsApi struct {
	baseUrl string
	client  *internalClient
}

func newProjectsV2(client *internalClient) ProjectsApi {
	return ProjectsApi{
		baseUrl: baseProjectsV2Url,
		client:  client,
	}
}

// Retrieves the a list of Projects.
func (api ProjectsApi) GetAll(params HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.DoGetV2(api.baseUrl, params)
}

// Retrieves a Project with the given ProjectID.
func (api ProjectsApi) Get(projectId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.baseUrl, projectId))
}

func (api ProjectsApi) Create(req CreateProjectRequest) (HarvestResponse, error) {
	return api.client.DoPost(api.baseUrl, req)
}

func (api ProjectsApi) Update(projectId uint, req ProjectPatchRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.baseUrl, projectId), req)
}

func (api ProjectsApi) Delete(projectId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d", api.baseUrl, projectId))
}

func (api ProjectsApi) GetAllUserAssigments(params HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.DoGetV2("v2/user_assignments", params)
}

func (api ProjectsApi) GetAllUserAssigmentsForProject(projectId uint, params HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.DoGetV2(fmt.Sprintf("%s/%d/user_assignments", api.baseUrl, projectId), params)
}

func (api ProjectsApi) GetUserAssigment(projectId, userAssignmentId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d/user_assignments/%d", api.baseUrl, projectId, userAssignmentId))
}

func (api ProjectsApi) CreateUserAssignment(projectId uint, req CreateUserAssignmentRequest) (HarvestResponse, error) {
	return api.client.DoPost(fmt.Sprintf("%s/%d/user_assignments", api.baseUrl, projectId), req)
}

func (api ProjectsApi) UpdateUserAssignment(projectId, userAssignmentId uint, req PatchUserAssignmentRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d/user_assignments/%d", api.baseUrl, projectId, userAssignmentId), req)
}

func (api ProjectsApi) DeleteUserAssignment(projectId, userAssignmentId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d/user_assignments/%d", api.baseUrl, projectId, userAssignmentId))
}

func (api ProjectsApi) GetAllTaskAssigments(params HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.DoGetV2("v2/task_assignments", params)
}

func (api ProjectsApi) GetAllTaskAssigmentsForProject(projectId uint, params HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.DoGetV2(fmt.Sprintf("%s/%d/task_assignments", api.baseUrl, projectId), params)
}

func (api ProjectsApi) GetTaskAssigment(projectId, taskAssignmentId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d/task_assignments/%d", api.baseUrl, projectId, taskAssignmentId))
}

func (api ProjectsApi) CreateTaskAssignment(projectId uint, req CreateTaskAssignmentRequest) (HarvestResponse, error) {
	return api.client.DoPost(fmt.Sprintf("%s/%d/task_assignments", api.baseUrl, projectId), req)
}

func (api ProjectsApi) UpdateTaskAssignment(projectId, userAssignmentId uint, req PatchTaskAssignmentRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d/task_assignments/%d", api.baseUrl, projectId, userAssignmentId), req)
}

func (api ProjectsApi) DeleteTaskAssignment(projectId, taskAssignmentId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d/task_assignments/%d", api.baseUrl, projectId, taskAssignmentId))
}
