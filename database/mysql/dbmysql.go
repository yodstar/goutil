package dbmysql

import (
	"database/sql"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/yodstar/goutil/database/sqlbuilder"
)

// dbmysql
var dbmysql *sqlbuilder.Dao

// MustOpen
func MustOpen(options any, args ...string) {
	dbmysql = sqlbuilder.MustOpen("mysql", options)
	if len(args) > 0 {
		sqlbuilder.SetDaoSuffix(args[0])
	}
}

// Writer
func Writer() *sqlbuilder.Db {
	return dbmysql.Writer()
}

// Reader
func Reader() *sqlbuilder.Db {
	return dbmysql.Reader()
}

// Count
func Count(dst any, where string, args ...any) (int64, error) {
	return dbmysql.Count(dst, where, args...)
}

// Select
func Select(dst any, where string, args ...any) error {
	return dbmysql.Select(dst, where, args...)
}

// UnsafeSelect
func UnsafeSelect(dst any, where string, args ...any) error {
	return dbmysql.UnsafeSelect(dst, where, args...)
}

// Delete
func Delete(dst any, where string, args ...any) (sql.Result, error) {
	return dbmysql.Delete(dst, where, args...)
}

// Update
func Update(dst any, where string, args ...any) (sql.Result, error) {
	return dbmysql.Update(dst, where, args...)
}

// Insert
func Insert(dst any) (sql.Result, error) {
	return dbmysql.Insert(dst)
}

// Selectx
func Selectx(dst any, query string, args ...any) error {
	return dbmysql.Reader().Selectx(dst, query, args...)
}

// UnsafeSelectx
func UnsafeSelectx(dst any, query string, args ...any) error {
	return dbmysql.Reader().UnsafeSelectx(dst, query, args...)
}

// Queryx
func Queryx(query string, args ...any) (*sqlx.Rows, error) {
	return dbmysql.Reader().Queryx(query, args...)
}

// UnsafeQueryx
func UnsafeQueryx(query string, args ...any) (*sqlx.Rows, error) {
	return dbmysql.Reader().UnsafeQueryx(query, args...)
}

// QueryRowx
func QueryRowx(query string, args ...any) *sqlx.Row {
	return dbmysql.Reader().QueryRowx(query, args...)
}

// UnsafeQueryRowx
func UnsafeQueryRowx(query string, args ...any) *sqlx.Row {
	return dbmysql.Reader().UnsafeQueryRowx(query, args...)
}

// Exec
func Exec(query string, args ...any) (sql.Result, error) {
	return dbmysql.Writer().Exec(query, args...)
}

// Begin
func Begin() (*sqlbuilder.Db, error) {
	return dbmysql.Writer().Begin()
}

// Transaction
func Transaction(f func(*sqlbuilder.Db) error) error {
	return dbmysql.Writer().Transaction(f)
}
