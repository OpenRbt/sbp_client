package conversions

import (
	"fmt"
	"sbp/internal/entities"
	rabbitEntities "sbp/internal/entities/rabbit"
	"sbp/internal/repository/models"
	apiModels "sbp/openapi/models"
	"sbp/openapi/restapi/operations/transactions"
	payModels "sbp/tinkoffapi/models"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func InitPaymentFromRest(response payModels.ResponseInit) entities.PaymentInit {
	return entities.PaymentInit{
		PaymentInfo: entities.PaymentInfo{
			Success:   response.Success,
			OrderID:   response.OrderID,
			PaymentID: response.PaymentID,
		},
		Status: response.Status,
		Url:    response.PaymentURL,
	}
}

func GetQRFromRest(resp payModels.ResponseGetQr) entities.PaymentGetQr {

	return entities.PaymentGetQr{
		PaymentInfo: entities.PaymentInfo{
			Success:   resp.Success,
			OrderID:   resp.OrderID,
			PaymentID: fmt.Sprint(resp.PaymentID),
		},
		ErrorCode: resp.ErrorCode,
		Message:   resp.Message,
		UrlPay:    resp.Data,
	}
}

func CancelPaymentFromRest(resp payModels.ResponseCancel) entities.PaymentCancel {
	return entities.PaymentCancel{
		PaymentInfo: entities.PaymentInfo{
			Success:   resp.Success,
			OrderID:   resp.OrderID,
			PaymentID: fmt.Sprint(resp.PaymentID),
		},
		Status:    resp.Status,
		ErrorCode: resp.ErrorCode,
	}
}

func PaymentNotificationFromRest(req apiModels.Notification) entities.PaymentNotification {
	return entities.PaymentNotification{
		Success:     req.Success,
		Amount:      req.Amount,
		ErrorCode:   req.ErrorCode,
		OrderID:     req.OrderID,
		Pan:         req.Pan,
		PaymentID:   req.PaymentID,
		Status:      req.Status,
		TerminalKey: req.TerminalKey,
		Token:       req.Token,
		ExpDate:     req.ExpDate,
		CardID:      req.CardID,
	}
}

func PaymentResponseToLea(e entities.PaymentResponse) rabbitEntities.PaymentResponse {
	return rabbitEntities.PaymentResponse{
		WashID:     e.WashID,
		PostID:     e.PostID,
		OrderID:    e.OrderID,
		UrlPayment: e.UrlPay,
		Failed:     e.Failed,
		Error:      e.Error,
	}
}

func PaymentNotifcationToLea(e entities.PaymentNotificationForLea) rabbitEntities.PaymentNotifcation {
	return rabbitEntities.PaymentNotifcation{
		WashID:  e.WashID,
		PostID:  e.PostID,
		OrderID: e.OrderID,
		Status:  e.Status,
	}
}

func PaymentRequestToSbp(e rabbitEntities.PaymentRequest) entities.PaymentRequest {
	return entities.PaymentRequest{
		WashID:  e.WashID,
		PostID:  e.PostID,
		OrderID: e.OrderID,
		Amount:  e.Amount,
	}
}

func PaymentСancellationRequestToSbp(e rabbitEntities.PaymentСancellationRequest) entities.PaymentСancellationRequest {
	return entities.PaymentСancellationRequest{
		WashID:  e.WashID,
		PostID:  e.PostID,
		OrderID: e.OrderID,
	}
}

func TransactionStatusFromRest(status apiModels.TransactionStatus) entities.TransactionStatus {
	switch status {
	case apiModels.TransactionStatusNew:
		return entities.TransactionStatusNew
	case apiModels.TransactionStatusAuthorized:
		return entities.TransactionStatusAuthorized
	case apiModels.TransactionStatusConfirmedNotSynced:
		return entities.TransactionStatusConfirmedNotSynced
	case apiModels.TransactionStatusConfirmed:
		return entities.TransactionStatusConfirmed
	case apiModels.TransactionStatusCanceling:
		return entities.TransactionStatusCanceling
	case apiModels.TransactionStatusCanceled:
		return entities.TransactionStatusCanceled
	case apiModels.TransactionStatusRefunded:
		return entities.TransactionStatusRefunded
	case apiModels.TransactionStatusUnknown:
		return entities.TransactionStatusUnknown
	default:
		panic(fmt.Sprintf("unable to parse status '%s' to app layer", status))
	}
}

func TransactionStatusToRest(status entities.TransactionStatus) apiModels.TransactionStatus {
	switch status {
	case entities.TransactionStatusNew:
		return apiModels.TransactionStatusNew
	case entities.TransactionStatusAuthorized:
		return apiModels.TransactionStatusAuthorized
	case entities.TransactionStatusConfirmedNotSynced:
		return apiModels.TransactionStatusConfirmedNotSynced
	case entities.TransactionStatusConfirmed:
		return apiModels.TransactionStatusConfirmed
	case entities.TransactionStatusCanceling:
		return apiModels.TransactionStatusCanceling
	case entities.TransactionStatusCanceled:
		return apiModels.TransactionStatusCanceled
	case entities.TransactionStatusRefunded:
		return apiModels.TransactionStatusRefunded
	case entities.TransactionStatusUnknown:
		return apiModels.TransactionStatusUnknown
	default:
		panic(fmt.Sprintf("unable to parse status '%s' to app layer", status))
	}
}

func TransactionToRest(transaction entities.TransactionForPage) apiModels.Transaction {
	id := strfmt.UUID(transaction.ID.String())
	status := TransactionStatusToRest(transaction.Status)
	washId := strfmt.UUID(transaction.Wash.ID.String())
	groupId := strfmt.UUID(transaction.Group.ID.String())
	organizationId := strfmt.UUID(transaction.Organization.ID.String())
	return apiModels.Transaction{
		ID:        &id,
		Amount:    &transaction.Amount,
		CreatedAt: (*strfmt.DateTime)(&transaction.CreatedAt),
		Status:    &status,
		Wash: &apiModels.SimpleWash{
			Name:    &transaction.Wash.Title,
			Deleted: &transaction.Wash.Deleted,
			ID:      &washId,
		},
		Group: &apiModels.Group{
			Name:    &transaction.Group.Name,
			Deleted: &transaction.Group.Deleted,
			ID:      &groupId,
		},
		Organization: &apiModels.Organization{
			Name:    &transaction.Organization.Name,
			Deleted: &transaction.Organization.Deleted,
			ID:      &organizationId,
		},
	}
}

func TransactionPageToRest(page entities.Page[entities.TransactionForPage]) *apiModels.TransactionPage {
	items := []*apiModels.Transaction{}
	for _, i := range page.Items {
		t := TransactionToRest(i)
		items = append(items, &t)
	}

	return &apiModels.TransactionPage{
		Items:      items,
		Page:       &page.Page,
		PageSize:   &page.PageSize,
		TotalItems: &page.TotalPages,
		TotalPages: &page.TotalPages,
	}
}

func TransactionFilterFromRest(params transactions.GetTransactionsParams) (entities.TransactionFilter, error) {
	var organizationID *uuid.UUID
	if params.OrganizationID != nil {
		id, err := uuid.FromString(params.OrganizationID.String())
		if err != nil {
			return entities.TransactionFilter{}, err
		}
		organizationID = &id
	}
	var groupID *uuid.UUID
	if params.GroupID != nil {
		id, err := uuid.FromString(params.GroupID.String())
		if err != nil {
			return entities.TransactionFilter{}, err
		}
		groupID = &id
	}
	var washServerID *uuid.UUID
	if params.WashID != nil {
		id, err := uuid.FromString(params.WashID.String())
		if err != nil {
			return entities.TransactionFilter{}, err
		}
		washServerID = &id
	}
	var status *entities.TransactionStatus
	if params.Status != nil {
		s := TransactionStatusFromRest(apiModels.TransactionStatus(*params.Status))
		status = &s
	}

	return entities.TransactionFilter{
		Filter:         entities.NewFilter(*params.Page, *params.PageSize),
		OrganizationID: organizationID,
		GroupID:        groupID,
		WashID:         washServerID,
		PostID:         params.PostID,
		Status:         status,
	}, nil
}

func TransactionFromDB(transaction models.Transaction) entities.TransactionForPage {
	return entities.TransactionForPage{
		ID:        transaction.ID,
		PostID:    transaction.PostID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
		Wash: entities.SimpleWash{
			ID:      transaction.Wash.ID,
			Title:   transaction.Wash.Title,
			Deleted: transaction.Wash.Deleted,
		},
		Group: entities.SimpleGroup{
			ID:      transaction.Group.ID,
			Name:    transaction.Group.Name,
			Deleted: transaction.Group.Deleted,
		},
		Organization: entities.SimpleOrganization{
			ID:      transaction.Organization.ID,
			Name:    transaction.Organization.Name,
			Deleted: transaction.Organization.Deleted,
		},
	}
}

func TransactionsFromDB(transactions []models.Transaction) []entities.TransactionForPage {
	items := []entities.TransactionForPage{}
	for _, v := range transactions {
		items = append(items, TransactionFromDB(v))
	}

	return items
}
