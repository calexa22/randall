package randall

import (
	"fmt"
	"strconv"
	"time"
)

const baseTimeEntriesV2Url = "v2/time_entries"

// Encapsulates the Harvest API methods under /time_entries
type TimeEntriesApi struct {
	baseUrl string
	client  *internalClient
}

type GetTimeEntriesParams struct {
	UserId              *int
	ClientId            *int
	ProjectId           *int
	TaskId              *int
	ExternalReferenceId *int
	IsBilled            *bool
	IsRunning           *bool
	UpdateSince         *time.Time
	FromDate            *time.Time
	StartDate           *time.Time
	Page                *int
	PerPage             *int
}

type TimeEntryViaDurationRequest struct {
	UserId    uint        `json:"user_id,omitempty"`
	ProjectId uint        `json:"project_id"`
	TaskId    uint        `json:"task_id"`
	SpentDate HarvestDate `json:"spent_date"`
	Hours     float32     `json:"hours,omitempty"`
	Notes     string      `json:"notes,omitempty"`
	//todo external_reference
}

type TimeEntry struct {
	Id                uint         `json:"id,omitempty"`
	UserId            uint         `json:"user_id"`
	SpentDate         HarvestDate  `json:"spent_date"`
	Hours             float32      `json:"hours"`
	HoursWithoutTimer float32      `json:"hours_without_timer"`
	RoundedHours      float32      `json:"rounded_hours,omitempty"`
	Notes             string       `json:"notes,omitempty"`
	IsLocked          *bool        `json:"is_locked,omitempty"`
	LockedReason      string       `json:"locked_reason,omitempty"`
	IsClosed          *bool        `json:"is_closed,omitempty"`
	IsBilled          *bool        `json:"is_billed,omitempty"`
	TimerStartedAt    *time.Time   `json:"timer_started_at,omitempty"`
	StartedTime       *HarvestTime `json:"started_time,omitempty"`
	EndedTime         *HarvestTime `json:"ended_time,omitempty"`
	IsRunning         *bool        `json:"is_running,omitempty"`
	Billable          *bool        `json:"billable,omitempty"`
	Budgeted          *bool        `json:"budgeted,omitempty"`
	BillableRate      *float32     `json:"billable_rate,omitempty"`
	CreatedAt         time.Time    `json:"created_at,omitempty"`
	UpdatedAt         *time.Time   `json:"updated_at,omitempty"`
}

func newTimeEntriesV2(client *internalClient) TimeEntriesApi {
	return TimeEntriesApi{
		baseUrl: baseTimeEntriesV2Url,
		client:  client,
	}
}

// Retrieves the time entries accessible to th currently authenticated user.
// Returns a company object and a 200 OK response code.
func (api TimeEntriesApi) GetAll(params GetTimeEntriesParams) (HarvestResponse, error) {
	return api.client.DoGet(api.baseUrl, TimeEntriesQuery(params))
}

func (api TimeEntriesApi) Get(timeEntryId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.baseUrl, timeEntryId))
}

func (api TimeEntriesApi) CreateViaDuration(req TimeEntryViaDurationRequest) (HarvestResponse, error) {
	return api.client.DoPost(api.baseUrl, req)
}

func (api TimeEntriesApi) Delete(timeEntryId uint) (HarvestResponse, error) {
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

func TimeEntriesQuery(p GetTimeEntriesParams) map[string]string {
	query := make(map[string]string)

	if p.UserId != nil {
		query["user_id"] = strconv.Itoa(*p.UserId)
	}

	if p.ClientId != nil {
		query["client_id"] = strconv.Itoa(*p.ClientId)
	}

	if p.ProjectId != nil {
		query["project_id"] = strconv.Itoa(*p.ProjectId)
	}

	if p.TaskId != nil {
		query["task_id"] = strconv.Itoa(*p.TaskId)
	}

	if p.ExternalReferenceId != nil {
		query["external_reference_id"] = strconv.Itoa(*p.ExternalReferenceId)
	}

	if p.IsBilled != nil {
		query["is_billed"] = strconv.FormatBool(*p.IsBilled)
	}

	if p.IsRunning != nil {
		query["is_running"] = strconv.FormatBool(*p.IsRunning)
	}

	if p.UpdateSince != nil {
		query["updated_since"] = (*p.UpdateSince).Format(time.RFC3339)
	}

	if p.FromDate != nil {
		query["from"] = (*p.FromDate).Format("2006-01-02")
	}

	if p.StartDate != nil {
		query["start"] = (*p.FromDate).Format("2006-01-02")
	}

	if p.Page != nil {
		query["page"] = strconv.Itoa(*p.Page)
	}

	if p.PerPage != nil {
		query["per_page"] = strconv.Itoa(*p.PerPage)
	}

	return query
}
