package randall

import "fmt"

const baseExpensesV2Url = "v2/expenses"

// Encapsulates the Harvest API methods under /expenses
type ExpensesApi struct {
	baseUrl string
	client  *internalClient
}

type CreateExpenseRequest struct {
	ProjectId         uint        `json:"project_id"`
	ExpenseCategoryId uint        `json:"expense_category_id"`
	SpentDate         HarvestDate `json:"spent_date"`
	UserId            *uint       `json:"user_id,omitempty"`
	Units             *uint       `json:"units,omitempty"`
	TotalCost         *float32    `json:"total_cost,omitempty"`
	Notes             *string     `json:"notes,omitempty"`
	Billable          *bool       `json:"billable,omitempty"`
	Receipt           *string     `json:"receipt,omitempty"` // TODO add support for attachment
}

type UpdateExpenseRequest struct {
	ProjectId         uint        `json:"project_id"`
	ExpenseCategoryId uint        `json:"expense_category_id"`
	SpentDate         HarvestDate `json:"spent_date"`
	Units             *uint       `json:"units,omitempty"`
	TotalCost         *float32    `json:"total_cost,omitempty"`
	Notes             *string     `json:"notes,omitempty"`
	Billable          *bool       `json:"billable,omitempty"`
	Receipt           *string     `json:"receipt,omitempty"` // TODO add support for attachment
	DeleteReceipt     *bool       `json:"delete_receipt,omitempty"`
}

func newExpensesV2(client *internalClient) ExpensesApi {
	return ExpensesApi{
		baseUrl: baseExpensesV2Url,
		client:  client,
	}
}

func (api *ExpensesApi) GetAll(params HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.DoGetV2(api.baseUrl, params)
}

func (api *ExpensesApi) Get(expenseId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.baseUrl, expenseId))
}

func (api *ExpensesApi) Create(req CreateExpenseRequest) (HarvestResponse, error) {
	return api.client.DoPost(api.baseUrl, req)
}

func (api *ExpensesApi) Update(expenseId uint, req UpdateExpenseRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.baseUrl, expenseId), req)
}

func (api *ExpensesApi) Delete(expenseId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d", api.baseUrl, expenseId))
}
