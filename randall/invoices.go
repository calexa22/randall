package randall

import (
	"fmt"
	"time"
)

const (
	InvoicePaymentTermUponReceipt = "upon receipt"
	InvoicePaymentTermNet15       = "net 15"
	InvoicePaymentTermNet30       = "net 30"
	InvoicePaymentTermNet45       = "net 45"
	InvoicePaymentTermNet60       = "net 60"
	InvoicePaymentTermCustom      = "custom"

	TimeImportSummaryTypeProject  = "project"
	TimeImportSummaryTypeTask     = "task"
	TimeImportSummaryTypePeople   = "people"
	TimeImportSummaryTypeDetailed = "detailed"

	ExpenseImportSummaryTypeProject  = "project"
	ExpenseImportSummaryTypeCategory = "category"
	ExpenseImportSummaryTypePeople   = "people"
	ExpenseImportSummaryTypeDetailed = "detailed"
)

// Encapsulates the Harvest API methods under /projects
type InvoicesApi struct {
	baseUrl               string
	itemCategoriesBaseUrl string
	client                *internalClient
}

type CreateFreeFormInvoiceRequest struct {
	ClientId      uint                            `json:"client_id"`
	RetainerId    *uint                           `json:"retainer_id,omitempty"`
	EstimateId    *uint                           `json:"estimate_id,omitempty"`
	Number        *string                         `json:"number,omitempty"`
	PurchaseOrder *string                         `json:"purchase_order,omitempty"`
	Tax           *float32                        `json:"tax,omitempty"`
	Tax2          *float32                        `json:"tax2,omitempty"`
	Discount      *float32                        `json:"discount,omitempty"`
	Subject       *string                         `json:"subject,omitempty"`
	Notes         *string                         `json:"notes,omitempty"`
	Currency      *string                         `json:"currency,omitempty"`
	IssueDate     *time.Time                      `json:"issue_date,omitempty" layout:"2006-01-02"`
	DueDate       *time.Time                      `json:"due_date,omitempty" layout:"2006-01-02"`
	PaymentTerm   *string                         `json:"payment_term,omitempty"`
	LineItems     []CreateEstimateLineItemRequest `json:"line_items,omitempty"`
}

type CreateInvoiceFromTrackedTimeAndExpenseRequest struct {
	ClientId        uint                          `json:"client_id"`
	RetainerId      *uint                         `json:"retainer_id,omitempty"`
	EstimateId      *uint                         `json:"estimate_id,omitempty"`
	Number          *string                       `json:"number,omitempty"`
	PurchaseOrder   *string                       `json:"purchase_order,omitempty"`
	Tax             *float32                      `json:"tax,omitempty"`
	Tax2            *float32                      `json:"tax2,omitempty"`
	Discount        *float32                      `json:"discount,omitempty"`
	Subject         *string                       `json:"subject,omitempty"`
	Notes           *string                       `json:"notes,omitempty"`
	Currency        *string                       `json:"currency,omitempty"`
	IssueDate       *time.Time                    `json:"issue_date,omitempty" layout:"2006-01-02"`
	DueDate         *time.Time                    `json:"due_date,omitempty" layout:"2006-01-02"`
	PaymentTerm     *string                       `json:"payment_term,omitempty"`
	LineItemsImport *CreateLineItemsImportRequest `json:"line_items_import,omitempty"`
}

type CreateInvoiceLineItemRequest struct {
	Kind        string   `json:"kind"`
	ProjectId   *uint    `json:"project_id,omitempty"`
	Description *string  `json:"description,omitempty"`
	Quantity    *uint    `json:"quantity,omitempty"`
	UnitPrice   *float32 `json:"unit_price,omitempty"`
	Taxed       *bool    `json:"taxed,omitempty"`
	Taxed2      *bool    `json:"taxed2,omitempty"`
}

type CreateLineItemsImportRequest struct {
	ProjectIds []uint         `json:"project_ids"`
	Time       *TimeImport    `json:"time,omitempty"`
	Expenses   ExpensesImport `json:"expenses,omitempty"`
}

type TimeImport struct {
	SummaryType string     `json:"summary_type"`
	From        *time.Time `json:"from,omitempty" layout:"2006-01-02"`
	To          *time.Time `json:"to,omitempty" layout:"2006-01-02"`
}

type ExpensesImport struct {
	SummaryType    string     `json:"summary_type"`
	From           *time.Time `json:"from,omitempty" layout:"2006-01-02"`
	To             *time.Time `json:"to,omitempty" layout:"2006-01-02"`
	AttachReceipts *bool      `json:"attach_receipts,omitempty"`
}

type UpdateInvoiceRequest struct {
	ClientId      *uint                          `json:"client_id,omitempty"`
	RetainerId    *uint                          `json:"retainer_id,omitempty"`
	EstimateId    *uint                          `json:"estimate_id,omitempty"`
	Number        *string                        `json:"number,omitempty"`
	PurchaseOrder *string                        `json:"purchase_order,omitempty"`
	Tax           *float32                       `json:"tax,omitempty"`
	Tax2          *float32                       `json:"tax2,omitempty"`
	Discount      *float32                       `json:"discount,omitempty"`
	Subject       *string                        `json:"subject,omitempty"`
	Notes         *string                        `json:"notes,omitempty"`
	Currency      *string                        `json:"currency,omitempty"`
	IssueDate     *time.Time                     `json:"issue_date,omitempty" layout:"2006-01-02"`
	DueDate       *time.Time                     `json:"due_date,omitempty" layout:"2006-01-02"`
	PaymentTerm   *string                        `json:"payment_term,omitempty"`
	LineItems     []UpdateInvoiceLineItemRequest `json:"line_items,omitempty"`
}

type UpdateInvoiceLineItemRequest struct {
	Id          *uint    `json:"id,omitempty"`
	ProjectId   *uint    `json:"project_id,omitempty"`
	Kind        *string  `json:"kind,omitempty"`
	Description *string  `json:"description,omitempty"`
	Quantity    *uint    `json:"quantity,omitempty"`
	UnitPrice   *float32 `json:"unit_price,omitempty"`
	Taxed       *bool    `json:"taxed,omitempty"`
	Taxed2      *bool    `json:"taxed2,omitempty"`
	Destroy     *bool    `json:"_destroy,omitempty"`
}

type CreateInvoiceMessageRequest struct {
	Recipients                 []MessageRecipient `json:"recipients"`
	Subject                    *string            `json:"subject,omitempty"`
	Body                       *string            `json:"body,omitempty"`
	IncldueLinkToClientInvoice *bool              `json:"include_link_to_client_invoice,omitempty"`
	AttachPdf                  *bool              `json:"attach_pdf,omitempty"`
	SendMeACopy                *bool              `json:"send_me_a_copy,omitempty"`
	ThankYou                   *bool              `json:"thank_you,omitempty"`
	EventType                  *string            `json:"event_type,omitempty"`
}

type CreateInvoicePaymentRequest struct {
	Amount   float32    `json:"amount"`
	PaidAt   *time.Time `json:"paid_at,omitempty"`
	PaidDate *time.Time `json:"paid_date,omitempty" layout:"2006-01-02"`
	Notes    *string    `json:"notes,omitempty"`
}

func newInvoicesV2(client *internalClient) InvoicesApi {
	return InvoicesApi{
		baseUrl:               "v2/invoices",
		itemCategoriesBaseUrl: "v2/invoice_item_categories",
		client:                client,
	}
}

func (api InvoicesApi) GetAll(params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(api.baseUrl, getOptionalCollectionParams(params))
}

func (api InvoicesApi) Get(invoiceId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d", api.baseUrl, invoiceId))
}

func (api InvoicesApi) CreateFreeForm(req CreateFreeFormInvoiceRequest) (HarvestResponse, error) {
	return api.client.doPost(api.baseUrl, req)
}

func (api InvoicesApi) CreateFromTrackedTimeAndExpenses(req CreateInvoiceFromTrackedTimeAndExpenseRequest) (HarvestResponse, error) {
	return api.client.doPost(api.baseUrl, req)
}

func (api InvoicesApi) Update(invoiceId uint, req UpdateInvoiceRequest) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d", api.baseUrl, invoiceId), req)
}

func (api InvoicesApi) Delete(invoiceId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d", api.baseUrl, invoiceId))
}

func (api InvoicesApi) GetAllInvoiceItemCategories(params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(api.itemCategoriesBaseUrl, getOptionalCollectionParams(params))
}

func (api InvoicesApi) GetInvoiceItemCategory(invoiceItemCategoryItemId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d", api.itemCategoriesBaseUrl, invoiceItemCategoryItemId))
}

func (api InvoicesApi) CreateInvoiceItemCategory(categoryName string) (HarvestResponse, error) {
	return api.client.doPost(api.itemCategoriesBaseUrl, upsertItemCategoryRequest{
		Name: categoryName,
	})
}

func (api InvoicesApi) UpdateInvoiceItemCategory(invoiceItemCategoryItemId uint, categoryName string) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d", api.itemCategoriesBaseUrl, invoiceItemCategoryItemId),
		upsertItemCategoryRequest{
			Name: categoryName,
		})
}

func (api EstimatesApi) DeleteInvoiceItemCategory(estimateCategoryItemId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d", api.estimateItemCategoriesBaseUrl, estimateCategoryItemId))
}

func (api InvoicesApi) GetAllInvoiceMessages(invoiceId uint, params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d/messages", api.baseUrl, invoiceId), getOptionalCollectionParams(params))
}

func (api InvoicesApi) GetInvoiceMessageandBody(invoiceId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d/messages/new", api.baseUrl, invoiceId))
}

func (api InvoicesApi) CreateInvoiceMessage(invoiceId uint, req CreateInvoiceMessageRequest) (HarvestResponse, error) {
	return api.client.doPost(fmt.Sprintf("%s/%d/messages", api.baseUrl, invoiceId), req)
}

func (api InvoicesApi) MarkDraftEstimateSent(invoiceId uint) (HarvestResponse, error) {
	return api.client.doPost(
		fmt.Sprintf("%s/%d/messages", api.baseUrl, invoiceId),
		getUpdateEventTypeRequest("send"))
}

func (api InvoicesApi) MarkOpenInvoiceClosed(invoiceId uint) (HarvestResponse, error) {
	return api.client.doPost(
		fmt.Sprintf("%s/%d/messages", api.client.baseUrl, invoiceId),
		getUpdateEventTypeRequest("close"))
}

func (api InvoicesApi) ReopenCloseInvoice(invoiceId uint) (HarvestResponse, error) {
	return api.client.doPost(
		fmt.Sprintf("%s/%d/messages", api.baseUrl, invoiceId),
		getUpdateEventTypeRequest("re-open"))
}

func (api InvoicesApi) MarkOpenInvoiceDraft(invoiceId uint) (HarvestResponse, error) {
	return api.client.doPost(
		fmt.Sprintf("%s/%d/messages", api.baseUrl, invoiceId),
		getUpdateEventTypeRequest("draft"))
}

func (api InvoicesApi) DeleteInvoiceMessage(invoiceId, invoiceMessageId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d/messages/%d", api.baseUrl, invoiceId, invoiceMessageId))
}

func (api InvoicesApi) GetAllInvoicePayments(invoiceId uint, params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d/payments", api.baseUrl, invoiceId), getOptionalCollectionParams(params))
}

func (api InvoicesApi) CreateInvoicePayment(invoiceId uint, req CreateInvoicePaymentRequest) (HarvestResponse, error) {
	return api.client.doPost(fmt.Sprintf("%s/%d/payments", api.baseUrl, invoiceId), req)
}

func (api InvoicesApi) DeleteInvoicePayment(invoiceId, paymentId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d/payments/%d", api.baseUrl, invoiceId, paymentId))
}
