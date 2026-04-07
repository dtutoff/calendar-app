package reminder

import (
	"fmt"
	"strings"
	"time"

	"github.com/SamiRemi/project/app/logger"
	"github.com/SamiRemi/project/app/validation"
)

type Reminder struct {
	Message string
	At      time.Time
	Sent    bool
	Timer   *time.Timer
}

func NewReminder(message string, At time.Time, notify func(msg string)) (*Reminder, error) {
	logger.Info("Запуск функции NewReminder")
	if len(strings.TrimSpace(message)) == 0 {
		return nil, fmt.Errorf("не удается создать напоминание: %w", validation.ErrEmptyMessage)
	}

	text := validation.IsValidTitle(message)
	if !text {
		return nil, validation.IncorrectHeaderFormat
	}
	return &Reminder{
		Message: message,
		At:      At,
		Sent:    false,
		Timer:   nil,
	}, nil
}
func (r *Reminder) Send() {
	logger.Info("Запуск функции  Send")
	if r.Sent {
		return
	}
	fmt.Println("Напоминание!", r.Message)
	r.Sent = true
}

func (r *Reminder) Start(t time.Duration) {
	logger.Info("Запуск функции  Start")
	if r.Timer != nil {
		r.Timer.Stop()
		r.Timer = nil
	}
	time.AfterFunc(t*time.Second, r.Send)
}

func (r *Reminder) Stop() bool {
	logger.Info("Запуск функции  Stop")
	if r.Timer != nil {
		fmt.Println("Напоминание удалено")
		return r.Timer.Stop()
	} else {
		fmt.Println("Таймер пуст")
	}
	return false
}
