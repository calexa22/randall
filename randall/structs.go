package randall

import (
	"encoding/json"
	"time"
)

// The general response object for any response from the Harvest API.
type HarvestResponse struct {
	// The HTTP status code of the response from Harvest.
	StatusCode int
	// The JSON payload of the response from Harvest.
	Data map[string]any
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
	return
}

type HarvestTime time.Time

func (t HarvestTime) MarshalJSON() ([]byte, error) {
	s := time.Time(t).Format("2006-01-02")

	return json.Marshal(s)
}

func (t *HarvestTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse("2006-01-02", string(b))
	if err != nil {
		return err
	}

	*t = HarvestTime(date.UTC())
	return
}
