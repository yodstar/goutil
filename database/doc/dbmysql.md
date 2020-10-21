

# dbmysql
`import "github.com/yodstar/goutil/database/mysql"`

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
* [func MustOpen(conf []*sqlbuilder.Conf, options ...bool)](#MustOpen)
* [func NewSQLBuilder(value interface{}) *sqlbuilder.SQLBuilder](#NewSQLBuilder)
* [func OrderBy(orderBy string) *sqlbuilder.SQLBuilder](#OrderBy)
* [func QueryRowx(query string, args ...interface{}) *sqlx.Row](#QueryRowx)
* [func Queryx(query string, args ...interface{}) (*sqlx.Rows, error)](#Queryx)
* [func Reader() *sqlbuilder.Conn](#Reader)
* [func Readerx() *sqlx.DB](#Readerx)
* [func Select(dest interface{}, where string, args ...interface{}) error](#Select)
* [func Selectx(dest interface{}, query string, args ...interface{}) error](#Selectx)
* [func Transaction(f func(*sqlbuilder.Conn) error) error](#Transaction)
* [func Update(dest interface{}, where string, args ...interface{}) (sql.Result, error)](#Update)
* [func Where(where string, args ...interface{}) *sqlbuilder.SQLBuilder](#Where)
* [func WhereNot(where string, args ...interface{}) *sqlbuilder.SQLBuilder](#WhereNot)
* [func WhereOr(where string, args ...interface{}) *sqlbuilder.SQLBuilder](#WhereOr)
* [func Writer() *sqlbuilder.Conn](#Writer)
* [func Writerx() *sqlx.DB](#Writerx)


#### <a name="pkg-files">Package files</a>
[dbmysql.go](/src/github.com/yodstar/goutil/database/mysql/dbmysql.go) [sqlbuilder.go](/src/github.com/yodstar/goutil/database/mysql/sqlbuilder.go) 





## <a name="Begin">func</a> [Begin](/src/target/dbmysql.go?s=2042:2080#L92)
``` go
func Begin() (*sqlbuilder.Conn, error)
```
Begin



## <a name="Count">func</a> [Count](/src/target/dbmysql.go?s=819:897#L47)
``` go
func Count(dest interface{}, where string, args ...interface{}) (int64, error)
```
Count



## <a name="Debug">func</a> [Debug](/src/target/sqlbuilder.go?s=260:304#L15)
``` go
func Debug(mode bool) *sqlbuilder.SQLBuilder
```
Debug



## <a name="Delete">func</a> [Delete](/src/target/dbmysql.go?s=1098:1182#L57)
``` go
func Delete(dest interface{}, where string, args ...interface{}) (sql.Result, error)
```
Delete



## <a name="Exec">func</a> [Exec](/src/target/dbmysql.go?s=1912:1976#L87)
``` go
func Exec(query string, args ...interface{}) (sql.Result, error)
```
Exec



## <a name="Fields">func</a> [Fields](/src/target/sqlbuilder.go?s=378:427#L20)
``` go
func Fields(fields string) *sqlbuilder.SQLBuilder
```
Fields



## <a name="GroupBy">func</a> [GroupBy](/src/target/sqlbuilder.go?s=973:1024#L40)
``` go
func GroupBy(groupBy string) *sqlbuilder.SQLBuilder
```
GroupBy



## <a name="Having">func</a> [Having](/src/target/sqlbuilder.go?s=1103:1152#L45)
``` go
func Having(having string) *sqlbuilder.SQLBuilder
```
Having



## <a name="Insert">func</a> [Insert](/src/target/dbmysql.go?s=1398:1447#L67)
``` go
func Insert(dest interface{}) (sql.Result, error)
```
Insert



## <a name="Limit">func</a> [Limit](/src/target/sqlbuilder.go?s=1359:1404#L55)
``` go
func Limit(offset int) *sqlbuilder.SQLBuilder
```
Limit



## <a name="Model">func</a> [Model](/src/target/sqlbuilder.go?s=1479:1531#L60)
``` go
func Model(value interface{}) *sqlbuilder.SQLBuilder
```
Model



## <a name="MustOpen">func</a> [MustOpen](/src/target/dbmysql.go?s=241:296#L16)
``` go
func MustOpen(conf []*sqlbuilder.Conf, options ...bool)
```
MustOpen



## <a name="NewSQLBuilder">func</a> [NewSQLBuilder](/src/target/sqlbuilder.go?s=140:200#L10)
``` go
func NewSQLBuilder(value interface{}) *sqlbuilder.SQLBuilder
```
NewSQLBuilder



## <a name="OrderBy">func</a> [OrderBy](/src/target/sqlbuilder.go?s=1230:1281#L50)
``` go
func OrderBy(orderBy string) *sqlbuilder.SQLBuilder
```
OrderBy



## <a name="QueryRowx">func</a> [QueryRowx](/src/target/dbmysql.go?s=1783:1842#L82)
``` go
func QueryRowx(query string, args ...interface{}) *sqlx.Row
```
QueryRowx



## <a name="Queryx">func</a> [Queryx](/src/target/dbmysql.go?s=1645:1711#L77)
``` go
func Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
```
Queryx



## <a name="Reader">func</a> [Reader](/src/target/dbmysql.go?s=578:608#L32)
``` go
func Reader() *sqlbuilder.Conn
```
Reader



## <a name="Readerx">func</a> [Readerx](/src/target/dbmysql.go?s=750:773#L42)
``` go
func Readerx() *sqlx.DB
```
Deprecated: Readerx



## <a name="Select">func</a> [Select](/src/target/dbmysql.go?s=962:1032#L52)
``` go
func Select(dest interface{}, where string, args ...interface{}) error
```
Select



## <a name="Selectx">func</a> [Selectx](/src/target/dbmysql.go?s=1498:1569#L72)
``` go
func Selectx(dest interface{}, query string, args ...interface{}) error
```
Selectx



## <a name="Transaction">func</a> [Transaction](/src/target/dbmysql.go?s=2139:2193#L97)
``` go
func Transaction(f func(*sqlbuilder.Conn) error) error
```
Transaction



## <a name="Update">func</a> [Update](/src/target/dbmysql.go?s=1248:1332#L62)
``` go
func Update(dest interface{}, where string, args ...interface{}) (sql.Result, error)
```
Update



## <a name="Where">func</a> [Where](/src/target/sqlbuilder.go?s=503:571#L25)
``` go
func Where(where string, args ...interface{}) *sqlbuilder.SQLBuilder
```
Where



## <a name="WhereNot">func</a> [WhereNot](/src/target/sqlbuilder.go?s=814:885#L35)
``` go
func WhereNot(where string, args ...interface{}) *sqlbuilder.SQLBuilder
```
WhereNot



## <a name="WhereOr">func</a> [WhereOr](/src/target/sqlbuilder.go?s=656:726#L30)
``` go
func WhereOr(where string, args ...interface{}) *sqlbuilder.SQLBuilder
```
WhereOr



## <a name="Writer">func</a> [Writer](/src/target/dbmysql.go?s=502:532#L27)
``` go
func Writer() *sqlbuilder.Conn
```
Writer



## <a name="Writerx">func</a> [Writerx](/src/target/dbmysql.go?s=667:690#L37)
``` go
func Writerx() *sqlx.DB
```
Deprecated: Writerx








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
