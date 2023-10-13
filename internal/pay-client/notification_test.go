package tinkoff

import (
	"encoding/json"
	"fmt"
	logicEntities "sbp/internal/logic/entities"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestIsNotificationCorrect(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	payClient, err := NewPayClient(logger.Sugar(), time.Hour)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Тестовые случаи
	testCases := []struct {
		request  []byte
		result   bool
		password string
	}{
		{
			request: []byte(`
			{
				"TerminalKey": "1683729350126",
				"OrderId": "31577bad-951c-e310-60f9-fd38f6910df0",
				"Success": true,
				"Status": "AUTHORIZED",
				"PaymentId": 3269844626,
				"ErrorCode": "0",
				"Amount": 1000,
				"Pan": "+7 (916) ***-**-21",
				"Token": "439b0b1305cd54852e39622619362100eae57347cb29e7eee8ea2f55140252af"
			}
			`),
			result:   true,
			password: "b8avo031uimk2fpw",
		},
		{
			request: []byte(`
			{
				"TerminalKey": "1696908061914DEMO",
				"OrderId": "a24a9ca2-e323-49aa-bff6-493e5deebc0c",
				"Success": true,
				"Status": "AUTHORIZED",
				"PaymentId": 3367515924,
				"ErrorCode": "0",
				"Amount": 5000,
				"Pan": "430000******0777",
				"Token": "188ef1adacba53d889aa427184aaa47b7ee7eb306427dc5ab3dc1a7f87c4b7d5"
			}
			`),
			result:   true,
			password: "7wqy8w821dqbldev",
		},
	}

	// Проверка каждого тестового случая
	for _, testCase := range testCases {
		var notification logicEntities.PaymentNotification
		err := json.Unmarshal(testCase.request, &notification)
		if err != nil {
			fmt.Println(err.Error())
		}
		result := payClient.IsNotificationCorrect(notification, testCase.password)
		if result != testCase.result {
			t.Errorf("IsNotificationCorrect('%#+v', '%s') = '%t', ожидалось '%t'", notification, testCase.password, result, testCase.result)
		}
	}
}
