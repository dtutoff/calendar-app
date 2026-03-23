package events

import (
	"time"

	"github.com/SamiRemi/project/app/reminder"
	"github.com/SamiRemi/project/app/validation"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID       string             `json:"id"`
	Title    string             `json:"title"`
	StartAt  time.Time          `json:"start_at"`
	Priority Priority           `json:"priority"`
	Reminder *reminder.Reminder `json:"reminder"`
}

func getNextID() string {
	return uuid.New().String()
}

func NewEvent(title, dateStr string, p Priority) (*Event, error) {
	isValid := validation.IsValidTitle(title)
	if !isValid {
		return &Event{}, validation.NewTitleError(title)
	}
	time, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return &Event{}, validation.NewDateError(dateStr)
	}
	return &Event{
		ID:       getNextID(),
		Title:    title,
		StartAt:  time,
		Priority: p,
		Reminder: nil,
	}, nil
}

func (e *Event) Update(title string, dateStr string) error {
	isValid := validation.IsValidTitle(title)
	if !isValid {
		return validation.NewTitleError(title)
	}
	time, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return validation.NewDateError(dateStr)
	}
	e.Title = title
	e.StartAt = time
	return nil
}

func (e *Event) AddReminder(message string, at time.Time) {
	e.Reminder = reminder.NewReminder(message, at)
}

func (e *Event) RemoveReminder() {
	e.Reminder = nil
}
