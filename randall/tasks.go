package randall

import (
	"fmt"
	"strconv"
)

const baseTasksV2Url = "v2/tasks"

// Encapsulates the Harvest API methods under /projects
type TasksApi struct {
	baseUrl string
	client  *internalClient
}

func newTasksV2(client *internalClient) TasksApi {
	return TasksApi{
		baseUrl: baseTasksV2Url,
		client:  client,
	}
}

// Retrieves the a list of Tasks.
func (api *TasksApi) Tasks(isActive *bool, page, perPage *int) (HarvestResponse, error) {
	query := make(map[string]string)

	if isActive != nil {
		query["is_active"] = strconv.FormatBool(*isActive)
	}

	if page != nil {
		query["page"] = strconv.Itoa(max(*page, 1))
	}

	if perPage != nil {
		query["per_page"] = strconv.Itoa(min(max(*perPage, 1), 2000))
	}

	return api.client.DoGet(api.baseUrl, query)
}

// Retrieves a Task with the given TaskID.
func (api *TasksApi) Task(taskId int) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.baseUrl, taskId))
}
