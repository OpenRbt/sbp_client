package repconverter

import (
	logicEntities "sbp/internal/logic/entities"
	repEntities "sbp/internal/repository/entities"
)

// ConvertUserFromDB ...
func ConvertUserFromDB(user repEntities.User) logicEntities.User {
	return logicEntities.User{
		ID:        user.ID.UUID,
		Identity:  user.Identity,
		Role:      logicEntities.UserRoleFromString(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
