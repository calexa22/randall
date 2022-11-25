package randall

const baseCompanyV2Url = "v2/company"

// Encapsulates the Harvest API methods under /company
type CompanyApi struct {
	baseUrl string
	client  *internalClient
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
