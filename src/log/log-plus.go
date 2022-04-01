package log

import (
	"log"
	"os"
)

var (
	INFO  func(v ...interface{})
	WARN  func(v ...interface{})
	ERROR func(v ...interface{})
)

func init() {
	_INFO := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	_WARN := log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	_ERROR := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	INFO = _INFO.Println
	WARN = _WARN.Println
	ERROR = _ERROR.Panicln
}
