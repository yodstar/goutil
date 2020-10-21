package dbmysql

import (
	"database/sql"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/yodstar/goutil/database/sqlbuilder"
)

// dbmysql
var dbmysql *sqlbuilder.DB

// MustOpen
func MustOpen(conf []*sqlbuilder.Conf, options ...bool) {
	dbmysql = sqlbuilder.MustOpen("mysql", conf)
	if len(options) > 0 {
		sqlbuilder.SetDebugMode(options[0])
		if len(options) > 1 {
			sqlbuilder.SetModelSuffix(options[1])
		}
	}
}

// Writer
func Writer() *sqlbuilder.Conn {
	return dbmysql.Writer()
}

// Reader
func Reader() *sqlbuilder.Conn {
	return dbmysql.Reader()
}

// Deprecated: Writerx
func Writerx() *sqlx.DB {
	return dbmysql.Writerx()
}

// Deprecated: Readerx
func Readerx() *sqlx.DB {
	return dbmysql.Readerx()
}

// Count
func Count(dest interface{}, where string, args ...interface{}) (int64, error) {
	return dbmysql.Count(dest, where, args...)
}

// Select
func Select(dest interface{}, where string, args ...interface{}) error {
	return dbmysql.Select(dest, where, args...)
}

// Delete
func Delete(dest interface{}, where string, args ...interface{}) (sql.Result, error) {
	return dbmysql.Delete(dest, where, args...)
}

// Update
func Update(dest interface{}, where string, args ...interface{}) (sql.Result, error) {
	return dbmysql.Update(dest, where, args...)
}

// Insert
func Insert(dest interface{}) (sql.Result, error) {
	return dbmysql.Insert(dest)
}

// Selectx
func Selectx(dest interface{}, query string, args ...interface{}) error {
	return dbmysql.Reader().Selectx(dest, query, args...)
}

// Queryx
func Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return dbmysql.Reader().Queryx(query, args...)
}

// QueryRowx
func QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return dbmysql.Reader().QueryRowx(query, args...)
}

// Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	return dbmysql.Writer().Exec(query, args...)
}

// Begin
func Begin() (*sqlbuilder.Conn, error) {
	return dbmysql.Writer().Begin()
}

// Transaction
func Transaction(f func(*sqlbuilder.Conn) error) error {
	return dbmysql.Writer().Transaction(f)
}
