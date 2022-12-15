package randall

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

// Encapsulates the Harvest API methods under /expenses
type EstimatesApi struct {
	estimatesBaseUrl              string
	estimateItemCategoriesBaseUrl string
	client                        *internalClient
}

type CreateEstimateRequest struct {
	ClientId      uint                            `json:"client_id"`
	Number        *string                         `json:"number,omitempty"`
	PurchaseOrder *string                         `json:"purchase_order,omitempty"`
	Tax           *decimal.Decimal                `json:"tax,omitempty"`
	Tax2          *decimal.Decimal                `json:"tax2,omitempty"`
	Discount      *decimal.Decimal                `json:"discount,omitempty"`
	Subject       *string                         `json:"subject,omitempty"`
	Notes         *string                         `json:"notes,omitempty"`
	Currency      *string                         `json:"currency,omitempty"`
	IssueDate     time.Time                       `json:"issue_date,omitempty"`
	LineItems     []CreateEstimateLineItemRequest `json:"line_items,omitempty"`
}

type CreateEstimateLineItemRequest struct {
	Kind        string           `json:"kind"`
	Description *string          `json:"description,omitempty"`
	Quantity    *uint            `json:"quantity,omitempty"`
	UnitPrice   *decimal.Decimal `json:"unit_price,omitempty"`
	Taxed       *bool            `json:"taxed,omitempty"`
	Taxed2      *bool            `json:"taxed2,omitempty"`
}

type UpdateEstimateRequest struct {
	ClientId      *uint                           `json:"client_id,omitempty"`
	Number        *string                         `json:"number,omitempty"`
	PurchaseOrder *string                         `json:"purchase_order,omitempty"`
	Tax           *decimal.Decimal                `json:"tax,omitempty"`
	Tax2          *decimal.Decimal                `json:"tax2,omitempty"`
	Discount      *decimal.Decimal                `json:"discount,omitempty"`
	Subject       *string                         `json:"subject,omitempty"`
	Notes         *string                         `json:"notes,omitempty"`
	Currency      *string                         `json:"currency,omitempty"`
	IssueDate     time.Time                       `json:"issue_date,omitempty"`
	LineItems     []UpdateEstimateLineItemRequest `json:"line_items,omitempty"`
}

type UpdateEstimateLineItemRequest struct {
	Id          *uint            `json:"id,omitempty"`
	Kind        *string          `json:"kind,omitempty"`
	Description *string          `json:"description,omitempty"`
	Quantity    *uint            `json:"quantity,omitempty"`
	UnitPrice   *decimal.Decimal `json:"unit_price,omitempty"`
	Taxed       *bool            `json:"taxed,omitempty"`
	Taxed2      *bool            `json:"taxed2,omitempty"`
	Destroy     *bool            `json:"_destroy,omitempty"`
}

type CreateEstimateMessageRequest struct {
	Recipients  []MessageRecipient `json:"recipients"`
	Subject     *string            `json:"subject,omitempty"`
	Body        *string            `json:"body,omitempty"`
	SendMeACopy *bool              `json:"send_me_a_copy,omitempty"`
	EventType   *string            `json:"event_type,omitempty"`
}

func newEstimatesV2(client *internalClient) EstimatesApi {
	return EstimatesApi{
		estimatesBaseUrl:              "v2/estimates",
		estimateItemCategoriesBaseUrl: "v2/estimate_item_categories",
		client:                        client,
	}
}

func (api EstimatesApi) GetAll(params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(api.estimatesBaseUrl, getOptionalCollectionParams(params))
}

func (api EstimatesApi) Get(estimateId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d", api.estimatesBaseUrl, estimateId))
}

func (api EstimatesApi) Create(req CreateEstimateRequest) (HarvestResponse, error) {
	return api.client.doPost(api.estimatesBaseUrl, req)
}

func (api EstimatesApi) Update(estimateId uint, req UpdateEstimateRequest) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d", api.estimatesBaseUrl, estimateId), req)
}

func (api EstimatesApi) Delete(estimateId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d", api.estimatesBaseUrl, estimateId))
}

func (api EstimatesApi) GetAllEstimateMessages(estimateId uint, params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(
		fmt.Sprintf("%s/%d/messages", api.estimatesBaseUrl, estimateId),
		getOptionalCollectionParams(params),
	)
}

func (api EstimatesApi) CreateEstimateMessage(estimateId uint, req CreateEstimateMessageRequest) (HarvestResponse, error) {
	return api.client.doPost(fmt.Sprintf("%s/%d/messages", api.estimatesBaseUrl, estimateId), req)
}

func (api EstimatesApi) MarkDraftEstimateSent(estimateId uint) (HarvestResponse, error) {
	return api.client.doPost(
		fmt.Sprintf("%s/%d/messages", api.estimatesBaseUrl, estimateId),
		getUpdateEventTypeRequest("send"))
}

func (api EstimatesApi) MarkEstimateAccepted(estimateId uint) (HarvestResponse, error) {
	return api.client.doPost(
		fmt.Sprintf("%s/%d/messages", api.estimatesBaseUrl, estimateId),
		getUpdateEventTypeRequest("accept"),
	)
}

func (api EstimatesApi) MarkEstimateDeclined(estimateId uint) (HarvestResponse, error) {
	return api.client.doPost(
		fmt.Sprintf("%s/%d/messages", api.estimatesBaseUrl, estimateId),
		getUpdateEventTypeRequest("decline"),
	)
}

func (api EstimatesApi) ReopenClosedEstimate(estimateId uint) (HarvestResponse, error) {
	return api.client.doPost(
		fmt.Sprintf("%s/%d/messages", api.estimatesBaseUrl, estimateId),
		getUpdateEventTypeRequest("re-open"),
	)
}

func (api EstimatesApi) DeleteEstimateMessage(estimateId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d/messages", api.estimatesBaseUrl, estimateId))
}

func (api EstimatesApi) GetAllEstimateItemCategories(params ...HarvestCollectionParams) (HarvestResponse, error) {
	return api.client.doGet(api.estimateItemCategoriesBaseUrl, getOptionalCollectionParams(params))
}

func (api EstimatesApi) GetEstimateItemCategory(estimateItemCategoryId uint) (HarvestResponse, error) {
	return api.client.doGet(fmt.Sprintf("%s/%d", api.estimateItemCategoriesBaseUrl, estimateItemCategoryId))
}

func (api EstimatesApi) CreateEstimateItemCategory(categoryName string) (HarvestResponse, error) {
	return api.client.doPost(api.estimateItemCategoriesBaseUrl, upsertItemCategoryRequest{
		Name: categoryName,
	})
}

func (api EstimatesApi) UpdateEstimateItemCategory(estimateCategoryItemId uint, categoryName string) (HarvestResponse, error) {
	return api.client.doPatch(fmt.Sprintf("%s/%d", api.estimateItemCategoriesBaseUrl, estimateCategoryItemId),
		upsertItemCategoryRequest{
			Name: categoryName,
		})
}

func (api EstimatesApi) DeleteEstimateItemCategory(estimateCategoryItemId uint) (HarvestResponse, error) {
	return api.client.doDelete(fmt.Sprintf("%s/%d", api.estimateItemCategoriesBaseUrl, estimateCategoryItemId))
}
