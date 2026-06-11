package cmd

import "fmt"

func Remove(c *Cmd, args []string) error {
	if len(args) < 1 {
		c.logger.Add(fmt.Sprintf("remove: %s", ErrInvalidInput))
		fmt.Println(`Syntax: remove "id"`)
		return ErrInvalidInput
	}

	eventId := args[0]

	err := c.calendar.DeleteEvent(eventId)
	if err != nil {
		c.logger.Add(fmt.Sprintf("%s: %v", ErrEventDelete, err))
		return ErrEventDelete
	}
	err1 := c.calendar.Save()
	if err1 != nil {
		c.logger.Add(fmt.Sprintf("%s: %v", ErrFailedToSave, err1))
		return ErrFailedToSave
	}

	fmt.Println("Event with ID:", eventId, "was deleted")
	c.logger.Add(fmt.Sprintf("Event with id %s was removed", eventId))
	return nil
}
