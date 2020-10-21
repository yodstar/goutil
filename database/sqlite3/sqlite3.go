package sqlite3

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/yodstar/goutil/database/sqlbuilder"
)

// sqlite3
var sqlite3 *sqlbuilder.DB

// MustOpen
func MustOpen(conf []*sqlbuilder.Conf, options ...bool) {
	sqlite3 = sqlbuilder.MustOpen("sqlite3", conf)
	if len(options) > 0 {
		sqlbuilder.SetDebugMode(options[0])
		if len(options) > 1 {
			sqlbuilder.SetModelSuffix(options[1])
		}
	}
}

// Count
func Count(dest interface{}, where string, args ...interface{}) (int64, error) {
	return sqlite3.Count(dest, where, args...)
}

// Select
func Select(dest interface{}, where string, args ...interface{}) error {
	return sqlite3.Select(dest, where, args...)
}

// Delete
func Delete(dest interface{}, where string, args ...interface{}) (sql.Result, error) {
	return sqlite3.Delete(dest, where, args...)
}

// Update
func Update(dest interface{}, where string, args ...interface{}) (sql.Result, error) {
	return sqlite3.Update(dest, where, args...)
}

// Insert
func Insert(dest interface{}) (sql.Result, error) {
	return sqlite3.Insert(dest)
}

// Selectx
func Selectx(dest interface{}, query string, args ...interface{}) error {
	return sqlite3.Reader().Selectx(dest, query, args...)
}

// Queryx
func Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return sqlite3.Reader().Queryx(query, args...)
}

// QueryRowx
func QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return sqlite3.Reader().QueryRowx(query, args...)
}

// Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	return sqlite3.Writer().Exec(query, args...)
}

// Begin
func Begin() (*sqlbuilder.Conn, error) {
	return sqlite3.Writer().Begin()
}

// Transaction
func Transaction(f func(*sqlbuilder.Conn) error) error {
	return sqlite3.Writer().Transaction(f)
}
