package validation

import (
	"regexp"
)

func IsValidTitle(title string) bool {
	pattern := "^[а-яА-Яa-zA-Z0-9 ,/.]{3,50}$"
	matched, err := regexp.MatchString(pattern, title)
	if err != nil {
		return false
	}
	return matched
}
