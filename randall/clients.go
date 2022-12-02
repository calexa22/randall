package randall

import "fmt"

const _baseClientsV2Url = "v2/clients"

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
		baseUrl: _baseClientsV2Url,
		client:  client,
	}
}

func (api ClientsApi) GetAll(params ...HarvestCollectionParams) (HarvestResponse, error) {
	var param *HarvestCollectionParams

	if len(params) > 0 {
		param = &params[0]
	}
	return api.client.DoGet(api.baseUrl, param)
}

func (api ClientsApi) Get(clientId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.baseUrl, clientId))
}

func (api ClientsApi) Create(req CreateClientRequest) (HarvestResponse, error) {
	return api.client.DoPost(api.baseUrl, req)
}

func (api ClientsApi) Update(clientId uint, req PatchClientRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.baseUrl, clientId), req)
}

func (api ClientsApi) Delete(clientId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d", api.baseUrl, clientId))
}
