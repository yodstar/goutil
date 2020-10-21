package sqlite3

import (
	"github.com/yodstar/goutil/database/sqlbuilder"
)

var _ = sqlbuilder.Init("sqlite3")

// NewSQLBuilder
func NewSQLBuilder(value interface{}) *sqlbuilder.SQLBuilder {
	return sqlbuilder.NewSQLBuilder(value)
}

// Debug
func Debug(mode bool) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(sqlite3).Debug(mode)
}

// Fields
func Fields(fields string) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(sqlite3).Fields(fields)
}

// Where
func Where(where string, args ...interface{}) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(sqlite3).Where(where, args...)
}

// WhereOr
func WhereOr(where string, args ...interface{}) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(sqlite3).WhereOr(where, args...)
}

// WhereNot
func WhereNot(where string, args ...interface{}) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(sqlite3).WhereNot(where, args...)
}

// GroupBy
func GroupBy(groupBy string) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(sqlite3).GroupBy(groupBy)
}

// Having
func Having(having string) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(sqlite3).Having(having)
}

// OrderBy
func OrderBy(orderBy string) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(sqlite3).OrderBy(orderBy)
}

// Limit
func Limit(offset int) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(sqlite3).Limit(offset)
}

// Model
func Model(value interface{}) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(sqlite3).Model(value)
}
