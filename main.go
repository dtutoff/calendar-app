package main

import (
	"fmt"

	"github.com/SamiRemi/project/app/calendar"
	"github.com/SamiRemi/project/app/cmd"
	"github.com/SamiRemi/project/app/logger"
	"github.com/SamiRemi/project/app/storage"
)

func main() {
	logger.CreateLogger("app.log")
	defer logger.CloseFile()
	logger.System("Запуск системы")
	s := storage.NewJsonStorage("calendar.Json")
	logger.System("Создано хранилище")
	c := calendar.NewCalendar(s)
	err1 := c.Load()
	if err1 != nil {
		fmt.Println("Warning :", err1)
		return
	}
	defer c.Save()

	cli := cmd.NewCmd(c)
	cli.Run()
}
