package main

import (
	"fmt"

	"github.com/SamiRemi/project/app/calendar"
	"github.com/SamiRemi/project/app/events"
	"github.com/SamiRemi/project/app/storage"
)

func main() {
	zs := storage.NewZipStorage("calendar.zip")
	c := calendar.NewCalendar(zs)

	err := c.Load()
	if err != nil {
		fmt.Println("Warning :", err)
		return
	}

	defer c.Save()

	event1, err1 := c.AddEvent("Поспать", "2026/03/17", events.PriorityHigh)
	if err1 != nil {
		fmt.Println("Ошибка :", err1)
	} else {
		fmt.Println("Событие", "'", event1.Title, "'", "добавлено")
	}

	event2, err2 := c.AddEvent("Встреча", "2026/03/19", events.PriorityLow)
	if err2 != nil {
		fmt.Println("Ошибка :", err2)
	} else {
		fmt.Println("Событие", "'", event2.Title, "'", "добавлено")
	}
	event3, err3 := c.AddEvent("Прогулка", "2026/03/18", events.PriorityMedium)
	if err3 != nil {
		fmt.Println("Ошибка :", err3)
	} else {
		fmt.Println("Событие", "'", event3.Title, "'", "добавлено")
	}
	err4 := c.EditEvent(event2.ID, "Созвон", "2026/03/19 20:00")
	if err4 != nil {
		fmt.Println("Ошибка :", err4)
	}
	c.DeleteEvent(event3.ID)
	c.ShowEvent()
}
