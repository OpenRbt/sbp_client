package handlers

import (
	"sbp/internal/conversions"
	"sbp/internal/entities"
	"sbp/openapi/restapi/operations/transactions"
)

func (handler *Handler) GetTransactions(params transactions.GetTransactionsParams, auth *entities.Auth) transactions.GetTransactionsResponder {
	op := "Get transactions:"
	resp := transactions.NewGetTransactionsDefault(500)

	filter, err := conversions.TransactionFilterFromRest(params)
	if err != nil {
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	list, err := handler.svc.TransactionsList(params.HTTPRequest.Context(), auth, filter)
	if err != nil {
		setAPIError(handler.logger, op, err, resp)
		return resp
	}

	return transactions.NewGetTransactionsOK().WithPayload(conversions.TransactionPageToRest(list))
}
