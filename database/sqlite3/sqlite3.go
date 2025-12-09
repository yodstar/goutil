package sqlite3

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/yodstar/goutil/database/sqlbuilder"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// sqlite3
var sqlite3 *sqlbuilder.Dao

// MustOpen
func MustOpen(options any, args ...string) {
	sqlite3 = sqlbuilder.MustOpen("sqlite3", options)
	if len(args) > 0 {
		sqlbuilder.SetDaoSuffix(args[0])
	}
}

// Count
func Count(dst any, where string, args ...any) (int64, error) {
	return sqlite3.Count(dst, where, args...)
}

// Select
func Select(dst any, where string, args ...any) error {
	return sqlite3.Select(dst, where, args...)
}

// UnsafeSelect
func UnsafeSelect(dst any, where string, args ...any) error {
	return sqlite3.UnsafeSelect(dst, where, args...)
}

// Delete
func Delete(dst any, where string, args ...any) (sql.Result, error) {
	return sqlite3.Delete(dst, where, args...)
}

// Update
func Update(dst any, where string, args ...any) (sql.Result, error) {
	return sqlite3.Update(dst, where, args...)
}

// Insert
func Insert(dst any) (sql.Result, error) {
	return sqlite3.Insert(dst)
}

// Selectx
func Selectx(dst any, query string, args ...any) error {
	return sqlite3.Reader().Selectx(dst, query, args...)
}

// UnsafeSelectx
func UnsafeSelectx(dst any, query string, args ...any) error {
	return sqlite3.Reader().UnsafeSelectx(dst, query, args...)
}

// Queryx
func Queryx(query string, args ...any) (*sqlx.Rows, error) {
	return sqlite3.Reader().Queryx(query, args...)
}

// UnsafeQueryx
func UnsafeQueryx(query string, args ...any) (*sqlx.Rows, error) {
	return sqlite3.Reader().UnsafeQueryx(query, args...)
}

// QueryRowx
func QueryRowx(query string, args ...any) *sqlx.Row {
	return sqlite3.Reader().QueryRowx(query, args...)
}

// UnsafeQueryRowx
func UnsafeQueryRowx(query string, args ...any) *sqlx.Row {
	return sqlite3.Reader().UnsafeQueryRowx(query, args...)
}

// Exec
func Exec(query string, args ...any) (sql.Result, error) {
	return sqlite3.Writer().Exec(query, args...)
}

// Begin
func Begin() (*sqlbuilder.Db, error) {
	return sqlite3.Writer().Begin()
}

// Transaction
func Transaction(f func(*sqlbuilder.Db) error) error {
	return sqlite3.Writer().Transaction(f)
}
