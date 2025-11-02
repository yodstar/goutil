package oracle

import (
	"github.com/yodstar/goutil/database/sqlbuilder"
)

var _ = sqlbuilder.Init("oracle")

// NewSqlBuilder
func NewSqlBuilder(value any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.NewSqlBuilder(value)
}

// Fields
func Fields(fields string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(oracle).Fields(fields)
}

// Where
func Where(where string, args ...any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(oracle).Where(where, args...)
}

// WhereOr
func WhereOr(where string, args ...any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(oracle).WhereOr(where, args...)
}

// WhereNot
func WhereNot(where string, args ...any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(oracle).WhereNot(where, args...)
}

// GroupBy
func GroupBy(groupBy string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(oracle).GroupBy(groupBy)
}

// Having
func Having(having string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(oracle).Having(having)
}

// OrderBy
func OrderBy(orderBy string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(oracle).OrderBy(orderBy)
}

// Limit
func Limit(offset int) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(oracle).Limit(offset)
}

// Dao
func Dao(value any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(oracle).Dao(value)
}
