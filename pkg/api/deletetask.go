package api

import (
	"log"
	"net/http"

	"go-final-project/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	_, err := db.GetTask(id)
	if err != nil {
		log.Println(err)
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	err = db.DeleteTask(id)
	if err != nil {
		log.Println(err.Error())
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJson(w, http.StatusOK, struct{}{})
}
