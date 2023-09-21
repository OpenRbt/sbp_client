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
	password := "b8avo031uimk2fpw"
	logger, _ := zap.NewDevelopment()
	payClient, err := NewPayClient(logger.Sugar(), time.Hour)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Тестовые случаи
	testCases := []struct {
		request []byte
		result  bool
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
			result: true,
		},
	}

	// Проверка каждого тестового случая
	for _, testCase := range testCases {
		var notification logicEntities.PaymentNotification
		err := json.Unmarshal(testCase.request, &notification)
		if err != nil {
			fmt.Println(err.Error())
		}
		result := payClient.IsNotificationCorrect(notification, password)
		if result != testCase.result {
			t.Errorf("IsNotificationCorrect('%#+v', '%s') = '%t', ожидалось '%t'", notification, password, result, testCase.result)
		}
	}
}
