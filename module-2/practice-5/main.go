package main

import (
	"fmt"

	"github.com/dtutoff/app/calendar"
	"github.com/dtutoff/app/cmd"
	"github.com/dtutoff/app/storage"
)

func main() {
	s, err := storage.NewStorage("./calendar.json")
	if err != nil {
		fmt.Println("Error creating calendar storage:", err)
		return
	}
	c := calendar.NewCalendar(s)

	l, err := storage.NewLogStorage("./logs.txt")
	if err != nil {
		fmt.Println("Error creating log storage:", err)
		return
	}

	logger := cmd.NewLogger(l)
	cli := cmd.NewCmd(c, logger)

	err = c.Load()
	if err != nil {
		fmt.Println("Loading error:", err)
		return
	}

	cli.Run()
}
