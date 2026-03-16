package calendar

import (
	"errors"
	"fmt"

	"github.com/SamiRemi/project/app/events"
)

type Calendar struct {
	calendarEvents map[string]*events.Event
}

func NewCalendar() *Calendar {
	return &Calendar{
		calendarEvents: make(map[string]*events.Event),
	}
}
func (c *Calendar) AddEvent(title string, date string) (*events.Event, error) {

	e, err := events.NewEvent(title, date)
	if err != nil {
		return e, err
	}
	if _, ok := c.calendarEvents[title]; ok {
		return e, errors.New("Событие с именем " + title + " уже существует!")
	}
	if len(title) == 0 {
		return e, errors.New("Нельзя ввести пустое имя")
	}
	c.calendarEvents[e.ID] = e
	return e, nil
}

func (c *Calendar) ShowEvent() error {
	if len(c.calendarEvents) == 0 {
		return errors.New("Список пуст")
	}
	for _, v := range c.calendarEvents {
		utcTime := v.StartAt.UTC()
		fmt.Println(v.Title, "", utcTime.Format("02.01.2006 15:04"))
	}
	return nil
}

func (c *Calendar) DeleteEvent(title string) error {
	e := c.calendarEvents[title]
	if _, ok := c.calendarEvents[title]; !ok {
		return errors.New("Событие с именем " + title + " не существует")
	}
	delete(c.calendarEvents, title)
	fmt.Println("=========================")
	fmt.Println("Событие :", e.Title)
	fmt.Println("С ID :", e.ID)
	fmt.Println("Удалено")
	fmt.Println("=========================")
	fmt.Println("")
	return nil
}

func (c *Calendar) EditEvent(id, newTitle, dateStr string) error {
	e, exist := c.calendarEvents[id]
	if !exist {
		return fmt.Errorf("Событие с ключом %q не найдено", id)
	}
	err := e.Update(newTitle, dateStr)
	if err != nil {
		return err
	}
	fmt.Println("=========================")
	fmt.Println("Событие :", newTitle)
	fmt.Println("С ID :", e.ID)
	fmt.Println("УСпешно изменено")
	fmt.Println("=========================")
	fmt.Println("")
	return nil
}
