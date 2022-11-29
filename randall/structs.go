package randall

import (
	"encoding/json"
	"errors"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

type HarvestCollectionParams struct {
	Page         *int       `url:"page,omitempty"`
	PerPage      *int       `url:"per_page,omitempty"`
	UserId       *uint      `url:"user_id,omitempty"`
	ClientId     *uint      `url:"client_id,omitempty"`
	ProjectId    *uint      `url:"project_id,omitempty"`
	IsActive     *bool      `url:"is_active,omitempty"`
	IsBilled     *bool      `url:"is_billed,omitempty"`
	UpdatedSince *time.Time `url:"updated_since,omitempty"`
	/// from, to (*HarvestDate)
}

func (p HarvestCollectionParams) AddQuery(v url.Values) (url.Values, error) {
	if p.Page != nil && *p.Page < 1 {
		return v, errors.New("\"page\" query param must be greater than 0")
	}

	if p.PerPage != nil && (1 > *p.PerPage || *p.PerPage > 2000) {
		return v, errors.New("\"per_page\" query param must be between 1 and 2000")
	}

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

// The general response object for any response payload sent by the Harvest API.
type HarvestResponse struct {
	// The HTTP status code of the response from Harvest.
	StatusCode int
	// The JSON payload of the response from Harvest.
	Data map[string]interface{}
}

type HarvestDate time.Time

func (d HarvestDate) MarshalJSON() ([]byte, error) {
	s := time.Time(d).Format("2006-01-02")

	return json.Marshal(s)
}

func (d *HarvestDate) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse("2006-01-02", string(b))
	if err != nil {
		return err
	}

	*d = HarvestDate(date.UTC())
	return nil
}
