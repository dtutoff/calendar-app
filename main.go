package main

import (
	"fmt"

	"github.com/SamiRemi/project/app/calendar"
	"github.com/SamiRemi/project/app/cmd"
	"github.com/SamiRemi/project/app/storage"
)

// func Add(a, b int) int {
// 	return a + b
// }

// func CheckPositive(num int) error {
// 	if num <= 0 {
// 		return errors.New("лещь от программы")
// 	}
// 	return nil
// }

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
