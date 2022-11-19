package randall

// Collection of header values required to be sent on every request to Harvest.
type HarvestHeaders struct {
	// A Harvest account identifier, required for header Harvest-Account-ID.
	// See https://help.getharvest.com/api-v2/introduction/overview/general/
	AccountId string
	// A Harvest personal access token, required for Authentication.
	// See https://help.getharvest.com/api-v2/authentication-api/authentication/authentication/
	AccessToken string
	// The name of the app calling the Harvest API, required for the header User-Agent.
	// Seehttps://help.getharvest.com/api-v2/introduction/overview/general/
	UserAgentApp string
	// An email associated with the app/developer calling the Harvest API, required for the header User-Agent.
	// See https://help.getharvest.com/api-v2/introduction/overview/general/
	UserAgentEmail string
}

type HarvestResponse struct {
	// The Http status code of the response from Harvest.
	StatusCode int
	// The JSON payload of the response from Harvest.
	Data map[string]any
}
