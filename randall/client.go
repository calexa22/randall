package randall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/shopspring/decimal"
)

const domain = "https://api.harvestapp.com"

var client *http.Client
var mtx sync.Mutex

// The interface used to interact with the entire Harvest API.
type Client struct {
	// The interface used to make calls to /company endpoints under the given Client.
	// The interface used to make calls to /users endpoints under the given Client.

	Company     CompanyApi
	Users       UsersApi
	Projects    ProjectsApi
	Tasks       TasksApi
	TimeEntries TimeEntriesApi
}

type internalClient struct {
	httpClient     *http.Client
	baseUrl        string
	accountId      string
	accessToken    string
	userAgentApp   string
	userAgentEmail string
}

// Initializes a new instance of Client. Requests through the Client will have the headers
// required by the Harvest API with the passed in values
func New(accountId, accessToken, userAgentApp, userAgentEmail string) *Client {

	decimal.MarshalJSONWithoutQuotes = true
	internal := &internalClient{
		httpClient:     getHttpClient(),
		baseUrl:        domain,
		accountId:      accountId,
		accessToken:    accessToken,
		userAgentApp:   userAgentApp,
		userAgentEmail: userAgentApp,
	}

	return &Client{
		Company:     newCompanyV2(internal),
		Users:       newUsersV2(internal),
		Projects:    newProjectsV2(internal),
		Tasks:       newTasksV2(internal),
		TimeEntries: newTimeEntriesV2(internal),
	}
}

func (client *internalClient) DoGet(url string, queryStr ...map[string]string) (HarvestResponse, error) {
	r, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s", client.baseUrl, url), nil)

	if len(queryStr) > 0 {
		client.SetQuery(r, queryStr[0])
	}

	client.SetHeaders(r, false)

	return client.readResponse(r)
}

func (client *internalClient) DoPost(url string, body interface{}) (HarvestResponse, error) {
	buff, err := json.Marshal(body)

	if err != nil {
		return HarvestResponse{}, err
	}

	r, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", client.baseUrl, url), bytes.NewBuffer(buff))

	if err != nil {
		return HarvestResponse{}, err
	}

	client.SetHeaders(r, true)

	return client.readResponse(r)
}

func (client *internalClient) DoPatch(url string, body ...interface{}) (HarvestResponse, error) {

	var b *bytes.Buffer

	if len(body) > 1 && body[0] != nil {
		buff, err := json.Marshal(body[0])

		if err != nil {
			return HarvestResponse{}, err
		}

		b = bytes.NewBuffer(buff)
	}

	r, err := http.NewRequest("PATCH", fmt.Sprintf("%s/%s", client.baseUrl, url), b)

	if err != nil {
		return HarvestResponse{}, err
	}

	client.SetHeaders(r, false)

	return client.readResponse(r)
}

func (client *internalClient) DoDelete(url string) (HarvestResponse, error) {
	r, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", client.baseUrl, url), nil)

	if err != nil {
		return HarvestResponse{}, err
	}

	client.SetHeaders(r, false)

	return client.readResponse(r)
}

func (client *internalClient) SetQuery(r *http.Request, queryStr map[string]string) {
	query := r.URL.Query()
	for key, value := range queryStr {
		query.Set(key, value)
	}

	r.URL.RawQuery = query.Encode()
}

func (client *internalClient) SetHeaders(r *http.Request, includeContentType bool) {
	r.Header.Set("User-Agent", fmt.Sprintf("%s (%s)", client.userAgentApp, client.userAgentEmail))
	r.Header.Set("Harvest-Account-ID", client.accountId)
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.accessToken))

	if includeContentType {
		r.Header.Set("Content-Type", "application/json")
	}
}

func (client *internalClient) readResponse(req *http.Request) (HarvestResponse, error) {
	resp, err := client.httpClient.Do(req)

	if err != nil {
		return HarvestResponse{}, err
	}

	defer resp.Body.Close()

	var data map[string]any

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		return HarvestResponse{}, err
	}

	return HarvestResponse{
		StatusCode: resp.StatusCode,
		Data:       data,
	}, nil
}

func getHttpClient() *http.Client {
	mtx.Lock()
	defer mtx.Unlock()

	if client == nil {
		client = &http.Client{}
	}
	return client
}
