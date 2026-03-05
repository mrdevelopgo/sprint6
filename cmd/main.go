package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "MORSE: ", log.Ldate|log.Ltime|log.Lshortfile)
	srv := server.NewServer(logger)
	logger.Println("Запуск сервера...")
	if err := srv.Start(); err != nil {
		logger.Fatal("Ошибка запуска сервера:", err)
	}
}
