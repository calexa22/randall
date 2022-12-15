package randall

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const (
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
	ClientId                         uint             `json:"client_id"`
	Name                             string           `json:"name"`
	IsBillable                       bool             `json:"is_billable"`
	BillBy                           string           `json:"bill_by"`
	BudgetBy                         string           `json:"budget_by"`
	Code                             *string          `json:"code,omitempty"`
	IsActive                         *bool            `json:"is_active,omitempty"`
	IsFixedFee                       *bool            `json:"is_fixed_fee,omitempty"`
	HourlyRate                       *decimal.Decimal `json:"hourly_rate,omitempty"`
	Budget                           *decimal.Decimal `json:"budget,omitempty"`
	BudgetIsMonthly                  *bool            `json:"budget_is_monthly,omitempty"`
	NotifyWhenOverBudget             *bool            `json:"notify_when_over_budget,omitempty"`
	OverBudgetNotificationPercentage *decimal.Decimal `json:"over_budget_notification_percentage,omitempty"`
	ShowBudgetToAll                  *bool            `json:"show_budget_to_all,omitempty"`
	CostBudget                       *decimal.Decimal `json:"cost_budget,omitempty"`
	CostBudgetIncludeExpenses        *bool            `json:"cost_budget_include_expenses,omitempty"`
	Fee                              *decimal.Decimal `json:"fee,omitempty"`
	Notes                            *string          `json:"notes,omitempty"`
	StartsOn                         time.Time        `json:"starts_on,omitempty" layout:"2006-01-02"`
	EndsOn                           time.Time        `json:"ends_on,omitempty" layout:"2006-01-02"`
}

type UpdateProjectRequest struct {
	ClientId                         *uint            `json:"client_id,omitempty"`
	Name                             *string          `json:"name,omitempty"`
	IsBillable                       *bool            `json:"is_billable,omitempty"`
	BillBy                           *string          `json:"bill_by,omitempty"`
	BudgetBy                         *string          `json:"budget_by,omitempty"`
	Code                             *string          `json:"code,omitempty"`
	IsActive                         *bool            `json:"is_active,omitempty"`
	IsFixedFee                       *bool            `json:"is_fixed_fee,omitempty"`
	HourlyRate                       *decimal.Decimal `json:"hourly_rate,omitempty"`
	Budget                           *decimal.Decimal `json:"budget,omitempty"`
	BudgetIsMonthly                  *bool            `json:"budget_is_monthly,omitempty"`
	NotifyWhenOverBudget             *bool            `json:"notify_when_over_budget,omitempty"`
	OverBudgetNotificationPercentage *decimal.Decimal `json:"over_budget_notification_percentage,omitempty"`
	ShowBudgetToAll                  *bool            `json:"show_budget_to_all,omitempty"`
	CostBudget                       *decimal.Decimal `json:"cost_budget,omitempty"`
	CostBudgetIncludeExpenses        *bool            `json:"cost_budget_include_expenses,omitempty"`
	Fee                              *decimal.Decimal `json:"fee,omitempty"`
	Notes                            *string          `json:"notes,omitempty"`
	StartsOn                         time.Time        `json:"starts_on,omitempty" layout:"2006-01-02"`
	EndsOn                           time.Time        `json:"ends_on,omitempty" layout:"2006-01-02"`
}

type CreateUserAssignmentRequest struct {
	UserId           uint             `json:"user_id"`
	Name             string           `json:"name"`
	IsActive         *bool            `json:"is_active,omitempty"`
	IsProjectManager *bool            `json:"is_project_manager,omitempty"`
	UseDefaultRates  *bool            `json:"use_default_rates,omitempty"`
	HourlyRate       *decimal.Decimal `json:"hourly_rate,omitempty"`
	Budget           *decimal.Decimal `json:"budget,omitempty"`
}

type PatchUserAssignmentRequest struct {
	IsActive         *bool            `json:"is_active,omitempty"`
	IsProjectManager *bool            `json:"is_project_manager,omitempty"`
	UseDefaultRates  *bool            `json:"use_default_rates,omitempty"`
	HourlyRate       *decimal.Decimal `json:"hourly_rate,omitempty"`
	Budget           *decimal.Decimal `json:"budget,omitempty"`
}

type CreateTaskAssignmentRequest struct {
	TaskId     uint             `json:"task_id"`
	IsActive   *bool            `json:"is_active,omitempty"`
	Billable   *bool            `json:"billable,omitempty"`
	HourlyRate *decimal.Decimal `json:"hourly_rate,omitempty"`
	Budget     *decimal.Decimal `json:"budget,omitempty"`
}

type PatchTaskAssignmentRequest struct {
	IsActive   *bool            `json:"is_active,omitempty"`
	Billable   *bool            `json:"billable,omitempty"`
	HourlyRate *decimal.Decimal `json:"hourly_rate,omitempty"`
	Budget     *decimal.Decimal `json:"budget,omitempty"`
}

// Encapsulates the Harvest API methods under /projects
type ProjectsApi struct {
	baseUrl string
	client  *internalClient
}

func newProjectsV2(client *internalClient) ProjectsApi {
	return ProjectsApi{
		baseUrl: "v2/projects",
		client:  client,
	}
}

// Retrieves the a list of Projects.
func (api ProjectsApi) GetAll(params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(api.baseUrl, getOptionalCollectionParams(params))
}

// Retrieves a Project with the given ProjectID.
func (api ProjectsApi) Get(projectId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d", api.baseUrl, projectId))
}

func (api ProjectsApi) Create(req CreateProjectRequest) (HarvestResponse, error) {
	return api.client.doPost(api.baseUrl, req)
}

func (api ProjectsApi) Update(projectId uint, req UpdateProjectRequest) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d", api.baseUrl, projectId), req)
}

func (api ProjectsApi) Delete(projectId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d", api.baseUrl, projectId))
}

func (api ProjectsApi) GetAllUserAssigments(params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet("v2/user_assignments", getOptionalCollectionParams(params))
}

func (api ProjectsApi) GetAllUserAssigmentsForProject(projectId uint, params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d/user_assignments", api.baseUrl, projectId), getOptionalCollectionParams(params))
}

func (api ProjectsApi) GetUserAssigment(projectId, userAssignmentId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d/user_assignments/%d", api.baseUrl, projectId, userAssignmentId))
}

func (api ProjectsApi) CreateUserAssignment(projectId uint, req CreateUserAssignmentRequest) (HarvestResponse, error) {
	return api.client.doPost(fmt.Sprintf("%s/%d/user_assignments", api.baseUrl, projectId), req)
}

func (api ProjectsApi) UpdateUserAssignment(projectId, userAssignmentId uint, req PatchUserAssignmentRequest) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d/user_assignments/%d", api.baseUrl, projectId, userAssignmentId), req)
}

func (api ProjectsApi) DeleteUserAssignment(projectId, userAssignmentId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d/user_assignments/%d", api.baseUrl, projectId, userAssignmentId))
}

func (api ProjectsApi) GetAllTaskAssigments(params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet("v2/task_assignments", getOptionalCollectionParams(params))
}

func (api ProjectsApi) GetAllTaskAssigmentsForProject(projectId uint, params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d/task_assignments", api.baseUrl, projectId), getOptionalCollectionParams(params))
}

func (api ProjectsApi) GetTaskAssigment(projectId, taskAssignmentId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d/task_assignments/%d", api.baseUrl, projectId, taskAssignmentId))
}

func (api ProjectsApi) CreateTaskAssignment(projectId uint, req CreateTaskAssignmentRequest) (HarvestResponse, error) {
	return api.client.doPost(fmt.Sprintf("%s/%d/task_assignments", api.baseUrl, projectId), req)
}

func (api ProjectsApi) UpdateTaskAssignment(projectId, userAssignmentId uint, req PatchTaskAssignmentRequest) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d/task_assignments/%d", api.baseUrl, projectId, userAssignmentId), req)
}

func (api ProjectsApi) DeleteTaskAssignment(projectId, taskAssignmentId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d/task_assignments/%d", api.baseUrl, projectId, taskAssignmentId))
}
