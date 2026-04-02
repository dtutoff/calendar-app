package main

import (
	"fmt"

	"github.com/SamiRemi/project/app/calendar"
	"github.com/SamiRemi/project/app/cmd"
	"github.com/SamiRemi/project/app/storage"
)

func main() {

	s := storage.NewJsonStorage("calendar.Json")
	c := calendar.NewCalendar(s)
	err := c.Load()
	if err != nil {
		fmt.Println("Warning :", err)
		return
	}

	defer c.Save()

	cli := cmd.NewCmd(c)
	cli.Run()
}
