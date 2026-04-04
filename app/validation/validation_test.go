package validation

import "testing"

func TestIsValidTitle_ValidTitles(t *testing.T) {
	validTitles := []string{
		"Valid Title",
		"Заголовок на русском",
		"Mixed 123 Текст",
		"Short",
		"This is a very long but valid title within 50 characters",
		"Title, with punctuation",
		"Title / with slash",
		"Letters and numbers 12345",
	}

	for _, title := range validTitles {
		t.Run(title, func(t *testing.T) {
			result := IsValidTitle(title)
			if !result {
				t.Errorf("Expected true for '%s', got false", title)
			}
		})
	}
}
func TestIsValidTitle_LengthValidation(t *testing.T) {
	tests := []struct {
		title       string
		expected    bool
		description string
	}{
		{"Ab", false, "too short (2 characters)"},
		{"ABC", true, "minimum length (3 characters)"},
		{"A very long title that exceeds fifty characters exactly here!", false, "too long (51+ characters)"},
		{"This is exactly fifty characters long title here now", true, "maximum length (50 characters)"},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := IsValidTitle(tt.title)
			if result != tt.expected {
				t.Errorf("For '%s' (%s): expected %v, got %v",
					tt.title, tt.description, tt.expected, result)
			}
		})
	}
}
