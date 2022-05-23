package log

import (
	"log"
	"os"
	"sync"
)

type logFile struct {
	name string
	m    sync.Mutex
}

var errorLog, infoLog, warningLog *logFile

func init() {
	infoLog = &logFile{name: "info_log.txt"}
	warningLog = &logFile{name: "warning_log.txt"}
	errorLog = &logFile{name: "error_log.txt"}
}

func (logFile *logFile) write(s ...interface{}) {
	logFile.m.Lock()
	defer logFile.m.Unlock()

	file, err := os.OpenFile(logFile.name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(s...)
}

func Error(s ...interface{}) {
	errorLog.write(s...)
}

func Warning(s ...interface{}) {
	warningLog.write(s...)
}

func Info(s ...interface{}) {
	infoLog.write(s...)
}

func Fatal(s ...interface{}) {
	errorLog.write(s...)
	log.Fatal(s...)
}
