package calendar

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/SamiRemi/project/app/events"
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

func (c *Calendar) AddEvent(title, date string, priority events.Priority) (*events.Event, error) {
	e, err := events.NewEvent(title, date, priority)
	if err != nil {
		return nil, err
	}
	if _, ok := c.calendarEvents[title]; ok {
		return nil, validation.TitleError
	}
	if len(title) == 0 {
		return nil, validation.ListError
	}
	c.calendarEvents[e.ID] = e
	return e, nil
}

func (c *Calendar) Notify(msg string) {
	c.Notification <- msg
}

func (c *Calendar) SetEventReminder(ID, message, dateStr string) error {
	if c == nil {
		return validation.EqualError
	}
	event, exists := c.calendarEvents[ID]
	if !exists {
		return validation.EventNotFoundError
	}
	if event.Reminder != nil {
		return validation.ReminderAlreadyExistError
	}
	startAt, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return validation.DateFormatError
	}
	if startAt.Before(time.Now()) {
		return validation.ReminderDateError
	}

	err = event.AddReminder(message, startAt, c.Notify)
	if err != nil {
		return validation.ReminderAddEventError
	}
	return nil
}

func (c *Calendar) CancelEventReminder(ID string) error {
	event, exists := c.calendarEvents[ID]
	if !exists {
		return validation.EventNotFoundError
	}
	event.RemoveReminder()
	c.Save()
	fmt.Println("Напоминание удалено")
	return nil
}

func (c *Calendar) ShowEvent() error {
	if len(c.calendarEvents) == 0 {
		return validation.EmptyListError
	}
	for _, v := range c.calendarEvents {
		utcTime := v.StartAt.UTC()
		fmt.Println(v.Title, "", utcTime.Format("02.01.2006 15:04"), "", v.Priority, v.ID)

		err := v.Reminder
		if err != nil {
			fmt.Println("Есть напоминание", v.Reminder.Message, " ", v.Reminder.Timer)
		}
	}
	return nil
}

func (c *Calendar) DeleteEvent(ID string) error {
	e := c.calendarEvents[ID]
	if _, ok := c.calendarEvents[ID]; !ok {
		return validation.EventNotFoundError
	}
	delete(c.calendarEvents, e.ID)
	c.Save()
	fmt.Println("=========================")
	fmt.Println("Событие :", e.Title)
	fmt.Println("С ID :", e.ID)
	fmt.Println("Удалено")
	fmt.Println("=========================")
	fmt.Println("")
	return nil
}

func (c *Calendar) EditEvent(id, newTitle, dateStr string, p events.Priority) error {
	e, exist := c.calendarEvents[id]
	if !exist {
		return validation.EventNotFoundError
	}
	err := e.Update(newTitle, dateStr, p)
	if err != nil {
		return err
	}
	c.Save()
	fmt.Println("=========================")
	fmt.Println("Событие :", newTitle)
	fmt.Println("С ID :", e.ID)
	fmt.Println("Уcпешно изменено")
	fmt.Println("=========================")
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
