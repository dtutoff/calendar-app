package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/dtutoff/app/calendar"
	"github.com/dtutoff/app/events"
	"github.com/google/shlex"
)

type Cmd struct {
	calendar *calendar.Calendar
}

func NewCmd(c *calendar.Calendar) *Cmd {
	return &Cmd{calendar: c}
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
			fmt.Println("invalid input")
			fmt.Println(`Syntax: add "id" "name event" "date" "priority"`)
			return
		}

		title := parts[1]
		date := parts[2]
		priority := events.Priority(parts[3])

		e, err := c.calendar.AddEvent(title, date, priority)
		if err != nil {
			fmt.Println("Error adding event:", err)
		} else {
			err := c.calendar.Save()
			if err != nil {
				fmt.Println("Error saving calendar:", err)
			}
			fmt.Println("Event:", e.Title, "added")
		}
	case "list":
		c.calendar.ShowEvents()

	case "remove":
		eventId := parts[1]

		err := c.calendar.DeleteEvent(parts[1])
		if err != nil {
			fmt.Println("Error deleting event:", err)
		} else {
			err := c.calendar.Save()
			if err != nil {
				fmt.Println("Error saving calendar:", err)
			}
			fmt.Println("Event with ID:", eventId, "was deleted")
		}

	case "update":

		if len(parts) < 5 {
			fmt.Println("invalid input")
			fmt.Println(`Syntax: update "id" "name event" "date" "priority"`)
			return
		}

		id := parts[1]
		title := parts[2]
		date := parts[3]
		priority := events.Priority(parts[4])

		err := c.calendar.EditEvent(id, title, date, priority)
		if err != nil {
			fmt.Println("Error updating event:", err)
		} else {
			err := c.calendar.Save()
			if err != nil {
				fmt.Println("Error saving calendar:", err)
			}
		}
	case "remind":

		id := parts[1]
		message := parts[2]
		at := parts[3]

		err := c.calendar.SetEventReminder(id, message, at)
		if err != nil {
			fmt.Println("Error reminding event:", err)
		} else {
			err := c.calendar.Save()
			if err != nil {
				fmt.Println("Error saving calendar:", err)
			}
		}

	case "exit":
		err := c.calendar.Save()
		if err != nil {
			fmt.Println("Saving error:", err)
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
