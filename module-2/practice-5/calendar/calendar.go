package calendar

import (
	"fmt"

	"github.com/dtutoff/app/events"
)

type Calendar struct {
	calendarEvents map[string]*events.Event
}

func NewCalendar() *Calendar {
	return &Calendar{calendarEvents: make(map[string]*events.Event)}
}

func (c *Calendar) AddEvent(title string, date string) (*events.Event, error) {
	e, err := events.NewEvent(title, date)
	if err != nil {
		return nil, fmt.Errorf("error creating event: %w", err)
	}

	return e, nil
}

func (c *Calendar) EditEvent(id string, title string, date string) error {
	e, exists := c.calendarEvents[id]
	if !exists {
		return fmt.Errorf("event with key %q not found", id)
	}

	err := e.Update(title, date)
	if err != nil {
		return err
	}

	return nil
}

func (c *Calendar) DeleteEvent(id string) {
	for key, _ := range c.calendarEvents {
		if key == id {
			delete(c.calendarEvents, key)
		}
	}
}

func (c *Calendar) ShowEvents() {
	for key := range c.calendarEvents {
		fmt.Println(c.calendarEvents[key])
	}
}
