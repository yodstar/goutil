package dbmysql

import (
	"github.com/yodstar/goutil/database/sqlbuilder"
)

var _ = sqlbuilder.Init("dbmysql")

// NewSqlBuilder
func NewSqlBuilder(value any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.NewSqlBuilder("mysql", value)
}

// Fields
func Fields(fields string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dbmysql).Fields(fields)
}

// Where
func Where(where string, args ...any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dbmysql).Where(where, args...)
}

// WhereOr
func WhereOr(where string, args ...any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dbmysql).WhereOr(where, args...)
}

// WhereNot
func WhereNot(where string, args ...any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dbmysql).WhereNot(where, args...)
}

// GroupBy
func GroupBy(groupBy string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dbmysql).GroupBy(groupBy)
}

// Having
func Having(having string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dbmysql).Having(having)
}

// OrderBy
func OrderBy(orderBy string) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dbmysql).OrderBy(orderBy)
}

// Limit
func Limit(offset int) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dbmysql).Limit(offset)
}

// Dao
func Dao(value any) *sqlbuilder.SqlBuilder {
	return sqlbuilder.DbSqlBuilder(dbmysql).Dao(value)
}
