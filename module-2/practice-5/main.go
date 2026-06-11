package main

import (
	"fmt"

	"github.com/dtutoff/app/calendar"
	"github.com/dtutoff/app/cmd"
	"github.com/dtutoff/app/logger"
	"github.com/dtutoff/app/storage"
)

func main() {
	if err := logger.InitLogger(); err != nil {
		panic("Logger init error: " + err.Error())
	}
	defer logger.Close()

	logger.Info("Application started")

	s, err := storage.NewJSONFileStorage("./calendar.json")
	if err != nil {
		logger.Error("Error opening calendar.json file")
		fmt.Println("Error creating calendar storage:", err)
		return
	}
	c := calendar.NewCalendar(s)
	logger.Info("Calendar initialized")

	l, err := storage.NewTextFileStorage("./logs.txt")
	if err != nil {
		logger.Error("Cannot open logs.txt")
		fmt.Println("Error creating log storage:", err)
		return
	}

	cmdLogger := cmd.NewLogger(l)
	logger.Info("Cmd Logger initialized")
	cli := cmd.NewCmd(c, cmdLogger)

	err = c.Load()
	if err != nil {
		logger.Error("Cannot load calendar.json")
		fmt.Println("Loading error:", err)
		return
	}

	logger.Info("CLI started")
	cli.Run()
}
