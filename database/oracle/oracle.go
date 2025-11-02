package oracle

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/yodstar/goutil/database/sqlbuilder"

	// oci8 driver
	_ "github.com/mattn/go-oci8"
)

// oracle
var oracle *sqlbuilder.Dao

// MustOpen
func MustOpen(options any, args ...string) {
	oracle = sqlbuilder.MustOpen("oci8", options)
	if len(args) > 0 {
		sqlbuilder.SetDaoSuffix(args[0])
	}
}

// Writer
func Writer() *sqlbuilder.Db {
	return oracle.Writer()
}

// Reader
func Reader() *sqlbuilder.Db {
	return oracle.Reader()
}

// Count
func Count(dst any, where string, args ...any) (int64, error) {
	return oracle.Count(dst, where, args...)
}

// Select
func Select(dst any, where string, args ...any) error {
	return oracle.Select(dst, where, args...)
}

// Delete
func Delete(dst any, where string, args ...any) (sql.Result, error) {
	return oracle.Delete(dst, where, args...)
}

// Update
func Update(dst any, where string, args ...any) (sql.Result, error) {
	return oracle.Update(dst, where, args...)
}

// Insert
func Insert(dst any) (sql.Result, error) {
	return oracle.Insert(dst)
}

// Selectx
func Selectx(dst any, query string, args ...any) error {
	return oracle.Reader().Selectx(dst, query, args...)
}

// Queryx
func Queryx(query string, args ...any) (*sqlx.Rows, error) {
	return oracle.Reader().Queryx(query, args...)
}

// QueryRowx
func QueryRowx(query string, args ...any) *sqlx.Row {
	return oracle.Reader().QueryRowx(query, args...)
}

// Exec
func Exec(query string, args ...any) (sql.Result, error) {
	return oracle.Writer().Exec(query, args...)
}

// Begin
func Begin() (*sqlbuilder.Db, error) {
	return oracle.Writer().Begin()
}

// Transaction
func Transaction(f func(*sqlbuilder.Db) error) error {
	return oracle.Writer().Transaction(f)
}
