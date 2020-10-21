package oracle

import (
	"github.com/yodstar/goutil/database/sqlbuilder"
)

var _ = sqlbuilder.Init("oracle")

// NewSQLBuilder
func NewSQLBuilder(value interface{}) *sqlbuilder.SQLBuilder {
	return sqlbuilder.NewSQLBuilder(value)
}

// Debug
func Debug(mode bool) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(oracle).Debug(mode)
}

// Fields
func Fields(fields string) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(oracle).Fields(fields)
}

// Where
func Where(where string, args ...interface{}) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(oracle).Where(where, args...)
}

// WhereOr
func WhereOr(where string, args ...interface{}) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(oracle).WhereOr(where, args...)
}

// WhereNot
func WhereNot(where string, args ...interface{}) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(oracle).WhereNot(where, args...)
}

// GroupBy
func GroupBy(groupBy string) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(oracle).GroupBy(groupBy)
}

// Having
func Having(having string) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(oracle).Having(having)
}

// OrderBy
func OrderBy(orderBy string) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(oracle).OrderBy(orderBy)
}

// Limit
func Limit(offset int) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(oracle).Limit(offset)
}

// Model
func Model(value interface{}) *sqlbuilder.SQLBuilder {
	return sqlbuilder.DBSQLBuilder(oracle).Model(value)
}
