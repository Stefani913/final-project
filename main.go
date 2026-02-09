package main

import (
	"log"

	"go-final-project/pkg/db"
	"go-final-project/pkg/server"
)

func main() {
	err := db.Init("scheduler.db")
	if err != nil {
		db.Close()
		log.Fatal(err)
	}
	server.Init()
	server.Start()
}
