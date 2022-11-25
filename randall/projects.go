package randall

import (
	"fmt"
	"strconv"
)

const baseProjectsV2Url = "v2/projects"

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
func (api ProjectsApi) Projects(isActive *bool, clientId, page, perPage *int) (HarvestResponse, error) {
	query := make(map[string]string)

	if isActive != nil {
		query["is_active"] = strconv.FormatBool(*isActive)
	}

	if clientId != nil {
		query["client_id"] = strconv.Itoa(*clientId)
	}

	if page != nil {
		query["page"] = strconv.Itoa(max(*page, 1))
	}

	if perPage != nil {
		query["per_page"] = strconv.Itoa(min(max(*perPage, 1), 2000))
	}

	return api.client.DoGet(api.baseUrl, query)
}

// Retrieves a Project with the given ProjectID.
func (api ProjectsApi) Project(projectId int) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.baseUrl, projectId))
}
