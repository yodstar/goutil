package logger_test

import (
	"bytes"
	"fmt"
	"log"

	"github.com/yodstar/goutil/logger"
)

var LOG = logger.LOG

func ExampleLogger() {
	var buf bytes.Buffer

	LOG.SetLevel("ALL")
	LOG.SetOutput(&buf)
	LOG.SetFlags(log.Lshortfile)

	LOG.Fatal("Fatal: %s", "FATAL")
	LOG.Alert("Alert: %s", "ALERT")
	LOG.Error("Error: %s", "ERROR")
	LOG.Warn("Warn: %s", "WARN")
	LOG.Notice("Notice: %s", "NOTICE")
	LOG.Info("Info: %s", "INFO")
	LOG.Debug("Debug: %s", "DEBUG")
	LOG.Trace("Trace: %s", "TRACE")

	fmt.Print(&buf)
	// Output:
	// example_test.go:20: [FATAL] Fatal: FATAL
	// example_test.go:21: [ALERT] Alert: ALERT
	// example_test.go:22: [ERROR] Error: ERROR
	// example_test.go:23: [WARN] Warn: WARN
	// example_test.go:24: [NOTICE] Notice: NOTICE
	// example_test.go:25: [INFO] Info: INFO
	// example_test.go:26: [DEBUG] Debug: DEBUG
	// example_test.go:27: [TRACE] Trace: TRACE
}

func ExampleLoggerSetFilter() {
	var (
		buf1 bytes.Buffer
		log1 = log.New(&buf1, "", log.Lshortfile)

		buf2 bytes.Buffer
		log2 = logger.New(&buf2, "", log.Lshortfile)
	)

	log2.SetLevel("INFO")
	log2.SetFilter("DEBUG", func(s string) { log1.Output(4, s) })

	log2.Fatal("Fatal: %s", "FATAL")
	log2.Alert("Alert: %s", "ALERT")
	log2.Error("Error: %s", "ERROR")
	log2.Warn("Warn: %s", "WARN")
	log2.Notice("Notice: %s", "NOTICE")
	log2.Info("Info: %s", "INFO")
	log2.Debug("Debug: %s", "DEBUG")
	log2.Trace("Trace: %s", "TRACE")

	fmt.Printf("LOG1:\n%s", &buf1)
	fmt.Printf("LOG2:\n%s", &buf2)
	// Output:
	// LOG1:
	// example_test.go:53: [FATAL] Fatal: FATAL
	// example_test.go:54: [ALERT] Alert: ALERT
	// example_test.go:55: [ERROR] Error: ERROR
	// example_test.go:56: [WARN] Warn: WARN
	// example_test.go:57: [NOTICE] Notice: NOTICE
	// example_test.go:58: [INFO] Info: INFO
	// example_test.go:59: [DEBUG] Debug: DEBUG
	// LOG2:
	// example_test.go:53: [FATAL] Fatal: FATAL
	// example_test.go:54: [ALERT] Alert: ALERT
	// example_test.go:55: [ERROR] Error: ERROR
	// example_test.go:56: [WARN] Warn: WARN
	// example_test.go:57: [NOTICE] Notice: NOTICE
	// example_test.go:58: [INFO] Info: INFO
}
