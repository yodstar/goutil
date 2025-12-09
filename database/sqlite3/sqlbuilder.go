package sqlite3

import (
	"github.com/yodstar/goutil/database/sqlbuilder"
)

var _ = sqlbuilder.Init("sqlite3")

// NewSqlBuilder
func NewSqlBuilder(value any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.NewSqlBuilder("sqlite3", value)
}

// Fields
func Fields(fields string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(sqlite3).Fields(fields)
}

// Where
func Where(where string, args ...any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(sqlite3).Where(where, args...)
}

// WhereOr
func WhereOr(where string, args ...any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(sqlite3).WhereOr(where, args...)
}

// WhereNot
func WhereNot(where string, args ...any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(sqlite3).WhereNot(where, args...)
}

// GroupBy
func GroupBy(groupBy string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(sqlite3).GroupBy(groupBy)
}

// Having
func Having(having string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(sqlite3).Having(having)
}

// OrderBy
func OrderBy(orderBy string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(sqlite3).OrderBy(orderBy)
}

// Limit
func Limit(offset int) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(sqlite3).Limit(offset)
}

// Dao
func Dao(value any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(sqlite3).Dao(value)
}
