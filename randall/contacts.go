package randall

import "fmt"

// Encapsulates the Harvest API methods under /contacts
type ContactsApi struct {
	baseUrl string
	client  *internalClient
}

type CreateContactRequest struct {
	ClientId    uint    `json:"client_id"`
	Firstname   string  `json:"first_name"`
	Title       *string `json:"title,omitempty"`
	Lastname    *string `json:"last_name,omitempty"`
	Email       *string `json:"email,omitempty"`
	OfficePhone *string `json:"phone_office,omitempty"`
	MobilePhone *string `json:"phone_mobile,omitempty"`
	Fax         *string `json:"fax,omitempty"`
}

type PatchContactRequest struct {
	ClientId    *uint   `json:"client_id,omitempty"`
	Firstname   *string `json:"first_name,omitempty"`
	Title       *string `json:"title,omitempty"`
	Lastname    *string `json:"last_name,omitempty"`
	Email       *string `json:"email,omitempty"`
	OfficePhone *string `json:"phone_office,omitempty"`
	MobilePhone *string `json:"phone_mobile,omitempty"`
	Fax         *string `json:"fax,omitempty"`
}

func newContactsV2(client *internalClient) ContactsApi {
	return ContactsApi{
		baseUrl: "v2/contacts",
		client:  client,
	}
}

func (api ContactsApi) GetAll(params ...HarvestCollectionParams) (HarvestResponse, error) {
	var param *HarvestCollectionParams

	if len(params) > 0 {
		param = &params[0]
	}
	return api.client.DoGet(api.baseUrl, param)
}

func (api ContactsApi) Get(contactId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.baseUrl, contactId))
}

func (api ContactsApi) CreateContact(req CreateContactRequest) (HarvestResponse, error) {
	return api.client.DoPost(api.baseUrl, req)
}

func (api ContactsApi) UpdateContact(contactId uint, req PatchContactRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.baseUrl, contactId), req)
}

func (api ContactsApi) DeleteClient(contactId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d", api.baseUrl, contactId))
}
