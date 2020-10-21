package sqlbuilder

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"unicode"

	"github.com/jmoiron/sqlx"
)

var (
	isDebugMode    = false
	hasModelSuffix = false
	sqlbuilderName = "sqlbuilder"
	errNotBindDB   = fmt.Errorf("not bind DB yet")
)

// SetDebugMode
func SetDebugMode(mode bool) {
	isDebugMode = mode
}

// SetModelSuffix
func SetModelSuffix(has bool) {
	hasModelSuffix = has
}

// Init
func Init(name string) bool {
	sqlbuilderName = name
	return true
}

// SQLBuilder
type SQLBuilder struct {
	*dbSQLBuilder
	db           *DB
	bindValue    interface{}
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
	sqlArgs      []interface{}
	IsSliceValue bool
	IsDebugMode  bool
	Result       sql.Result
	Error        error
}

// SQLBuilder.Debug
func (sb *SQLBuilder) Debug(mode bool) *SQLBuilder {
	sb.IsDebugMode = mode
	return sb
}

// SQLBuilder.Table
func (sb *SQLBuilder) Table(tableName string) *SQLBuilder {
	sb.tableName = tableName
	return sb
}

// SQLBuilder.Fields
func (sb *SQLBuilder) Fields(fields string) *SQLBuilder {
	uf := []string{}
	vf := []string{}
	sf := strings.Split(fields, ",")
	for i, _ := range sf {
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

// SQLBuilder.Where
func (sb *SQLBuilder) Where(where string, args ...interface{}) *SQLBuilder {
	if sb.sqlWhere == "" {
		sb.sqlWhere = fmt.Sprintf("(%s)", where)
	} else {
		sb.sqlWhere = fmt.Sprintf("%s AND (%s)", sb.sqlWhere, where)
	}
	sb.sqlArgs = append(sb.sqlArgs, args...)
	return sb
}

// SQLBuilder.WhereOr
func (sb *SQLBuilder) WhereOr(where string, args ...interface{}) *SQLBuilder {
	if sb.sqlWhere == "" {
		sb.sqlWhere = fmt.Sprintf("(%s)", where)
	} else {
		sb.sqlWhere = fmt.Sprintf("%s OR (%s)", sb.sqlWhere, where)
	}
	sb.sqlArgs = append(sb.sqlArgs, args...)
	return sb
}

// SQLBuilder.WhereNot
func (sb *SQLBuilder) WhereNot(where string, args ...interface{}) *SQLBuilder {
	if sb.sqlWhere == "" {
		sb.sqlWhere = fmt.Sprintf("NOT (%s)", where)
	} else {
		sb.sqlWhere = fmt.Sprintf("%s NOT (%s)", sb.sqlWhere, where)
	}
	sb.sqlArgs = append(sb.sqlArgs, args...)
	return sb
}

// SQLBuilder.GroupBy
func (sb *SQLBuilder) GroupBy(groupBy string) *SQLBuilder {
	sb.sqlGroupBy = fmt.Sprintf(" GROUP BY %s", groupBy)
	return sb
}

// SQLBuilder.Having
func (sb *SQLBuilder) Having(having string) *SQLBuilder {
	sb.sqlHaving = fmt.Sprintf(" HAVING %s", having)
	return sb
}

// SQLBuilder.OrderBy
func (sb *SQLBuilder) OrderBy(orderBy string) *SQLBuilder {
	sb.sqlOrderBy = fmt.Sprintf(" ORDER BY %s", orderBy)
	return sb
}

// SQLBuilder.Limit
func (sb *SQLBuilder) Limit(offset int) *SQLBuilder {
	sb.sqlLimit = fmt.Sprintf(" LIMIT %d", offset)
	return sb
}

// SQLBuilder.Count
func (sb *SQLBuilder) Count(count *int64) *SQLBuilder {
	if sb.db == nil {
		sb.Error = sqlError(errNotBindDB)
		return sb
	}
	where := fmt.Sprintf("%s%s%s%s%s", sb.sqlWhere, sb.sqlGroupBy, sb.sqlHaving, sb.sqlOrderBy, sb.sqlLimit)
	query, args, err := sb.BuildCountSQL(where, sb.sqlArgs...)
	if err != nil {
		sb.Error = sqlError(err)
		return sb
	}
	sb.Error = sqlError(sb.db.QueryRowx(query, args...).Scan(count))
	if sb.IsDebugMode {
		log.Println("[DEBUG]", query, args)
	}
	return sb
}

// SQLBuilder.Select
func (sb *SQLBuilder) Select(dest interface{}) *SQLBuilder {
	if sb.db == nil {
		sb.Error = sqlError(errNotBindDB)
		return sb
	}
	if dest != nil {
		sb.Model(dest)
	}
	where := fmt.Sprintf("%s%s%s%s%s", sb.sqlWhere, sb.sqlGroupBy, sb.sqlHaving, sb.sqlOrderBy, sb.sqlLimit)
	query, args, err := sb.BuildSelectSQL(where, sb.sqlArgs...)
	if err != nil {
		sb.Error = err
		return sb
	}
	if sb.IsSliceValue {
		sb.Error = sqlError(sb.db.Selectx(sb.bindValue, query, args...))
	} else {
		sb.Error = sqlError(sb.db.QueryRowx(query, args...).StructScan(sb.bindValue))
	}
	if sb.IsDebugMode {
		log.Println("[DEBUG]", query, args, sb.Error)
	}
	return sb
}

// SQLBuilder.Update
func (sb *SQLBuilder) Update(dest ...interface{}) *SQLBuilder {
	if sb.db == nil {
		sb.Error = sqlError(errNotBindDB)
		return sb
	}
	if len(dest) > 0 {
		sb.Model(dest[0])
	}
	query, args, err := sb.BuildUpdateSQL(sb.sqlWhere, sb.sqlArgs...)
	if err != nil {
		sb.Error = err
		return sb
	}
	sb.Result, sb.Error = sb.db.Exec(query, args...)
	if sb.Error != nil {
		sb.Error = sqlError(sb.Error)
	}
	if sb.IsDebugMode {
		log.Println("[DEBUG]", query, args, sb.Error)
	}
	return sb
}

// SQLBuilder.Delete
func (sb *SQLBuilder) Delete(dest ...interface{}) *SQLBuilder {
	if sb.db == nil {
		sb.Error = sqlError(errNotBindDB)
		return sb
	}
	if len(dest) > 0 {
		sb.Model(dest[0])
	}
	query, args, err := sb.BuildDeleteSQL(sb.sqlWhere, sb.sqlArgs...)
	if err != nil {
		sb.Error = err
		return sb
	}
	sb.Result, sb.Error = sb.db.Exec(query, args...)
	if sb.Error != nil {
		sb.Error = sqlError(sb.Error)
	}
	if sb.IsDebugMode {
		log.Println("[DEBUG]", query, args, sb.Error)
	}
	return sb
}

// SQLBuilder.Insert
func (sb *SQLBuilder) Insert(dest ...interface{}) *SQLBuilder {
	if sb.db == nil {
		sb.Error = sqlError(errNotBindDB)
		return sb
	}
	if len(dest) > 0 {
		sb.Model(dest[0])
	}
	query, args, err := sb.BuildInsertSQL()
	if err != nil {
		sb.Error = err
		return sb
	}
	sb.Result, sb.Error = sb.db.Exec(query, args...)
	if sb.Error != nil {
		sb.Error = sqlError(sb.Error)
	}
	if sb.IsDebugMode {
		log.Println("[DEBUG]", query, args, sb.Error)
	}
	return sb
}

// SQLBuilder.BuildCountSQL
func (sb *SQLBuilder) BuildCountSQL(where string, args ...interface{}) (query string, argv []interface{}, err error) {
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

// SQLBuilder.BuildSelectSQL
func (sb *SQLBuilder) BuildSelectSQL(where string, args ...interface{}) (query string, argv []interface{}, err error) {
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

// SQLBuilder.BuildDeleteSQL
func (sb *SQLBuilder) BuildDeleteSQL(where string, args ...interface{}) (query string, argv []interface{}, err error) {
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

// SQLBuilder.BuildUpdateSQL
func (sb *SQLBuilder) BuildUpdateSQL(where string, args ...interface{}) (query string, argv []interface{}, err error) {
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

// SQLBuilder.BuildInsertSQL
func (sb *SQLBuilder) BuildInsertSQL() (query string, argv []interface{}, err error) {
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

// SQLBuilder.Model
func (sb *SQLBuilder) Model(value interface{}) *SQLBuilder {
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
	if sb.dbSQLBuilder != nil {
		return sb
	}
	sb.dbSQLBuilder = cachedDBSQLBuilder(rt)
	if sb.tableName == "" {
		sb.tableName = sb.dbSQLBuilder.tableName
	}
	if sb.sqlFields == "" {
		sb.sqlFields = sb.dbSQLBuilder.sqlFields
		sb.sqlUpdates = sb.dbSQLBuilder.sqlUpdates
		sb.sqlColumns = sb.dbSQLBuilder.sqlColumns
		sb.sqlValues = sb.dbSQLBuilder.sqlValues
	}
	return sb
}

// SQLBuilder.Database
func (sb *SQLBuilder) Database(db *DB) *SQLBuilder {
	sb.db = db
	return sb
}

// NewSQLBuilder
func NewSQLBuilder(value interface{}) (sb *SQLBuilder) {
	sb = &SQLBuilder{
		IsDebugMode: isDebugMode,
		sqlArgs:     []interface{}{},
	}
	if value != nil {
		sb.Model(value)
	}
	return
}

// DBSQLBuilder
func DBSQLBuilder(db *DB) (sb *SQLBuilder) {
	sb = &SQLBuilder{
		IsDebugMode: isDebugMode,
		sqlArgs:     []interface{}{},
	}
	if db != nil {
		sb.db = db
	}
	return
}

// dbSQLBuilder
type dbSQLBuilder struct {
	tableName  string
	sqlFields  string
	sqlUpdates string
	sqlColumns string
	sqlValues  string
}

var dbSQLBuilderCache sync.Map // map[reflect.Type]dbSQLBuilder

// cachedDBSQLBuilder
func cachedDBSQLBuilder(rt reflect.Type) *dbSQLBuilder {
	if sb, ok := dbSQLBuilderCache.Load(rt); ok {
		return sb.(*dbSQLBuilder)
	}
	sb, _ := dbSQLBuilderCache.LoadOrStore(rt, reflectDBSQLBuilder(rt))
	return sb.(*dbSQLBuilder)
}

// reflectDBSQLBuilder
func reflectDBSQLBuilder(rt reflect.Type) (sb *dbSQLBuilder) {
	sb = &dbSQLBuilder{}
	if rm := reflect.New(rt).MethodByName("TableName"); rm.IsValid() {
		sb.tableName = rm.Call(nil)[0].Interface().(string)
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
		if hasModelSuffix {
			sb.tableName = strings.TrimSuffix(string(r), "_model")
		} else {
			sb.tableName = string(r)
		}
	}
	var fields, columns, updates, values []string
	if rt.Kind() == reflect.Struct {
		for i := 0; i < rt.NumField(); i++ {
			f := rt.Field(i)
			k := f.Tag.Get("db")
			v := f.Tag.Get(sqlbuilderName)
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
	sb.sqlFields = strings.Join(fields, ", ")
	sb.sqlColumns = strings.Join(columns, ", ")
	sb.sqlUpdates = strings.Join(updates, ", ")
	sb.sqlValues = strings.Join(values, ", ")
	return sb
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
		file = file[strings.LastIndex(file, "/findgo")+1:]
		err = fmt.Errorf("%s (%s:%d)", err, file, line)
	}
	return err
}
