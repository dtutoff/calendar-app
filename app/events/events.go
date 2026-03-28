package events

import (
	"fmt"
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

func validateAndParse(title, dateStr string, p Priority) (string, time.Time, Priority, error) {
	if !validation.IsValidTitle(title) {
		return "", time.Time{}, p, fmt.Errorf("Неверный формат заголовка '%s'", title)
	}
	parsedDate, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return "", time.Time{}, p, fmt.Errorf("Неверный формат даты '%s'", dateStr)
	}

	if err := p.Validate(); err != nil {
		return "", time.Time{}, p, err
	}

	return title, parsedDate, p, nil
}

func NewEvent(title, dateStr string, p Priority) (*Event, error) {
	validatedTitle, validatedDate, validatedPriority, err := validateAndParse(title, dateStr, p)
	if err != nil {
		return nil, err
	}
	return &Event{
		ID:       getNextID(),
		Title:    validatedTitle,
		StartAt:  validatedDate,
		Priority: validatedPriority,
		Reminder: nil,
	}, nil
}

func (e *Event) Update(title, dateStr string, p Priority) error {
	validatedTitle, validatedDate, validatedPriority, err := validateAndParse(title, dateStr, p)
	if err != nil {
		return err
	}
	e.Title = validatedTitle
	e.StartAt = validatedDate
	e.Priority = validatedPriority
	return nil
}

func (e *Event) AddReminder(message string, at time.Time) {
	e.Reminder = reminder.NewReminder(message, at)
}

func (e *Event) RemoveReminder() {
	e.Reminder = nil
}
