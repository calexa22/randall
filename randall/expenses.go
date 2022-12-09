package randall

import (
	"fmt"
	"time"
)

// Encapsulates the Harvest API methods under /expenses and /expense_categories
type ExpensesApi struct {
	expensesBaseUrl          string
	expenseCategoriesBaseUrl string
	client                   *internalClient
}

type CreateExpenseRequest struct {
	ProjectId         uint      `json:"project_id"`
	ExpenseCategoryId uint      `json:"expense_category_id"`
	SpentDate         time.Time `json:"spent_date" layout:"2006-01-02"`
	UserId            *uint     `json:"user_id,omitempty"`
	Units             *uint     `json:"units,omitempty"`
	TotalCost         *float32  `json:"total_cost,omitempty"`
	Notes             *string   `json:"notes,omitempty"`
	Billable          *bool     `json:"billable,omitempty"`
	Receipt           *string   `json:"receipt,omitempty" layout:"2006-01-02"` // TODO add support for attachment
}

type UpdateExpenseRequest struct {
	ProjectId         *uint     `json:"project_id,omitempty"`
	ExpenseCategoryId *uint     `json:"expense_category_id,omitempty"`
	SpentDate         time.Time `json:"spent_date,omitempty" layout:"2006-01-02"`
	Units             *uint     `json:"units,omitempty"`
	TotalCost         *float32  `json:"total_cost,omitempty"`
	Notes             *string   `json:"notes,omitempty"`
	Billable          *bool     `json:"billable,omitempty"`
	Receipt           *string   `json:"receipt,omitempty"` // TODO add support for attachment
	DeleteReceipt     *bool     `json:"delete_receipt,omitempty"`
}

type CreateExpenseCategoryRequest struct {
	Name      string   `json:"name"`
	UnitName  *string  `json:"unit_name,omitempty"`
	UnitPrice *float32 `json:"unit_price,omitempty"`
	IsActive  *bool    `json:"is_active,omitempty"`
}

type UpdateExpenseCategoryRequest struct {
	Name      *string  `json:"name,omitempty"`
	UnitName  *string  `json:"unit_name,omitempty"`
	UnitPrice *float32 `json:"unit_price,omitempty"`
	IsActive  *bool    `json:"is_active,omitempty"`
}

func newExpensesV2(client *internalClient) ExpensesApi {
	return ExpensesApi{
		expensesBaseUrl:          "v2/expenses",
		expenseCategoriesBaseUrl: "v2/expense_categories",
		client:                   client,
	}
}

func (api ExpensesApi) GetAll(params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.DoGet(api.expensesBaseUrl, getOptionalCollectionParams(params))
}

func (api ExpensesApi) Get(expenseId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.expensesBaseUrl, expenseId))
}

func (api ExpensesApi) Create(req CreateExpenseRequest) (HarvestResponse, error) {
	return api.client.DoPost(api.expensesBaseUrl, req)
}

func (api ExpensesApi) Update(expenseId uint, req UpdateExpenseRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.expensesBaseUrl, expenseId), req)
}

func (api ExpensesApi) Delete(expenseId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d", api.expensesBaseUrl, expenseId))
}

func (api ExpensesApi) GetAllExpenseCategories(params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.DoGet(api.expenseCategoriesBaseUrl, getOptionalCollectionParams(params))
}

func (api ExpensesApi) GetExpenseCategory(expenseCategoryId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.expenseCategoriesBaseUrl, expenseCategoryId))
}

func (api ExpensesApi) CreateExpenseCategory(req CreateExpenseRequest) (HarvestResponse, error) {
	return api.client.DoPost(api.expenseCategoriesBaseUrl, req)
}

func (api ExpensesApi) UpdateExpenseCategory(expenseCategoryId uint, req UpdateExpenseRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.expenseCategoriesBaseUrl, expenseCategoryId), req)
}

func (api ExpensesApi) DeleteExpenseCategory(expenseCategoryId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d", api.expenseCategoriesBaseUrl, expenseCategoryId))
}
