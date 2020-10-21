# logger
Simple Logger for Golang

## Example:

- Example 1

```
import (
	...

	"github.com/yodstar/goutil/logger"
)

var (
	...

	LOG  = logger.LOG
)

func main() {
	...

	// logger
	LOG.SetLevel("INFO")
	LOG.SetOutFile("./logs/example1_%.log", "200601")
	LOG.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	LOG.SetFilter("WARN", func(s string) { log.Output(4, s) })

	...

	// error
	LOG.Error("error: %s", err.Error())

/*
	LOG.Fatal(fmt, v...)
	LOG.Alert(fmt, v...)
	LOG.Error(fmt, v...)
	LOG.Warn(fmt, v...)
	LOG.Notice(fmt, v...)
	LOG.Info(fmt, v...)
	LOG.Debug(fmt, v...)
	LOG.Trace(fmt, v...)
*/

}


```

- Example 2

```
import (
	...

	"github.com/yodstar/goutil/logger"
)

var (
	...
	log = logger.New(os.Stderr, "", log.LstdFlags)
)

func main() {
	...

	// logger
	log.SetLevel("INFO")
	log.SetOutFile("./logs/example2_%.log", "200601")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	...

	// error
	log.Error("error: %s", err.Error())

/*
	log.Fatal(fmt, v...)
	log.Alert(fmt, v...)
	log.Error(fmt, v...)
	log.Warn(fmt, v...)
	log.Notice(fmt, v...)
	log.Info(fmt, v...)
	log.Debug(fmt, v...)
	log.Trace(fmt, v...)
*/

}


```
