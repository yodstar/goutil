package dbmysql

import (
	"github.com/yodstar/goutil/database/sqlbuilder"
)

var _ = sqlbuilder.Init("dbmysql")

// NewSQLBuilder
func NewSQLBuilder(value interface{}) *sqlbuilder.SQLBuilder {
	return sqlbuilder.NewSQLBuilder(value)
}

// Debug
func Debug(mode bool) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(dbmysql).Debug(mode)
}

// Fields
func Fields(fields string) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(dbmysql).Fields(fields)
}

// Where
func Where(where string, args ...interface{}) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(dbmysql).Where(where, args...)
}

// WhereOr
func WhereOr(where string, args ...interface{}) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(dbmysql).WhereOr(where, args...)
}

// WhereNot
func WhereNot(where string, args ...interface{}) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(dbmysql).WhereNot(where, args...)
}

// GroupBy
func GroupBy(groupBy string) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(dbmysql).GroupBy(groupBy)
}

// Having
func Having(having string) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(dbmysql).Having(having)
}

// OrderBy
func OrderBy(orderBy string) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(dbmysql).OrderBy(orderBy)
}

// Limit
func Limit(offset int) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(dbmysql).Limit(offset)
}

// Model
func Model(value interface{}) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(dbmysql).Model(value)
}
