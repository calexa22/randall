package randall

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

const baseTimeEntriesV2Url = "v2/time_entries"

// Encapsulates the Harvest API methods under /time_entries
type TimeEntriesApi struct {
	baseUrl string
	client  *internalClient
}

type GetTimeEntriesParams struct {
	UserId              int        `url:"user_id,omitempty"`
	ClientId            int        `url:"client_id,omitempty"`
	ProjectId           int        `url:"project_id,omitempty"`
	TaskId              int        `url:"task_id,omitempty"`
	ExternalReferenceId int        `url:"external_reference_id,omitempty"`
	IsBilled            *bool      `url:"is_billed,omitempty"`
	IsRunning           *bool      `url:"is_running,omitempty"`
	UpdateSince         *time.Time `url:"updated_since,omitempty"`
	FromDate            *time.Time `url:"from,omitempty"`
	StartDate           *time.Time `url:"start,omitempty"`
}

func (p GetTimeEntriesParams) AddQuery(v url.Values) (url.Values, error) {
	newV, err := query.Values(p)

	if err != nil {
		return v, err
	}

	for key, values := range newV {
		for _, value := range values {
			v.Add(key, value)
		}
	}

	return v, nil
}

type CreateTimeEntryViaDurationRequest struct {
	UserId      uint               `json:"user_id,omitempty"`
	ProjectId   uint               `json:"project_id"`
	TaskId      uint               `json:"task_id"`
	SpentDate   HarvestDate        `json:"spent_date"`
	Hours       float32            `json:"hours,omitempty"`
	Notes       string             `json:"notes,omitempty"`
	ExternalRef *ExternalReference `json:"external_reference,omitempty"`
}

type CreateTimeEntryViaStartEndRequest struct {
	UserId      uint               `json:"user_id,omitempty"`
	ProjectId   uint               `json:"project_id"`
	TaskId      uint               `json:"task_id"`
	SpentDate   HarvestDate        `json:"spent_date"`
	StartedTime string             `json:"started_time,omitempty"`
	EndTime     string             `json:"end_time,omitempty"`
	Notes       string             `json:"notes,omitempty"`
	ExternalRef *ExternalReference `json:"external_reference,omitempty"`
}

type TimeEntryPatchRequest struct {
	ProjectId   uint               `json:"project_id,omitempty"`
	TaskId      uint               `json:"task_id,omitempty"`
	SpentDate   *HarvestDate       `json:"spent_date,omitempty"`
	StartedTime string             `json:"started_time,omitempty"`
	EndTime     string             `json:"end_time,omitempty"`
	Hours       float32            `json:"hours,omitempty"`
	Notes       string             `json:"notes,omitempty"`
	ExternalRef *ExternalReference `json:"external_reference,omitempty"`
}

type ExternalReference struct {
	Id        uint   `json:"id"`
	GroupId   uint   `json:"group_id"`
	AccountId uint   `json:"account_id"`
	Permalink string `json:"permalink"`
}

func newTimeEntriesV2(client *internalClient) TimeEntriesApi {
	return TimeEntriesApi{
		baseUrl: baseTimeEntriesV2Url,
		client:  client,
	}
}

// Retrieves the time entries accessible to th currently authenticated user.
// Returns a company object and a 200 OK response code.
func (api TimeEntriesApi) GetAll(params ...QueryStringProvider) (HarvestResponse, error) {
	query, err := getUrlValues(params...)

	if err != nil {
		return HarvestResponse{}, err
	}

	return api.client.DoGet(api.baseUrl, query)
}

func (api TimeEntriesApi) GetTimeEntry(timeEntryId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.baseUrl, timeEntryId))
}

func (api TimeEntriesApi) CreateViaDuration(req CreateTimeEntryViaDurationRequest) (HarvestResponse, error) {
	return api.client.DoPost(api.baseUrl, req)
}

func (api TimeEntriesApi) CreateViaStartEnd(req CreateTimeEntryViaStartEndRequest) (HarvestResponse, error) {
	return api.client.DoPost(api.baseUrl, req)
}

func (api TimeEntriesApi) UpdateTimeEntry(timeEntryId uint, req TimeEntryPatchRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.baseUrl, timeEntryId), req)
}

func (api TimeEntriesApi) DeleteTimeEntry(timeEntryId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d", api.baseUrl, timeEntryId))
}

func (api TimeEntriesApi) DeleteExternalReference(timeEntryId int) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d/external_reference", api.baseUrl, timeEntryId))
}

func (api TimeEntriesApi) RestartTimeEntry(timeEntryId uint) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d/restart", api.baseUrl, timeEntryId))
}

func (api TimeEntriesApi) StopTimeEntry(timeEntryId uint) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d/stop", api.baseUrl, timeEntryId))
}

// func TimeEntriesQuery(p GetTimeEntriesParams) map[string]string {
// 	query := make(map[string]string)

// 	if p.UserId != nil {
// 		query["user_id"] = strconv.Itoa(*p.UserId)
// 	}

// 	if p.ClientId != nil {
// 		query["client_id"] = strconv.Itoa(*p.ClientId)
// 	}

// 	if p.ProjectId != nil {
// 		query["project_id"] = strconv.Itoa(*p.ProjectId)
// 	}

// 	if p.TaskId != nil {
// 		query["task_id"] = strconv.Itoa(*p.TaskId)
// 	}

// 	if p.ExternalReferenceId != nil {
// 		query["external_reference_id"] = strconv.Itoa(*p.ExternalReferenceId)
// 	}

// 	if p.IsBilled != nil {
// 		query["is_billed"] = strconv.FormatBool(*p.IsBilled)
// 	}

// 	if p.IsRunning != nil {
// 		query["is_running"] = strconv.FormatBool(*p.IsRunning)
// 	}

// 	if p.UpdateSince != nil {
// 		query["updated_since"] = (*p.UpdateSince).Format(time.RFC3339)
// 	}

// 	if p.FromDate != nil {
// 		query["from"] = (*p.FromDate).Format("2006-01-02")
// 	}

// 	if p.StartDate != nil {
// 		query["start"] = (*p.FromDate).Format("2006-01-02")
// 	}

// 	if p.Page != nil {
// 		query["page"] = strconv.Itoa(*p.Page)
// 	}

// 	if p.PerPage != nil {
// 		query["per_page"] = strconv.Itoa(*p.PerPage)
// 	}

// 	return query
// }