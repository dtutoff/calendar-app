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
	logg     *Log
}

func NewCmd(c *calendar.Calendar) *Cmd {
	return &Cmd{
		calendar: c,
		logg:     NewLogger("app.logger"),
	}
}

func (c *Cmd) Run() {
	err := c.logg.LoadFromFile()
	if err != nil {
		warningMsg := ("Предупреждение: не удалось загрузить лог из файла: " + err.Error())
		fmt.Println(warningMsg)
		c.logg.Log(warningMsg)
	} else {
		loadMsg := ("Лог загружен из файла: " + c.logg.filePath)
		c.logg.Log(loadMsg)
	}
	c.logg.Log("Программа запущена")

	go func() {
		for msg := range c.calendar.Notification {
			notificationMsg := "Напоминание :" + msg
			fmt.Println(notificationMsg)
			c.logg.Log(notificationMsg)
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
	c.logg.Log("Ввод пользователя :" + input)
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
			errorMsg := ("Ошибка : неверный формат команды add")
			fmt.Println("Неверный формат команды")
			fmt.Println("Формат: add \"название события\" \"дата и время\" \"приоритет\"")
			c.logg.Log(errorMsg)
			return
		}
		title := parts[1]
		date := parts[2]
		priority := events.Priority(parts[3])
		e, err := c.calendar.AddEvent(title, date, priority)
		if err != nil {
			errorMessage := ("Ошибка добавления: " + err.Error())
			fmt.Println(errorMessage)
			c.logg.Log(errorMessage)
		} else {
			successMsg := ("Событие: " + e.Title + " добавлено")
			fmt.Println(successMsg)
			c.logg.Log(successMsg)
		}

	case "remove":
		if len(parts) != 2 {
			errorMsg := ("Ошибка: неверный формат команды remove")
			fmt.Println("Неверный формат команды")
			fmt.Println("Формат: remove \"ID события\"")
			c.logg.Log(errorMsg)
			return
		}
		ID := parts[1]
		err := c.calendar.DeleteEvent(ID)
		if err != nil {
			errorMessage := ("Ошибка удаления: " + err.Error())
			fmt.Println(errorMessage)
			c.logg.Log(errorMessage)
		} else {
			successMsg := ("Событие с ID " + ID + " удалено")
			fmt.Println(successMsg)
			c.logg.Log(successMsg)
		}

	case "update":
		if len(parts) != 5 {
			errorMsg := ("Ошибка: неверный формат команды update")
			fmt.Println("Неверный формат команды")
			fmt.Println("Формат: update \"ID\" \"Новое название события\" \"Новая дата и время\" \"Новый приоритет\"")
			c.logg.Log(errorMsg)
			return
		}
		ID := parts[1]
		title := parts[2]
		date := parts[3]
		priority := events.Priority(parts[4])
		err := c.calendar.EditEvent(ID, title, date, priority)
		if err != nil {
			errorMessage := ("Ошибка изменения: " + err.Error())
			fmt.Println(errorMessage)
			c.logg.Log(errorMessage)
		} else {
			successMsg := ("Событие с ID " + ID + " обновлено: " + title)
			fmt.Println(successMsg)
			c.logg.Log(successMsg)
		}

	case "list":
		c.logg.Log("Команда list: получение списка событий")
		err := c.calendar.ShowEvent()
		if err != nil {
			errorMessage := ("Ошибка при получении списка событий: " + err.Error())
			fmt.Println("Ошибка при выполнении list")
			fmt.Println(errorMessage)
			c.logg.Log(errorMessage)
		} else {
			successMsg := ("Список событий успешно выведен")
			fmt.Println(successMsg)
			c.logg.Log(successMsg)
		}

	case "setreminder":
		if len(parts) != 4 {
			errorMsg := ("Ошибка: неверный формат команды setreminder")
			fmt.Println("Неверный формат команды")
			fmt.Println("Формат: setreminder \"ID события\" \"сообщение\" \"дата и время\"")
			c.logg.Log(errorMsg)
			return
		}
		ID := parts[1]
		Message := parts[2]
		DateStr := parts[3]
		err := c.calendar.SetEventReminder(ID, Message, DateStr)
		if err != nil {
			errorMessage := ("Ошибка установки напоминания: " + err.Error())
			fmt.Println(errorMessage)
			c.logg.Log(errorMessage)
		} else {
			successMsg := ("Напоминание для события ID " + ID + " успешно установлено")
			fmt.Println("Напоминание успешно добавлено!")
			fmt.Println("")
			c.logg.Log(successMsg)
		}

	case "cancelreminder":
		if len(parts) != 2 {
			errorMsg := ("Ошибка: неверный формат команды cancelreminder")
			fmt.Println("Неверный формат команды")
			fmt.Println("Формат: cancelreminder \"ID события\"")
			c.logg.Log(errorMsg)
			return
		}
		ID := parts[1]
		err := c.calendar.CancelEventReminder(ID)
		if err != nil {
			errorMessage := ("Ошибка отмены напоминания: " + err.Error())
			fmt.Println(errorMessage)
			c.logg.Log(errorMessage)
		} else {
			successMsg := ("Напоминание для события ID " + ID + " отменено")
			fmt.Println(successMsg)
			c.logg.Log(successMsg)
		}

	case "log":
		if len(parts) != 1 {
			fmt.Println("Формат : log")
			c.logg.Log("Ошибка : неверный формат команды log")
			return
		}
		fmt.Println("=== ЛОГ ПРОГРАММЫ ===")
		for _, entry := range c.logg.entries {
			fmt.Println(entry)
		}
		fmt.Println("===================")
		c.logg.Log("Выведен лог программы")

	case "savelog":
		err := c.logg.SaveToFile()
		if err != nil {
			errorMessage := ("Ошибка сохранения лога: " + err.Error())
			fmt.Println(errorMessage)
			c.logg.Log(errorMessage)
		} else {
			successMsg := ("Лог успешно сохранён в файл: " + c.logg.filePath)
			fmt.Println(successMsg)
			c.logg.Log(successMsg)
		}

	case "loadlog":
		err := c.logg.LoadFromFile()
		if err != nil {
			errorMessage := ("Ошибка загрузки лога: " + err.Error())
			fmt.Println(errorMessage)
			c.logg.Log(errorMessage)
		} else {
			successMsg := ("Лог успешно загружен из файла: " + c.logg.filePath)
			fmt.Println(successMsg)
			c.logg.Log(successMsg)
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
		fmt.Println("setreminder - Формат: setreminder \"ID события\" \"сообщение\" \"дата и время\"")
		fmt.Println("Дабовляет напоминание событию")
		fmt.Println("")
		fmt.Println("cancelreminder - Формат: cancelreminder \"ID события\"")
		fmt.Println("Удаляет напоминание у события")
		fmt.Println("")
		fmt.Println("log - Формат: log")
		fmt.Println("Показывает весь лог программы")
		fmt.Println("")
		fmt.Println("savelog - Формат: savelog")
		fmt.Println("Сохраняет лог в файл")
		fmt.Println("")
		fmt.Println("loadlog - Формат: loadlog")
		fmt.Println("Загружает лог из файла")
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
		saveErr := c.logg.SaveToFile()
		if saveErr != nil {
			errorMessage := ("Ошибка сохранения лога: " + saveErr.Error())
			fmt.Println(errorMessage)
			c.logg.Log(errorMessage)
		} else {
			successMsg := ("Лог успешно сохранён в файл: " + c.logg.filePath)
			c.logg.Log(successMsg)
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
		{Text: "setreminder", Description: "Дабовить напоминание событию"},
		{Text: "cancelreminder", Description: "Удалить напоминание у события"},
		{Text: "log", Description: "Показать лог программы"},
		{Text: "savelog", Description: "Сохранить лог в файл"},
		{Text: "loadlog", Description: "Загрузить лог из файла"},
		{Text: "exinot", Description: "Выйти без сохранения из программы"},
		{Text: "exit", Description: "Выйти и cохранить программу"},
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}
