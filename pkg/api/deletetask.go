package api

import (
	"go-final-project/pkg/db"
	"log"
	"net/http"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	_, err := db.GetTask(id)
	if err != nil {
		log.Println(err)
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	err = db.DeleteTask(id)
	if err != nil {
		log.Println(err.Error())
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}
	writeJson(w, struct{}{})
}
