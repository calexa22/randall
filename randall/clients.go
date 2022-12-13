package randall

import "fmt"

// Encapsulates the Harvest API methods under /company
type ClientsApi struct {
	baseUrl string
	client  *internalClient
}

type CreateClientRequest struct {
	Name     string  `json:"name"`
	IsActive *bool   `json:"is_active,omitempty"`
	Address  *string `json:"address,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

type PatchClientRequest struct {
	Name     *string `json:"name,omitempty"`
	IsActive *bool   `json:"is_active,omitempty"`
	Address  *string `json:"address,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

func newClientsV2(client *internalClient) ClientsApi {
	return ClientsApi{
		baseUrl: "v2/clients",
		client:  client,
	}
}

func (api ClientsApi) GetAll(params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(api.baseUrl, getOptionalCollectionParams(params))
}

func (api ClientsApi) Get(clientId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d", api.baseUrl, clientId))
}

func (api ClientsApi) Create(req CreateClientRequest) (HarvestResponse, error) {
	return api.client.doPost(api.baseUrl, req)
}

func (api ClientsApi) Update(clientId uint, req PatchClientRequest) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d", api.baseUrl, clientId), req)
}

func (api ClientsApi) Delete(clientId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d", api.baseUrl, clientId))
}
