package randall

const baseCompanyV2Url = "v2/company"

// Encapsulates the Harvest API methods under /company
type CompanyApi struct {
	baseUrl string
	client  *internalClient
}

type CompanyPatchRequest struct {
	WeeklyCapacity       *bool `json:"weekly_capacity,omitempty"`
	WantsTimestampTimers *bool `json:"wants_timestamp_timers,omitempty"`
}

func newCompanyV2(client *internalClient) CompanyApi {
	return CompanyApi{
		baseUrl: baseCompanyV2Url,
		client:  client,
	}
}

// Retrieves the Company of currently authenticated user. Returns a company object and a 200 OK response code.
func (api CompanyApi) MyCompany() (HarvestResponse, error) {
	return api.client.DoGet(api.baseUrl)
}

func (api CompanyApi) Update(req CompanyPatchRequest) (HarvestResponse, error) {
	return api.client.DoPatch(api.baseUrl, req)
}
