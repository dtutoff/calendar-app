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
	timer   *time.Timer
}

func NewReminder(message string, at string) (*Reminder, error) {
	if message == "" {
		return nil, errors.New("reminder message is empty")
	}

	d, err := dateparse.ParseIn(at, time.Local)
	if err != nil {
		return nil, err
	}

	fmt.Println("reminder:", d, "added")

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
	fmt.Println(r.Message)
	r.Sent = true
}

func (r *Reminder) Stop() {
	if r.timer != nil {
		r.timer.Stop()
	}
}

func (r *Reminder) Start() {
	delay := time.Until(r.At)

	if delay <= 0 {
		fmt.Println("Date reminder is invalid")
		return
	}

	r.timer = time.AfterFunc(delay, func() {
		r.Send()
	})
}
