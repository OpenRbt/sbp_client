package tinkoff

import (
	"fmt"
	"testing"
)

// TestRequest1
type TestRequest1 struct {
	A int    `json:"A,omitempty"`
	B string `json:"Token,omitempty"`
}

// TestRequest2
type TestRequest2 struct {
	B int64 `json:"B,omitempty"`
	A bool  `json:"A,omitempty"`
}

// TestRequest3
type TestRequest3 struct {
	X string `json:"X,omitempty"`
	TestRequest2
}

// TestGenerateToken ...
func TestGenerateToken(t *testing.T) {
	password := "password"

	// Тестовые случаи
	testCases := []struct {
		request interface{}
		result  string
	}{
		{
			request: TestRequest1{
				A: 1,
				B: "2",
			},
			result: checksumSha256(fmt.Sprintf("1%s", password)),
		},
		{
			request: TestRequest2{
				B: 2,
				A: true,
			},
			result: checksumSha256(fmt.Sprintf("true2%s", password)),
		},
		{
			request: TestRequest3{
				X: "3",
				TestRequest2: TestRequest2{
					B: 2,
					A: true,
				},
			},
			result: checksumSha256(fmt.Sprintf("true2%s3", password)),
		},
	}
	tokenGenerator, err := NewTokkenGenerator(password)
	if err != nil {
		t.Errorf("tokenGenerator failed init: %s", err.Error())
	}

	// Проверка каждого тестового случая
	for _, testCase := range testCases {
		result := tokenGenerator.generateToken(testCase.request, "json")
		if result != testCase.result {
			t.Errorf("generateToken('%#+v', '%s') = '%s', ожидалось '%s'", testCase.request, password, result, testCase.result)
		}
	}
}
