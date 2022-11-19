package randall

type HeaderValues struct {
	AccountId      string
	AccessToken    string
	UserAgentApp   string
	UserAgentEmail string
}

type JsonResponse struct {
	StatusCode int
	Data       map[string]any
}
