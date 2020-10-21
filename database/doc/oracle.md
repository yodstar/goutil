

# oracle
`import "github.com/yodstar/goutil/database/oracle"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func Begin() (*sqlbuilder.Conn, error)](#Begin)
* [func Count(dest interface{}, where string, args ...interface{}) (int64, error)](#Count)
* [func Debug(mode bool) *sqlbuilder.SQLBuilder](#Debug)
* [func Delete(dest interface{}, where string, args ...interface{}) (sql.Result, error)](#Delete)
* [func Exec(query string, args ...interface{}) (sql.Result, error)](#Exec)
* [func Fields(fields string) *sqlbuilder.SQLBuilder](#Fields)
* [func GroupBy(groupBy string) *sqlbuilder.SQLBuilder](#GroupBy)
* [func Having(having string) *sqlbuilder.SQLBuilder](#Having)
* [func Insert(dest interface{}) (sql.Result, error)](#Insert)
* [func Limit(offset int) *sqlbuilder.SQLBuilder](#Limit)
* [func Model(value interface{}) *sqlbuilder.SQLBuilder](#Model)
* [func MustOpen(dbConfig []*sqlbuilder.Conf, options ...bool)](#MustOpen)
* [func NewSQLBuilder(value interface{}) *sqlbuilder.SQLBuilder](#NewSQLBuilder)
* [func OrderBy(orderBy string) *sqlbuilder.SQLBuilder](#OrderBy)
* [func QueryRowx(query string, args ...interface{}) *sqlx.Row](#QueryRowx)
* [func Queryx(query string, args ...interface{}) (*sqlx.Rows, error)](#Queryx)
* [func Reader() *sqlbuilder.Conn](#Reader)
* [func Select(dest interface{}, where string, args ...interface{}) error](#Select)
* [func Selectx(dest interface{}, query string, args ...interface{}) error](#Selectx)
* [func Transaction(f func(*sqlbuilder.Conn) error) error](#Transaction)
* [func Update(dest interface{}, where string, args ...interface{}) (sql.Result, error)](#Update)
* [func Where(where string, args ...interface{}) *sqlbuilder.SQLBuilder](#Where)
* [func WhereNot(where string, args ...interface{}) *sqlbuilder.SQLBuilder](#WhereNot)
* [func WhereOr(where string, args ...interface{}) *sqlbuilder.SQLBuilder](#WhereOr)
* [func Writer() *sqlbuilder.Conn](#Writer)


#### <a name="pkg-files">Package files</a>
[oracle.go](/src/github.com/yodstar/goutil/database/oracle/oracle.go) [sqlbuilder.go](/src/github.com/yodstar/goutil/database/oracle/sqlbuilder.go) 





## <a name="Begin">func</a> [Begin](/src/target/oracle.go?s=1861:1899#L82)
``` go
func Begin() (*sqlbuilder.Conn, error)
```
Begin



## <a name="Count">func</a> [Count](/src/target/oracle.go?s=647:725#L37)
``` go
func Count(dest interface{}, where string, args ...interface{}) (int64, error)
```
Count



## <a name="Debug">func</a> [Debug](/src/target/sqlbuilder.go?s=258:302#L15)
``` go
func Debug(mode bool) *sqlbuilder.SQLBuilder
```
Debug



## <a name="Delete">func</a> [Delete](/src/target/oracle.go?s=924:1008#L47)
``` go
func Delete(dest interface{}, where string, args ...interface{}) (sql.Result, error)
```
Delete



## <a name="Exec">func</a> [Exec](/src/target/oracle.go?s=1732:1796#L77)
``` go
func Exec(query string, args ...interface{}) (sql.Result, error)
```
Exec



## <a name="Fields">func</a> [Fields](/src/target/sqlbuilder.go?s=375:424#L20)
``` go
func Fields(fields string) *sqlbuilder.SQLBuilder
```
Fields



## <a name="GroupBy">func</a> [GroupBy](/src/target/sqlbuilder.go?s=966:1017#L40)
``` go
func GroupBy(groupBy string) *sqlbuilder.SQLBuilder
```
GroupBy



## <a name="Having">func</a> [Having](/src/target/sqlbuilder.go?s=1095:1144#L45)
``` go
func Having(having string) *sqlbuilder.SQLBuilder
```
Having



## <a name="Insert">func</a> [Insert](/src/target/oracle.go?s=1222:1271#L57)
``` go
func Insert(dest interface{}) (sql.Result, error)
```
Insert



## <a name="Limit">func</a> [Limit](/src/target/sqlbuilder.go?s=1349:1394#L55)
``` go
func Limit(offset int) *sqlbuilder.SQLBuilder
```
Limit



## <a name="Model">func</a> [Model](/src/target/sqlbuilder.go?s=1468:1520#L60)
``` go
func Model(value interface{}) *sqlbuilder.SQLBuilder
```
Model



## <a name="MustOpen">func</a> [MustOpen](/src/target/oracle.go?s=231:290#L16)
``` go
func MustOpen(dbConfig []*sqlbuilder.Conf, options ...bool)
```
MustOpen



## <a name="NewSQLBuilder">func</a> [NewSQLBuilder](/src/target/sqlbuilder.go?s=138:198#L10)
``` go
func NewSQLBuilder(value interface{}) *sqlbuilder.SQLBuilder
```
NewSQLBuilder



## <a name="OrderBy">func</a> [OrderBy](/src/target/sqlbuilder.go?s=1221:1272#L50)
``` go
func OrderBy(orderBy string) *sqlbuilder.SQLBuilder
```
OrderBy



## <a name="QueryRowx">func</a> [QueryRowx](/src/target/oracle.go?s=1604:1663#L72)
``` go
func QueryRowx(query string, args ...interface{}) *sqlx.Row
```
QueryRowx



## <a name="Queryx">func</a> [Queryx](/src/target/oracle.go?s=1467:1533#L67)
``` go
func Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
```
Queryx



## <a name="Reader">func</a> [Reader](/src/target/oracle.go?s=573:603#L32)
``` go
func Reader() *sqlbuilder.Conn
```
Reader



## <a name="Select">func</a> [Select](/src/target/oracle.go?s=789:859#L42)
``` go
func Select(dest interface{}, where string, args ...interface{}) error
```
Select



## <a name="Selectx">func</a> [Selectx](/src/target/oracle.go?s=1321:1392#L62)
``` go
func Selectx(dest interface{}, query string, args ...interface{}) error
```
Selectx



## <a name="Transaction">func</a> [Transaction](/src/target/oracle.go?s=1957:2011#L87)
``` go
func Transaction(f func(*sqlbuilder.Conn) error) error
```
Transaction



## <a name="Update">func</a> [Update](/src/target/oracle.go?s=1073:1157#L52)
``` go
func Update(dest interface{}, where string, args ...interface{}) (sql.Result, error)
```
Update



## <a name="Where">func</a> [Where](/src/target/sqlbuilder.go?s=499:567#L25)
``` go
func Where(where string, args ...interface{}) *sqlbuilder.SQLBuilder
```
Where



## <a name="WhereNot">func</a> [WhereNot](/src/target/sqlbuilder.go?s=808:879#L35)
``` go
func WhereNot(where string, args ...interface{}) *sqlbuilder.SQLBuilder
```
WhereNot



## <a name="WhereOr">func</a> [WhereOr](/src/target/sqlbuilder.go?s=651:721#L30)
``` go
func WhereOr(where string, args ...interface{}) *sqlbuilder.SQLBuilder
```
WhereOr



## <a name="Writer">func</a> [Writer](/src/target/oracle.go?s=498:528#L27)
``` go
func Writer() *sqlbuilder.Conn
```
Writer








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
