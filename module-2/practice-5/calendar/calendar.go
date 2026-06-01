package calendar

import (
	"errors"
	"fmt"

	"github.com/araddon/dateparse"
	"github.com/dtutoff/app/events"
)

var calendarEvents = make(map[string]events.Event)

func AddEvent(title string, date string) (events.Event, error) {
	e, err := events.NewEvent(title, date) // создаем событие прямо в процессе добавления
	if err != nil {
		return events.Event{}, errors.New("")
	}

	calendarEvents[e.ID] = e
	return e, nil
}

func EditEvent(id string, title string, date string) error {
	for key, _ := range calendarEvents {
		if key == id {
			t, err := dateparse.ParseAny(date)
			if err != nil {
				return errors.New("неверный формат даты")
			}
			newEvent := events.Event{
				ID:      id,
				Title:   title,
				StartAt: t,
			}
			calendarEvents[key] = newEvent
		}
	}
	return nil
}

func DeleteEvent(id string) {
	for key, _ := range calendarEvents {
		if key == id {
			delete(calendarEvents, key)
		}
	}
}

func ShowEvents() {
	for key := range calendarEvents {
		fmt.Println(calendarEvents[key])
	}
}
