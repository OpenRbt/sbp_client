package app

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func (svc *WashServerSvc) Pay(ctx context.Context, amount int, serverID string, postID string) (string, string, error) {

	transaction, err := svc.repo.NewTransaction(ctx, serverID, postID, amount)
	if err != nil {
		return "", "", err
	}

	id, err := uuid.FromString(serverID)

	if err != nil {
		return "", "", err
	}
	server, err := svc.repo.GetWashServer(ctx, id)
	if err != nil {
		return "", "", err
	}

	mod, err := svc.t.Init(server.TerminalKey, transaction.Amount, transaction.ID.String())

	fmt.Println("Resp model ", mod)
	if err != nil {
		return "", "", err
	}

	err = svc.repo.UpdateTransaction(ctx, transaction.ID, &mod.PaymentID, &mod.Status)
	if err != nil {
		return "", "", err
	}

	token := generateToken(server.TerminalKey, server.TerminalPassword, mod.PaymentID)

	m, err := svc.t.GetQr(server.TerminalKey, mod.PaymentID, token)

	if err != nil {
		return "", "", err
	}
	fmt.Println("Model Response GetQr - ", m)

	return mod.Url, mod.OrderID, nil
}

func (svc *WashServerSvc) Notification(ctx context.Context, notification RegisterNotification) error {

	id, err := uuid.FromString(notification.OrderID)

	if err != nil {
		return err
	}

	transaction, err := svc.repo.GetTransaction(ctx, id)

	if err != nil {
		return err
	}

	serverID, err := uuid.FromString(transaction.ServerID)

	if err != nil {
		return err
	}
	server, err := svc.repo.GetWashServer(ctx, serverID)
	if err != nil {
		return err
	}

	if !checkToken(notification, server.TerminalPassword) {
		return ErrNotification
	}

	svc.repo.UpdateTransaction(ctx, id, nil, &notification.Status)

	return nil
}

func checkToken(n RegisterNotification, password string) bool {
	newToken := ""
	newToken += fmt.Sprint(n.Amount) + fmt.Sprint(n.CardId) + n.ErrorCode + n.ExpDate + n.OrderID +
		n.Pan + password + n.PaymentID + n.Status + fmt.Sprint(n.Success) + n.TerminalKey
	if newToken == n.Token {
		return true
	}
	return false
}

func generateToken(terminalKey string, password string, paymentID string) string {
	predToken := password + paymentID + terminalKey
	plainText := []byte(predToken)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}

func (svc *WashServerSvc) Cancel(ctx context.Context, orderID string) error {

	id, err := uuid.FromString(orderID)

	if err != nil {
		return err
	}

	transaction, err := svc.repo.GetTransaction(ctx, id)

	if err != nil {
		return err
	}

	serverID, err := uuid.FromString(transaction.ServerID)

	if err != nil {
		return err
	}
	server, err := svc.repo.GetWashServer(ctx, serverID)
	if err != nil {
		return err
	}

	token := generateToken(server.TerminalKey, server.TerminalPassword, transaction.PaymentID)

	mod, err := svc.t.Cancel(server.TerminalKey, transaction.PaymentID, token)
	if err != nil {
		return err
	}

	fmt.Println("Response model cancel - ", mod)

	svc.repo.UpdateTransaction(ctx, transaction.ID, nil, &mod.Status)

	return nil
}
