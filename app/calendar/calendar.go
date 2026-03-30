package calendar

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/SamiRemi/project/app/events"
	"github.com/SamiRemi/project/app/storage"
	"github.com/araddon/dateparse"
)

type Calendar struct {
	calendarEvents map[string]*events.Event
	storage        storage.Store
}

func NewCalendar(s storage.Store) *Calendar {
	return &Calendar{
		calendarEvents: make(map[string]*events.Event),
		storage:        s,
	}
}
func (c *Calendar) AddEvent(title, date string, priority events.Priority) (*events.Event, error) {

	e, err := events.NewEvent(title, date, priority)
	if err != nil {
		return nil, err
	}
	if _, ok := c.calendarEvents[title]; ok {
		return nil, errors.New("Событие с именем " + title + " уже существует!")
	}
	if len(title) == 0 {
		return nil, errors.New("Нельзя ввести пустое имя")
	}
	c.calendarEvents[e.ID] = e
	return e, nil
}

func (c *Calendar) SetEventReminder(ID, message, dateStr string) error {
	event, exists := c.calendarEvents[ID]
	if !exists {
		return fmt.Errorf("событие с ID"+ID+"не найдено", event)
	}

	if event.Reminder != nil {
		return fmt.Errorf("Напоминание уже существует")
	}

	startAt, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return err
	}
	err = event.AddReminder(message, startAt)
	if err != nil {
		return err
	}
	return nil
}

func (c *Calendar) CancelEventReminder(ID string) error {
	event, exists := c.calendarEvents[ID]
	if !exists {
		return fmt.Errorf("событие с ID"+ID+"не найдено", event)
	}
	event.RemoveReminder()
	c.Save()
	fmt.Println("Напоминание удалено")
	return nil
}

func (c *Calendar) ShowEvent() error {

	if len(c.calendarEvents) == 0 {
		return errors.New("Список пуст")

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
		return errors.New("Событие с ID " + ID + " не существует")
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
		return fmt.Errorf("Событие с ключом %q не найдено", id)
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
