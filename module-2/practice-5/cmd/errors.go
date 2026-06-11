package cmd

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrFailedToSave = errors.New("failed to save data")
	ErrEventUpdate  = errors.New("failed to update event")
	ErrEventAdd     = errors.New("error adding event")
	ErrEventDelete  = errors.New("error deleting event")
	ErrReminderAdd  = errors.New("error adding reminder")
)
