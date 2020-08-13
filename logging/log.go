package logging

import (
	"log"
	"os"
)

type logType int

const (
	INFO logType = 0 + iota
	WARNING
	ERROR
	FATAL
)

func WriteLog(logType logType, message string) {
	logFile, err := os.OpenFile("/var/log/madhyam/log1.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
	}

	switch logType {
	case INFO:
		logger := log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
		logger.Println(message)
	case WARNING:
		logger := log.New(logFile, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)
		logger.Println(message)
	case ERROR:
		logger := log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
		logger.Println(message)
	case FATAL:
		logger := log.New(logFile, "FATAL: ", log.Ldate|log.Ltime|log.Llongfile)
		logger.Println(message)
	}

}
