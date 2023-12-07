package repository

import "github.com/Masterminds/squirrel"

var SQ = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
