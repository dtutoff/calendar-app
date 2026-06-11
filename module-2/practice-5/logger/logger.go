package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	logFile  *os.File
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func InitLogger() error {
	var err error

	logFile, err = os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	InfoLog = log.New(logFile, "INFO: ", log.Ldate|log.Ltime)
	ErrorLog = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

func Info(msg string) {
	if InfoLog == nil {
		return
	}

	if err := InfoLog.Output(2, msg); err != nil {
		log.Printf("Error writing Info to file: %v", err)
	}
}

func Error(msg string) {
	if ErrorLog == nil {
		return
	}

	if err := ErrorLog.Output(2, msg); err != nil {
		log.Printf("Error writing Error to file: %v", err)
	}
}

func Close() {
	err := logFile.Close()
	if err != nil {
		fmt.Println(err)
	}
}
