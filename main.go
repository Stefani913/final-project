package main

import (
	"net/http"
)

/* TODO!!!
Как вариант, можете создать директорию pkg с тремя поддиректориями:
* api — будет содержать файлы с обработчиками API запросов;
* db — файлы с кодом, отвечающие за работу с базой данных;
* server — будет расположен файл с функцией запуска сервера.
Файлы main.go и go.mod располагаются непосредственно в директории проекта.
*/

func main() {
	// SERVER -> move into server package
	webDir := "./web"

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		panic(err)
	}
}
