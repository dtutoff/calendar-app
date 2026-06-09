package reminder

import (
	"errors"
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

type Reminder struct {
	Message string       `json:"message"`
	At      time.Time    `json:"at"`
	Sent    bool         `json:"-"`
	Timer   *time.Timer  `json:"-"`
	Notify  func(string) `json:"-"`
}

func NewReminder(message string, at string, notify func(string)) (*Reminder, error) {
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
		Notify:  notify,
	}, nil
}

func (r *Reminder) Send() {
	if r.Sent {
		return
	}
	r.Notify(r.Message)
	r.Sent = true
}

func (r *Reminder) Stop() {
	if r.Timer != nil {
		r.Timer.Stop()
	}
}

func (r *Reminder) Start() {
	delay := time.Until(r.At)

	if delay <= 0 {
		fmt.Println("Date reminder is invalid")
		return
	}

	r.Timer = time.AfterFunc(delay, func() {
		r.Send()
	})
}
