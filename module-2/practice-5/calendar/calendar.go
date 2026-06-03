package calendar

import (
	"encoding/json"
	"fmt"

	"github.com/dtutoff/app/events"
	"github.com/dtutoff/app/storage"
)

type Calendar struct {
	calendarEvents map[string]*events.Event
	storage        *storage.Storage
}

func NewCalendar(s *storage.Storage) *Calendar {
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
		return fmt.Errorf("unmarshalling error: %w", err)
	}
	return err
}

func (c *Calendar) autoSave() {
	if err := c.Save(); err != nil {
		fmt.Println("Auto-save error:", err)
	}
}

func (c *Calendar) AddEvent(title string, date string) (*events.Event, error) {
	if title == "" {
		return nil, fmt.Errorf("title cannot be empty")
	}

	defer c.autoSave()

	e, err := events.NewEvent(title, date)
	if err != nil {
		return nil, fmt.Errorf("error creating event: %w", err)
	}

	c.calendarEvents[e.ID] = e
	return e, nil
}

func (c *Calendar) EditEvent(id string, title string, date string) error {
	defer c.autoSave()

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

	defer c.autoSave()
	delete(c.calendarEvents, id)
	return nil
}

func (c *Calendar) ShowEvents() {
	for _, e := range c.calendarEvents {
		fmt.Printf("ID: %s | Title: %s | Start: %s\n", e.ID, e.Title, e.StartAt.Format("02.01.2006 15:04"))
	}
}
