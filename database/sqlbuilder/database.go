package sqlbuilder

import (
	"database/sql"
	"log"
	"math/rand"

	"github.com/jmoiron/sqlx"
)

// Conf
type Conf struct {
	DataSourceName string
	MaxOpenConns   int
	MaxIdleConns   int
}

// DB
type DB struct {
	conn []*Conn
}

// MustOpen
func MustOpen(driverName string, conf []*Conf) *DB {
	db := &DB{}
	db.conn = make([]*Conn, len(conf))
	for i, v := range conf {
		c := &Conn{r: sqlx.MustOpen(driverName, v.DataSourceName)}
		if v.MaxOpenConns > 0 {
			c.r.SetMaxOpenConns(v.MaxOpenConns)
		}
		if v.MaxIdleConns > 0 {
			c.r.SetMaxIdleConns(v.MaxIdleConns)
		}
		c.r.Ping()
		if i == 0 {
			c.w = c.r
		}
		db.conn[i] = c
	}
	return db
}

// DB.Writer
func (db *DB) Writer() *Conn {
	if err := db.conn[0].Ping(); err != nil {
		log.Println(err.Error())
	}
	return db.conn[0]
}

// DB.Reader
func (db *DB) Reader() *Conn {
	c := db.conn[rand.Intn(len(db.conn))]
	if err := c.Ping(); err != nil {
		log.Println(err.Error())
	}
	return c
}

// Deprecated: DB.Writerx
func (db *DB) Writerx() *sqlx.DB {
	if err := db.conn[0].Ping(); err != nil {
		log.Println(err.Error())
	}
	return db.conn[0].w
}

// Deprecated: DB.Readerx
func (db *DB) Readerx() *sqlx.DB {
	c := db.conn[rand.Intn(len(db.conn))]
	if err := c.Ping(); err != nil {
		log.Println(err.Error())
	}
	return c.r
}

// DB.Debug
func (db *DB) Debug(mode bool) *SQLBuilder {
	return DBSQLBuilder(db).Debug(mode)
}

// DB.Table
func (db *DB) Table(tableName string) *SQLBuilder {
	return DBSQLBuilder(db).Table(tableName)
}

// DB.Fields
func (db *DB) Fields(fields string) *SQLBuilder {
	return DBSQLBuilder(db).Fields(fields)
}

// DB.Where
func (db *DB) Where(where string, args ...interface{}) *SQLBuilder {
	return DBSQLBuilder(db).Where(where, args...)
}

// DB.WhereOr
func (db *DB) WhereOr(where string, args ...interface{}) *SQLBuilder {
	return DBSQLBuilder(db).WhereOr(where, args...)
}

// DB.WhereNot
func (db *DB) WhereNot(where string, args ...interface{}) *SQLBuilder {
	return DBSQLBuilder(db).WhereNot(where, args...)
}

// DB.GroupBy
func (db *DB) GroupBy(group string) *SQLBuilder {
	return DBSQLBuilder(db).GroupBy(group)
}

// DB.Having
func (db *DB) Having(having string) *SQLBuilder {
	return DBSQLBuilder(db).Having(having)
}

// DB.OrderBy
func (db *DB) OrderBy(order string) *SQLBuilder {
	return DBSQLBuilder(db).OrderBy(order)
}

// DB.Limit
func (db *DB) Limit(offset int) *SQLBuilder {
	return DBSQLBuilder(db).Limit(offset)
}

// DB.Count
func (db *DB) Count(dest interface{}, where string, args ...interface{}) (int64, error) {
	return db.Reader().Count(dest, where, args...)
}

// DB.Select
func (db *DB) Select(dest interface{}, where string, args ...interface{}) error {
	return db.Reader().Select(dest, where, args...)
}

// DB.Update
func (db *DB) Update(dest interface{}, where string, args ...interface{}) (sql.Result, error) {
	return db.Writer().Update(dest, where, args...)
}

// DB.Delete
func (db *DB) Delete(dest interface{}, where string, args ...interface{}) (sql.Result, error) {
	return db.Writer().Delete(dest, where, args...)
}

// DB.Insert
func (db *DB) Insert(dest interface{}) (sql.Result, error) {
	return db.Writer().Insert(dest)
}

// DB.Queryx
func (db *DB) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return db.Reader().Queryx(query, args...)
}

// DB.QueryRowx
func (db *DB) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return db.Reader().QueryRowx(query, args...)
}

// DB.Selectx
func (db *DB) Selectx(dest interface{}, query string, args ...interface{}) error {
	return db.Reader().Selectx(dest, query, args...)
}

// DB.Exec
func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Writer().Exec(query, args...)
}

// DB.Begin
func (db *DB) Begin() (*Conn, error) {
	return db.Writer().Begin()
}

// DB.Transaction
func (db *DB) Transaction(f func(*Conn) error) error {
	return db.Writer().Transaction(f)
}

// Conn
type Conn struct {
	r *sqlx.DB
	w *sqlx.DB
	x *sqlx.Tx
}

// Conn.Ping
func (c *Conn) Ping() error {
	return c.r.Ping()
}

// Conn.Count
func (c *Conn) Count(dest interface{}, where string, args ...interface{}) (int64, error) {
	sb := NewSQLBuilder(dest)
	query, args, err := sb.BuildCountSQL(where, args...)
	if err != nil {
		return 0, err
	}
	var count int64
	err = c.QueryRowx(query, args...).Scan(&count)
	if sb.IsDebugMode {
		log.Println("[DEBUG]", query, args)
	}
	return count, err
}

// Conn.Select
func (c *Conn) Select(dest interface{}, where string, args ...interface{}) error {
	sb := NewSQLBuilder(dest)
	query, args, err := sb.BuildSelectSQL(where, args...)
	if err != nil {
		return err
	}
	if sb.IsSliceValue {
		err = c.Selectx(dest, query, args...)
	} else {
		err = c.QueryRowx(query, args...).StructScan(dest)
	}
	if sb.IsDebugMode {
		log.Println("[DEBUG]", query, args)
	}
	return err
}

// Conn.Update
func (c *Conn) Update(dest interface{}, where string, args ...interface{}) (sql.Result, error) {
	sb := NewSQLBuilder(dest)
	query, args, err := sb.BuildUpdateSQL(where, args...)
	if err != nil {
		return nil, err
	}
	if sb.IsDebugMode {
		log.Println("[DEBUG]", query, args)
	}
	return c.Exec(query, args...)
}

// Conn.Delete
func (c *Conn) Delete(dest interface{}, where string, args ...interface{}) (sql.Result, error) {
	sb := NewSQLBuilder(dest)
	query, args, err := sb.BuildDeleteSQL(where, args...)
	if err != nil {
		return nil, err
	}
	if sb.IsDebugMode {
		log.Println("[DEBUG]", query, args)
	}
	return c.Exec(query, args...)
}

// Conn.Insert
func (c *Conn) Insert(dest interface{}) (sql.Result, error) {
	sb := NewSQLBuilder(dest)
	query, args, err := sb.BuildInsertSQL()
	if err != nil {
		return nil, err
	}
	if sb.IsDebugMode {
		log.Println("[DEBUG]", query, args)
	}
	return c.Exec(query, args...)
}

// Conn.Begin
func (c *Conn) Begin() (*Conn, error) {
	tx, err := c.w.Beginx()
	if err != nil {
		return nil, err
	}
	x := &Conn{}
	x.r = c.r
	x.w = c.w
	x.x = tx
	return x, nil
}

// Conn.Rollback
func (c *Conn) Rollback() error {
	x := c.x
	c.x = nil
	return x.Rollback()
}

// Conn.Commit
func (c *Conn) Commit() error {
	x := c.x
	c.x = nil
	return x.Commit()
}

// Conn.Transaction
func (c *Conn) Transaction(f func(*Conn) error) error {
	x, err := c.Begin()
	if err != nil {
		return err
	}
	if err = f(x); err != nil {
		x.Rollback()
		return err
	}
	x.Commit()
	return nil
}

// Conn.Selectx
func (c *Conn) Selectx(dest interface{}, query string, args ...interface{}) (err error) {
	if c.x != nil {
		err = c.x.Select(dest, query, args...)
	} else {
		err = c.r.Select(dest, query, args...)
	}
	return
}

// Conn.Queryx
func (c *Conn) Queryx(query string, args ...interface{}) (rows *sqlx.Rows, err error) {
	if c.x != nil {
		rows, err = c.x.Queryx(query, args...)
	} else {
		rows, err = c.r.Queryx(query, args...)
	}
	return
}

// Conn.QueryRowx
func (c *Conn) QueryRowx(query string, args ...interface{}) (row *sqlx.Row) {
	if c.x != nil {
		row = c.x.QueryRowx(query, args...)
	} else {
		row = c.r.QueryRowx(query, args...)
	}
	return
}

// Conn.Exec
func (c *Conn) Exec(query string, args ...interface{}) (res sql.Result, err error) {
	if c.x != nil {
		res, err = c.x.Exec(query, args...)
	} else {
		res, err = c.w.Exec(query, args...)
	}
	return
}
