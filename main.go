package main

import (
	"github.com/SamiRemi/project/app/calendar"
	"github.com/SamiRemi/project/app/storage"
)

func main() {
	filename := storage.NewStorage("calendar.json")
	c := calendar.NewCalendar(filename)
	defer c.Save()
}
