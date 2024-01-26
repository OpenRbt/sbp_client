package handlers

import (
	"net/http"
	"sbp/internal/entities"
	"sbp/openapi/models"

	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type errorSetter interface {
	SetPayload(payload *models.Error)
	SetStatusCode(code int)
}

func errorStatusCode(err error) int {
	switch err {
	case entities.ErrBadRequest:
		return http.StatusBadRequest
	case entities.ErrForbidden:
		return http.StatusForbidden
	case entities.ErrNotFound:
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}

var errorMapping = map[error]int{
	entities.ErrBadRequest: http.StatusBadRequest,
	entities.ErrForbidden:  http.StatusForbidden,
	entities.ErrNotFound:   http.StatusNotFound,
}

func setAPIError(l *zap.SugaredLogger, op string, err error, responder interface{}) {
	r, ok := responder.(errorSetter)
	if !ok {
		return
	}

	statusCode, exists := getStatusCodeForError(err)

	msg := err.Error()
	if !exists {
		statusCode = http.StatusInternalServerError
		msg = "internal server error"

		l.Errorln(op, err)
	}

	r.SetPayload(&models.Error{Code: swag.Int32(int32(statusCode)), Message: swag.String(msg)})
	r.SetStatusCode(statusCode)
}

func getStatusCodeForError(err error) (int, bool) {
	for knownErr, code := range errorMapping {
		if errors.Is(err, knownErr) {
			return code, true
		}
	}
	code, exists := errorMapping[err]
	return code, exists
}
