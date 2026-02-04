package server

import (
	"fmt"
	"go-final-project/pkg/api"
	"net/http"
)

func Run() {
	api.Init()
}

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
