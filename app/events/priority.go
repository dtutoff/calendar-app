package events

import (
	"fmt"

	"github.com/SamiRemi/project/app/logger"
	"github.com/SamiRemi/project/app/validation"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

func (p Priority) Validate() error {
	logger.Info("Запуск функции Validate")
	switch p {
	case PriorityLow, PriorityMedium, PriorityHigh:
		return nil
	default:
		return fmt.Errorf("Не удается установить приоритет:%w", validation.IncorrectPriorityError)
	}
}
