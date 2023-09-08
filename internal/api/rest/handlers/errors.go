package handlers

import (
	logicEntities "sbp/internal/logic/entities"
)

var (
	ErrAccessDenied           = logicEntities.ErrAccessDenied.Error()
	ErrAccessDeniedCode int32 = 403
)
