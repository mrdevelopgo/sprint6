package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// ConvertByType определяет тип входных данных (текст или азбука Морзе) и конвертирует их в противоположный формат.
func ConvertByType(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", nil
	}

	// Пытаемся интерпретировать как Морзе
	asText := morse.ToText(trimmed)

	// Если результат не пустой и отличается от входа, значит это был Морзе
	if asText != "" && asText != trimmed {
		return asText, nil
	}

	// Конвертируем в Морзе
	return morse.ToMorse(trimmed), nil
}
