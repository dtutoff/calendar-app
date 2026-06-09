package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/dtutoff/app/calendar"
	"github.com/dtutoff/app/events"
	"github.com/google/shlex"
)

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrFailedToSave = errors.New("failed to save data")
	ErrEventUpdate  = errors.New("failed to update event")
	ErrEventAdd     = errors.New("error adding event")
	ErrEventDelete  = errors.New("error deleting event")
	ErrReminderAdd  = errors.New("error adding reminder")
)

type Cmd struct {
	calendar *calendar.Calendar
	logger   *Logger
}

func NewCmd(c *calendar.Calendar, logger *Logger) *Cmd {
	return &Cmd{
		calendar: c,
		logger:   logger,
	}
}

func (c *Cmd) executor(input string) {
	parts, err := shlex.Split(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd := strings.ToLower(parts[0])

	switch cmd {
	case "add":
		if len(parts) < 4 {
			fmt.Println(ErrInvalidInput)
			c.logger.Add(fmt.Sprintf("add: %s", ErrInvalidInput))
			fmt.Println(`Syntax: add "name event" "date" "priority"`)
			return
		}

		title := parts[1]
		date := parts[2]
		priority := events.Priority(parts[3])

		e, err := c.calendar.AddEvent(title, date, priority)
		if err != nil {
			fmt.Println(ErrEventAdd, err)
			c.logger.Add(fmt.Sprintf("%s: %v", ErrEventAdd, err))
			return
		}
		err1 := c.calendar.Save()
		if err1 != nil {
			fmt.Println(ErrFailedToSave, err1)
			c.logger.Add(fmt.Sprintf("%s: %v", ErrFailedToSave, err1))
			return
		}
		fmt.Println("Event:", e.Title, "added")
		c.logger.Add(fmt.Sprintf("Event with id %s was added", e.ID))

	case "list":
		c.calendar.ShowEvents()
		c.logger.Add("User viewed all events")

	case "remove":
		if len(parts) < 2 {
			fmt.Println(ErrInvalidInput)
			c.logger.Add(fmt.Sprintf("remove: %s", ErrInvalidInput))
			fmt.Println(`Syntax: remove "id"`)
			return
		}

		eventId := parts[1]

		err := c.calendar.DeleteEvent(parts[1])
		if err != nil {
			fmt.Println(ErrEventDelete, err)
			c.logger.Add(fmt.Sprintf("%s: %v", ErrEventDelete, err))
			return
		}
		err1 := c.calendar.Save()
		if err1 != nil {
			fmt.Println(ErrFailedToSave, err1)
			c.logger.Add(fmt.Sprintf("%s: %v", ErrFailedToSave, err1))
			return
		}
		fmt.Println("Event with ID:", eventId, "was deleted")
		c.logger.Add(fmt.Sprintf("Event with id %s was removed", eventId))

	case "update":

		if len(parts) < 5 {
			fmt.Println(ErrEventUpdate)
			c.logger.Add(fmt.Sprintf("update: %s", ErrFailedToSave))
			fmt.Println(`Syntax: update "id" "name event" "date" "priority"`)
			return
		}

		id := parts[1]
		title := parts[2]
		date := parts[3]
		priority := events.Priority(parts[4])

		err := c.calendar.EditEvent(id, title, date, priority)
		if err != nil {
			fmt.Println(ErrEventUpdate, err)
			c.logger.Add(fmt.Sprintf("%s: %v", ErrEventUpdate, err))
			return
		}
		err1 := c.calendar.Save()
		if err1 != nil {
			fmt.Println(ErrFailedToSave, err1)
			c.logger.Add(ErrFailedToSave.Error())
			return
		}
		fmt.Println("Event:", id, "was updated")
		c.logger.Add(fmt.Sprintf("Event with id %s was updated", id))

	case "remind":
		if len(parts) < 2 {
			fmt.Println(ErrReminderAdd)
			c.logger.Add(fmt.Sprintf("update: %s", ErrReminderAdd))
			fmt.Println(`Syntax: remind "id" "reminder message" "date"`)
			return
		}

		id := parts[1]
		message := parts[2]
		at := parts[3]

		err := c.calendar.SetEventReminder(id, message, at)
		if err != nil {
			fmt.Println(ErrReminderAdd, err)
			c.logger.Add(fmt.Sprintf("%s: %v", ErrReminderAdd, err))
			return
		}
		err1 := c.calendar.Save()
		if err1 != nil {
			fmt.Println(ErrFailedToSave, err1)
			c.logger.Add(ErrFailedToSave.Error())
			return
		}
		fmt.Println("Reminder added")
		c.logger.Add(fmt.Sprintf("Reminder %s was added", message))

	case "exit":
		err := c.calendar.Save()
		if err != nil {
			fmt.Println(ErrFailedToSave, err)
			c.logger.Add(fmt.Sprintf("%s: %v", ErrFailedToSave, err))
		}
		close(c.calendar.Notification)
		os.Exit(0)

	case "help":
		commandList := map[string]string{
			"add":    "Add event - Syntax: add \"name event\" \"date\" \"priority\"",
			"list":   "Show all events - Syntax: list",
			"remove": "Delete event - Syntax: remove \"id\"",
			"update": "Update event - Syntax: update \"id\" \"name event\" \"date\" \"priority\"",
			"remind": "Set event reminder - Syntax: remind \"id\" \"reminder message\" \"date\"",
			"help":   "Show list of commands - Syntax: help",
			"exit":   "Exit program - Syntax: exit",
		}

		fmt.Println("\nAvailable commands:")
		fmt.Println("-------------------")
		for command, desc := range commandList {
			fmt.Printf("  %-8s - %s\n", command, desc)
		}

	default:
		fmt.Println("Unknown command")
		fmt.Println("Enter help for list of commands")
	}
}

func (c *Cmd) completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "add", Description: "Add new event"},
		{Text: "list", Description: "Show events"},
		{Text: "remove", Description: "Delete event"},
		{Text: "update", Description: "Update event"},
		{Text: "remind", Description: "Add event reminder"},
		{Text: "help", Description: "Show commands"},
		{Text: "exit", Description: "Exit program"},
	}

	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func (c *Cmd) Run() {
	p := prompt.New(
		c.executor,
		c.completer,
		prompt.OptionPrefix("> "),
	)
	go func() {
		for message := range c.calendar.Notification {
			fmt.Println(message)
		}
	}()
	p.Run()
}
