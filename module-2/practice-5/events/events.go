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

func NewEvent(title string, dateStr string) (Event, error) {
	t, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return Event{}, errors.New("неверный формат даты")
	}
	return Event{
		ID:      getNextID(),
		Title:   title,
		StartAt: t,
	}, nil
}

func getNextID() string {
	return uuid.New().String()
}
