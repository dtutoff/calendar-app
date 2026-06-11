package cmd

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/dtutoff/app/calendar"
	"github.com/google/shlex"
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

	if len(parts) == 0 {
		return
	}

	cmd := strings.ToLower(parts[0])
	args := parts[1:]

	switch cmd {
	case "add":
		err := Add(c, args)
		if err != nil {
			fmt.Println(err)
			return
		}

	case "list":
		c.calendar.ShowEvents()
		c.logger.Add("User viewed all events")

	case "remove":
		err := Remove(c, args)
		if err != nil {
			fmt.Println(err)
			return
		}

	case "update":
		err := Update(c, args)
		if err != nil {
			fmt.Println(err)
			return
		}

	case "remind":
		err := Remind(c, args)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Reminder added")

	case "exit":
		err := Exit(c)
		if err != nil {
			fmt.Println(err)
			return
		}

	case "help":
		Help()

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
