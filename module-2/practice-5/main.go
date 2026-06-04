package main

import (
	"fmt"

	"github.com/dtutoff/app/calendar"
	"github.com/dtutoff/app/storage"
)

func main() {
	s, err4 := storage.NewStorage("./calendar.json")
	if err4 != nil {
		fmt.Println("Error:", err4)
	}
	c := calendar.NewCalendar(s)

	z := storage.NewZipStorage("./calendar.zip")
	c1 := calendar.NewCalendar(z)
	err5 := c1.Save()
	if err5 != nil {
		return
	}

	err := c.Load()
	if err != nil {
		fmt.Println("Loading error:", err)
		return
	}

	event1, err1 := c.AddEvent("Meeting will be here!!!!!", "2025/06/12", "high")
	if err1 != nil {
		fmt.Println("Error:", err1)
		return
	} else {
		err := c.SetEventReminder(event1.ID, event1.Title, "2025/07/12")
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(event1.Title, "added")
	}

	if event1.Reminder != nil {
		event1.Reminder.Send()
	} else {
		fmt.Println("No Reminder")
	}

	err3 := c.Save()
	if err3 != nil {
		return
	}

	c.ShowEvents()
}
