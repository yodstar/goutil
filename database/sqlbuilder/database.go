package sqlbuilder

import (
	"database/sql"
	"log"
	"math/rand"
	"strings"

	"github.com/jmoiron/sqlx"
)

// Option
type Option struct {
	DataSourceName string
	MaxOpenConns   int
	MaxIdleConns   int
	IsDebugMode    bool
}

// Options
type Options []*Option

// Dao
type Dao struct {
	unsafe bool
	db     []*Db
	n      int
}

// MustOpen
func MustOpen(driverName string, option any) *Dao {
	dao := &Dao{}
	if options, ok := option.(Options); ok {
		dao.db = make([]*Db, len(options))
		for i, v := range options {
			x := &Db{r: sqlx.MustOpen(driverName, v.DataSourceName)}
			if v.MaxOpenConns > 0 {
				x.r.SetMaxOpenConns(v.MaxOpenConns)
			}
			if v.MaxIdleConns > 0 {
				x.r.SetMaxIdleConns(v.MaxIdleConns)
			}
			x.isDebugMode = v.IsDebugMode
			x.r.Ping()
			if i == 0 {
				x.w = x.r
			}
			dao.db[i] = x
		}
	} else if v, ok := option.(Option); ok {
		dao.db = make([]*Db, 1)
		x := &Db{r: sqlx.MustOpen(driverName, v.DataSourceName)}
		if v.MaxOpenConns > 0 {
			x.r.SetMaxOpenConns(v.MaxOpenConns)
		}
		if v.MaxIdleConns > 0 {
			x.r.SetMaxIdleConns(v.MaxIdleConns)
		}
		x.isDebugMode = v.IsDebugMode
		x.r.Ping()
		x.w = x.r
		dao.db[0] = x
	}
	dao.n = len(dao.db)
	return dao
}

// Dao.Writer
func (d *Dao) Unsafe(ok bool) *Dao {
	d.unsafe = ok
	return d
}

// Dao.Writer
func (d *Dao) Writer() *Db {
	if d == nil || d.n == 0 {
		log.Panicln(errNotBindDb)
	}
	return d.db[0]
}

// Dao.Reader
func (d *Dao) Reader() *Db {
	if d == nil || d.n == 0 {
		log.Panicln(errNotBindDb)
	}
	return d.db[rand.Intn(d.n)]
}

// Dao.Table
func (d *Dao) Table(tableName string) *SqlBuilder {
	return DbSqlBuilder(d).Table(tableName)
}

// Dao.Fields
func (d *Dao) Fields(fields string) *SqlBuilder {
	return DbSqlBuilder(d).Fields(fields)
}

// Dao.Where
func (d *Dao) Where(where string, args ...any) *SqlBuilder {
	return DbSqlBuilder(d).Where(where, args...)
}

// Dao.WhereOr
func (d *Dao) WhereOr(where string, args ...any) *SqlBuilder {
	return DbSqlBuilder(d).WhereOr(where, args...)
}

// Dao.WhereNot
func (d *Dao) WhereNot(where string, args ...any) *SqlBuilder {
	return DbSqlBuilder(d).WhereNot(where, args...)
}

// Dao.GroupBy
func (d *Dao) GroupBy(group string) *SqlBuilder {
	return DbSqlBuilder(d).GroupBy(group)
}

// Dao.Having
func (d *Dao) Having(having string) *SqlBuilder {
	return DbSqlBuilder(d).Having(having)
}

// Dao.OrderBy
func (d *Dao) OrderBy(order string) *SqlBuilder {
	return DbSqlBuilder(d).OrderBy(order)
}

// Dao.Limit
func (d *Dao) Limit(offset int) *SqlBuilder {
	return DbSqlBuilder(d).Limit(offset)
}

// Dao.Count
func (d *Dao) Count(dst any, where string, args ...any) (int64, error) {
	return d.Reader().Count(dst, where, args...)
}

// Dao.Select
func (d *Dao) Select(dst any, where string, args ...any) error {
	return d.Reader().Select(dst, where, args...)
}

// Dao.Update
func (d *Dao) Update(dst any, where string, args ...any) (sql.Result, error) {
	return d.Writer().Update(dst, where, args...)
}

// Dao.Delete
func (d *Dao) Delete(dst any, where string, args ...any) (sql.Result, error) {
	return d.Writer().Delete(dst, where, args...)
}

// Dao.Insert
func (d *Dao) Insert(dst any) (sql.Result, error) {
	return d.Writer().Insert(dst)
}

// Dao.Queryx
func (d *Dao) Queryx(query string, args ...any) (*sqlx.Rows, error) {
	return d.Reader().Queryx(query, args...)
}

// Dao.QueryRowx
func (d *Dao) QueryRowx(query string, args ...any) *sqlx.Row {
	return d.Reader().QueryRowx(query, args...)
}

// Dao.Selectx
func (d *Dao) Selectx(dst any, query string, args ...any) error {
	return d.Reader().Selectx(dst, query, args...)
}

// Dao.Exec
func (d *Dao) Exec(query string, args ...any) (sql.Result, error) {
	return d.Writer().Exec(query, args...)
}

// Dao.Begin
func (d *Dao) Begin() (*Db, error) {
	return d.Writer().Begin()
}

// Dao.Transaction
func (d *Dao) Transaction(f func(*Db) error) error {
	return d.Writer().Transaction(f)
}

// Db
type Db struct {
	r *sqlx.DB
	w *sqlx.DB
	t *sqlx.Tx

	isDebugMode bool
}

// Db.Ping
func (x *Db) Ping() error {
	return x.r.Ping()
}

// Db.Count
func (x *Db) Count(dst any, where string, args ...any) (int64, error) {
	sb := NewSqlBuilder(dst)
	query, args, err := sb.buildCountSQL(where, args...)
	if err != nil {
		return 0, err
	}
	var count int64
	err = x.QueryRowx(query, args...).Scan(&count)
	return count, err
}

// Db.Select
func (x *Db) Select(dst any, where string, args ...any) error {
	sb := NewSqlBuilder(dst)
	query, args, err := sb.buildSelectSQL(where, args...)
	if err != nil {
		return err
	}
	if sb.IsSliceValue {
		err = x.Selectx(dst, query, args...)
	} else {
		err = x.QueryRowx(query, args...).StructScan(dst)
	}
	return err
}

// Db.Update
func (x *Db) Update(dst any, where string, args ...any) (sql.Result, error) {
	sb := NewSqlBuilder(dst)
	query, args, err := sb.buildUpdateSQL(where, args...)
	if err != nil {
		return nil, err
	}
	return x.Exec(query, args...)
}

// Db.Delete
func (x *Db) Delete(dst any, where string, args ...any) (sql.Result, error) {
	sb := NewSqlBuilder(dst)
	query, args, err := sb.buildDeleteSQL(where, args...)
	if err != nil {
		return nil, err
	}
	return x.Exec(query, args...)
}

// Db.Insert
func (x *Db) Insert(dst any) (sql.Result, error) {
	sb := NewSqlBuilder(dst)
	query, args, err := sb.buildInsertSQL()
	if err != nil {
		log.Println(err, query, args)
		return nil, err
	}
	return x.Exec(query, args...)
}

// Db.Begin
func (x *Db) Begin() (*Db, error) {
	tx, err := x.w.Beginx()
	if err != nil {
		return nil, err
	}
	return &Db{x.r, x.w, tx, x.isDebugMode}, nil
}

// Db.Rollback
func (x *Db) Rollback() error {
	t := x.t
	x.t = nil
	return t.Rollback()
}

// Db.Commit
func (x *Db) Commit() error {
	t := x.t
	x.t = nil
	return t.Commit()
}

// Db.Transaction
func (x *Db) Transaction(f func(*Db) error) error {
	t, err := x.Begin()
	if err != nil {
		return err
	}
	if err = f(t); err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}

// Db.Selectx
func (x *Db) Selectx(dst any, query string, args ...any) (err error) {
	if strings.HasPrefix(strings.ToLower(query), "select *") {
		if x.t != nil {
			err = x.t.Unsafe().Select(dst, query, args...)
		} else {
			err = x.r.Unsafe().Select(dst, query, args...)
		}
	} else {
		if x.t != nil {
			err = x.t.Select(dst, query, args...)
		} else {
			err = x.r.Select(dst, query, args...)
		}
	}
	if x.isDebugMode {
		log.Println("[DEBUG]", query, args)
	}
	return
}

// Db.Queryx
func (x *Db) Queryx(query string, args ...any) (rows *sqlx.Rows, err error) {
	if strings.HasPrefix(strings.ToLower(query), "select *") {
		if x.t != nil {
			rows, err = x.t.Unsafe().Queryx(query, args...)
		} else {
			rows, err = x.r.Unsafe().Queryx(query, args...)
		}
	} else {
		if x.t != nil {
			rows, err = x.t.Queryx(query, args...)
		} else {
			rows, err = x.r.Queryx(query, args...)
		}
	}
	if x.isDebugMode {
		log.Println("[DEBUG]", query, args)
	}
	return
}

// Db.QueryRowx
func (x *Db) QueryRowx(query string, args ...any) (row *sqlx.Row) {
	if strings.HasPrefix(strings.ToLower(query), "select *") {
		if x.t != nil {
			row = x.t.Unsafe().QueryRowx(query, args...)
		} else {
			row = x.r.Unsafe().QueryRowx(query, args...)
		}
	} else {
		if x.t != nil {
			row = x.t.QueryRowx(query, args...)
		} else {
			row = x.r.QueryRowx(query, args...)
		}
	}
	if x.isDebugMode {
		log.Println("[DEBUG]", query, args)
	}
	return
}

// Db.Exec
func (x *Db) Exec(query string, args ...any) (res sql.Result, err error) {
	if x.t != nil {
		res, err = x.t.Exec(query, args...)
	} else {
		res, err = x.w.Exec(query, args...)
	}
	if x.isDebugMode {
		log.Println("[DEBUG]", query, args)
	}
	return
}
