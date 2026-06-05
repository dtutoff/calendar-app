package main

import (
	"fmt"

	"github.com/dtutoff/app/calendar"
	"github.com/dtutoff/app/cmd"
	"github.com/dtutoff/app/storage"
)

func main() {
	s, err4 := storage.NewStorage("./calendar.json")
	if err4 != nil {
		fmt.Println("Error:", err4)
	}
	c := calendar.NewCalendar(s)

	err := c.Load()
	if err != nil {
		fmt.Println("Loading error:", err)
		return
	}

	cli := cmd.NewCmd(c)
	cli.Run()
}
