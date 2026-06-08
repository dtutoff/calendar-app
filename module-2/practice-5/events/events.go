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

func NewEvent(title string, dateStr string, priority Priority) (*Event, error) {
	t := validateDate(dateStr)
	err1 := priority.Validate()
	if err1 != nil {
		return nil, err1
	}
	return &Event{
		ID:       getNextID(),
		Title:    title,
		StartAt:  t,
		Priority: priority,
		Reminder: nil,
	}, nil
}

func (e *Event) Update(title string, date string, priority Priority) error {
	t := validateDate(date)
	err1 := priority.Validate()
	if err1 != nil {
		return err1
	}

	e.Title = title
	e.StartAt = t
	e.Priority = priority
	return nil
}

func (e *Event) AddReminder(message string, at string) error {
	r, err := reminder.NewReminder(message, at)
	if err != nil {
		return fmt.Errorf(`error adding "%s" as a reminder`, message)
	}

	e.Reminder = r
	r.Start()
	return nil
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

func validateDate(date string) time.Time {
	t, err := dateparse.ParseIn(date, time.Local)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func getNextID() string {
	return uuid.New().String()
}
