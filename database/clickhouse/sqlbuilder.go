package clickhouse

import (
	"github.com/yodstar/goutil/database/sqlbuilder"
)

var _ = sqlbuilder.Init("clickhouse")

// NewSqlBuilder
func NewSqlBuilder(value any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.NewSqlBuilder("clickhouse", value)
}

// Fields
func Fields(fields string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(clickhouse).Fields(fields)
}

// Where
func Where(where string, args ...any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(clickhouse).Where(where, args...)
}

// WhereOr
func WhereOr(where string, args ...any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(clickhouse).WhereOr(where, args...)
}

// WhereNot
func WhereNot(where string, args ...any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(clickhouse).WhereNot(where, args...)
}

// GroupBy
func GroupBy(groupBy string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(clickhouse).GroupBy(groupBy)
}

// Having
func Having(having string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(clickhouse).Having(having)
}

// OrderBy
func OrderBy(orderBy string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(clickhouse).OrderBy(orderBy)
}

// Limit
func Limit(offset int) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(clickhouse).Limit(offset)
}

// Dao
func Dao(value any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(clickhouse).Dao(value)
}
