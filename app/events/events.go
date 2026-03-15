package events

import (
	"time"

	"github.com/SamiRemi/project/app/validation"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID      string
	Title   string
	StartAt time.Time
}

func getNextID() string {
	return uuid.New().String()
}

func NewEvent(title string, dateStr string) (Event, error) {
	isValid := validation.IsValidTitle(title)
	if !isValid {
		return Event{}, validation.NewTitleError(title)
	}
	t, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return Event{}, validation.NewDateError(dateStr)
	}
	return Event{
		ID:      getNextID(),
		Title:   title,
		StartAt: t,
	}, nil
}
