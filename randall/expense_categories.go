package randall

import "fmt"

const baseExpenseCategoriesV2Url = "v2/expense_categories"

// Encapsulates the Harvest API methods under /expense_categories
type ExpenseCategoriesApi struct {
	baseUrl string
	client  *internalClient
}

type CreateExpenseCategoryRequest struct {
	Name      string   `json:"name"`
	UnitName  *string  `json:"unit_name,omitempty"`
	UnitPrice *float32 `json:"unit_price,omitempty"`
	IsActive  *bool    `json:"is_active,omitempty"`
}

type UpdateExpenseCategoryRequest struct {
	Name      *string  `json:"name"`
	UnitName  *string  `json:"unit_name,omitempty"`
	UnitPrice *float32 `json:"unit_price,omitempty"`
	IsActive  *bool    `json:"is_active,omitempty"`
}

func newExpenseCategoriesV2(client *internalClient) ExpenseCategoriesApi {
	return ExpenseCategoriesApi{
		baseUrl: baseExpenseCategoriesV2Url,
		client:  client,
	}
}

func (api *ExpenseCategoriesApi) GetAll(params HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.DoGetV2(api.baseUrl, params)
}

func (api *ExpenseCategoriesApi) Get(expenseCategoryId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.baseUrl, expenseCategoryId))
}

func (api *ExpenseCategoriesApi) Create(req CreateExpenseRequest) (HarvestResponse, error) {
	return api.client.DoPost(api.baseUrl, req)
}

func (api *ExpenseCategoriesApi) Update(expenseCategoryId uint, req UpdateExpenseRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.baseUrl, expenseCategoryId), req)
}

func (api *ExpenseCategoriesApi) Delete(expenseCategoryId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d", api.baseUrl, expenseCategoryId))
}
