package randall

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/shopspring/decimal"
)

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
	ProjectId   uint               `json:"project_id"`
	TaskId      uint               `json:"task_id"`
	SpentDate   time.Time          `json:"spent_date" layout:"2006-01-02"`
	UserId      *uint              `json:"user_id,omitempty"`
	Hours       *decimal.Decimal   `json:"hours,omitempty"`
	Notes       *string            `json:"notes,omitempty"`
	ExternalRef *ExternalReference `json:"external_reference,omitempty"`
}

type CreateTimeEntryViaStartEndRequest struct {
	ProjectId   uint               `json:"project_id"`
	TaskId      uint               `json:"task_id"`
	SpentDate   time.Time          `json:"spent_date" layout:"2006-01-02"`
	UserId      *uint              `json:"user_id,omitempty"`
	StartedTime *string            `json:"started_time,omitempty"`
	EndTime     *string            `json:"end_time,omitempty"`
	Notes       *string            `json:"notes,omitempty"`
	ExternalRef *ExternalReference `json:"external_reference,omitempty"`
}

type UpdateTimeEntryRequest struct {
	ProjectId   *uint              `json:"project_id,omitempty"`
	TaskId      *uint              `json:"task_id,omitempty"`
	SpentDate   time.Time          `json:"spent_date,omitempty" layout:"2006-01-02"`
	StartedTime *string            `json:"started_time,omitempty"`
	EndTime     *string            `json:"end_time,omitempty"`
	Hours       *decimal.Decimal   `json:"hours,omitempty"`
	Notes       *string            `json:"notes,omitempty"`
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
		baseUrl: "v2/time_entries",
		client:  client,
	}
}

// Retrieves the time entries accessible to th currently authenticated user.
// Returns a company object and a 200 OK response code.
func (api TimeEntriesApi) GetAll(params ...GetTimeEntriesParams) (HarvestResponse, error) {
	var param *GetTimeEntriesParams

	if len(params) > 0 {
		param = &params[0]
	}
	return api.client.doGet(api.baseUrl, param)
}

func (api TimeEntriesApi) GetTimeEntry(timeEntryId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d", api.baseUrl, timeEntryId))
}

func (api TimeEntriesApi) CreateViaDuration(req CreateTimeEntryViaDurationRequest) (HarvestResponse, error) {
	return api.client.doPost(api.baseUrl, req)
}

func (api TimeEntriesApi) CreateViaStartEnd(req CreateTimeEntryViaStartEndRequest) (HarvestResponse, error) {
	return api.client.doPost(api.baseUrl, req)
}

func (api TimeEntriesApi) UpdateTimeEntry(timeEntryId uint, req UpdateTimeEntryRequest) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d", api.baseUrl, timeEntryId), req)
}

func (api TimeEntriesApi) DeleteTimeEntry(timeEntryId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d", api.baseUrl, timeEntryId))
}

func (api TimeEntriesApi) DeleteExternalReference(timeEntryId int) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d/external_reference", api.baseUrl, timeEntryId))
}

func (api TimeEntriesApi) RestartTimeEntry(timeEntryId uint) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d/restart", api.baseUrl, timeEntryId))
}

func (api TimeEntriesApi) StopTimeEntry(timeEntryId uint) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d/stop", api.baseUrl, timeEntryId))
}
