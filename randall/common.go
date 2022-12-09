package randall

import (
	"time"
)

type HarvestCollectionParams struct {
	Page         *int      `url:"page,omitempty"`
	PerPage      *int      `url:"per_page,omitempty"`
	UserId       *uint     `url:"user_id,omitempty"`
	ClientId     *uint     `url:"client_id,omitempty"`
	ProjectId    *uint     `url:"project_id,omitempty"`
	IsActive     *bool     `url:"is_active,omitempty"`
	IsBilled     *bool     `url:"is_billed,omitempty"`
	UpdatedSince time.Time `url:"updated_since,omitempty"`
	State        string    `url:"state,omitempty"`
	From         time.Time `url:"from,omitempty" layout:"2006-01-02"`
	To           time.Time `url:"to,omitempty" layout:"2006-01-02"`
}

// The general response object for any response payload sent by the Harvest API.
type HarvestResponse struct {
	// The HTTP status code of the response from Harvest.
	StatusCode int
	// The JSON payload of the response from Harvest.
	Data map[string]interface{}
}

type MessageRecipient struct {
	Email string  `json:"email"`
	Name  *string `json:"name,omitempty"`
}

type updateEventTypeRequest struct {
	EventType string `json:"event_type"`
}

type upsertItemCategoryRequest struct {
	Name string `json:"name"`
}

func getUpdateEventTypeRequest(eventType string) updateEventTypeRequest {
	return updateEventTypeRequest{
		EventType: eventType,
	}
}

func getOptionalCollectionParams(params []HarvestCollectionParams) *HarvestCollectionParams {
	if len(params) > 0 {
		return &params[0]
	}

	return nil
}
