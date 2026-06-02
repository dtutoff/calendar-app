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

	c.calendarEvents[e.ID] = e

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

func (c *Calendar) DeleteEvent(id string) error {
	if _, exists := c.calendarEvents[id]; !exists {
		return fmt.Errorf("event with id %q not found", id)
	}

	delete(c.calendarEvents, id)
	return nil
}

func (c *Calendar) ShowEvents() {
	for _, e := range c.calendarEvents {
		fmt.Printf("ID: %s | Title: %s | Start: %s\n", e.ID, e.Title, e.StartAt.Format("02.01.2006 15:04"))
	}
}
