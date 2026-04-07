package events

import (
	"fmt"
	"time"

	"github.com/SamiRemi/project/app/logger"
	"github.com/SamiRemi/project/app/reminder"
	"github.com/SamiRemi/project/app/validation"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID       string             `json:"id"`
	Title    string             `json:"title"`
	StartAt  time.Time          `json:"start_at"`
	Priority Priority           `json:"priority"`
	Reminder *reminder.Reminder `json:"reminder"`
}

func getNextID() string {
	return uuid.New().String()
}

func validateAndParse(title, dateStr string, p Priority) (string, time.Time, Priority, error) {
	if !validation.IsValidTitle(title) {
		return "", time.Time{}, p, validation.IncorrectHeaderFormat
	}

	parsedDate, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return "", time.Time{}, p, validation.DateFormatError
	}

	if err := p.Validate(); err != nil {
		return "", time.Time{}, p, err
	}

	return title, parsedDate, p, nil
}

func NewEvent(title, dateStr string, p Priority) (*Event, error) {
	logger.Info("Запуск фукции NewEvent")
	validatedTitle, validatedDate, validatedPriority, err := validateAndParse(title, dateStr, p)
	if err != nil {
		return nil, err
	}
	return &Event{
		ID:       getNextID(),
		Title:    validatedTitle,
		StartAt:  validatedDate,
		Priority: validatedPriority,
		Reminder: nil,
	}, nil
}

func (e *Event) Update(title, dateStr string, p Priority) error {
	logger.Info("Запуск функции Update")
	validatedTitle, validatedDate, validatedPriority, err := validateAndParse(title, dateStr, p)
	if err != nil {
		return err
	}
	e.Title = validatedTitle
	e.StartAt = validatedDate
	e.Priority = validatedPriority
	return nil
}

func (e *Event) AddReminder(message string, at time.Time, notify func(msg string)) error {
	logger.Info("Запуск функции AddReminder")
	reminder, err := reminder.NewReminder(message, at, notify)
	if err != nil {
		return err
	}
	e.Reminder = reminder
	actualTime := time.Now()
	reminder.Start(
		time.Duration((at.Hour()-actualTime.Hour())*3600) +
			time.Duration((at.Minute()-actualTime.Minute())*60) +
			time.Duration((at.Second() - actualTime.Second())),
	)
	return nil
}

func (e *Event) RemoveReminder() error {
	logger.Info("Запуск функции RemoveReminder")
	if e.Reminder == nil {
		return fmt.Errorf("Не удается удалить напоминание :%w ", validation.ReminderNotExistError)
	}
	stopped := e.Reminder.Stop()
	if stopped {
		fmt.Println("Таймер остановлен до срабатывания")
	} else {
		fmt.Println("Таймер уже сработал или остановлен")
	}
	e.Reminder = nil
	return nil
}
