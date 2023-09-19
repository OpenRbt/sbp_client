package logic

import (
	"context"
	"fmt"
	logicEntities "sbp/internal/logic/entities"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

// PaymentLogic ...
type PaymentLogic struct {
	logger                       *zap.SugaredLogger
	notificationExpirationPeriod time.Duration
	payClient                    PayClient
	repository                   PayRepository
	leaWashPublisher             LeaWashPublisher
	washLogic                    *WashLogic
}

// PayClient ... tinkoff now
type PayClient interface {
	Init(req logicEntities.PaymentCreate) (logicEntities.PaymentInit, error)
	GetQr(req logicEntities.PaymentCreds, password string) (logicEntities.PaymentGetQr, error)
	Cancel(req logicEntities.PaymentCreds, password string) (logicEntities.PaymentCancel, error)
	IsNotificationCorrect(req logicEntities.PaymentRegisterNotification, password string) bool
}

// LeaWashPublisher
type LeaWashPublisher interface {
	SendToLeaPaymentResponse(logicEntities.PaymentResponse) error
	SendToLeaPaymentNotification(logicEntities.PaymentNotifcation) error
	SendToLeaPaymentFailedResponse(washID string, postID string, orderID string) error
}

// PayRepository ...
type PayRepository interface {
	// Transaction ...
	GetTransaction(ctx context.Context, id uuid.UUID) (logicEntities.Transaction, error)
	CreateTransaction(ctx context.Context, createTransaction logicEntities.TransactionCreate) error
	UpdateTransaction(ctx context.Context, updateTransaction logicEntities.TransactionUpdate) error

	// GetTransactionsByStatus ...
	GetTransactionsByStatus(ctx context.Context, transactionsGet logicEntities.TransactionsGet) ([]logicEntities.Transaction, error)
}

// newPaymentLogic ...
func newPaymentLogic(
	ctx context.Context,
	logger *zap.SugaredLogger,
	notificationExpirationPeriod time.Duration,
	payClient PayClient,
	repository PayRepository,
	leaWashPublisher LeaWashPublisher,
	washLogic *WashLogic,
) (*PaymentLogic, error) {

	return &PaymentLogic{
		logger:                       logger,
		notificationExpirationPeriod: notificationExpirationPeriod,
		payClient:                    payClient,
		repository:                   repository,
		washLogic:                    washLogic,
		leaWashPublisher:             leaWashPublisher,
	}, nil
}

// Pay ...
func (logic *PaymentLogic) Pay(ctx context.Context, payRequest logicEntities.PaymentRequest) (*logicEntities.PaymentResponse, error) {
	// get wash uuid
	id, err := uuid.FromString(payRequest.WashID)
	if err != nil {
		return nil, err
	}

	// get wash wash terminal
	wash, err := logic.washLogic.GetWash(ctx, id)
	if err != nil {
		return nil, err
	}

	var orderID uuid.UUID
	if payRequest.OrderID != "" {
		orderID = uuid.FromStringOrNil(payRequest.OrderID)
	}

	if orderID == uuid.Nil {
		return nil, fmt.Errorf("payment failed: orderID = nil (wash_id: %s, post_id: %s )", payRequest.WashID, payRequest.PostID)
	}

	// payment init
	paymentCreate := logicEntities.PaymentCreate{
		TerminalKey: wash.TerminalKey,
		Amount:      payRequest.Amount,
		OrderID:     orderID.String(),
	}
	paymentInit, err := logic.payClient.Init(paymentCreate)
	if err != nil {
		return nil, err
	}

	// get QR code
	paymentCreds := logicEntities.PaymentCreds{
		TerminalKey: wash.TerminalKey,
		PaymentID:   paymentInit.PaymentID,
	}
	resp, err := logic.payClient.GetQr(paymentCreds, wash.TerminalPassword)
	if err != nil {
		return nil, err
	}
	if resp.ErrorCode != "0" {
		return nil, fmt.Errorf("pay client internal error: %s, %s", resp.ErrorCode, resp.Message)
	}

	// add payment to db
	transactionStatus := logicEntities.TransactionStatusFromString(paymentInit.Status)
	if err != nil {
		logic.logger.Error(err)
	}
	transactionCreate := logicEntities.TransactionCreate{
		ID:        orderID,
		WashID:    payRequest.WashID,
		PostID:    payRequest.PostID,
		Amount:    payRequest.Amount,
		PaymentID: paymentInit.PaymentID,
		Status:    transactionStatus,
	}
	err = logic.repository.CreateTransaction(ctx, transactionCreate)
	if err != nil {
		return nil, err
	}

	// send broker message
	payResponse := logicEntities.PaymentResponse{
		WashID:  transactionCreate.WashID,
		PostID:  payRequest.PostID,
		OrderID: resp.OrderID,
		UrlPay:  resp.UrlPay,
	}

	// for tests whithout qr
	// payResponse := logicEntities.PaymentResponse{
	// 	WashID:  transactionCreate.WashID,
	// 	PostID:  payRequest.PostID,
	// 	OrderID: orderID.String(),
	// 	UrlPay:  paymentInit.Url,
	// }
	//
	err = logic.leaWashPublisher.SendToLeaPaymentResponse(payResponse)
	if err != nil {
		return nil, err
	}

	return &payResponse, nil
}

// Notification ...
func (logic *PaymentLogic) Notification(ctx context.Context, notification logicEntities.PaymentRegisterNotification) error {
	// get transaction
	id, err := uuid.FromString(notification.OrderID)
	if err != nil {
		return err
	}
	transaction, err := logic.repository.GetTransaction(ctx, id)
	if err != nil {
		return err
	}

	// get terminal
	washID, err := uuid.FromString(transaction.WashID)
	if err != nil {
		return err
	}
	wash, err := logic.washLogic.GetWash(ctx, washID)
	if err != nil {
		return err
	}

	// check notification
	if !logic.payClient.IsNotificationCorrect(notification, wash.TerminalPassword) {
		return logicEntities.ErrNotification
	}

	// update transaction
	transactionStatus := logicEntities.TransactionStatusFromString(notification.Status)
	if transactionStatus == logicEntities.TransactionStatusUnknown {
		logic.logger.Errorf("Notification error: notification status '%s' is unknown", notification.Status)
	}
	err = logic.repository.UpdateTransaction(ctx, logicEntities.TransactionUpdate{
		ID:        id,
		Status:    transactionStatus,
		PaymentID: nil,
	})
	if err != nil {
		return err
	}

	// send broker message
	paymentNotifcation := logicEntities.PaymentNotifcation{
		WashID:  transaction.WashID,
		PostID:  transaction.PostID,
		OrderID: transaction.ID.String(),
		Status:  notification.Status,
	}
	err = logic.leaWashPublisher.SendToLeaPaymentNotification(paymentNotifcation)
	if err != nil {
		logic.logger.Error(err)
	}

	return nil
}

// Cancel ...
func (logic *PaymentLogic) Cancel(ctx context.Context, req logicEntities.PaymentСancellationRequest) (resendNeaded bool, err error) {

	resendNeaded = true
	// get transaction by order_id
	id, err := uuid.FromString(req.OrderID)
	if err != nil {
		return resendNeaded, err
	}
	transaction, err := logic.repository.GetTransaction(ctx, id)
	if err != nil {
		return resendNeaded, err
	}

	// get wash by wash_id
	washID, err := uuid.FromString(transaction.WashID)
	if err != nil {
		return resendNeaded, err
	}
	wash, err := logic.washLogic.GetWash(ctx, washID)
	if err != nil {
		return resendNeaded, err
	}

	// update transaction status canceling
	if transaction.Status != logicEntities.TransactionStatusСanceling {
		err = logic.repository.UpdateTransaction(ctx, logicEntities.TransactionUpdate{
			ID:        transaction.ID,
			Status:    logicEntities.TransactionStatusСanceling,
			PaymentID: nil,
		})
		if err != nil {
			return resendNeaded, err
		}
	}

	// cancel by pay_сlient
	paymentRegisterNotification := logicEntities.PaymentCreds{
		TerminalKey: wash.TerminalKey,
		PaymentID:   transaction.PaymentID,
	}
	_, err = logic.payClient.Cancel(paymentRegisterNotification, wash.TerminalPassword)
	if err != nil {
		return !resendNeaded, err
	}

	//  update transaction status canceled
	err = logic.repository.UpdateTransaction(ctx, logicEntities.TransactionUpdate{
		ID:        transaction.ID,
		Status:    logicEntities.TransactionStatusСanceled,
		PaymentID: nil,
	})
	if err != nil {
		return !resendNeaded, err
	}

	return !resendNeaded, nil
}

// SyncAllPayments ...
func (logic *PaymentLogic) SyncAllPayments(ctx context.Context) error {

	// confirmed not synced
	confirmedNotSyncedStatus := logicEntities.TransactionStatusConfirmedNotSynced
	confirmedNotSyncedTransactions, err := logic.repository.GetTransactionsByStatus(ctx, logicEntities.TransactionsGet{
		Status: confirmedNotSyncedStatus,
	})
	if err != nil {
		return err
	}

	for _, pt := range confirmedNotSyncedTransactions {
		// check expiration
		if time.Until(pt.CreatedAt) >= logic.notificationExpirationPeriod {
			// cancel
			err = logic.repository.UpdateTransaction(ctx, logicEntities.TransactionUpdate{
				ID:        pt.ID,
				Status:    logicEntities.TransactionStatusСanceling,
				PaymentID: nil,
			})
			if err != nil {
				return err
			}
		}

		// get wash
		wash, err := logic.washLogic.GetWash(ctx, uuid.FromStringOrNil(pt.WashID))
		if err != nil {
			return err
		}

		// send to lea
		paidStatus := string(logicEntities.TransactionStatusConfirmed)
		paymentNotifcation := logicEntities.PaymentNotifcation{
			WashID:  wash.ID.String(),
			PostID:  pt.PostID,
			OrderID: pt.ID.String(),
			Status:  paidStatus,
		}

		err = logic.leaWashPublisher.SendToLeaPaymentNotification(paymentNotifcation)
		if err != nil {
			return err
		}

		// update status in db
		err = logic.repository.UpdateTransaction(ctx, logicEntities.TransactionUpdate{
			ID:        pt.ID,
			Status:    logicEntities.TransactionStatusConfirmed,
			PaymentID: nil,
		})
		if err != nil {
			return err
		}
	}

	// canceling not canceled
	cancelingStatus := logicEntities.TransactionStatusСanceling
	cancelingTransactions, err := logic.repository.GetTransactionsByStatus(ctx, logicEntities.TransactionsGet{
		Status: cancelingStatus,
	})
	if err != nil {
		return err
	}
	for _, ct := range cancelingTransactions {
		_, err := logic.Cancel(ctx, logicEntities.PaymentСancellationRequest{
			OrderID: ct.ID.String(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
