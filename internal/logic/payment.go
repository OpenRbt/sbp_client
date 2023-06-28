package logic

import (
	"context"
	"crypto/rand"
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
	washServerLogic              *WashServerLogic
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
	SendToLea(serviceKey string, messageType string, messageStruct interface{}) error
	SendToLeaError(serviceKey string, serverID string, postID string, orderID string, errorDesc string, errorCode int64) error
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
	washServerLogic *WashServerLogic,
) (*PaymentLogic, error) {

	return &PaymentLogic{
		logger:                       logger,
		notificationExpirationPeriod: notificationExpirationPeriod,
		payClient:                    payClient,
		repository:                   repository,
		washServerLogic:              washServerLogic,
		leaWashPublisher:             leaWashPublisher,
	}, nil
}

// Pay ...
func (logic *PaymentLogic) Pay(ctx context.Context, payRequest logicEntities.PayRequest) (*logicEntities.PayResponse, error) {
	// get server uuid
	id, err := uuid.FromString(payRequest.ServerID)
	if err != nil {
		return nil, err
	}

	// get wash server terminal
	server, err := logic.washServerLogic.GetWashServer(ctx, id)
	if err != nil {
		return nil, err
	}

	var transactionID uuid.UUID
	if payRequest.OrderID != "" {
		transactionID = uuid.FromStringOrNil(payRequest.OrderID)
	}

	if transactionID == uuid.Nil {
		// generate transaction_id
		transactionID, err = generateTransactionID()
		if err != nil {
			return nil, err
		}
	}

	// payment init
	paymentCreate := logicEntities.PaymentCreate{
		TerminalKey: server.TerminalKey,
		Amount:      payRequest.Amount,
		OrderID:     transactionID.String(),
	}
	paymentInit, err := logic.payClient.Init(paymentCreate)
	if err != nil {
		return nil, err
	}

	// get QR code
	// paymentCreds := logicEntities.PaymentCreds{
	// 	TerminalKey: server.TerminalKey,
	// 	PaymentID:   paymentInit.PaymentID,
	// }
	// resp, err := logic.payClient.GetQr(paymentCreds, server.TerminalPassword)
	// if err != nil {
	// 	return nil, err
	// }
	// if resp.ErrorCode != "0" {
	// 	return nil, fmt.Errorf("pay client internal error: %s, %s", resp.ErrorCode, resp.Message)
	// }

	// add payment to db
	transactionStatus, err := logicEntities.TransactionStatusFromString(paymentInit.Status)
	if err != nil {
		logic.logger.Error(err)
	}
	transactionCreate := logicEntities.TransactionCreate{
		ID:        transactionID,
		ServerID:  payRequest.ServerID,
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
	// payResponse := &logicEntities.PayResponse{
	//  PostID:  payRequest.PostID,
	// 	OrderID: resp.OrderID,
	// 	UrlPay:  resp.UrlPay,
	// }
	payResponse := &logicEntities.PayResponse{
		PostID:  payRequest.PostID,
		OrderID: transactionID.String(),
		UrlPay:  paymentInit.Url,
	}
	//
	messageTypePaymentResponse := string(logicEntities.MessageTypePaymentResponse)
	err = logic.leaWashPublisher.SendToLea(server.ServiceKey, messageTypePaymentResponse, payResponse)
	if err != nil {
		return nil, err
	}

	return payResponse, nil
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
	serverID, err := uuid.FromString(transaction.ServerID)
	if err != nil {
		return err
	}
	server, err := logic.washServerLogic.GetWashServer(ctx, serverID)
	if err != nil {
		return err
	}

	// check notification
	if !logic.payClient.IsNotificationCorrect(notification, server.TerminalPassword) {
		return logicEntities.ErrNotification
	}

	// update transaction
	transactionStatus, err := logicEntities.TransactionStatusFromString(notification.Status)
	if err != nil {
		logic.logger.Error(err)
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
	payNotifcation := logicEntities.PayNotifcation{
		ServerID: transaction.ServerID,
		PostID:   transaction.PostID,
		OrderID:  transaction.ID.String(),
		Status:   notification.Status,
	}
	messageTypePaymentNotification := string(logicEntities.MessageTypePaymentNotification)
	err = logic.leaWashPublisher.SendToLea(server.ServiceKey, messageTypePaymentNotification, payNotifcation)
	if err != nil {
		logic.logger.Error(err)
	}

	return nil
}

// Cancel ...
func (logic *PaymentLogic) Cancel(ctx context.Context, req logicEntities.PayСancellationRequest) (resendNeaded bool, err error) {

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

	// get server by server_id
	serverID, err := uuid.FromString(transaction.ServerID)
	if err != nil {
		return resendNeaded, err
	}
	server, err := logic.washServerLogic.GetWashServer(ctx, serverID)
	if err != nil {
		return resendNeaded, err
	}

	// update transaction status canceling
	if transaction.Status != string(logicEntities.TransactionStatusСanceling) {
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
		TerminalKey: server.TerminalKey,
		PaymentID:   transaction.PaymentID,
	}
	_, err = logic.payClient.Cancel(paymentRegisterNotification, server.TerminalPassword)
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
		if time.Until(pt.DataCreate) >= logic.notificationExpirationPeriod {
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

		// get server
		server, err := logic.washServerLogic.GetWashServer(ctx, uuid.FromStringOrNil(pt.ServerID))
		if err != nil {
			return err
		}

		// send to lea
		paidStatus := string(logicEntities.TransactionStatusConfirmed)
		payNotifcation := logicEntities.PayNotifcation{
			ServerID: pt.ServerID,
			PostID:   pt.PostID,
			OrderID:  pt.ID.String(),
			Status:   paidStatus,
		}
		serviceKey := server.ServiceKey
		messageTypePaymentNotification := string(logicEntities.MessageTypePaymentNotification)
		err = logic.leaWashPublisher.SendToLea(serviceKey, messageTypePaymentNotification, payNotifcation)
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
		_, err := logic.Cancel(ctx, logicEntities.PayСancellationRequest{
			OrderID: ct.ID.String(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// generateTransactionID ...
func generateTransactionID() (uuid.UUID, error) {
	// Генерируем случайные байты для UUID
	randomBytes := make([]byte, 15)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return uuid.UUID{}, err
	}

	// Получаем текущее время
	now := time.Now()

	// Преобразуем текущее время в байты
	timeBytes := now.UTC().UnixNano()

	// Копируем байты времени в начало случайных байтов
	randomBytes = append(randomBytes, byte(timeBytes))

	// Создаем UUID из байтов
	uuid, err := uuid.FromBytes(randomBytes)
	if err != nil {
		return uuid, err
	}

	return uuid, nil
}
