package randall

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// Encapsulates the Harvest API methods under /projects
type TasksApi struct {
	baseUrl string
	client  *internalClient
}

type CreateTaskRequest struct {
	Name              string           `json:"name"`
	DefaultHourlyRate *decimal.Decimal `json:"default_hourly_rate,omitempty"`
	IsDefault         *bool            `json:"is_default,omitempty"`
	IsActive          *bool            `json:"is_active,omitempty"`
	BillableByDefault *bool            `json:"billable_by_default,omitempty"`
}

type UpdateTaskRequest struct {
	Name              *string          `json:"name,omitempty"`
	DefaultHourlyRate *decimal.Decimal `json:"default_hourly_rate,omitempty"`
	IsDefault         *bool            `json:"is_default,omitempty"`
	IsActive          *bool            `json:"is_active,omitempty"`
	BillableByDefault *bool            `json:"billable_by_default,omitempty"`
}

func newTasksV2(client *internalClient) TasksApi {
	return TasksApi{
		baseUrl: "v2/tasks",
		client:  client,
	}
}

// Retrieves the a list of Tasks.
func (api TasksApi) GetAllTasks(params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(api.baseUrl, getOptionalCollectionParams(params))
}

// Retrieves a Task with the given TaskID.
func (api TasksApi) GetTask(taskId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d", api.baseUrl, taskId))
}

func (api TasksApi) CreateTask(req CreateTaskRequest) (HarvestResponse, error) {
	return api.client.doPost(api.baseUrl, req)
}

func (api TasksApi) UpdateTask(taskId uint, req UpdateTaskRequest) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d", api.baseUrl, taskId), req)
}

func (api TasksApi) DeleteTask(taskId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d", api.baseUrl, taskId))
}
