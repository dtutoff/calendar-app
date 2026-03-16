package main

import (
	"fmt"

	"github.com/SamiRemi/project/app/calendar"
)

func main() {
	c := calendar.NewCalendar()

	event1, err1 := c.AddEvent("Поспать", "2026/03/17")
	if err1 != nil {
		fmt.Println("Ошибка :", err1)
	} else {
		fmt.Println("Событие", "'", event1.Title, "'", "добавлено")
	}

	event2, err2 := c.AddEvent("Встреча", "2026/03/19")
	if err2 != nil {
		fmt.Println("Ошибка :", err2)
	} else {
		fmt.Println("Событие", "'", event2.Title, "'", "добавлено")
	}
	event3, err3 := c.AddEvent("Прогулка", "2026/03/18")
	if err3 != nil {
		fmt.Println("Ошибка :", err3)
	} else {
		fmt.Println("Событие", "'", event3.Title, "'", "добавлено")
	}
	err := c.EditEvent(event2.ID, "Созвон", "2026/03/19 20:00")
	if err != nil {
		fmt.Println("Ошибка :", err)
	}
	c.DeleteEvent(event3.ID)
	c.ShowEvent()
}
