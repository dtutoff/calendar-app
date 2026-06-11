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

	d, err1 := dateparse.ParseIn(at, time.Local)
	if err1 != nil {
		return nil, err1
	}

	err := validateReminderDate(d)
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

func (r *Reminder) Start() error {

	err := validateReminderDate(r.At)
	if err != nil {
		return err
	}

	delay := time.Until(r.At)

	r.Timer = time.AfterFunc(delay, func() {
		r.Send()
	})

	return nil
}

func validateReminderDate(date time.Time) error {
	if time.Until(date) <= 0 {
		return errors.New("reminder's date is invalid")
	}

	return nil
}
