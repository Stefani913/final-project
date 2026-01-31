package server

import (
	"fmt"
	"net/http"
)

func Start() {
	fmt.Println("Start working")
	webDir := "./web"

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Stop working")
}
