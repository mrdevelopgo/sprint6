package handlers

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// Отдаем HTML-форму на главной странице
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим HTML файл из корня проекта
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		// Если не можем прочитать файл - 500 ошибка
		http.Error(w, "Ошибка загрузки страницы", http.StatusInternalServerError)
		return
	}

	// Отправляем HTML клиенту
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Ошибка отображения страницы", http.StatusInternalServerError)
		return
	}
}

// Обрабатываем загруженный файл
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Добавим логирование вызова
	log.Println("UploadHandler called")

	// Проверяем, что метод запроса - POST
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Парсим multipart форму до 10 MB
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Ошибка обработки формы", http.StatusInternalServerError)
		return
	}

	// Логируем все поля с файлами, которые пришли в форме
	if r.MultipartForm != nil && r.MultipartForm.File != nil {
		for key := range r.MultipartForm.File {
			log.Printf("Found file field with name: %q", key)
		}
	} else {
		log.Println("No file fields found in form")
	}

	// Получаем файл из формы
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("FormFile error for key 'file': %v", err)
		http.Error(w, "Файл не найден", http.StatusBadRequest)
		return
	}
	defer file.Close()

	log.Printf("Successfully got file: %s", header.Filename)

	// Читаем содержимое файла
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	// Вызываем функцию автоопределения и конвертации
	converted, err := service.ConvertByType(string(content))
	if err != nil {
		http.Error(w, "Ошибка конвертации", http.StatusInternalServerError)
		return
	}

	// Генерируем уникальное имя для выходного файла: текущее время + расширение оригинального файла
	timestamp := time.Now().UTC().Format("20060102_150405")
	ext := filepath.Ext(header.Filename)
	outputFilename := timestamp + "_converted" + ext

	// Создаем новый файл для результата
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Записываем результат в файл
	_, err = outputFile.WriteString(converted)
	if err != nil {
		http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
		return
	}

	// Отправляем результат пользователю
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Конвертация выполнена успешно!\n\n"))
	w.Write([]byte("Результат:\n"))
	w.Write([]byte(converted))
	w.Write([]byte("\n\nФайл сохранен как: " + outputFilename))
}
