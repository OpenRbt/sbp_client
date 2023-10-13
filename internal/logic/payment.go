package logic

import (
	"context"
	"fmt"
	logicEntities "sbp/internal/logic/entities"
	"strings"
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
	IsNotificationCorrect(notification logicEntities.PaymentNotification, password string) bool
}

// LeaWashPublisher
type LeaWashPublisher interface {
	SendToLeaPaymentResponse(logicEntities.PaymentResponse) error
	SendToLeaPaymentNotification(logicEntities.PaymentNotificationForLea) error
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
	errorPrefix := "Pay error:"
	// get wash uuid
	id, err := uuid.FromString(payRequest.WashID)
	if err != nil {
		return nil, fmt.Errorf("%s wash_id=%s is not correct, error:%s", errorPrefix, payRequest.WashID, err.Error())
	}

	// get wash wash terminal
	wash, err := logic.washLogic.GetWash(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s wash not found by id=%s, error:%s", errorPrefix, payRequest.WashID, err.Error())
	}

	var orderID uuid.UUID
	if payRequest.OrderID != "" {
		orderID = uuid.FromStringOrNil(payRequest.OrderID)
	}

	if orderID == uuid.Nil {
		return nil, fmt.Errorf(
			"%s orderID = nil (wash_id=%s, post_id=%s, transaction_id=%s)",
			errorPrefix,
			payRequest.WashID,
			payRequest.OrderID,
			payRequest.PostID)
	}

	// payment init
	paymentCreate := logicEntities.PaymentCreate{
		TerminalKey: wash.TerminalKey,
		Amount:      payRequest.Amount,
		OrderID:     orderID.String(),
	}
	paymentInit, err := logic.payClient.Init(paymentCreate)
	if err != nil {
		return nil, fmt.Errorf(
			"%s payment init failed (wash_id=%s, post_id=%s, transaction_id=%s), error: %s",
			errorPrefix,
			payRequest.WashID,
			payRequest.PostID,
			paymentInit.PaymentID,
			err.Error())
	}

	if !paymentInit.Success {
		return nil, fmt.Errorf(
			"%s payment init failed (wash_id=%s, post_id=%s, transaction_id=%s), errorCode: %s, message: %s, details: %s",
			errorPrefix,
			payRequest.WashID,
			payRequest.PostID,
			paymentInit.PaymentID,
			paymentInit.ErrorCode,
			paymentInit.Message,
			paymentInit.Details)
	}

	urlPay := paymentInit.Url
	if !strings.HasSuffix(wash.TerminalKey, "DEMO") {
		// get QR code
		paymentCreds := logicEntities.PaymentCreds{
			TerminalKey: wash.TerminalKey,
			PaymentID:   paymentInit.PaymentID,
		}
		resp, err := logic.payClient.GetQr(paymentCreds, wash.TerminalPassword)
		if err != nil {
			return nil, fmt.Errorf(
				"%s get qr failed (wash_id=%s, post_id=%s, transaction_id=%s), error: %s",
				errorPrefix,
				payRequest.WashID,
				payRequest.PostID,
				paymentInit.PaymentID,
				err.Error())
		}
		if !resp.Success {
			return nil, fmt.Errorf(
				"%s get qr failed (wash_id=%s, post_id=%s, transaction_id=%s), errorCode: %s, message: %s, details: %s",
				errorPrefix,
				payRequest.WashID,
				payRequest.PostID,
				resp.PaymentID,
				resp.ErrorCode,
				resp.Message,
				resp.Details)
		}
		urlPay = resp.UrlPay
	}
	// add payment to db
	transactionStatus := logicEntities.TransactionStatusFromString(paymentInit.Status)
	if err != nil {
		logic.logger.Errorf(
			"%s failed get transaction status=%s from string (transaction_id=%s), error: %s",
			errorPrefix,
			paymentInit.Status,
			paymentInit.PaymentID,
			err.Error())
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
		return nil, fmt.Errorf(
			"%s create transaction failed (wash_id=%s, post_id=%s, transaction_id=%s), error: %s",
			errorPrefix,
			payRequest.WashID,
			payRequest.PostID,
			paymentInit.PaymentID,
			err.Error())
	}

	// send broker message
	payResponse := logicEntities.PaymentResponse{
		WashID:  transactionCreate.WashID,
		PostID:  payRequest.PostID,
		OrderID: orderID.String(),
		UrlPay:  urlPay,
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
		return nil, fmt.Errorf(
			"%s send to lea payment response failed (wash_id=%s, post_id=%s, transaction_id=%s), error: %s",
			errorPrefix,
			payRequest.WashID,
			payRequest.PostID,
			paymentInit.PaymentID,
			err.Error())
	}

	return &payResponse, nil
}

// Notification ...
func (logic *PaymentLogic) Notification(ctx context.Context, notification logicEntities.PaymentNotification) error {
	errorPrefix := "Notification error:"
	// get transaction
	id, err := uuid.FromString(notification.OrderID)
	if err != nil {
		return fmt.Errorf("%s order_id=%s is not correct, error:%s", errorPrefix, notification.OrderID, err.Error())
	}
	transaction, err := logic.repository.GetTransaction(ctx, id)
	if err != nil {
		return fmt.Errorf("%s transaction not found by id=%s, error:%s", errorPrefix, id, err.Error())
	}

	// get terminal
	washID, err := uuid.FromString(transaction.WashID)
	if err != nil {
		return fmt.Errorf("%s wash_id=%s is not correct, error:%s", errorPrefix, transaction.WashID, err.Error())
	}
	wash, err := logic.washLogic.GetWash(ctx, washID)
	if err != nil {
		return fmt.Errorf("%s wash not found by id=%s", errorPrefix, washID)
	}

	// check notification
	if !logic.payClient.IsNotificationCorrect(notification, wash.TerminalPassword) {
		return fmt.Errorf("%s notification is not correct (wash_id=%s, notification=%+#v)", errorPrefix, washID, notification)
	}

	// update transaction
	transactionStatus := logicEntities.TransactionStatusFromString(notification.Status)
	if transactionStatus == logicEntities.TransactionStatusUnknown {
		logic.logger.Errorf("%s notification status '%s' is unknown", errorPrefix, notification.Status)
	}
	err = logic.repository.UpdateTransaction(ctx, logicEntities.TransactionUpdate{
		ID:        id,
		Status:    transactionStatus,
		PaymentID: nil,
	})
	if err != nil {
		return fmt.Errorf("%s update trasaction failed (trasaction_id=%s) error: %s", errorPrefix, notification.OrderID, err.Error())
	}

	// send broker message
	paymentNotifcation := logicEntities.PaymentNotificationForLea{
		WashID:  transaction.WashID,
		PostID:  transaction.PostID,
		OrderID: transaction.ID.String(),
		Status:  notification.Status,
	}
	err = logic.leaWashPublisher.SendToLeaPaymentNotification(paymentNotifcation)
	if err != nil {
		logic.logger.Errorf("%s send notification to lea failed (transaction_id=%s) error: %s", errorPrefix, notification.OrderID, err.Error())
	}

	return nil
}

// Cancel ...
func (logic *PaymentLogic) Cancel(ctx context.Context, req logicEntities.PaymentСancellationRequest) (resendNeaded bool, err error) {
	errorPrefix := "Cancel error:"
	resendNeaded = true
	// get transaction by order_id
	id, err := uuid.FromString(req.OrderID)
	if err != nil {
		return resendNeaded, fmt.Errorf("%s transaction_id not correct (trasaction_id=%s) error: %s", errorPrefix, req.OrderID, err.Error())
	}
	transaction, err := logic.repository.GetTransaction(ctx, id)
	if err != nil {
		return resendNeaded, fmt.Errorf("%s get trasaction failed (trasaction_id=%s) error: %s", errorPrefix, req.OrderID, err.Error())
	}

	// get wash by wash_id
	washID, err := uuid.FromString(transaction.WashID)
	if err != nil {
		return resendNeaded, fmt.Errorf("%s wash_id not correct (wash_id=%s) error: %s", errorPrefix, transaction.WashID, err.Error())
	}
	wash, err := logic.washLogic.GetWash(ctx, washID)
	if err != nil {
		return resendNeaded, fmt.Errorf("%s get wash failed (wash_id=%s) error: %s", errorPrefix, washID, err.Error())
	}

	// update transaction status canceling
	if transaction.Status != logicEntities.TransactionStatusСanceling {
		err = logic.repository.UpdateTransaction(ctx, logicEntities.TransactionUpdate{
			ID:        transaction.ID,
			Status:    logicEntities.TransactionStatusСanceling,
			PaymentID: nil,
		})
		if err != nil {
			return resendNeaded, fmt.Errorf("%s set canceling status for trasaction failed (trasaction_id=%s) error: %s", errorPrefix, transaction.ID, err.Error())
		}
	}

	// cancel by pay_сlient
	paymentRegisterNotification := logicEntities.PaymentCreds{
		TerminalKey: wash.TerminalKey,
		PaymentID:   transaction.PaymentID,
	}
	_, err = logic.payClient.Cancel(paymentRegisterNotification, wash.TerminalPassword)
	if err != nil {
		return !resendNeaded, fmt.Errorf("%s cancel pay_client_payment failed (trasaction_id=%s) error: %s", errorPrefix, transaction.ID, err.Error())
	}

	//  update transaction status canceled
	err = logic.repository.UpdateTransaction(ctx, logicEntities.TransactionUpdate{
		ID:        transaction.ID,
		Status:    logicEntities.TransactionStatusСanceled,
		PaymentID: nil,
	})
	if err != nil {
		return !resendNeaded, fmt.Errorf("%s set canceled status for trasaction failed (trasaction_id=%s) error: %s", errorPrefix, transaction.ID, err.Error())
	}

	return !resendNeaded, nil
}

// SyncAllPayments ...
func (logic *PaymentLogic) SyncAllPayments(ctx context.Context) error {
	errorPrefix := "SyncAllPayments error:"
	// confirmed not synced
	confirmedNotSyncedStatus := logicEntities.TransactionStatusConfirmedNotSynced
	confirmedNotSyncedTransactions, err := logic.repository.GetTransactionsByStatus(ctx, logicEntities.TransactionsGet{
		Status: confirmedNotSyncedStatus,
	})
	if err != nil {
		return fmt.Errorf("%s get trasactions whith status 'confirmed_not_synced' failed, error: %s", errorPrefix, err.Error())
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
				return fmt.Errorf("%s set canceling status for trasaction failed (trasaction_id=%s) error: %s", errorPrefix, pt.ID.String(), err.Error())
			}
		}

		// get wash
		wash, err := logic.washLogic.GetWash(ctx, uuid.FromStringOrNil(pt.WashID))
		if err != nil {
			return fmt.Errorf("%s get wash failed (wash_id=%s) error: %s", errorPrefix, pt.WashID, err.Error())
		}

		// send to lea
		paidStatus := string(logicEntities.TransactionStatusConfirmed)
		paymentNotifcation := logicEntities.PaymentNotificationForLea{
			WashID:  wash.ID.String(),
			PostID:  pt.PostID,
			OrderID: pt.ID.String(),
			Status:  paidStatus,
		}

		err = logic.leaWashPublisher.SendToLeaPaymentNotification(paymentNotifcation)
		if err != nil {
			return fmt.Errorf("%s send notification to lea failed (transaction_id=%s) error: %s", errorPrefix, pt.ID.String(), err.Error())
		}

		// update status in db
		err = logic.repository.UpdateTransaction(ctx, logicEntities.TransactionUpdate{
			ID:        pt.ID,
			Status:    logicEntities.TransactionStatusConfirmed,
			PaymentID: nil,
		})
		if err != nil {
			return fmt.Errorf("%s set confirmed status for trasaction failed (trasaction_id=%s) error: %s", errorPrefix, pt.ID.String(), err.Error())
		}
	}

	// canceling not canceled
	cancelingStatus := logicEntities.TransactionStatusСanceling
	cancelingTransactions, err := logic.repository.GetTransactionsByStatus(ctx, logicEntities.TransactionsGet{
		Status: cancelingStatus,
	})
	if err != nil {
		return fmt.Errorf("%s get trasactions whith status 'canceling' failed, error: %s", errorPrefix, err.Error())
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
