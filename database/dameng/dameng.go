package dameng

import (
	"database/sql"

	// dameng driver
	_ "github.com/yodstar/goutil/database/dm"

	"github.com/jmoiron/sqlx"
	"github.com/yodstar/goutil/database/sqlbuilder"
)

// dameng
var dameng *sqlbuilder.Dao

// MustOpen
func MustOpen(options any, args ...string) {
	dameng = sqlbuilder.MustOpen("dm", options)
	if len(args) > 0 {
		sqlbuilder.SetDaoSuffix(args[0])
	}
}

// Writer
func Writer() *sqlbuilder.Db {
	return dameng.Writer()
}

// Reader
func Reader() *sqlbuilder.Db {
	return dameng.Reader()
}

// Count
func Count(dst any, where string, args ...any) (int64, error) {
	return dameng.Count(dst, where, args...)
}

// Select
func Select(dst any, where string, args ...any) error {
	return dameng.Select(dst, where, args...)
}

// Delete
func Delete(dst any, where string, args ...any) (sql.Result, error) {
	return dameng.Delete(dst, where, args...)
}

// Update
func Update(dst any, where string, args ...any) (sql.Result, error) {
	return dameng.Update(dst, where, args...)
}

// Insert
func Insert(dst any) (sql.Result, error) {
	return dameng.Insert(dst)
}

// Selectx
func Selectx(dst any, query string, args ...any) error {
	return dameng.Reader().Selectx(dst, query, args...)
}

// Queryx
func Queryx(query string, args ...any) (*sqlx.Rows, error) {
	return dameng.Reader().Queryx(query, args...)
}

// QueryRowx
func QueryRowx(query string, args ...any) *sqlx.Row {
	return dameng.Reader().QueryRowx(query, args...)
}

// Exec
func Exec(query string, args ...any) (sql.Result, error) {
	return dameng.Writer().Exec(query, args...)
}

// Begin
func Begin() (*sqlbuilder.Db, error) {
	return dameng.Writer().Begin()
}

// Transaction
func Transaction(f func(*sqlbuilder.Db) error) error {
	return dameng.Writer().Transaction(f)
}
