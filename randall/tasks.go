package randall

import (
	"fmt"
)

const baseTasksV2Url = "v2/tasks"

// Encapsulates the Harvest API methods under /projects
type TasksApi struct {
	baseUrl string
	client  *internalClient
}

type CreateTaskRequest struct {
	Name              string   `json:"name"`
	DefaultHourlyRate *float32 `json:"default_hourly_rate,omitempty"`
	IsDefault         *bool    `json:"is_default,omitempty"`
	IsActive          *bool    `json:"is_active,omitempty"`
	BillableByDefault *bool    `json:"billable_by_default,omitempty"`
}

type UpdateTaskPatchRequest struct {
	Name              string   `json:"name,omitempty"`
	DefaultHourlyRate *float32 `json:"default_hourly_rate,omitempty"`
	IsDefault         *bool    `json:"is_default,omitempty"`
	IsActive          *bool    `json:"is_active,omitempty"`
	BillableByDefault *bool    `json:"billable_by_default,omitempty"`
}

func newTasksV2(client *internalClient) TasksApi {
	return TasksApi{
		baseUrl: baseTasksV2Url,
		client:  client,
	}
}

// Retrieves the a list of Tasks.
func (api *TasksApi) GetAllTasks(params ...QueryStringProvider) (HarvestResponse, error) {
	query, err := getUrlValues(params...)

	if err != nil {
		return HarvestResponse{}, err
	}

	return api.client.DoGet(api.baseUrl, query)
}

// Retrieves a Task with the given TaskID.
func (api *TasksApi) GetTask(taskId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.baseUrl, taskId))
}

func (api *TasksApi) CreateTask(req CreateTaskRequest) (HarvestResponse, error) {
	return api.client.DoPost(api.baseUrl, req)
}

func (api *TasksApi) UpdateTask(taskId uint, req UpdateTaskPatchRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.baseUrl, taskId), req)
}

func (api *TasksApi) DeleteTask(taskId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d", api.baseUrl, taskId))
}
