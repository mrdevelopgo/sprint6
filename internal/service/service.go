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
	if asText != trimmed {
		// Если ToText вернул другой результат, значит входные данные были Морзе,а результат — текст. Возвращаем текст.
		return asText, nil
	}

	// Конвертируем как текст в Морзе.
	return morse.ToMorse(trimmed), nil
}
