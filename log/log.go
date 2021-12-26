package log

import (
	"log"
	"os"
)

var (
	WarningLog *log.Logger
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
)

func init() {
	InfoLog = openLogFile("info_log.txt")
	WarningLog = openLogFile("warning_log.txt")
	ErrorLog = openLogFile("error_log.txt")
}

func openLogFile(filename string) *log.Logger {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func Error(s ...interface{}) {
	ErrorLog.Println(s...)
}

func Warning(s ...interface{}) {
	WarningLog.Println(s...)
}

func Info(s ...interface{}) {
	InfoLog.Println(s...)
}
