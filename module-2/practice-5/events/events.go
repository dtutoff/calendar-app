package events

import (
	"errors"
	"time"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID      string
	Title   string
	StartAt time.Time
}

func NewEvent(title string, dateStr string) (*Event, error) {
	t, err := GetDate(dateStr)
	if err != nil {
		return nil, err
	}
	return &Event{
		ID:      getNextID(),
		Title:   title,
		StartAt: t,
	}, nil
}

func (e *Event) Update(title string, date string) error {
	t, err := GetDate(date)
	if err != nil {
		return err
	}

	e.Title = title
	e.StartAt = t
	return nil
}

func getNextID() string {
	return uuid.New().String()
}

func GetDate(date string) (time.Time, error) {
	t, err := dateparse.ParseAny(date)
	if err != nil {
		return time.Time{}, errors.New("incorrect date format")
	}
	return t, nil
}
