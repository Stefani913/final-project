package main

import (
	"go-final-project/pkg/db"
	"go-final-project/server"
	"log"
)

func main() {
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatal(err)
	}

	server.Start()

}
