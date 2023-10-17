package tinkoff

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// tokenGenerator ...
type tokenGenerator struct {
	password string
}

func newTokkenGenerator(password string) (*tokenGenerator, error) {
	if password == "" {
		return nil, errors.New("password is empty")
	}
	return &tokenGenerator{
		password: password,
	}, nil
}

// parseRequest ...
func parseRequest(req interface{}, tag string) map[string]string {
	// Получаем тип структуры
	t := reflect.TypeOf(req)

	// Получаем значение структуры
	v := reflect.ValueOf(req)

	// Проходим по всем полям структуры
	resp := make(map[string]string)
	numField := t.NumField()
	for i := 0; i < numField; i++ {
		// Если поле является структурой рекурсивно идем по ней, результат суммируем
		if t.Field(i).Type.Kind() == reflect.Struct {
			subStruct := v.Field(i).Interface()
			r := parseRequest(subStruct, tag)
			for k, v := range r {
				resp[k] = fmt.Sprintf("%v", v)
			}
		}
		// Получаем тег каждого поля
		tag := t.Field(i).Tag.Get(tag)

		// Если тег не пустой, то добавляем его в список полей
		if tag != "" {
			tagAndOpts := strings.Split(tag, ",")
			onlyTag := strings.ToLower(tagAndOpts[0])
			resp[onlyTag] = fmt.Sprint(v.Field(i))
			fmt.Println(onlyTag, ":", fmt.Sprint(v.Field(i)))
		}
	}

	return resp
}

func (g tokenGenerator) generateToken(req interface{}, tag string) string {
	request := parseRequest(req, tag)
	// add password
	request["password"] = g.password
	// delete token attr
	delete(request, "token")
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

// checksumSha256 ...
func checksumSha256(s string) string {
	plainText := []byte(s)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}

// checkToken ...
func (g tokenGenerator) CheckToken(tag string, resp interface{}, token string) bool {
	return g.generateToken(resp, tag) == token
}
