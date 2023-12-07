package services

import (
	"context"
	"fmt"
	"sbp/internal/app"
	"sbp/internal/entities"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type paymentService struct {
	logger                       *zap.SugaredLogger
	notificationExpirationPeriod time.Duration
	payClient                    app.PaymentClient
	repository                   app.Repository
	leaWashPublisher             app.LeaWashPublisher
	washSvc                      *washService
}

func NewPaymentService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	notificationExpirationPeriod time.Duration,
	payClient app.PaymentClient,
	repository app.Repository,
	leaWashPublisher app.LeaWashPublisher,
	washLogic *washService,
) (*paymentService, error) {

	return &paymentService{
		logger:                       logger,
		notificationExpirationPeriod: notificationExpirationPeriod,
		payClient:                    payClient,
		repository:                   repository,
		washSvc:                      washLogic,
		leaWashPublisher:             leaWashPublisher,
	}, nil
}

func (svc *paymentService) InitPayment(ctx context.Context, payRequest entities.PaymentRequest) (*entities.PaymentResponse, error) {
	errorPrefix := "Pay error:"
	// get wash uuid
	id, err := uuid.FromString(payRequest.WashID)
	if err != nil {
		return nil, fmt.Errorf("%s wash_id=%s is not correct, error:%s", errorPrefix, payRequest.WashID, err.Error())
	}

	// get wash wash terminal
	wash, err := svc.repository.GetWashByID(ctx, id)
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
	paymentCreate := entities.PaymentCreate{
		TerminalKey: wash.TerminalKey,
		Amount:      payRequest.Amount,
		OrderID:     orderID.String(),
	}
	paymentInit, err := svc.payClient.InitPayment(paymentCreate)
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
	if !strings.HasSuffix(wash.TerminalKey, "DEMO") && payRequest.Amount >= 1000 {
		// get QR code
		paymentCreds := entities.PaymentCreds{
			TerminalKey: wash.TerminalKey,
			PaymentID:   paymentInit.PaymentID,
		}
		resp, err := svc.payClient.GetQr(paymentCreds, wash.TerminalPassword)
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
	transactionStatus := entities.TransactionStatusFromString(paymentInit.Status)
	if err != nil {
		svc.logger.Errorf(
			"%s failed get transaction status=%s from string (transaction_id=%s), error: %s",
			errorPrefix,
			paymentInit.Status,
			paymentInit.PaymentID,
			err.Error())
	}
	transactionCreate := entities.TransactionCreate{
		ID:            orderID,
		WashID:        payRequest.WashID,
		PostID:        payRequest.PostID,
		Amount:        payRequest.Amount,
		PaymentIDBank: paymentInit.PaymentID,
		Status:        transactionStatus,
	}
	err = svc.repository.CreateTransaction(ctx, transactionCreate)
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
	payResponse := entities.PaymentResponse{
		WashID:  transactionCreate.WashID,
		PostID:  payRequest.PostID,
		OrderID: orderID.String(),
		UrlPay:  urlPay,
	}

	// for tests whithout qr
	// payResponse := entities.PaymentResponse{
	// 	WashID:  transactionCreate.WashID,
	// 	PostID:  payRequest.PostID,
	// 	OrderID: orderID.String(),
	// 	UrlPay:  paymentInit.Url,
	// }
	//
	err = svc.leaWashPublisher.SendToLeaPaymentResponse(payResponse)
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
func (svc *paymentService) ReceiveNotification(ctx context.Context, notification entities.PaymentNotification) error {
	errorPrefix := "Notification error:"
	// get transaction
	orderID, err := uuid.FromString(notification.OrderID)
	if err != nil {
		return fmt.Errorf("%s order_id=%s is not correct, error:%s", errorPrefix, notification.OrderID, err.Error())
	}
	transaction, err := svc.repository.GetTransactionByOrderID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("%s transaction not found by id=%s, error:%s", errorPrefix, orderID, err.Error())
	}

	// get terminal
	washID, err := uuid.FromString(transaction.WashID)
	if err != nil {
		return fmt.Errorf("%s wash_id=%s is not correct, error:%s", errorPrefix, transaction.WashID, err.Error())
	}
	wash, err := svc.repository.GetWashByID(ctx, washID)
	if err != nil {
		return fmt.Errorf("%s wash not found by id=%s", errorPrefix, washID)
	}

	// check notification
	if valid, err := IsNotificationCorrect(notification, wash.TerminalPassword); !valid {
		return fmt.Errorf("%s notification is not correct (wash_id=%s, notification=%+#v), err=%s", errorPrefix, washID, notification, err)
	}

	// update transaction
	transactionStatus := entities.TransactionStatusFromString(notification.Status)
	if transactionStatus == entities.TransactionStatusUnknown {
		svc.logger.Errorf("%s notification status '%s' is unknown", errorPrefix, notification.Status)
	}
	err = svc.repository.UpdateTransaction(ctx, entities.TransactionUpdate{
		ID:            orderID,
		Status:        transactionStatus,
		PaymentIDBank: nil,
	})
	if err != nil {
		return fmt.Errorf("%s update trasaction failed (trasaction_id=%s) error: %s", errorPrefix, notification.OrderID, err.Error())
	}

	// send broker message
	paymentNotifcation := entities.PaymentNotificationForLea{
		WashID:  transaction.WashID,
		PostID:  transaction.PostID,
		OrderID: transaction.ID.String(),
		Status:  notification.Status,
	}
	err = svc.leaWashPublisher.SendToLeaPaymentNotification(paymentNotifcation)
	if err != nil {
		svc.logger.Errorf("%s send notification to lea failed (transaction_id=%s) error: %s", errorPrefix, notification.OrderID, err.Error())
	}

	return nil
}

// Cancel ...
func (svc *paymentService) CancelPayment(ctx context.Context, req entities.PaymentСancellationRequest) (resendNeaded bool, err error) {
	errorPrefix := "Cancel error:"
	resendNeaded = true
	// get transaction by order_id
	orderID, err := uuid.FromString(req.OrderID)
	if err != nil {
		return resendNeaded, fmt.Errorf("%s transaction_id not correct (trasaction_id=%s) error: %s", errorPrefix, req.OrderID, err.Error())
	}
	transaction, err := svc.repository.GetTransactionByOrderID(ctx, orderID)
	if err != nil {
		return resendNeaded, fmt.Errorf("%s get trasaction failed (trasaction_id=%s) error: %s", errorPrefix, req.OrderID, err.Error())
	}

	// get wash by wash_id
	washID, err := uuid.FromString(transaction.WashID)
	if err != nil {
		return resendNeaded, fmt.Errorf("%s wash_id not correct (wash_id=%s) error: %s", errorPrefix, transaction.WashID, err.Error())
	}
	wash, err := svc.repository.GetWashByID(ctx, washID)
	if err != nil {
		return resendNeaded, fmt.Errorf("%s get wash failed (wash_id=%s) error: %s", errorPrefix, washID, err.Error())
	}

	// update transaction status canceling
	if transaction.Status != entities.TransactionStatusСanceling {
		err = svc.repository.UpdateTransaction(ctx, entities.TransactionUpdate{
			ID:            transaction.ID,
			Status:        entities.TransactionStatusСanceling,
			PaymentIDBank: nil,
		})
		if err != nil {
			return resendNeaded, fmt.Errorf("%s set canceling status for trasaction failed (trasaction_id=%s) error: %s", errorPrefix, transaction.ID, err.Error())
		}
	}

	// cancel by pay_сlient
	paymentRegisterNotification := entities.PaymentCreds{
		TerminalKey: wash.TerminalKey,
		PaymentID:   transaction.PaymentIDBank,
	}
	_, err = svc.payClient.CancelPayment(paymentRegisterNotification, wash.TerminalPassword)
	if err != nil {
		return !resendNeaded, fmt.Errorf("%s cancel pay_client_payment failed (trasaction_id=%s) error: %s", errorPrefix, transaction.ID, err.Error())
	}

	//  update transaction status canceled
	err = svc.repository.UpdateTransaction(ctx, entities.TransactionUpdate{
		ID:            transaction.ID,
		Status:        entities.TransactionStatusСanceled,
		PaymentIDBank: nil,
	})
	if err != nil {
		return !resendNeaded, fmt.Errorf("%s set canceled status for trasaction failed (trasaction_id=%s) error: %s", errorPrefix, transaction.ID, err.Error())
	}

	return !resendNeaded, nil
}

// SyncAllPayments ...
func (svc *paymentService) SyncAllPayments(ctx context.Context) error {
	errorPrefix := "SyncAllPayments error:"
	// confirmed not synced
	confirmedNotSyncedStatus := entities.TransactionStatusConfirmedNotSynced
	confirmedNotSyncedTransactions, err := svc.repository.GetTransactionsByStatus(ctx, entities.TransactionsGet{
		Status: confirmedNotSyncedStatus,
	})
	if err != nil {
		return fmt.Errorf("%s get trasactions whith status 'confirmed_not_synced' failed, error: %s", errorPrefix, err.Error())
	}

	for _, pt := range confirmedNotSyncedTransactions {
		// check expiration
		if time.Until(pt.CreatedAt) >= svc.notificationExpirationPeriod {
			// cancel
			err = svc.repository.UpdateTransaction(ctx, entities.TransactionUpdate{
				ID:            pt.ID,
				Status:        entities.TransactionStatusСanceling,
				PaymentIDBank: nil,
			})
			if err != nil {
				return fmt.Errorf("%s set canceling status for trasaction failed (trasaction_id=%s) error: %s", errorPrefix, pt.ID.String(), err.Error())
			}
		}

		// get wash
		wash, err := svc.repository.GetWashByID(ctx, uuid.FromStringOrNil(pt.WashID))
		if err != nil {
			return fmt.Errorf("%s get wash failed (wash_id=%s) error: %s", errorPrefix, pt.WashID, err.Error())
		}

		// send to lea
		paidStatus := string(entities.TransactionStatusConfirmed)
		paymentNotifcation := entities.PaymentNotificationForLea{
			WashID:  wash.ID.String(),
			PostID:  pt.PostID,
			OrderID: pt.ID.String(),
			Status:  paidStatus,
		}

		err = svc.leaWashPublisher.SendToLeaPaymentNotification(paymentNotifcation)
		if err != nil {
			return fmt.Errorf("%s send notification to lea failed (transaction_id=%s) error: %s", errorPrefix, pt.ID.String(), err.Error())
		}

		// update status in db
		err = svc.repository.UpdateTransaction(ctx, entities.TransactionUpdate{
			ID:            pt.ID,
			Status:        entities.TransactionStatusConfirmed,
			PaymentIDBank: nil,
		})
		if err != nil {
			return fmt.Errorf("%s set confirmed status for trasaction failed (trasaction_id=%s) error: %s", errorPrefix, pt.ID.String(), err.Error())
		}
	}

	// canceling not canceled
	cancelingStatus := entities.TransactionStatusСanceling
	cancelingTransactions, err := svc.repository.GetTransactionsByStatus(ctx, entities.TransactionsGet{
		Status: cancelingStatus,
	})
	if err != nil {
		return fmt.Errorf("%s get trasactions whith status 'canceling' failed, error: %s", errorPrefix, err.Error())
	}
	for _, ct := range cancelingTransactions {
		_, err := svc.CancelPayment(ctx, entities.PaymentСancellationRequest{
			OrderID: ct.ID.String(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
