

# sqlbuilder
`import "github.com/yodstar/goutil/database/sqlbuilder"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func Init(name string) bool](#Init)
* [func SetDebugMode(mode bool)](#SetDebugMode)
* [func SetModelSuffix(has bool)](#SetModelSuffix)
* [type Conf](#Conf)
* [type Conn](#Conn)
  * [func (c *Conn) Begin() (*Conn, error)](#Conn.Begin)
  * [func (c *Conn) Commit() error](#Conn.Commit)
  * [func (c *Conn) Count(dest interface{}, where string, args ...interface{}) (int64, error)](#Conn.Count)
  * [func (c *Conn) Delete(dest interface{}, where string, args ...interface{}) (sql.Result, error)](#Conn.Delete)
  * [func (c *Conn) Exec(query string, args ...interface{}) (res sql.Result, err error)](#Conn.Exec)
  * [func (c *Conn) Insert(dest interface{}) (sql.Result, error)](#Conn.Insert)
  * [func (c *Conn) Ping() error](#Conn.Ping)
  * [func (c *Conn) QueryRowx(query string, args ...interface{}) (row *sqlx.Row)](#Conn.QueryRowx)
  * [func (c *Conn) Queryx(query string, args ...interface{}) (rows *sqlx.Rows, err error)](#Conn.Queryx)
  * [func (c *Conn) Rollback() error](#Conn.Rollback)
  * [func (c *Conn) Select(dest interface{}, where string, args ...interface{}) error](#Conn.Select)
  * [func (c *Conn) Selectx(dest interface{}, query string, args ...interface{}) (err error)](#Conn.Selectx)
  * [func (c *Conn) Transaction(f func(*Conn) error) error](#Conn.Transaction)
  * [func (c *Conn) Update(dest interface{}, where string, args ...interface{}) (sql.Result, error)](#Conn.Update)
* [type DB](#DB)
  * [func MustOpen(driverName string, conf []*Conf) *DB](#MustOpen)
  * [func (db *DB) Begin() (*Conn, error)](#DB.Begin)
  * [func (db *DB) Count(dest interface{}, where string, args ...interface{}) (int64, error)](#DB.Count)
  * [func (db *DB) Debug(mode bool) *SQLBuilder](#DB.Debug)
  * [func (db *DB) Delete(dest interface{}, where string, args ...interface{}) (sql.Result, error)](#DB.Delete)
  * [func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error)](#DB.Exec)
  * [func (db *DB) Fields(fields string) *SQLBuilder](#DB.Fields)
  * [func (db *DB) GroupBy(group string) *SQLBuilder](#DB.GroupBy)
  * [func (db *DB) Having(having string) *SQLBuilder](#DB.Having)
  * [func (db *DB) Insert(dest interface{}) (sql.Result, error)](#DB.Insert)
  * [func (db *DB) Limit(offset int) *SQLBuilder](#DB.Limit)
  * [func (db *DB) OrderBy(order string) *SQLBuilder](#DB.OrderBy)
  * [func (db *DB) QueryRowx(query string, args ...interface{}) *sqlx.Row](#DB.QueryRowx)
  * [func (db *DB) Queryx(query string, args ...interface{}) (*sqlx.Rows, error)](#DB.Queryx)
  * [func (db *DB) Reader() *Conn](#DB.Reader)
  * [func (db *DB) Readerx() *sqlx.DB](#DB.Readerx)
  * [func (db *DB) Select(dest interface{}, where string, args ...interface{}) error](#DB.Select)
  * [func (db *DB) Selectx(dest interface{}, query string, args ...interface{}) error](#DB.Selectx)
  * [func (db *DB) Table(tableName string) *SQLBuilder](#DB.Table)
  * [func (db *DB) Transaction(f func(*Conn) error) error](#DB.Transaction)
  * [func (db *DB) Update(dest interface{}, where string, args ...interface{}) (sql.Result, error)](#DB.Update)
  * [func (db *DB) Where(where string, args ...interface{}) *SQLBuilder](#DB.Where)
  * [func (db *DB) WhereNot(where string, args ...interface{}) *SQLBuilder](#DB.WhereNot)
  * [func (db *DB) WhereOr(where string, args ...interface{}) *SQLBuilder](#DB.WhereOr)
  * [func (db *DB) Writer() *Conn](#DB.Writer)
  * [func (db *DB) Writerx() *sqlx.DB](#DB.Writerx)
* [type SQLBuilder](#SQLBuilder)
  * [func DBSQLBuilder(db *DB) (sb *SQLBuilder)](#DBSQLBuilder)
  * [func NewSQLBuilder(value interface{}) (sb *SQLBuilder)](#NewSQLBuilder)
  * [func (sb *SQLBuilder) BuildCountSQL(where string, args ...interface{}) (query string, argv []interface{}, err error)](#SQLBuilder.BuildCountSQL)
  * [func (sb *SQLBuilder) BuildDeleteSQL(where string, args ...interface{}) (query string, argv []interface{}, err error)](#SQLBuilder.BuildDeleteSQL)
  * [func (sb *SQLBuilder) BuildInsertSQL() (query string, argv []interface{}, err error)](#SQLBuilder.BuildInsertSQL)
  * [func (sb *SQLBuilder) BuildSelectSQL(where string, args ...interface{}) (query string, argv []interface{}, err error)](#SQLBuilder.BuildSelectSQL)
  * [func (sb *SQLBuilder) BuildUpdateSQL(where string, args ...interface{}) (query string, argv []interface{}, err error)](#SQLBuilder.BuildUpdateSQL)
  * [func (sb *SQLBuilder) Count(count *int64) *SQLBuilder](#SQLBuilder.Count)
  * [func (sb *SQLBuilder) Database(db *DB) *SQLBuilder](#SQLBuilder.Database)
  * [func (sb *SQLBuilder) Debug(mode bool) *SQLBuilder](#SQLBuilder.Debug)
  * [func (sb *SQLBuilder) Delete(dest ...interface{}) *SQLBuilder](#SQLBuilder.Delete)
  * [func (sb *SQLBuilder) Fields(fields string) *SQLBuilder](#SQLBuilder.Fields)
  * [func (sb *SQLBuilder) GroupBy(groupBy string) *SQLBuilder](#SQLBuilder.GroupBy)
  * [func (sb *SQLBuilder) Having(having string) *SQLBuilder](#SQLBuilder.Having)
  * [func (sb *SQLBuilder) Insert(dest ...interface{}) *SQLBuilder](#SQLBuilder.Insert)
  * [func (sb *SQLBuilder) Limit(offset int) *SQLBuilder](#SQLBuilder.Limit)
  * [func (sb *SQLBuilder) Model(value interface{}) *SQLBuilder](#SQLBuilder.Model)
  * [func (sb *SQLBuilder) OrderBy(orderBy string) *SQLBuilder](#SQLBuilder.OrderBy)
  * [func (sb *SQLBuilder) Select(dest interface{}) *SQLBuilder](#SQLBuilder.Select)
  * [func (sb *SQLBuilder) Table(tableName string) *SQLBuilder](#SQLBuilder.Table)
  * [func (sb *SQLBuilder) Update(dest ...interface{}) *SQLBuilder](#SQLBuilder.Update)
  * [func (sb *SQLBuilder) Where(where string, args ...interface{}) *SQLBuilder](#SQLBuilder.Where)
  * [func (sb *SQLBuilder) WhereNot(where string, args ...interface{}) *SQLBuilder](#SQLBuilder.WhereNot)
  * [func (sb *SQLBuilder) WhereOr(where string, args ...interface{}) *SQLBuilder](#SQLBuilder.WhereOr)


#### <a name="pkg-files">Package files</a>
[database.go](/src/github.com/yodstar/goutil/database/sqlbuilder/database.go) [sqlbuilder.go](/src/github.com/yodstar/goutil/database/sqlbuilder/sqlbuilder.go) 





## <a name="Init">func</a> [Init](/src/target/sqlbuilder.go?s=464:491#L34)
``` go
func Init(name string) bool
```
Init



## <a name="SetDebugMode">func</a> [SetDebugMode](/src/target/sqlbuilder.go?s=317:345#L24)
``` go
func SetDebugMode(mode bool)
```
SetDebugMode



## <a name="SetModelSuffix">func</a> [SetModelSuffix](/src/target/sqlbuilder.go?s=394:423#L29)
``` go
func SetModelSuffix(has bool)
```
SetModelSuffix




## <a name="Conf">type</a> [Conf](/src/target/database.go?s=115:202#L12)
``` go
type Conf struct {
    DataSourceName string
    MaxOpenConns   int
    MaxIdleConns   int
}

```
Conf










## <a name="Conn">type</a> [Conn](/src/target/database.go?s=4092:4152#L184)
``` go
type Conn struct {
    // contains filtered or unexported fields
}

```
Conn










### <a name="Conn.Begin">func</a> (\*Conn) [Begin](/src/target/database.go?s=6037:6074#L268)
``` go
func (c *Conn) Begin() (*Conn, error)
```
Conn.Begin




### <a name="Conn.Commit">func</a> (\*Conn) [Commit](/src/target/database.go?s=6335:6364#L288)
``` go
func (c *Conn) Commit() error
```
Conn.Commit




### <a name="Conn.Count">func</a> (\*Conn) [Count](/src/target/database.go?s=4241:4329#L196)
``` go
func (c *Conn) Count(dest interface{}, where string, args ...interface{}) (int64, error)
```
Conn.Count




### <a name="Conn.Delete">func</a> (\*Conn) [Delete](/src/target/database.go?s=5405:5499#L242)
``` go
func (c *Conn) Delete(dest interface{}, where string, args ...interface{}) (sql.Result, error)
```
Conn.Delete




### <a name="Conn.Exec">func</a> (\*Conn) [Exec](/src/target/database.go?s=7359:7441#L339)
``` go
func (c *Conn) Exec(query string, args ...interface{}) (res sql.Result, err error)
```
Conn.Exec




### <a name="Conn.Insert">func</a> (\*Conn) [Insert](/src/target/database.go?s=5746:5805#L255)
``` go
func (c *Conn) Insert(dest interface{}) (sql.Result, error)
```
Conn.Insert




### <a name="Conn.Ping">func</a> (\*Conn) [Ping](/src/target/database.go?s=4170:4197#L191)
``` go
func (c *Conn) Ping() error
```
Conn.Ping




### <a name="Conn.QueryRowx">func</a> (\*Conn) [QueryRowx](/src/target/database.go?s=7141:7216#L329)
``` go
func (c *Conn) QueryRowx(query string, args ...interface{}) (row *sqlx.Row)
```
Conn.QueryRowx




### <a name="Conn.Queryx">func</a> (\*Conn) [Queryx](/src/target/database.go?s=6902:6987#L319)
``` go
func (c *Conn) Queryx(query string, args ...interface{}) (rows *sqlx.Rows, err error)
```
Conn.Queryx




### <a name="Conn.Rollback">func</a> (\*Conn) [Rollback](/src/target/database.go?s=6234:6265#L281)
``` go
func (c *Conn) Rollback() error
```
Conn.Rollback




### <a name="Conn.Select">func</a> (\*Conn) [Select](/src/target/database.go?s=4628:4708#L211)
``` go
func (c *Conn) Select(dest interface{}, where string, args ...interface{}) error
```
Conn.Select




### <a name="Conn.Selectx">func</a> (\*Conn) [Selectx](/src/target/database.go?s=6664:6751#L309)
``` go
func (c *Conn) Selectx(dest interface{}, query string, args ...interface{}) (err error)
```
Conn.Selectx




### <a name="Conn.Transaction">func</a> (\*Conn) [Transaction](/src/target/database.go?s=6437:6490#L295)
``` go
func (c *Conn) Transaction(f func(*Conn) error) error
```
Conn.Transaction




### <a name="Conn.Update">func</a> (\*Conn) [Update](/src/target/database.go?s=5064:5158#L229)
``` go
func (c *Conn) Update(dest interface{}, where string, args ...interface{}) (sql.Result, error)
```
Conn.Update




## <a name="DB">type</a> [DB](/src/target/database.go?s=213:247#L19)
``` go
type DB struct {
    // contains filtered or unexported fields
}

```
DB







### <a name="MustOpen">func</a> [MustOpen](/src/target/database.go?s=264:314#L24)
``` go
func MustOpen(driverName string, conf []*Conf) *DB
```
MustOpen





### <a name="DB.Begin">func</a> (\*DB) [Begin](/src/target/database.go?s=3893:3929#L174)
``` go
func (db *DB) Begin() (*Conn, error)
```
DB.Begin




### <a name="DB.Count">func</a> (\*DB) [Count](/src/target/database.go?s=2562:2649#L129)
``` go
func (db *DB) Count(dest interface{}, where string, args ...interface{}) (int64, error)
```
DB.Count




### <a name="DB.Debug">func</a> (\*DB) [Debug](/src/target/database.go?s=1372:1414#L79)
``` go
func (db *DB) Debug(mode bool) *SQLBuilder
```
DB.Debug




### <a name="DB.Delete">func</a> (\*DB) [Delete](/src/target/database.go?s=3039:3132#L144)
``` go
func (db *DB) Delete(dest interface{}, where string, args ...interface{}) (sql.Result, error)
```
DB.Delete




### <a name="DB.Exec">func</a> (\*DB) [Exec](/src/target/database.go?s=3756:3829#L169)
``` go
func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error)
```
DB.Exec




### <a name="DB.Fields">func</a> (\*DB) [Fields](/src/target/database.go?s=1589:1636#L89)
``` go
func (db *DB) Fields(fields string) *SQLBuilder
```
DB.Fields




### <a name="DB.GroupBy">func</a> (\*DB) [GroupBy](/src/target/database.go?s=2124:2171#L109)
``` go
func (db *DB) GroupBy(group string) *SQLBuilder
```
DB.GroupBy




### <a name="DB.Having">func</a> (\*DB) [Having](/src/target/database.go?s=2235:2282#L114)
``` go
func (db *DB) Having(having string) *SQLBuilder
```
DB.Having




### <a name="DB.Insert">func</a> (\*DB) [Insert](/src/target/database.go?s=3205:3263#L149)
``` go
func (db *DB) Insert(dest interface{}) (sql.Result, error)
```
DB.Insert




### <a name="DB.Limit">func</a> (\*DB) [Limit](/src/target/database.go?s=2457:2500#L124)
``` go
func (db *DB) Limit(offset int) *SQLBuilder
```
DB.Limit




### <a name="DB.OrderBy">func</a> (\*DB) [OrderBy](/src/target/database.go?s=2347:2394#L119)
``` go
func (db *DB) OrderBy(order string) *SQLBuilder
```
DB.OrderBy




### <a name="DB.QueryRowx">func</a> (\*DB) [QueryRowx](/src/target/database.go?s=3465:3533#L159)
``` go
func (db *DB) QueryRowx(query string, args ...interface{}) *sqlx.Row
```
DB.QueryRowx




### <a name="DB.Queryx">func</a> (\*DB) [Queryx](/src/target/database.go?s=3320:3395#L154)
``` go
func (db *DB) Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
```
DB.Queryx




### <a name="DB.Reader">func</a> (\*DB) [Reader](/src/target/database.go?s=850:878#L53)
``` go
func (db *DB) Reader() *Conn
```
DB.Reader




### <a name="DB.Readerx">func</a> (\*DB) [Readerx](/src/target/database.go?s=1198:1230#L70)
``` go
func (db *DB) Readerx() *sqlx.DB
```
Deprecated: DB.Readerx




### <a name="DB.Select">func</a> (\*DB) [Select](/src/target/database.go?s=2721:2800#L134)
``` go
func (db *DB) Select(dest interface{}, where string, args ...interface{}) error
```
DB.Select




### <a name="DB.Selectx">func</a> (\*DB) [Selectx](/src/target/database.go?s=3604:3684#L164)
``` go
func (db *DB) Selectx(dest interface{}, query string, args ...interface{}) error
```
DB.Selectx




### <a name="DB.Table">func</a> (\*DB) [Table](/src/target/database.go?s=1474:1523#L84)
``` go
func (db *DB) Table(tableName string) *SQLBuilder
```
DB.Table




### <a name="DB.Transaction">func</a> (\*DB) [Transaction](/src/target/database.go?s=3986:4038#L179)
``` go
func (db *DB) Transaction(f func(*Conn) error) error
```
DB.Transaction




### <a name="DB.Update">func</a> (\*DB) [Update](/src/target/database.go?s=2873:2966#L139)
``` go
func (db *DB) Update(dest interface{}, where string, args ...interface{}) (sql.Result, error)
```
DB.Update




### <a name="DB.Where">func</a> (\*DB) [Where](/src/target/database.go?s=1699:1765#L94)
``` go
func (db *DB) Where(where string, args ...interface{}) *SQLBuilder
```
DB.Where




### <a name="DB.WhereNot">func</a> (\*DB) [WhereNot](/src/target/database.go?s=1980:2049#L104)
``` go
func (db *DB) WhereNot(where string, args ...interface{}) *SQLBuilder
```
DB.WhereNot




### <a name="DB.WhereOr">func</a> (\*DB) [WhereOr](/src/target/database.go?s=1837:1905#L99)
``` go
func (db *DB) WhereOr(where string, args ...interface{}) *SQLBuilder
```
DB.WhereOr




### <a name="DB.Writer">func</a> (\*DB) [Writer](/src/target/database.go?s=703:731#L45)
``` go
func (db *DB) Writer() *Conn
```
DB.Writer




### <a name="DB.Writerx">func</a> (\*DB) [Writerx](/src/target/database.go?s=1032:1064#L62)
``` go
func (db *DB) Writerx() *sqlx.DB
```
Deprecated: DB.Writerx




## <a name="SQLBuilder">type</a> [SQLBuilder](/src/target/sqlbuilder.go?s=553:978#L40)
``` go
type SQLBuilder struct {
    IsSliceValue bool
    IsDebugMode  bool
    Result       sql.Result
    Error        error
    // contains filtered or unexported fields
}

```
SQLBuilder







### <a name="DBSQLBuilder">func</a> [DBSQLBuilder](/src/target/sqlbuilder.go?s=10154:10196#L419)
``` go
func DBSQLBuilder(db *DB) (sb *SQLBuilder)
```
DBSQLBuilder


### <a name="NewSQLBuilder">func</a> [NewSQLBuilder](/src/target/sqlbuilder.go?s=9936:9990#L407)
``` go
func NewSQLBuilder(value interface{}) (sb *SQLBuilder)
```
NewSQLBuilder





### <a name="SQLBuilder.BuildCountSQL">func</a> (\*SQLBuilder) [BuildCountSQL](/src/target/sqlbuilder.go?s=6044:6160#L265)
``` go
func (sb *SQLBuilder) BuildCountSQL(where string, args ...interface{}) (query string, argv []interface{}, err error)
```
SQLBuilder.BuildCountSQL




### <a name="SQLBuilder.BuildDeleteSQL">func</a> (\*SQLBuilder) [BuildDeleteSQL](/src/target/sqlbuilder.go?s=7035:7152#L301)
``` go
func (sb *SQLBuilder) BuildDeleteSQL(where string, args ...interface{}) (query string, argv []interface{}, err error)
```
SQLBuilder.BuildDeleteSQL




### <a name="SQLBuilder.BuildInsertSQL">func</a> (\*SQLBuilder) [BuildInsertSQL](/src/target/sqlbuilder.go?s=8272:8356#L347)
``` go
func (sb *SQLBuilder) BuildInsertSQL() (query string, argv []interface{}, err error)
```
SQLBuilder.BuildInsertSQL




### <a name="SQLBuilder.BuildSelectSQL">func</a> (\*SQLBuilder) [BuildSelectSQL](/src/target/sqlbuilder.go?s=6531:6648#L283)
``` go
func (sb *SQLBuilder) BuildSelectSQL(where string, args ...interface{}) (query string, argv []interface{}, err error)
```
SQLBuilder.BuildSelectSQL




### <a name="SQLBuilder.BuildUpdateSQL">func</a> (\*SQLBuilder) [BuildUpdateSQL](/src/target/sqlbuilder.go?s=7505:7622#L319)
``` go
func (sb *SQLBuilder) BuildUpdateSQL(where string, args ...interface{}) (query string, argv []interface{}, err error)
```
SQLBuilder.BuildUpdateSQL




### <a name="SQLBuilder.Count">func</a> (\*SQLBuilder) [Count](/src/target/sqlbuilder.go?s=3237:3290#L148)
``` go
func (sb *SQLBuilder) Count(count *int64) *SQLBuilder
```
SQLBuilder.Count




### <a name="SQLBuilder.Database">func</a> (\*SQLBuilder) [Database](/src/target/sqlbuilder.go?s=9834:9884#L401)
``` go
func (sb *SQLBuilder) Database(db *DB) *SQLBuilder
```
SQLBuilder.Database




### <a name="SQLBuilder.Debug">func</a> (\*SQLBuilder) [Debug](/src/target/sqlbuilder.go?s=1003:1053#L62)
``` go
func (sb *SQLBuilder) Debug(mode bool) *SQLBuilder
```
SQLBuilder.Debug




### <a name="SQLBuilder.Delete">func</a> (\*SQLBuilder) [Delete](/src/target/sqlbuilder.go?s=5001:5062#L217)
``` go
func (sb *SQLBuilder) Delete(dest ...interface{}) *SQLBuilder
```
SQLBuilder.Delete




### <a name="SQLBuilder.Fields">func</a> (\*SQLBuilder) [Fields](/src/target/sqlbuilder.go?s=1246:1301#L74)
``` go
func (sb *SQLBuilder) Fields(fields string) *SQLBuilder
```
SQLBuilder.Fields




### <a name="SQLBuilder.GroupBy">func</a> (\*SQLBuilder) [GroupBy](/src/target/sqlbuilder.go?s=2636:2693#L124)
``` go
func (sb *SQLBuilder) GroupBy(groupBy string) *SQLBuilder
```
SQLBuilder.GroupBy




### <a name="SQLBuilder.Having">func</a> (\*SQLBuilder) [Having](/src/target/sqlbuilder.go?s=2791:2846#L130)
``` go
func (sb *SQLBuilder) Having(having string) *SQLBuilder
```
SQLBuilder.Having




### <a name="SQLBuilder.Insert">func</a> (\*SQLBuilder) [Insert](/src/target/sqlbuilder.go?s=5532:5593#L241)
``` go
func (sb *SQLBuilder) Insert(dest ...interface{}) *SQLBuilder
```
SQLBuilder.Insert




### <a name="SQLBuilder.Limit">func</a> (\*SQLBuilder) [Limit](/src/target/sqlbuilder.go?s=3095:3146#L142)
``` go
func (sb *SQLBuilder) Limit(offset int) *SQLBuilder
```
SQLBuilder.Limit




### <a name="SQLBuilder.Model">func</a> (\*SQLBuilder) [Model](/src/target/sqlbuilder.go?s=8772:8830#L365)
``` go
func (sb *SQLBuilder) Model(value interface{}) *SQLBuilder
```
SQLBuilder.Model




### <a name="SQLBuilder.OrderBy">func</a> (\*SQLBuilder) [OrderBy](/src/target/sqlbuilder.go?s=2941:2998#L136)
``` go
func (sb *SQLBuilder) OrderBy(orderBy string) *SQLBuilder
```
SQLBuilder.OrderBy




### <a name="SQLBuilder.Select">func</a> (\*SQLBuilder) [Select](/src/target/sqlbuilder.go?s=3770:3828#L167)
``` go
func (sb *SQLBuilder) Select(dest interface{}) *SQLBuilder
```
SQLBuilder.Select




### <a name="SQLBuilder.Table">func</a> (\*SQLBuilder) [Table](/src/target/sqlbuilder.go?s=1119:1176#L68)
``` go
func (sb *SQLBuilder) Table(tableName string) *SQLBuilder
```
SQLBuilder.Table




### <a name="SQLBuilder.Update">func</a> (\*SQLBuilder) [Update](/src/target/sqlbuilder.go?s=4470:4531#L193)
``` go
func (sb *SQLBuilder) Update(dest ...interface{}) *SQLBuilder
```
SQLBuilder.Update




### <a name="SQLBuilder.Where">func</a> (\*SQLBuilder) [Where](/src/target/sqlbuilder.go?s=1700:1774#L91)
``` go
func (sb *SQLBuilder) Where(where string, args ...interface{}) *SQLBuilder
```
SQLBuilder.Where




### <a name="SQLBuilder.WhereNot">func</a> (\*SQLBuilder) [WhereNot](/src/target/sqlbuilder.go?s=2320:2397#L113)
``` go
func (sb *SQLBuilder) WhereNot(where string, args ...interface{}) *SQLBuilder
```
SQLBuilder.WhereNot




### <a name="SQLBuilder.WhereOr">func</a> (\*SQLBuilder) [WhereOr](/src/target/sqlbuilder.go?s=2009:2085#L102)
``` go
func (sb *SQLBuilder) WhereOr(where string, args ...interface{}) *SQLBuilder
```
SQLBuilder.WhereOr








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
