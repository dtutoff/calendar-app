package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/SamiRemi/project/app/calendar"
	"github.com/SamiRemi/project/app/events"
	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
)

type Cmd struct {
	calendar *calendar.Calendar
}

func NewCmd(c *calendar.Calendar) *Cmd {
	return &Cmd{
		calendar: c,
	}
}

func (c *Cmd) Run() {
	go func() {
		for msg := range c.calendar.Notification {
			fmt.Println("Напоминание :", msg)
		}
	}()
	p := prompt.New(
		c.executor,
		c.completer,
		prompt.OptionPrefix("> "),
	)
	p.Run()
}

func (c *Cmd) executor(input string) {
	err := c.calendar.Load()
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
	}
	parts, err := shlex.Split(input)
	if err != nil {
		fmt.Println("Ошибка добавления:", err)
	}

	cmd := strings.ToLower(parts[0])

	switch cmd {
	case "add":
		if len(parts) < 4 {
			fmt.Println("Формат: add \"название события\" \"дата и время\" \"приоритет\"")
			return
		}

		title := parts[1]
		date := parts[2]
		priority := events.Priority(parts[3])

		e, err := c.calendar.AddEvent(title, date, priority)
		if err != nil {
			fmt.Println("Ошибка добавления:", err)
		} else {
			fmt.Println("Событие:", e.Title, "добавлено")
		}

	case "remove":
		if len(parts) != 2 {
			fmt.Println("Формат: remove \"ID события\"")
			return
		}

		ID := parts[1]
		err := c.calendar.DeleteEvent(ID)
		if err != nil {
			fmt.Println("Ошибка", err)
			return
		}
	case "update":
		if len(parts) != 5 {
			fmt.Println("Формат: update \"ID\" \"Новое название события\" \"Новоя дата и время\" \"Новый приоритет\"")
			return
		}

		ID := parts[1]
		title := parts[2]
		date := parts[3]
		priority := events.Priority(parts[4])

		err := c.calendar.EditEvent(ID, title, date, priority)
		if err != nil {
			fmt.Println("Ошибка изминения :", err)
		}
		fmt.Println("Событие:", title, "Изминено")

	case "list":
		err := c.calendar.ShowEvent()
		if err != nil {
			fmt.Println("Формат: list")
			fmt.Println(err)
		}

	case "setreminder":
		if len(parts) != 4 {
			fmt.Println("warning")
			fmt.Println("Format")
			return
		}
		ID := parts[1]
		Message := parts[2]
		DateStr := parts[3]
		err := c.calendar.SetEventReminder(ID, Message, DateStr)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Напоминание успешно добавлено!")
		fmt.Println("")

	case "cancelreminder":
		if len(parts) != 2 {
			fmt.Println("warning")
			fmt.Println("Format")
		}
		ID := parts[1]
		err := c.calendar.CancelEventReminder(ID)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "help":
		if len(parts) != 1 {
			fmt.Println("Формат: help")
			return
		}
		fmt.Println("=============================================================================")
		fmt.Println("Cписок команд:")
		fmt.Println("")
		fmt.Println("add - Формат: add \"название события\" \"дата и время\" \"приоритет\"")
		fmt.Println("Добавляет событие")
		fmt.Println("")
		fmt.Println("list - Формат: list")
		fmt.Println("Показывает список всех событий")
		fmt.Println("")
		fmt.Println("remove - Формат: remove \"ID события\"")
		fmt.Println("Удаляет событие по ID")
		fmt.Println("")
		fmt.Println("update - Формат: update \"ID\" \"Новое название события\" \"Новоя дата и время\" \"Новый приоритет\"")
		fmt.Println("Редактирует событие по ID")
		fmt.Println("")
		fmt.Println("exit - Формат: exit")
		fmt.Println("Закрывает и сохраняет программу")
		fmt.Println("")
		fmt.Println("exinot - формат: exinot")
		fmt.Println("Закрывает программу без сохранения")
		fmt.Println("=============================================================================")

	case "exinot":
		os.Exit(0)

	case "exit":
		err := c.calendar.Save()
		if err != nil {
			fmt.Println(err)
			return
		}
		close(c.calendar.Notification)
		os.Exit(0)

	default:
		fmt.Println("Неизвестная команда:")
		fmt.Println("Введите 'help' для списка команд")
	}
}

func (c *Cmd) completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "add", Description: "Добавить событие"},
		{Text: "list", Description: "Показать все события"},
		{Text: "remove", Description: "Удалить событие"},
		{Text: "update", Description: "Редактировать собите"},
		{Text: "help", Description: "Показать справку"},
		{Text: "exit", Description: "Выйти и cохранить программу"},
		{Text: "exinot", Description: "Выйти без сохранения из программы"},
		{Text: "setreminder", Description: "xz"},
		{Text: "cancelreminder", Description: "xz"},
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}
