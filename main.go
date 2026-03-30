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

	// event1, err1 := c.AddEvent("Поспать", "2026/03/17", events.PriorityHigh)
	// if err1 != nil {
	// 	fmt.Println("Ошибка :", err1)
	// } else {
	// 	fmt.Println("Событие", "'", event1.Title, "'", "добавлено")
	// }
	// event3, err3 := c.AddEvent("Прогулка", "2026/03/18", events.PriorityMedium)
	// if err3 != nil {
	// 	fmt.Println("Ошибка :", err3)
	// } else {
	// 	fmt.Println("Событие", "'", event3.Title, "'", "добавлено")
	// }

	// c.ShowEvent()
}
