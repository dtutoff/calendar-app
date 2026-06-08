package calendar

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/araddon/dateparse"
	"github.com/dtutoff/app/events"
	"github.com/dtutoff/app/storage"
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

func (c *Calendar) Save() error {
	data, err := json.Marshal(c.calendarEvents)
	if err != nil {
		return fmt.Errorf("marshalling error: %w", err)
	}

	err = c.storage.Save(data)
	return err
}

func (c *Calendar) Load() error {
	data, err := c.storage.Load()
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &c.calendarEvents)
	if err != nil {
		return err
	}

	for _, e := range c.calendarEvents {
		if e.Reminder != nil && !e.Reminder.Sent {
			e.Reminder.Start()
		}
	}

	return nil
}

func (c *Calendar) AddEvent(title string, date string, priority events.Priority) (*events.Event, error) {
	if title == "" {
		return nil, fmt.Errorf("title cannot be empty")
	}
	d, err := dateparse.ParseIn(date, time.Local)
	if err != nil {
		return nil, err
	}
	if c.eventExists(title, d) {
		return nil, fmt.Errorf(`title "%s" already exists`, title)
	}

	e, err1 := events.NewEvent(title, date, priority)
	if err1 != nil {
		return nil, fmt.Errorf("error creating event: %w", err1)
	}

	c.calendarEvents[e.ID] = e
	return e, nil
}

func (c *Calendar) SetEventReminder(eventID string, message string, at string) error {
	event, exists := c.calendarEvents[eventID]
	if !exists {
		return fmt.Errorf("event with ID %s not found", eventID)
	}

	return event.AddReminder(message, at)
}

func (c *Calendar) RemoveEventReminder(eventID string) error {
	event, exists := c.calendarEvents[eventID]
	if !exists {
		return fmt.Errorf("event not found")
	}

	return event.RemoveReminder()
}

func (c *Calendar) EditEvent(id string, title string, date string, priority events.Priority) error {

	e, exists := c.calendarEvents[id]
	if !exists {
		return fmt.Errorf("event with key %q not found", id)
	}

	return e.Update(title, date, priority)
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
		fmt.Printf("ID: %s | Title: %s | Start: %s | Priority: %s\n", e.ID, e.Title, e.StartAt.Format("02.01.2006 15:04"), e.Priority)
	}
}

func (c *Calendar) eventExists(title string, startAt time.Time) bool {
	for _, event := range c.calendarEvents {
		if event.Title == title && event.StartAt.Equal(startAt) {
			return true
		}
	}
	return false
}
