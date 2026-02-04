package main

import (
	"go-final-project/pkg/db"
	"go-final-project/pkg/server"
	"log"
)

func main() {
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatal(err)
	}
	server.Run()
	server.Start()

}
