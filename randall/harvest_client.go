package randall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/go-querystring/query"
)

// The interface used to interact with the entire Harvest API.
type HarvestClient struct {
	Clients     ClientsApi
	Company     CompanyApi
	Contacts    ContactsApi
	Estimates   EstimatesApi
	Expenses    ExpensesApi
	Invoices    InvoicesApi
	Projects    ProjectsApi
	Roles       RolesApi
	Tasks       TasksApi
	TimeEntries TimeEntriesApi
	Users       UsersApi
}

type internalClient struct {
	httpClient     *http.Client
	baseUrl        string
	accountId      string
	accessToken    string
	userAgentApp   string
	userAgentEmail string
}

type multipartData struct {
	data  map[string]string
	files map[string]string
}

// Initializes a new instance of Client. Requests through the Client will have the headers
// required by the Harvest API with the passed in values
func NewClient(accountId, accessToken, userAgentApp, userAgentEmail string) *HarvestClient {
	internal := &internalClient{
		httpClient:     &http.Client{},
		baseUrl:        "https://api.harvestapp.com",
		accountId:      accountId,
		accessToken:    accessToken,
		userAgentApp:   userAgentApp,
		userAgentEmail: userAgentApp,
	}

	return &HarvestClient{
		Clients:     newClientsV2(internal),
		Company:     newCompanyV2(internal),
		Contacts:    newContactsV2(internal),
		Estimates:   newEstimatesV2(internal),
		Expenses:    newExpensesV2(internal),
		Invoices:    newInvoicesV2(internal),
		Projects:    newProjectsV2(internal),
		Roles:       newRolesV2(internal),
		Tasks:       newTasksV2(internal),
		TimeEntries: newTimeEntriesV2(internal),
		Users:       newUsersV2(internal),
	}
}

func (client *internalClient) DoGet(url string, params ...interface{}) (HarvestResponse, error) {
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", client.baseUrl, url), nil)

	if err != nil {
		return HarvestResponse{}, nil
	}

	if len(params) > 0 {
		values, err := query.Values(params[0])

		if err != nil {
			return HarvestResponse{}, nil
		}

		r.URL.RawQuery = values.Encode()
	}

	client.setHeaders(r)

	return client.readResponse(r)
}

func (client *internalClient) DoPost(url string, body ...interface{}) (HarvestResponse, error) {
	b, err := client.getJsonBody(body)

	if err != nil {
		return HarvestResponse{}, err
	}

	r, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", client.baseUrl, url), b)

	if err != nil {
		return HarvestResponse{}, err
	}

	client.setHeaders(r, "application/json")

	return client.readResponse(r)
}

func (client *internalClient) DoPostMultipart(url string, formData multipartData) (HarvestResponse, error) {
	ct, b, err := client.getMultipartBody(formData)

	if err != nil {
		return HarvestResponse{}, err
	}

	r, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", client.baseUrl, url), b)

	if err != nil {
		return HarvestResponse{}, err
	}

	client.setHeaders(r, ct)

	return client.readResponse(r)
}

func (client *internalClient) DoPatch(url string, body ...interface{}) (HarvestResponse, error) {
	b, err := client.getJsonBody(body)

	if err != nil {
		return HarvestResponse{}, err
	}

	r, err := http.NewRequest("PATCH", fmt.Sprintf("%s/%s", client.baseUrl, url), b)

	if err != nil {
		return HarvestResponse{}, err
	}

	client.setHeaders(r, "application/json")

	return client.readResponse(r)
}

func (client *internalClient) DoDelete(url string) (HarvestResponse, error) {
	r, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", client.baseUrl, url), nil)

	if err != nil {
		return HarvestResponse{}, err
	}

	client.setHeaders(r)

	return client.readResponse(r)
}

func (client *internalClient) SetQuery(r *http.Request, queryStr map[string]string) {
	query := r.URL.Query()
	for key, value := range queryStr {
		query.Set(key, value)
	}

	r.URL.RawQuery = query.Encode()
}

func (client *internalClient) setHeaders(r *http.Request, contentType ...string) {
	r.Header.Set("User-Agent", fmt.Sprintf("%s (%s)", client.userAgentApp, client.userAgentEmail))
	r.Header.Set("Harvest-Account-ID", client.accountId)
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.accessToken))

	if len(contentType) > 0 && len(contentType[0]) > 0 {
		r.Header.Set("Content-Type", contentType[0])
	}
}

func (client *internalClient) getJsonBody(body []interface{}) (*bytes.Buffer, error) {
	var b *bytes.Buffer

	if len(body) > 0 && body[0] != nil {
		buff, err := json.Marshal(body[0])

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(buff)
	}

	return b, nil
}

func (client *internalClient) getMultipartBody(formData multipartData) (string, io.Reader, error) {
	buf := &bytes.Buffer{}
	bw := multipart.NewWriter(buf)

	for field, value := range formData.data {
		pw, err := bw.CreateFormField(field)

		if err != nil {
			return "", nil, err
		}

		_, err = pw.Write([]byte(value))

		if err != nil {
			return "", nil, err
		}
	}

	for field, path := range formData.files {
		f, err := os.Open(path)
		if err != nil {
			return "", nil, err
		}
		defer f.Close()

		_, filename := filepath.Split(path)
		fw, err := bw.CreateFormFile(field, filename)
		if err != nil {
			return "", nil, err
		}

		_, err = io.Copy(fw, f)
		if err != nil {
			return "", nil, err
		}
	}

	bw.Close()
	return bw.FormDataContentType(), buf, nil
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
