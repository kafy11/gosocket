package log

import (
	"log"
	"os"
)

func writeLog(filename string, s ...interface{}) *log.Logger {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(s...)
	file.Close()
}

func Error(s ...interface{}) {
	writeLog("error_log.txt", s...)
}

func Warning(s ...interface{}) {
	writeLog("warning_log.txt", s...)
}

func Info(s ...interface{}) {
	writeLog("info_log.txt", s...)
}

func Fatal(s ...interface{}) {
	writeLog("error_log.txt", s...)
	log.Fatal(s...)
}
