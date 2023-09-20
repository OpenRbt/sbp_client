package tinkoff

import (
	"fmt"
	logicEntities "sbp/internal/logic/entities"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestIsNotificationCorrect(t *testing.T) {
	password := ""
	logger := &zap.SugaredLogger{}
	payClient, err := NewPayClient(logger, time.Hour)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Тестовые случаи
	testCases := []struct {
		request logicEntities.PaymentNotification
		result  bool
	}{
		{
			request: logicEntities.PaymentNotification{},
			result:  true,
		},
	}

	// Проверка каждого тестового случая
	for _, testCase := range testCases {

		result := payClient.IsNotificationCorrect(testCase.request, password)
		if result != testCase.result {
			t.Errorf("IsNotificationCorrect('%#+v', '%s') = '%t', ожидалось '%t'", testCase.request, password, result, testCase.result)
		}
	}
}
