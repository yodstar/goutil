package dameng

import (
	"github.com/yodstar/goutil/database/sqlbuilder"
)

var _ = sqlbuilder.Init("dameng")

// NewSqlBuilder
func NewSqlBuilder(value any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.NewSqlBuilder(value)
}

// Fields
func Fields(fields string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dameng).Fields(fields)
}

// Where
func Where(where string, args ...any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dameng).Where(where, args...)
}

// WhereOr
func WhereOr(where string, args ...any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dameng).WhereOr(where, args...)
}

// WhereNot
func WhereNot(where string, args ...any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dameng).WhereNot(where, args...)
}

// GroupBy
func GroupBy(groupBy string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dameng).GroupBy(groupBy)
}

// Having
func Having(having string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dameng).Having(having)
}

// OrderBy
func OrderBy(orderBy string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dameng).OrderBy(orderBy)
}

// Limit
func Limit(offset int) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dameng).Limit(offset)
}

// Model
func Model(value any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dameng).Dao(value)
}
