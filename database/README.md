# database
Simple SQLBuilder for Golang

## Example:

- Example 1

```
import (
	...

	"github.com/yodstar/goutil/database/sqlbuilder"
	"github.com/yodstar/goutil/database/mysql"
)

func main() {
	...

	dbmysql.MustOpen(dbconf)

	dbmysql.Select(&dest, where, args...)
	
	dbmysql.Update(&dest, where, args...)

	dbmysql.Delete(&dest, where, args...)

	dbmysql.Insert(&dest)

	dbmysql.Selectx(&dest, query, args...)

	dbmysql.Queryx(query, args...)

	dbmysql.QueryRowx(query, args...)

	dbmysql.Exec(query, args...)

	dbmysql.Where(where, args...).
		GroupBy(groupBy).
		Having(having).
		OrderBy(orderBy).
		Limit(offset).
		Fields(fields).
		Select(&dest).Error

	dbmysql.Where(where, args...).
		Fields(fields).
		Update(&dest).Error

	dbmysql.Where(where, args...).
		Delete(&dest).Error

	dbmysql.Fields(fields).
		Insert(&dest).Error

	dbmysql.Transaction(func(c *sqlbuilder.Conn) error {
		...
	})
}


```

- Example 2

```
import (
	...

	"github.com/yodstar/goutil/database/sqlbuilder"
	"github.com/yodstar/goutil/database/oracle"
)

func main() {
	...

	oracle.MustOpen(dbconf)

	oracle.Select(&dest, where, args...)
	
	oracle.Update(&dest, where, args...)

	oracle.Delete(&dest, where, args...)
	
	oracle.Insert(&dest)

	oracle.Selectx(&dest, query, args...)

	oracle.Queryx(query, args...)

	oracle.QueryRowx(query, args...)
	
	oracle.Exec(query, args...)

	oracle.Transaction(func(c *sqlbuilder.Conn) error {
		...
	})
}


```