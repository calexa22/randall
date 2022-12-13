package randall

// Encapsulates the Harvest API methods under /company
type CompanyApi struct {
	baseUrl string
	client  *internalClient
}

type UpdateCompanyRequest struct {
	WeeklyCapacity       *bool `json:"weekly_capacity,omitempty"`
	WantsTimestampTimers *bool `json:"wants_timestamp_timers,omitempty"`
}

func newCompanyV2(client *internalClient) CompanyApi {
	return CompanyApi{
		baseUrl: "v2/company",
		client:  client,
	}
}

// Retrieves the Company of currently authenticated user. Returns a company object and a 200 OK response code.
func (api CompanyApi) MyCompany() (HarvestResponse, error) {
	return api.client.doGet(api.baseUrl)
}

func (api CompanyApi) Update(req UpdateCompanyRequest) (HarvestResponse, error) {
	return api.client.doPatch(api.baseUrl, req)
}
