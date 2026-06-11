package cmd

import (
	"fmt"

	"github.com/dtutoff/app/events"
)

func Add(c *Cmd, args []string) error {
	fmt.Println(args)
	if len(args) < 3 {
		c.logger.Add(fmt.Sprintf("add: %s", ErrInvalidInput))
		fmt.Println(`Syntax: add "name event" "date" "priority"`)
		return ErrInvalidInput
	}

	title := args[0]
	date := args[1]
	priority := events.Priority(args[2])

	e, err := c.calendar.AddEvent(title, date, priority)
	if err != nil {
		c.logger.Add(fmt.Sprintf("%s: %v", ErrEventAdd, err))
		return ErrEventAdd
	}
	err1 := c.calendar.Save()
	if err1 != nil {
		c.logger.Add(fmt.Sprintf("%s: %v", ErrFailedToSave, err1))
		return ErrFailedToSave
	}

	fmt.Println("Event:", e.Title, "added")
	c.logger.Add(fmt.Sprintf("Event with id %s was added", e.ID))

	return nil
}
