package server

import (
	"go-final-project/pkg/api"
	"log"
	"net/http"
)

func Run() {
	api.Init()
}

func Start() {
	log.Println("Start working")
	webDir := "./web"

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		panic(err)
	}
	log.Println("Stop working")
}
