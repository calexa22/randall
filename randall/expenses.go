package randall

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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
	Receipt           *string   `json:"receipt,omitempty" layout:"2006-01-02"`
}

type UpdateExpenseRequest struct {
	ProjectId         *uint      `json:"project_id,omitempty"`
	ExpenseCategoryId *uint      `json:"expense_category_id,omitempty"`
	SpentDate         *time.Time `json:"spent_date,omitempty" layout:"2006-01-02"`
	Units             *uint      `json:"units,omitempty"`
	TotalCost         *float32   `json:"total_cost,omitempty"`
	Notes             *string    `json:"notes,omitempty"`
	Billable          *bool      `json:"billable,omitempty"`
	Receipt           *string    `json:"receipt,omitempty"`
	DeleteReceipt     *bool      `json:"delete_receipt,omitempty"`
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
	return api.client.doGet(api.expensesBaseUrl, getOptionalCollectionParams(params))
}

func (api ExpensesApi) Get(expenseId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d", api.expensesBaseUrl, expenseId))
}

func (api ExpensesApi) Create(req CreateExpenseRequest) (HarvestResponse, error) {
	if req.Receipt != nil {
		multipart, err := req.multipartData()
		if err != nil {
			return HarvestResponse{}, err
		}
		return api.client.doPostMultipart(api.expensesBaseUrl, multipart)
	}

	return api.client.doPost(api.expensesBaseUrl, req)
}

func (api ExpensesApi) Update(expenseId uint, req UpdateExpenseRequest) (HarvestResponse, error) {
	if req.Receipt != nil {
		multipart, err := req.multipartData()
		if err != nil {
			return HarvestResponse{}, err
		}
		return api.client.doPostMultipart(api.expensesBaseUrl, multipart)
	}
	return api.client.doPatch(fmt.Sprintf("%s/%d", api.expensesBaseUrl, expenseId), req)
}

func (api ExpensesApi) Delete(expenseId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d", api.expensesBaseUrl, expenseId))
}

func (api ExpensesApi) GetAllExpenseCategories(params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(api.expenseCategoriesBaseUrl, getOptionalCollectionParams(params))
}

func (api ExpensesApi) GetExpenseCategory(expenseCategoryId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d", api.expenseCategoriesBaseUrl, expenseCategoryId))
}

func (api ExpensesApi) CreateExpenseCategory(req CreateExpenseRequest) (HarvestResponse, error) {
	return api.client.doPost(api.expenseCategoriesBaseUrl, req)
}

func (api ExpensesApi) UpdateExpenseCategory(expenseCategoryId uint, req UpdateExpenseRequest) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d", api.expenseCategoriesBaseUrl, expenseCategoryId), req)
}

func (api ExpensesApi) DeleteExpenseCategory(expenseCategoryId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d", api.expenseCategoriesBaseUrl, expenseCategoryId))
}

func (r CreateExpenseRequest) multipartData() (multipartData, error) {
	data := make(map[string]string, 8)

	data["project_id"] = strconv.FormatUint(uint64(r.ProjectId), 10)
	data["expense_category_id"] = strconv.FormatUint(uint64(r.ExpenseCategoryId), 10)
	data["spent_date"] = time.Time(r.SpentDate).Format("2006-01-02")

	if r.UserId != nil {
		data["user_id"] = strconv.FormatUint(uint64(*r.UserId), 10)
	}

	if r.Units != nil {
		data["units"] = strconv.FormatUint(uint64(*r.Units), 10)
	}

	if r.TotalCost != nil {
		// TODO find the proper way to serialize decimals for multipart requests
		// might look into decimal pkg
		//data["total_cost"] = strconv.FormatFloat(uint64(*r.Units), 10)
	}

	if r.Notes != nil {
		data["notes"] = *r.Notes
	}

	if r.Billable != nil {
		data["billable"] = strconv.FormatBool(*r.Billable)
	}

	files := make(map[string]string, 1)
	if r.Receipt != nil {
		if err := isValidReceipt(*r.Receipt); err != nil {
			return multipartData{}, err
		}

		files["receipt"] = *r.Receipt
	}

	return multipartData{
		data:  data,
		files: files,
	}, nil
}

func (r UpdateExpenseRequest) multipartData() (multipartData, error) {
	data := make(map[string]string, 8)

	if r.ProjectId != nil {
		data["project_id"] = strconv.FormatUint(uint64(*r.ProjectId), 10)
	}

	if r.ExpenseCategoryId != nil {
		data["expense_category_id"] = strconv.FormatUint(uint64(*r.ExpenseCategoryId), 10)
	}

	if r.SpentDate != nil {
		data["spent_date"] = time.Time(*r.SpentDate).Format("2006-01-02")
	}

	if r.Units != nil {
		data["units"] = strconv.FormatUint(uint64(*r.Units), 10)
	}

	if r.TotalCost != nil {
		// TODO find the proper way to serialize decimals for multipart requests
		// might look into decimal pkg
		//data["total_cost"] = strconv.FormatFloat(uint64(*r.Units), 10)
	}

	if r.Notes != nil {
		data["notes"] = *r.Notes
	}

	if r.Billable != nil {
		data["billable"] = strconv.FormatBool(*r.Billable)
	}

	files := make(map[string]string, 1)
	if r.Receipt != nil {
		if err := isValidReceipt(*r.Receipt); err != nil {
			return multipartData{}, err
		}

		files["receipt"] = *r.Receipt
	}

	return multipartData{
		data:  data,
		files: files,
	}, nil
}

func isValidReceipt(receiptPath string) error {
	if _, err := os.Stat(receiptPath); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("unable to find file at specified location %s", receiptPath)
	}

	validExts := []string{"pdf", "png", "jpg", "gif"}
	ext := filepath.Ext(receiptPath)
	isValidExt := false

	for _, validExt := range validExts {
		if ext == validExt {
			isValidExt = true
			break
		}
	}

	if !isValidExt {
		return fmt.Errorf("%s: invalid file type, valid files types are %+v", receiptPath, validExts)
	}

	return nil
}
