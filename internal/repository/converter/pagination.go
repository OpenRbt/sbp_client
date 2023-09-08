package repconverter

import (
	logicEntities "sbp/internal/logic/entities"
	repEntities "sbp/internal/repository/entities"
)

// ConvertPaginationToDB ...
func ConvertPaginationToDB(entity logicEntities.Pagination) repEntities.Pagination {
	return repEntities.Pagination{
		Limit:  entity.Limit,
		Offset: entity.Offset,
	}
}
