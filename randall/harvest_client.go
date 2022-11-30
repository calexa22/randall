package randall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/shopspring/decimal"
)

// The interface used to interact with the entire Harvest API.
type HarvestClient struct {
	// The interface used to make calls to /company endpoints under the given Client.
	// The interface used to make calls to /users endpoints under the given Client.

	Clients           ClientsApi
	Company           CompanyApi
	Contacts          ContactsApi
	ExpenseCategories ExpenseCategoriesApi
	Expenses          ExpensesApi
	Projects          ProjectsApi
	Roles             RolesApi
	Tasks             TasksApi
	TimeEntries       TimeEntriesApi
	Users             UsersApi
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
func NewClient(accountId, accessToken, userAgentApp, userAgentEmail string) *HarvestClient {
	// todo
	decimal.MarshalJSONWithoutQuotes = true

	internal := &internalClient{
		httpClient:     &http.Client{},
		baseUrl:        "https://api.harvestapp.com",
		accountId:      accountId,
		accessToken:    accessToken,
		userAgentApp:   userAgentApp,
		userAgentEmail: userAgentApp,
	}

	return &HarvestClient{
		Clients:           newClientsV2(internal),
		Company:           newCompanyV2(internal),
		Contacts:          newContactsV2(internal),
		ExpenseCategories: newExpenseCategoriesV2(internal),
		Expenses:          newExpensesV2(internal),
		Projects:          newProjectsV2(internal),
		Roles:             newRolesV2(internal),
		Tasks:             newTasksV2(internal),
		TimeEntries:       newTimeEntriesV2(internal),
		Users:             newUsersV2(internal),
	}
}

func (client *internalClient) DoGet(url string, query ...url.Values) (HarvestResponse, error) {
	r, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s", client.baseUrl, url), nil)

	if len(query) > 0 && query[0] != nil {
		r.URL.RawQuery = query[0].Encode()
	}

	client.SetHeaders(r, false)

	return client.readResponse(r)
}

func (client *internalClient) DoGetV2(url string, params HarvestCollectionParams) (HarvestResponse, error) {
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", client.baseUrl, url), nil)

	if err != nil {
		return HarvestResponse{}, nil
	}

	values, err := query.Values(params)

	if err != nil {
		return HarvestResponse{}, nil
	}

	r.URL.RawQuery = values.Encode()

	client.SetHeaders(r, false)

	return client.readResponse(r)
}

func (client *internalClient) DoPost(url string, body ...interface{}) (HarvestResponse, error) {
	b, err := client.getBody(body)

	if err != nil {
		return HarvestResponse{}, err
	}

	r, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", client.baseUrl, url), b)

	if err != nil {
		return HarvestResponse{}, err
	}

	client.SetHeaders(r, true)

	return client.readResponse(r)
}

func (client *internalClient) DoPatch(url string, body ...interface{}) (HarvestResponse, error) {
	b, err := client.getBody(body)

	if err != nil {
		return HarvestResponse{}, err
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

func (client *internalClient) getBody(body []interface{}) (*bytes.Buffer, error) {
	var b *bytes.Buffer

	if len(body) > 1 && body[0] != nil {
		buff, err := json.Marshal(body[0])

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(buff)
	}

	return b, nil
}

func (client *internalClient) readResponse(req *http.Request) (HarvestResponse, error) {
	resp, err := client.httpClient.Do(req)

	if err != nil {
		return HarvestResponse{}, err
	}

	defer resp.Body.Close()

	var data map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		return HarvestResponse{}, err
	}

	return HarvestResponse{
		StatusCode: resp.StatusCode,
		Data:       data,
	}, nil
}
