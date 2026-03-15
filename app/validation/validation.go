package validation

import (
	"errors"
	"regexp"
)

func NewTitleError(title string) error {
	return errors.New("Неверный формат заголовка" + " '" + title + "'")
}

func NewDateError(dateStr string) error {
	return errors.New("Неверный формат даты" + " '" + dateStr + "'")
}

func IsValidTitle(title string) bool {
	pattern := "^[а-яА-Яa-zA-Z0-9 ,/.]{3,50}$"
	matched, err := regexp.MatchString(pattern, title)
	if err != nil {
		return false
	}
	return matched
}
