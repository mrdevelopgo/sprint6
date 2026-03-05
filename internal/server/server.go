package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Структура сервера
type Server struct {
	logger *log.Logger
	server *http.Server
}

// Функция создания нового сервера
func NewServer(logger *log.Logger) *Server {
	// Создание роутера
	mux := http.NewServeMux()

	// Регистрируем обработчики
	mux.HandleFunc("/", handlers.IndexHandler)        // главная страница
	mux.HandleFunc("/upload", handlers.UploadHandler) // загрузка файла

	// Настраиваем HTTP-сервер
	httpServer := &http.Server{
		Addr:         ":8080",          // порт 8080
		Handler:      mux,              // роутер
		ErrorLog:     logger,           // логгер для ошибок
		ReadTimeout:  5 * time.Second,  // таймаут на чтение запроса
		WriteTimeout: 10 * time.Second, // таймаут на запись ответа
		IdleTimeout:  15 * time.Second, // таймаут на простаивание соединения
	}

	return &Server{
		logger: logger,
		server: httpServer,
	}
}

// Метод для запуска сервера
func (s *Server) Start() error {
	s.logger.Println("Сервер запущен на http://localhost:8080")
	return s.server.ListenAndServe()
}
