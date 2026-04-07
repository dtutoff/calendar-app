package validation

import (
	"errors"
	"regexp"

	"github.com/SamiRemi/project/app/logger"
)

var (
	IncorrectPriorityError    = errors.New("Неверный приоритет")
	TitleError                = errors.New("Событие с таким именем уже существует!")
	ListError                 = errors.New("Нельзя ввести пустое имя")
	EqualError                = errors.New("Календарь равен нулю")
	EventNotFoundError        = errors.New("событие с ID не найдено")
	ReminderAlreadyExistError = errors.New("Напоминание уже существует")
	ReminderNotExistError     = errors.New("Напоминания не существует")
	DateFormatError           = errors.New("некорректный формат даты ")
	IncorrectHeaderFormat     = errors.New("Неверный формат заголовка ")
	ReminderDateError         = errors.New("дата напоминания уже прошла")
	ReminderAddEventError     = errors.New("ошибка при добавлении напоминания для события")
	EmptyListError            = errors.New("Список пуст")
	ArchiveEmptyError         = errors.New("Архив пуст")
	ErrEmptyMessage           = errors.New("сообщение пусто")
)

func IsValidTitle(title string) bool {
	logger.Info("Запуск функции IsValidTitle")
	pattern := "^[а-яА-Яa-zA-Z0-9 ,/.]{3,50}$"
	matched, err := regexp.MatchString(pattern, title)
	if err != nil {
		return false
	}
	return matched
}
