package calendar

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/SamiRemi/project/app/events"
	"github.com/SamiRemi/project/app/logger"
	"github.com/SamiRemi/project/app/storage"
	"github.com/SamiRemi/project/app/validation"
	"github.com/araddon/dateparse"
)

type Calendar struct {
	calendarEvents map[string]*events.Event
	storage        storage.Store
	Notification   chan string
	closed         bool
}

func NewCalendar(s storage.Store) *Calendar {
	return &Calendar{
		calendarEvents: make(map[string]*events.Event),
		storage:        s,
		Notification:   make(chan string),
		closed:         false,
	}
}

func show() {
	fmt.Println("=========================")
}

func (c *Calendar) AddEvent(title, date string, priority events.Priority) (*events.Event, error) {
	logger.Info("Запуск функции AddEvent")
	e, err := events.NewEvent(title, date, priority)
	if err != nil {
		return nil, err
	}
	if _, ok := c.calendarEvents[title]; ok {
		return nil, fmt.Errorf("Не удается создать событие : %w", validation.TitleError)
	}
	if len(title) == 0 {
		return nil, fmt.Errorf("Не удается создать событие : %w", validation.ListError)
	}
	c.calendarEvents[e.ID] = e
	return e, nil
}

func (c *Calendar) Notify(msg string) {
	c.Notification <- msg
}

func (c *Calendar) SetEventReminder(ID, message, dateStr string) error {
	logger.Info("Запуск функции SetEventReminder")
	if c == nil {
		return fmt.Errorf("Не удается установить напоминание о событие : %w", validation.EqualError)
	}
	event, exists := c.calendarEvents[ID]
	if !exists {
		return fmt.Errorf("Не удается установить напоминание о событие : %w", validation.EventNotFoundError)
	}
	if event.Reminder != nil {
		return fmt.Errorf("Не удается установить напоминание о событие : %w", validation.ReminderAlreadyExistError)
	}
	startAt, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return fmt.Errorf("Не удается установить напоминание о событие : %w", validation.DateFormatError)
	}
	if startAt.Before(time.Now()) {
		return fmt.Errorf("Не удается установить напоминание о событие : %w", validation.ReminderDateError)
	}

	err = event.AddReminder(message, startAt, c.Notify)
	if err != nil {
		return fmt.Errorf("Не удается установить напоминание о событие : %w", validation.ReminderAddEventError)
	}
	return nil
}

func (c *Calendar) CancelEventReminder(ID string) error {
	logger.Info("Запуск функции CancelEventReminder")
	event, exists := c.calendarEvents[ID]
	if !exists {
		return fmt.Errorf("Не удается удалить напоминание :%w ", validation.EventNotFoundError)
	}
	event.RemoveReminder()
	c.Save()
	fmt.Println("Напоминание удалено")
	return nil
}

func (c *Calendar) ShowEvent() error {
	logger.Info("Запуск функции ShowEvent")
	if len(c.calendarEvents) == 0 {
		return fmt.Errorf("Не удается показать список всех событий :%w ", validation.EmptyListError)
	}
	for _, v := range c.calendarEvents {
		utcTime := v.StartAt.UTC()
		fmt.Println(v.Title, "", utcTime.Format("02.01.2006 15:04"), "", v.Priority, v.ID)

		if v.Reminder != nil {
			fmt.Println("Есть напоминание", v.Reminder.Message, " ", v.Reminder.Timer)
		}
	}
	return nil
}

func (c *Calendar) DeleteEvent(ID string) error {
	logger.Info("Запуск функции DeleteEvent")
	event, exists := c.calendarEvents[ID]
	if !exists {
		return fmt.Errorf("Не удается удалить событие: %w", validation.EventNotFoundError)
	}
	delete(c.calendarEvents, ID)
	c.Save()
	show()
	fmt.Println("Событие:", event.Title, "удалено")
	show()
	fmt.Println("")
	return nil
}

func (c *Calendar) EditEvent(id, newTitle, dateStr string, p events.Priority) error {
	logger.Info("Запуск функции EditEvent")
	e, exist := c.calendarEvents[id]
	if !exist {
		return fmt.Errorf("Не удается изменить событие :%w ", validation.EventNotFoundError)
	}
	err := e.Update(newTitle, dateStr, p)
	if err != nil {
		return err
	}
	c.Save()
	show()
	fmt.Println("Событие :", newTitle, "Уcпешно изменено")
	show()
	fmt.Println("")
	return nil
}

func (c *Calendar) Save() error {
	date, err := json.Marshal(c.calendarEvents)
	if err != nil {
		return err
	}
	return c.storage.Save(date)
}
func (c *Calendar) Load() error {
	data, err := c.storage.Load()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &c.calendarEvents)
}
func (c *Calendar) Close() {
	close(c.Notification)
}
