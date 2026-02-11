package main

import (
	"log"

	"go-final-project/pkg/db"
	"go-final-project/pkg/server"
)

func main() {
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatal(err)
	}
	server.Init()
	defer db.Close()
	server.Start()
}
