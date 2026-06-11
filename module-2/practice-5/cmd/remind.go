package cmd

import (
	"fmt"
)

func Remind(c *Cmd, args []string) error {
	if len(args) < 3 {
		fmt.Println(`Syntax: remind "id" "reminder message" "date"`)
		c.logger.Add(fmt.Sprintf("update: %s", ErrReminderAdd))
		return ErrReminderAdd
	}

	id := args[0]
	message := args[1]
	at := args[2]

	err := c.calendar.SetEventReminder(id, message, at)
	if err != nil {
		c.logger.Add(fmt.Sprintf("%s: %v", ErrReminderAdd, err))
		return err
	}

	err1 := c.calendar.Save()
	if err1 != nil {
		c.logger.Add(ErrFailedToSave.Error())
		return ErrFailedToSave
	}
	c.logger.Add(fmt.Sprintf("Reminder %s was added", message))

	return nil
}
