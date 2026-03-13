package events

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID      string
	Title   string
	StartAt time.Time
}

func generateId() string {
	return uuid.New().String()
}

func ValidateTitle(title string) bool {
	pattern := `^[a-zA-Zа-яА-Я0-9 ]{3,100}$`
	matched, err := regexp.MatchString(pattern, title)
	if err != nil {
		return false
	}
	return matched
}

func NewEvent(title string, dateStr string) (Event, error) {
	//fmt.Printf("Creating new event\nTitle - %s\nDateTime - %s\n===\n", title, dateStr)
	if !ValidateTitle(title) {
		return Event{}, errors.New("Invalid title\n===")
	}
	t, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return Event{}, errors.New("Invalid date\n===")
	}
	event := Event{
		ID:      generateId(),
		Title:   title,
		StartAt: t,
	}
	fmt.Printf("Completed\n")
	return event, nil
}
