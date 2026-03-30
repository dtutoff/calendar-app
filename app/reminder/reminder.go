package reminder

import (
	"errors"
	"fmt"
	"time"

	"github.com/SamiRemi/project/app/validation"
)

type Reminder struct {
	Message string
	At      time.Time
	Sent    bool
	Timer   *time.Timer
}

func NewReminder(message string, startAt time.Time) (*Reminder, error) {
	text := validation.IsValidTitle(message)
	if !text {
		return nil, errors.New("Неверный формат загаловка")
	}
	return &Reminder{
		Message: message,
		At:      startAt,
		Sent:    false,
		Timer:   nil,
	}, nil
}
func (r *Reminder) Send() {
	if r.Sent {
		return
	}
	fmt.Println("Напоминание!", r.Message)
	r.Sent = true
}

func (r *Reminder) Start(t time.Duration) {
	if r.Timer != nil {
		r.Timer.Stop()
		r.Timer = nil
	}
	time.AfterFunc(t*time.Second, r.Send)
}

func (r *Reminder) Stop() bool {
	if r.Timer != nil {
		fmt.Println("Напоминание удалено")
		return r.Timer.Stop()
	} else {
		fmt.Println("Таймер пуст")
	}
	return false
}
