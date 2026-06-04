package reminder

import (
	"errors"
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

type Reminder struct {
	Message string
	At      time.Time
	Sent    bool
}

func NewReminder(message string, at string) (*Reminder, error) {
	if message == "" {
		return nil, errors.New("reminder message is empty")
	}

	d, err := dateparse.ParseAny(at)
	if err != nil {
		return nil, err
	}

	return &Reminder{
		Message: message,
		At:      d,
		Sent:    false,
	}, nil
}

func (r *Reminder) Send() {
	if r.Sent {
		return
	}
	fmt.Println("Reminder!", r.Message)
	r.Sent = true
}

func (r *Reminder) Stop() {
	// TODO: Здесь будет логика остановки напоминания
}
