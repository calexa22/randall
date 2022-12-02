package randall

import (
	"fmt"
	"time"
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
	Tax           *float32                        `json:"tax,omitempty"`
	Tax2          *float32                        `json:"tax2,omitempty"`
	Discount      *float32                        `json:"discount,omitempty"`
	Subject       *string                         `json:"subject,omitempty"`
	Notes         *string                         `json:"notes,omitempty"`
	Currency      *string                         `json:"currency,omitempty"`
	IssueDate     time.Time                       `json:"issue_date,omitempty"`
	LineItems     []CreateEstimateLineItemRequest `json:"line_items,omitempty"`
}

type CreateEstimateLineItemRequest struct {
	Kind        string   `json:"kind"`
	Description *string  `json:"description,omitempty"`
	Quantity    *uint    `json:"quantity,omitempty"`
	UnitPrice   *float32 `json:"unit_price,omitempty"`
	Taxed       *bool    `json:"taxed,omitempty"`
	Taxed2      *bool    `json:"taxed2,omitempty"`
}

type UpdateEstimateRequest struct {
	ClientId      *uint                           `json:"client_id,omitempty"`
	Number        *string                         `json:"number,omitempty"`
	PurchaseOrder *string                         `json:"purchase_order,omitempty"`
	Tax           *float32                        `json:"tax,omitempty"`
	Tax2          *float32                        `json:"tax2,omitempty"`
	Discount      *float32                        `json:"discount,omitempty"`
	Subject       *string                         `json:"subject,omitempty"`
	Notes         *string                         `json:"notes,omitempty"`
	Currency      *string                         `json:"currency,omitempty"`
	IssueDate     time.Time                       `json:"issue_date,omitempty"`
	LineItems     []UpdateEstimateLineItemRequest `json:"line_items,omitempty"`
}

type UpdateEstimateLineItemRequest struct {
	Id          *uint    `json:"id,omitempty"`
	Kind        *string  `json:"kind,omitempty"`
	Description *string  `json:"description,omitempty"`
	Quantity    *uint    `json:"quantity,omitempty"`
	UnitPrice   *float32 `json:"unit_price,omitempty"`
	Taxed       *bool    `json:"taxed,omitempty"`
	Taxed2      *bool    `json:"taxed2,omitempty"`
	Destroy     *bool    `json:"_destroy,omitempty"`
}

type CreateEstimateMessageRequest struct {
	Recipients  []EstimateMessageRecipient `json:"recipients"`
	Subject     *string                    `json:"subject,omitempty"`
	Body        *string                    `json:"body,omitempty"`
	SendMeACopy *bool                      `json:"send_me_a_copy,omitempty"`
	EventType   *string                    `json:"event_type,omitempty"`
}

type EstimateMessageRecipient struct {
	Email string  `json:"email"`
	Name  *string `json:"name,omitempty"`
}

type EstimateMessageEventType uint

const (
	Send EstimateMessageEventType = iota
	Accept
	Decline
	Reopen
)

type updateEventTypeRequest struct {
	EventType string `json:"event_type"`
}

type upsertEstimateItemCategory struct {
	Name string `json:"name"`
}

func newEstimatesV2(client *internalClient) EstimatesApi {
	return EstimatesApi{
		estimatesBaseUrl:              "v2/estimates",
		estimateItemCategoriesBaseUrl: "v2/estimate_item_categories",
		client:                        client,
	}
}

func (api EstimatesApi) GetAll(params ...HarvestCollectionParams) (HarvestResponse, error) {
	var param *HarvestCollectionParams
	if len(params) > 0 {
		param = &params[0]
	}
	return api.client.DoGet(api.estimatesBaseUrl, param)
}

func (api EstimatesApi) Get(estimateId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.estimatesBaseUrl, estimateId))
}

func (api EstimatesApi) Create(req CreateEstimateRequest) (HarvestResponse, error) {
	return api.client.DoPost(api.estimatesBaseUrl, req)
}

func (api EstimatesApi) Update(estimateId uint, req UpdateEstimateRequest) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.estimatesBaseUrl, estimateId), req)
}

func (api EstimatesApi) Delete(estimateId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d", api.estimatesBaseUrl, estimateId))
}

func (api EstimatesApi) GetAllEstimateMessages(estimateId uint, params ...HarvestCollectionParams) (HarvestResponse, error) {
	var param *HarvestCollectionParams

	if len(params) > 0 {
		param = &params[0]
	}
	return api.client.DoGet(fmt.Sprintf("%s/%d/messages", api.estimatesBaseUrl, estimateId), param)
}

func (api EstimatesApi) CreateEstimateMessage(estimateId uint, req CreateEstimateMessageRequest) (HarvestResponse, error) {
	return api.client.DoPost(fmt.Sprintf("%s/%d/messages", api.estimatesBaseUrl, estimateId), req)
}

func (api EstimatesApi) DeleteEstimateMessage(estimateId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d/messages", api.estimatesBaseUrl, estimateId))
}

func (api EstimatesApi) MarkDraftEstimateSent(estimateId uint) (HarvestResponse, error) {
	return api.client.DoPost(
		fmt.Sprintf("%s/%d/messages", api.estimatesBaseUrl, estimateId),
		getUpdateEventTypeRequest(Send))
}

func (api EstimatesApi) MarkEstimateAccepted(estimateId uint) (HarvestResponse, error) {
	return api.client.DoPost(
		fmt.Sprintf("%s/%d/messages", api.estimatesBaseUrl, estimateId),
		getUpdateEventTypeRequest(Accept))
}

func (api EstimatesApi) MarkEstimateDeclined(estimateId uint) (HarvestResponse, error) {
	return api.client.DoPost(
		fmt.Sprintf("%s/%d/messages", api.estimatesBaseUrl, estimateId),
		getUpdateEventTypeRequest(Decline))
}

func (api EstimatesApi) ReopenClosedEstimate(estimateId uint) (HarvestResponse, error) {
	return api.client.DoPost(
		fmt.Sprintf("%s/%d/messages", api.estimatesBaseUrl, estimateId),
		getUpdateEventTypeRequest(Reopen))
}

func (api EstimatesApi) GetAllEstimateItemCategories(params ...HarvestCollectionParams) (HarvestResponse, error) {
	var param *HarvestCollectionParams
	if len(params) > 0 {
		param = &params[0]
	}
	return api.client.DoGet(api.estimateItemCategoriesBaseUrl, param)
}

func (api EstimatesApi) GetEstimateItemCategory(estimateItemCategoryId uint) (HarvestResponse, error) {
	return api.client.DoGet(fmt.Sprintf("%s/%d", api.estimateItemCategoriesBaseUrl, estimateItemCategoryId))
}

func (api EstimatesApi) CreateEstimateItemCategory(categoryName string) (HarvestResponse, error) {
	return api.client.DoPost(api.estimateItemCategoriesBaseUrl, upsertEstimateItemCategory{
		Name: categoryName,
	})
}

func (api EstimatesApi) UpdateEstimateItemCategory(estimateCategoryItemId uint, categoryName string) (HarvestResponse, error) {
	return api.client.DoPatch(fmt.Sprintf("%s/%d", api.estimateItemCategoriesBaseUrl, estimateCategoryItemId),
		upsertEstimateItemCategory{
			Name: categoryName,
		})
}

func (api EstimatesApi) DeleteEstimateItemCategory(estimateCategoryItemId uint) (HarvestResponse, error) {
	return api.client.DoDelete(fmt.Sprintf("%s/%d", api.estimateItemCategoriesBaseUrl, estimateCategoryItemId))
}

func (e EstimateMessageEventType) toString() string {
	switch e {
	case Send:
		return "send"
	case Accept:
		return "accept"
	case Decline:
		return "decline"
	case Reopen:
		return "re-open"
	default:
		return ""
	}
}

func getUpdateEventTypeRequest(e EstimateMessageEventType) updateEventTypeRequest {
	return updateEventTypeRequest{
		EventType: e.toString(),
	}
}
