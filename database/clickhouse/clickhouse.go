package clickhouse

import (
	"database/sql"

	// clickhouse driver
	"github.com/yodstar/goutil/database/sqlbuilder"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/jmoiron/sqlx"
)

// clickhouse
var clickhouse *sqlbuilder.Dao

// MustOpen
func MustOpen(options any, args ...string) {
	clickhouse = sqlbuilder.MustOpen("clickhouse", options)
	if len(args) > 0 {
		sqlbuilder.SetDaoSuffix(args[0])
	}
}

// Writer
func Writer() *sqlbuilder.Db {
	return clickhouse.Writer()
}

// Reader
func Reader() *sqlbuilder.Db {
	return clickhouse.Reader()
}

// Count
func Count(dst any, where string, args ...any) (int64, error) {
	return clickhouse.Count(dst, where, args...)
}

// Select
func Select(dst any, where string, args ...any) error {
	return clickhouse.Select(dst, where, args...)
}

// UnsafeSelect
func UnsafeSelect(dst any, where string, args ...any) error {
	return clickhouse.UnsafeSelect(dst, where, args...)
}

// Delete
func Delete(dst any, where string, args ...any) (sql.Result, error) {
	return clickhouse.Delete(dst, where, args...)
}

// Update
func Update(dst any, where string, args ...any) (sql.Result, error) {
	return clickhouse.Update(dst, where, args...)
}

// Insert
func Insert(dst any) (sql.Result, error) {
	return clickhouse.Insert(dst)
}

// Selectx
func Selectx(dst any, query string, args ...any) error {
	return clickhouse.Reader().Selectx(dst, query, args...)
}

// UnsafeSelectx
func UnsafeSelectx(dst any, query string, args ...any) error {
	return clickhouse.Reader().UnsafeSelectx(dst, query, args...)
}

// Queryx
func Queryx(query string, args ...any) (*sqlx.Rows, error) {
	return clickhouse.Reader().Queryx(query, args...)
}

// UnsafeQueryx
func UnsafeQueryx(query string, args ...any) (*sqlx.Rows, error) {
	return clickhouse.Reader().UnsafeQueryx(query, args...)
}

// QueryRowx
func QueryRowx(query string, args ...any) *sqlx.Row {
	return clickhouse.Reader().QueryRowx(query, args...)
}

// UnsafeQueryRowx
func UnsafeQueryRowx(query string, args ...any) *sqlx.Row {
	return clickhouse.Reader().UnsafeQueryRowx(query, args...)
}

// Exec
func Exec(query string, args ...any) (sql.Result, error) {
	return clickhouse.Writer().Exec(query, args...)
}

// Begin
func Begin() (*sqlbuilder.Db, error) {
	return clickhouse.Writer().Begin()
}

// Transaction
func Transaction(f func(*sqlbuilder.Db) error) error {
	return clickhouse.Writer().Transaction(f)
}
