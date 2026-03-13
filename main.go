package main

import (
	"fmt"

	"github.com/smokizazzi/app/calendar"
)

func main() {
	event1, err := calendar.AddEvent("Встреча", "2025/06/12 16:33")
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}

	event2, err := calendar.AddEvent("Обед", "15 Jul 2024 13:00")
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}

	event3, err := calendar.AddEvent("Звонок", "2024-07-15 09:30")
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}

	calendar.ShowEvents()

	_, err = calendar.AddEvent("AB", "2025/06/12 16:33")
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}

	err = calendar.EditEvent(event1.ID, "Важная встреча", "2025/06/15 10:00")
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}

	calendar.ShowEvents()

	err = calendar.EditEvent("nevere-n0t-4exi-stin-guuid", "Тест", "2025/06/15 10:00")
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}

	err = calendar.DeleteEvent(event2.ID)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}

	calendar.ShowEvents()

	err = calendar.DeleteEvent(event2.ID)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}

	calendar.ShowEvents()

	fmt.Printf("ID событий для справки:\n")
	fmt.Printf("event1 ID: %s\n", event1.ID)
	fmt.Printf("event2 ID: %s\n", event2.ID)
	fmt.Printf("event3 ID: %s\n", event3.ID)
}
