package logger

import (
	"log"
	"os"
)

var (
	infoLogger   *log.Logger
	errorLogger  *log.Logger
	systemLogger *log.Logger
	file         *os.File
)

func CreateLogger(filename string) error {
	Info("Запуск функции CreateLogger")
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	systemLogger = log.New(file, "SYSTEM: ", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

func Info(msg string) {
	infoLogger.Output(2, msg)
}

func Error(msg string) {
	errorLogger.Output(2, msg)
}

func System(msg string) {
	systemLogger.Output(2, msg)
}

func CloseFile() error {
	return file.Close()
}
