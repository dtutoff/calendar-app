package main

import (
	"fmt"
	"time"

	"github.com/SamiRemi/project/app/calendar"
)

func main() {

	event1, err1 := calendar.AddEvent("Поспать", "2026/03/19 03:00")
	if err1 != nil {
		fmt.Println("Ошибка:", err1)
		return
	}

	event2, err2 := calendar.AddEvent("Еще одна встреча", "2026/03/26 20:00")
	if err2 != nil {
		fmt.Println("Ошибка:", err2)
		return
	}

	calendar.ShowEvent()
	calendar.DeleteEvent(event1.ID)

	err := calendar.EditEvent(event2.ID, "Созвон", "2026/03/26 20:00")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	calendar.ShowEvent()

	fmt.Println("Warniiiiiiiiiing")
	fmt.Println("Warniiiiiiiiiing")
	fmt.Println("Warniiiiiiiiiing")

	time.Sleep(5 * time.Second)
}
