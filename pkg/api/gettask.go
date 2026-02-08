package api

import (
	"go-final-project/pkg/db"
	"log"
	"net/http"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	task, err := db.GetTask(id)
	if err != nil {
		log.Println(err.Error())
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}
	writeJson(w, task)
}
