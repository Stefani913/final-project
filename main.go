package main

import (
	"go-final-project/pkg/db"
	"go-final-project/pkg/server"
	"log"
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
