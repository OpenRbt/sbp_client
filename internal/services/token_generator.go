package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"sbp/internal/entities"
	"sort"
	"strings"
)

type TokenGenerator struct {
	password string
}

func IsNotificationCorrect(param entities.PaymentNotification, password string) (bool, error) {
	tokkenGenerator, err := newTokenGenerator(password)
	if err != nil {
		return false, fmt.Errorf("IsNotificationCorrect error: %s", err.Error())
	}

	token := param.Token
	isTokenValid := tokkenGenerator.checkToken("json", param, token)
	return isTokenValid, nil
}

func newTokenGenerator(password string) (*TokenGenerator, error) {
	if password == "" {
		return nil, errors.New("password is empty")
	}
	return &TokenGenerator{
		password: password,
	}, nil
}

func (g TokenGenerator) generateToken(req entities.PaymentNotification, tag string) string {
	request := make(map[string]string)

	// add password
	request["password"] = g.password

	request["terminalKey"] = req.TerminalKey
	request["orderId"] = req.OrderID
	request["success"] = fmt.Sprintf("%t", req.Success)
	request["status"] = req.Status
	request["paymentId"] = fmt.Sprintf("%d", req.PaymentID)
	request["errorCode"] = req.ErrorCode
	request["amount"] = fmt.Sprintf("%d", req.Amount)
	request["pan"] = req.Pan
	if req.ExpDate != nil {
		request["expDate"] = *req.ExpDate
	}
	if req.CardID != nil {
		request["cardId"] = fmt.Sprintf("%d", *req.CardID)
	}

	// Получаем ключи из map
	keys := make([]string, 0, len(request))
	for key := range request {
		keys = append(keys, key)
	}

	// Сортируем ключи в алфавитном порядке
	sort.Strings(keys)

	// Выводим отсортированные ключи и значения
	var correctOrder []string
	for _, key := range keys {
		value := request[key]
		correctOrder = append(correctOrder, value)
	}

	concatenatedData := strings.Join(correctOrder, "")
	return checksumSha256(concatenatedData)
}

func checksumSha256(s string) string {
	plainText := []byte(s)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}

func (g TokenGenerator) checkToken(tag string, resp entities.PaymentNotification, token string) bool {
	return g.generateToken(resp, tag) == token
}
