package cmd

import (
	"fmt"

	"github.com/dtutoff/app/events"
)

func Update(c *Cmd, args []string) error {
	if len(args) < 4 {
		c.logger.Add(fmt.Sprintf("update: %s", ErrEventUpdate))
		fmt.Println(`Syntax: update "id" "name event" "date" "priority"`)
		return ErrEventUpdate
	}

	id := args[0]
	title := args[1]
	date := args[2]
	priority := events.Priority(args[3])

	err := c.calendar.EditEvent(id, title, date, priority)
	if err != nil {
		c.logger.Add(fmt.Sprintf("%s: %v", ErrEventUpdate, err))
		return ErrEventUpdate
	}
	err1 := c.calendar.Save()
	if err1 != nil {
		c.logger.Add(ErrFailedToSave.Error())
		return ErrFailedToSave
	}
	fmt.Println("Event:", id, "was updated")
	c.logger.Add(fmt.Sprintf("Event with id %s was updated", id))
	return nil
}
