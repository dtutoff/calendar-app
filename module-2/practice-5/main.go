package main

import (
	"fmt"

	"github.com/dtutoff/app/calendar"
)

func main() {
	event1, err1 := calendar.AddEvent("Встреча", "2025/06/12 16:33")
	if err1 != nil {
		fmt.Println("Ошибка:", err1)
		return
	}

	event2, err2 := calendar.AddEvent("Еще одна встреча", "2025/06/12 15:00")
	if err2 != nil {
		fmt.Println("Ошибка:", err2)
		return
	}

	calendar.ShowEvents()
	calendar.DeleteEvent(event1.ID)

	err := calendar.EditEvent(event2.ID, "Созвон", "2025/06/12 16:50")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	calendar.ShowEvents()
}
