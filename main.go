package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Указываем, что все файлы из папки "static" должны быть доступны
	// по URL-адресам, начинающимся с "/".
	// Например, файл static/index.html будет доступен по адресу http://localhost:8080/index.html
	// А файл static/mydocument.pdf по адресу http://localhost:8080/mydocument.pdf
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Задаем порт, на котором будет работать наше приложение
	port := "8080"
	fmt.Printf("Запускаем сервер на порту %s...\n", port)
	fmt.Printf("Перейдите по адресу http://localhost:%s\n", port)

	// Запускаем сервер. Если возникнет ошибка, приложение завершится.
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

