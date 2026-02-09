package server

import (
	"log"
	"net/http"

	"go-final-project/pkg/api"
)

func Init() {
	api.Init()
}

const port = ":7540"

func Start() {
	log.Printf("Start working port %s\n", port)
	webDir := "./web"

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
