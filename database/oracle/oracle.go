package oracle

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	// oci8 driver
	_ "github.com/mattn/go-oci8"
	"github.com/yodstar/goutil/database/sqlbuilder"
)

// oracle
var oracle *sqlbuilder.DB

// MustOpen
func MustOpen(dbConfig []*sqlbuilder.Conf, options ...bool) {
	oracle = sqlbuilder.MustOpen("oci8", dbConfig)
	if len(options) > 0 {
		sqlbuilder.SetDebugMode(options[0])
		if len(options) > 1 {
			sqlbuilder.SetModelSuffix(options[1])
		}
	}
}

// Writer
func Writer() *sqlbuilder.Conn {
	return oracle.Writer()
}

// Reader
func Reader() *sqlbuilder.Conn {
	return oracle.Reader()
}

// Count
func Count(dest interface{}, where string, args ...interface{}) (int64, error) {
	return oracle.Count(dest, where, args...)
}

// Select
func Select(dest interface{}, where string, args ...interface{}) error {
	return oracle.Select(dest, where, args...)
}

// Delete
func Delete(dest interface{}, where string, args ...interface{}) (sql.Result, error) {
	return oracle.Delete(dest, where, args...)
}

// Update
func Update(dest interface{}, where string, args ...interface{}) (sql.Result, error) {
	return oracle.Update(dest, where, args...)
}

// Insert
func Insert(dest interface{}) (sql.Result, error) {
	return oracle.Insert(dest)
}

// Selectx
func Selectx(dest interface{}, query string, args ...interface{}) error {
	return oracle.Reader().Selectx(dest, query, args...)
}

// Queryx
func Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return oracle.Reader().Queryx(query, args...)
}

// QueryRowx
func QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return oracle.Reader().QueryRowx(query, args...)
}

// Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	return oracle.Writer().Exec(query, args...)
}

// Begin
func Begin() (*sqlbuilder.Conn, error) {
	return oracle.Writer().Begin()
}

// Transaction
func Transaction(f func(*sqlbuilder.Conn) error) error {
	return oracle.Writer().Transaction(f)
}
