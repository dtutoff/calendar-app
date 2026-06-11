package events

import (
	"errors"
	"fmt"
	"time"

	"github.com/araddon/dateparse"
	"github.com/dtutoff/app/reminder"
	"github.com/google/uuid"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

type Event struct {
	ID       string             `json:"id"`
	Title    string             `json:"title"`
	StartAt  time.Time          `json:"start_at"`
	Priority Priority           `json:"priority"`
	Reminder *reminder.Reminder `json:"reminder"`
}

func NewEvent(title string, date string, priority Priority) (*Event, error) {
	t, err := validateAddParseEvent(date, priority)
	if err != nil {
		return nil, err
	}

	return &Event{
		ID:       getNextID(),
		Title:    title,
		StartAt:  *t,
		Priority: priority,
		Reminder: nil,
	}, nil
}

func (e *Event) Update(title string, date string, priority Priority) error {
	t, err := validateAddParseEvent(date, priority)
	if err != nil {
		return err
	}

	e.Title = title
	e.StartAt = *t
	e.Priority = priority
	return nil
}

func (e *Event) AddReminder(message string, at string, notify func(string)) error {
	r, err := reminder.NewReminder(message, at, notify)
	if err != nil {
		return err
	}

	e.Reminder = r
	return r.Start()
}

func (e *Event) RemoveReminder() error {
	if e.Reminder == nil {
		return fmt.Errorf("event has no reminder")
	}
	e.Reminder.Stop()
	e.Reminder = nil
	return nil
}

func (p Priority) Validate() error {
	switch p {
	case PriorityLow, PriorityMedium, PriorityHigh:
		return nil
	default:
		return errors.New("invalid priority")
	}
}

func validateAddParseEvent(date string, priority Priority) (*time.Time, error) {
	t, err := dateparse.ParseIn(date, time.Local)
	if err != nil {
		return nil, err
	}

	err = priority.Validate()
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func getNextID() string {
	return uuid.New().String()
}
