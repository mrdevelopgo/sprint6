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
	log.Println("UploadHandler called")

	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("ParseMultipartForm error: %v", err)
		http.Error(w, "Ошибка обработки формы", http.StatusInternalServerError)
		return
	}

	if r.MultipartForm != nil && r.MultipartForm.File != nil {
		for key := range r.MultipartForm.File {
			log.Printf("Found file field with name: %q", key)
		}
	} else {
		log.Println("No file fields found in form")
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		log.Printf("FormFile error for key 'myFile': %v", err)
		http.Error(w, "Файл не найден", http.StatusBadRequest)
		return
	}
	defer file.Close()

	log.Printf("Successfully got file: %s", header.Filename)

	content, err := io.ReadAll(file)
	if err != nil {
		log.Printf("ReadAll error: %v", err)
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	converted, err := service.ConvertByType(string(content))
	if err != nil {
		log.Printf("ConvertByType error: %v", err)
		http.Error(w, "Ошибка конвертации", http.StatusInternalServerError)
		return
	}

	// Сохраняем результат в новый файл (как требует задание)
	timestamp := time.Now().UTC().Format("20060102_150405")
	ext := filepath.Ext(header.Filename)
	outputFilename := timestamp + "_converted" + ext

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		log.Printf("Create error: %v", err)
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(converted)
	if err != nil {
		log.Printf("WriteString error: %v", err)
		http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
		return
	}

	// Отправляем только результат конвертации (без лишнего текста)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(converted))
}
