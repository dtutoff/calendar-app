package cmd

import (
	"fmt"
	"os"
)

func Exit(c *Cmd) error {
	err := c.calendar.Close()
	if err != nil {
		return err
	}

	c.logger.Add(fmt.Sprintf("%s: %v", ErrFailedToSave, err))
	os.Exit(0)

	return nil
}
