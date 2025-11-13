package sqlbuilder

import (
	"database/sql"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"unicode"

	"github.com/jmoiron/sqlx"
)

var (
	sqlDaoSuffix   = "_dao"
	sqlbuilderName = "sqlbuilder"
	errNotBindDb   = fmt.Errorf("not bind Db yet")
)

// SetDaoSuffix
func SetDaoSuffix(suffix string) {
	if strings.HasPrefix(suffix, "_") {
		sqlDaoSuffix = strings.ToLower(suffix)
	} else {
		sqlDaoSuffix = "_" + strings.ToLower(suffix)
	}
}

// Init
func Init(name string) bool {
	sqlbuilderName = name
	return true
}

// SqlBuilder
type SqlBuilder struct {
	*dbSqlBuilder
	dao          *Dao
	bindValue    any
	driverName   string
	tableName    string
	sqlFields    string
	sqlUpdates   string
	sqlColumns   string
	sqlValues    string
	sqlWhere     string
	sqlGroupBy   string
	sqlHaving    string
	sqlOrderBy   string
	sqlLimit     string
	sqlArgs      []any
	IsSliceValue bool
	Result       sql.Result
	Error        error
}

// SqlBuilder.Table
func (sb *SqlBuilder) Table(tableName string) *SqlBuilder {
	sb.tableName = tableName
	return sb
}

// SqlBuilder.Fields
func (sb *SqlBuilder) Fields(fields string) *SqlBuilder {
	uf := []string{}
	vf := []string{}
	sf := strings.Split(fields, ",")
	for i := range sf {
		f := strings.TrimSpace(sf[i])
		uf = append(uf, fmt.Sprintf("%s = :%s", f, f))
		vf = append(uf, fmt.Sprintf(":%s", f))
	}
	sb.sqlFields = fields
	sb.sqlUpdates = strings.Join(uf, ", ")
	sb.sqlColumns = fields
	sb.sqlValues = strings.Join(vf, ", ")
	return sb
}

// SqlBuilder.Where
func (sb *SqlBuilder) Where(where string, args ...any) *SqlBuilder {
	if sb.sqlWhere == "" {
		sb.sqlWhere = fmt.Sprintf("(%s)", where)
	} else {
		sb.sqlWhere = fmt.Sprintf("%s AND (%s)", sb.sqlWhere, where)
	}
	sb.sqlArgs = append(sb.sqlArgs, args...)
	return sb
}

// SqlBuilder.WhereOr
func (sb *SqlBuilder) WhereOr(where string, args ...any) *SqlBuilder {
	if sb.sqlWhere == "" {
		sb.sqlWhere = fmt.Sprintf("(%s)", where)
	} else {
		sb.sqlWhere = fmt.Sprintf("%s OR (%s)", sb.sqlWhere, where)
	}
	sb.sqlArgs = append(sb.sqlArgs, args...)
	return sb
}

// SqlBuilder.WhereNot
func (sb *SqlBuilder) WhereNot(where string, args ...any) *SqlBuilder {
	if sb.sqlWhere == "" {
		sb.sqlWhere = fmt.Sprintf("NOT (%s)", where)
	} else {
		sb.sqlWhere = fmt.Sprintf("%s NOT (%s)", sb.sqlWhere, where)
	}
	sb.sqlArgs = append(sb.sqlArgs, args...)
	return sb
}

// SqlBuilder.GroupBy
func (sb *SqlBuilder) GroupBy(groupBy string) *SqlBuilder {
	sb.sqlGroupBy = fmt.Sprintf(" GROUP BY %s", groupBy)
	return sb
}

// SqlBuilder.Having
func (sb *SqlBuilder) Having(having string) *SqlBuilder {
	sb.sqlHaving = fmt.Sprintf(" HAVING %s", having)
	return sb
}

// SqlBuilder.OrderBy
func (sb *SqlBuilder) OrderBy(orderBy string) *SqlBuilder {
	sb.sqlOrderBy = fmt.Sprintf(" ORDER BY %s", orderBy)
	return sb
}

// SqlBuilder.Limit
func (sb *SqlBuilder) Limit(offset int) *SqlBuilder {
	sb.sqlLimit = fmt.Sprintf(" LIMIT %d", offset)
	return sb
}

// SqlBuilder.Count
func (sb *SqlBuilder) Count(count *int64) *SqlBuilder {
	if sb.dao == nil {
		sb.Error = sqlError(errNotBindDb)
		return sb
	}
	where := fmt.Sprintf("%s%s%s%s%s", sb.sqlWhere, sb.sqlGroupBy, sb.sqlHaving, sb.sqlOrderBy, sb.sqlLimit)
	query, args, err := sb.buildCountSQL(where, sb.sqlArgs...)
	if err != nil {
		sb.Error = sqlError(err)
		return sb
	}
	sb.Error = sqlError(sb.dao.QueryRowx(query, args...).Scan(count))
	return sb
}

// SqlBuilder.Select
func (sb *SqlBuilder) Select(dst any) *SqlBuilder {
	if sb.dao == nil {
		sb.Error = sqlError(errNotBindDb)
		return sb
	}
	if dst != nil {
		sb.Dao(dst)
	}
	where := fmt.Sprintf("%s%s%s%s%s", sb.sqlWhere, sb.sqlGroupBy, sb.sqlHaving, sb.sqlOrderBy, sb.sqlLimit)
	query, args, err := sb.buildSelectSQL(where, sb.sqlArgs...)
	if err != nil {
		sb.Error = err
		return sb
	}
	if sb.IsSliceValue {
		sb.Error = sqlError(sb.dao.Selectx(sb.bindValue, query, args...))
	} else {
		sb.Error = sqlError(sb.dao.QueryRowx(query, args...).StructScan(sb.bindValue))
	}
	return sb
}

// SqlBuilder.Update
func (sb *SqlBuilder) Update(dst ...any) *SqlBuilder {
	if sb.dao == nil {
		sb.Error = sqlError(errNotBindDb)
		return sb
	}
	if len(dst) > 0 {
		sb.Dao(dst[0])
	}
	query, args, err := sb.buildUpdateSQL(sb.sqlWhere, sb.sqlArgs...)
	if err != nil {
		sb.Error = err
		return sb
	}
	sb.Result, sb.Error = sb.dao.Exec(query, args...)
	if sb.Error != nil {
		sb.Error = sqlError(sb.Error)
	}
	return sb
}

// SqlBuilder.Delete
func (sb *SqlBuilder) Delete(dst ...any) *SqlBuilder {
	if sb.dao == nil {
		sb.Error = sqlError(errNotBindDb)
		return sb
	}
	if len(dst) > 0 {
		sb.Dao(dst[0])
	}
	query, args, err := sb.buildDeleteSQL(sb.sqlWhere, sb.sqlArgs...)
	if err != nil {
		sb.Error = err
		return sb
	}
	sb.Result, sb.Error = sb.dao.Exec(query, args...)
	if sb.Error != nil {
		sb.Error = sqlError(sb.Error)
	}
	return sb
}

// SqlBuilder.Insert
func (sb *SqlBuilder) Insert(dst ...any) *SqlBuilder {
	if sb.dao == nil {
		sb.Error = sqlError(errNotBindDb)
		return sb
	}
	if len(dst) > 0 {
		sb.Dao(dst[0])
	}
	query, args, err := sb.buildInsertSQL()
	if err != nil {
		sb.Error = err
		return sb
	}
	sb.Result, sb.Error = sb.dao.Exec(query, args...)
	if sb.Error != nil {
		sb.Error = sqlError(sb.Error)
	}
	return sb
}

// SqlBuilder.buildCountSQL
func (sb *SqlBuilder) buildCountSQL(where string, args ...any) (query string, argv []any, err error) {
	if sb.Error != nil {
		err = sb.Error
		return
	}
	if where != "" {
		query = fmt.Sprintf("SELECT COUNT(1) FROM %s WHERE %s", sb.tableName, where)
	} else {
		query = fmt.Sprintf("SELECT COUNT(1) FROM %s", sb.tableName)
	}
	query, argv, err = sqlx.In(query, args...)
	if err != nil {
		err = sqlError(err)
	}
	return
}

// SqlBuilder.buildSelectSQL
func (sb *SqlBuilder) buildSelectSQL(where string, args ...any) (query string, argv []any, err error) {
	if sb.Error != nil {
		err = sb.Error
		return
	}
	if where != "" {
		query = fmt.Sprintf("SELECT %s FROM %s WHERE %s", sb.sqlFields, sb.tableName, where)
	} else {
		query = fmt.Sprintf("SELECT %s FROM %s", sb.sqlFields, sb.tableName)
	}
	query, argv, err = sqlx.In(query, args...)
	if err != nil {
		err = sqlError(err)
	}
	return
}

// SqlBuilder.buildDeleteSQL
func (sb *SqlBuilder) buildDeleteSQL(where string, args ...any) (query string, argv []any, err error) {
	if sb.Error != nil {
		err = sb.Error
		return
	}
	if where != "" {
		query = fmt.Sprintf("DELETE FROM %s WHERE %s", sb.tableName, where)
	} else {
		query = fmt.Sprintf("DELETE FROM %s", sb.tableName)
	}
	query, argv, err = sqlx.In(query, args...)
	if err != nil {
		err = sqlError(err)
	}
	return
}

// SqlBuilder.buildUpdateSQL
func (sb *SqlBuilder) buildUpdateSQL(where string, args ...any) (query string, argv []any, err error) {
	if sb.Error != nil {
		err = sb.Error
		return
	}
	if sb.IsSliceValue {
		err = sqlError(fmt.Errorf("expected %s but got %s", reflect.Struct, reflect.Slice))
		return
	}
	if where != "" {
		query = fmt.Sprintf("UPDATE %s SET %s WHERE %s", sb.tableName, sb.sqlUpdates, where)
	} else {
		query = fmt.Sprintf("UPDATE %s SET %s", sb.tableName, sb.sqlUpdates)
	}
	query, argv, err = sqlx.Named(query, sb.bindValue)
	if err != nil {
		err = sqlError(err)
		return
	}
	argv = append(argv, args...)
	query, argv, err = sqlx.In(query, argv...)
	if err != nil {
		err = sqlError(err)
	}
	return
}

// SqlBuilder.buildInsertSQL
func (sb *SqlBuilder) buildInsertSQL() (query string, argv []any, err error) {
	if sb.Error != nil {
		err = sb.Error
		return
	}
	if sb.IsSliceValue {
		err = sqlError(fmt.Errorf("expected %s but got %s", reflect.Struct, reflect.Slice))
		return
	}
	query = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", sb.tableName, sb.sqlColumns, sb.sqlValues)
	query, argv, err = sqlx.Named(query, sb.bindValue)
	if err != nil {
		err = sqlError(err)
	}
	return
}

// SqlBuilder.Dao
func (sb *SqlBuilder) Dao(value any) *SqlBuilder {
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Ptr {
		sb.Error = sqlError(fmt.Errorf("expected %s but got %s", reflect.Ptr, rv.Kind()))
		return sb
	}
	if rv.IsNil() {
		sb.Error = sqlError(fmt.Errorf("nil pointer passed to StructScan destination"))
		return sb
	}
	rt := reflectDeref(rv.Type())
	if sb.IsSliceValue = (rt.Kind() == reflect.Slice); sb.IsSliceValue {
		rt = reflectDeref(rt.Elem())
	}
	if rt.Kind() != reflect.Struct {
		sb.Error = sqlError(fmt.Errorf("expected %s but got %s", reflect.Struct, rt.Kind()))
		return sb
	}
	sb.bindValue = value
	if sb.dbSqlBuilder != nil {
		return sb
	}
	sb.dbSqlBuilder = sb.cachedDbSqlBuilder(rt)
	if sb.tableName == "" {
		sb.tableName = sb.dbSqlBuilder.tableName
	}
	if sb.sqlFields == "" {
		sb.sqlFields = sb.dbSqlBuilder.sqlFields
		sb.sqlUpdates = sb.dbSqlBuilder.sqlUpdates
		sb.sqlColumns = sb.dbSqlBuilder.sqlColumns
		sb.sqlValues = sb.dbSqlBuilder.sqlValues
	}
	return sb
}

// SqlBuilder.Database
func (sb *SqlBuilder) Database(dao *Dao) *SqlBuilder {
	sb.dao = dao
	return sb
}

// NewSqlBuilder
func NewSqlBuilder(driverName string, value any) *SqlBuilder {
	sb := &SqlBuilder{
		driverName: driverName,
		sqlArgs:    []any{},
	}
	if value != nil {
		sb.Dao(value)
	}
	return sb
}

// DbSqlBuilder
func DbSqlBuilder(dao *Dao) (sb *SqlBuilder) {
	sb = &SqlBuilder{
		sqlArgs: []any{},
	}
	if dao != nil {
		sb.driverName = dao.driverName
		sb.dao = dao
	}
	return
}

// dbSqlBuilder
type dbSqlBuilder struct {
	tableName  string
	sqlFields  string
	sqlUpdates string
	sqlColumns string
	sqlValues  string
}

var dbSqlBuilderCache sync.Map // map[reflect.Type]dbSqlBuilder

// cachedDbSqlBuilder
func (sb *SqlBuilder) cachedDbSqlBuilder(rt reflect.Type) *dbSqlBuilder {
	if db, ok := dbSqlBuilderCache.Load(rt); ok {
		return db.(*dbSqlBuilder)
	}
	db, _ := dbSqlBuilderCache.LoadOrStore(rt, sb.reflectDbSqlBuilder(rt))
	return db.(*dbSqlBuilder)
}

// reflectDbSqlBuilder
func (sb *SqlBuilder) reflectDbSqlBuilder(rt reflect.Type) *dbSqlBuilder {
	db := &dbSqlBuilder{}
	if rm := reflect.New(rt).MethodByName("TableName"); rm.IsValid() {
		db.tableName = rm.Call(nil)[0].Interface().(string)
	} else {
		var r []rune
		for i, c := range rt.Name() {
			if unicode.IsUpper(c) {
				c += 'a' - 'A'
				if i > 0 {
					r = append(r, '_')
				}
			}
			r = append(r, c)
		}
		db.tableName = strings.TrimSuffix(string(r), sqlDaoSuffix)
	}
	var fields, columns, updates, values []string
	if rt.Kind() == reflect.Struct {
		for i := 0; i < rt.NumField(); i++ {
			f := rt.Field(i)
			k := f.Tag.Get("db")
			v := f.Tag.Get(sb.driverName)
			if k == "" || k == "-" {
				continue
			}
			if v == "" {
				fields = append(fields, k)
				columns = append(columns, k)
				updates = append(updates, fmt.Sprintf("%s = :%s", k, k))
				values = append(values, fmt.Sprintf(":%s", k))
			} else {
				if v == "-" {
					fields = append(fields, k)
				} else {
					fields = append(fields, fmt.Sprintf("%s AS %s", v, k))
				}
			}
		}
	}
	db.sqlFields = strings.Join(fields, ", ")
	db.sqlColumns = strings.Join(columns, ", ")
	db.sqlUpdates = strings.Join(updates, ", ")
	db.sqlValues = strings.Join(values, ", ")
	return db
}

// reflectDeref is Indirect for reflect.Types
func reflectDeref(rt reflect.Type) reflect.Type {
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	return rt
}

func sqlError(err error) error {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			file = "???"
			line = 0
		}
		file = file[strings.LastIndex(file, "/goutil")+1:]
		err = fmt.Errorf("%s (%s:%d)", err, file, line)
	}
	return err
}
