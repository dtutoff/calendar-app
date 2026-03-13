package calendar

import (
	"errors"
	"fmt"

	"github.com/smokizazzi/app/events"
)

var calendarStore = make(map[string]events.Event)

func AddEvent(title string, dateStr string) (events.Event, error) {
	fmt.Printf("Adding event to calendar:\nTitle - %s\nDateTime - %s\n", title, dateStr)
	event, err := events.NewEvent(title, dateStr)
	if err != nil {
		return events.Event{}, err
	}
	calendarStore[event.ID] = event
	fmt.Printf("Event Added with ID: %s\n===\n", event.ID)
	return event, nil
}
func ShowEvents() {
	fmt.Printf("Showing events\n")
	if len(calendarStore) == 0 {
		fmt.Println("No calendar events found\n===")
	}

	for id, event := range calendarStore {
		formattedTime := event.StartAt.Format("2006-01-02 15:04")
		fmt.Printf("ID:%s: %s - %s\n", id, event.Title, formattedTime)
	}
	fmt.Println("===")
}

func GetEvent(id string) (events.Event, error) {
	event, exists := calendarStore[id]
	if !exists {
		return events.Event{}, errors.New("Event not found\n===")
	}
	return event, nil
}

func DeleteEvent(id string) error {
	fmt.Printf("Deleting %s\n", id)
	_, exists := calendarStore[id]
	if !exists {
		return errors.New("Event not found\n===")
	}
	delete(calendarStore, id)
	fmt.Printf("Event %s deleted\n===\n", id)
	return nil
}

func EditEvent(id string, newTitle string, newDateStr string) error {
	fmt.Printf("Editing event %s\n", id)
	_, exists := calendarStore[id]
	if !exists {
		return errors.New("Event does not exist\n===")
	}
	newEvent, err := events.NewEvent(newTitle, newDateStr)
	if err != nil {
		fmt.Printf("Error editing event: %s\n===\n", err)
		return err
	}
	calendarStore[id] = newEvent
	fmt.Printf("Event %s edited: %s - %s\n===\n",
		id, newEvent.Title, newEvent.StartAt.Format("2006-01-02 15:04"))

	return nil
}
