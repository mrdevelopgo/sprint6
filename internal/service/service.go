package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Определяем 'nj текст или морзе, конвертируем в противоположное
func ConvertByType(input string) (string, error) {
	// Убираем лишние пробелы
	trimmed := strings.TrimSpace(input)

	// Если строка пустая возвращаем
	if trimmed == "" {
		return "", nil
	}

	// Смотрим каждый символ строки
	isMorse := true

	for _, char := range trimmed {
		// Ищем точка, тире и пробел
		if char != '.' && char != '-' && char != ' ' {
			isMorse = false
			break
		}
	}

	// Выполняем конвертацию
	if isMorse {
		// Если морзе - переводим в текст
		return morse.ToText(trimmed), nil
	} else {
		// Если текст - переводим в морзе
		return morse.ToMorse(trimmed), nil
	}
}
