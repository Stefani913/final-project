package api

import (
	"log"
	"net/http"
	"time"

	"go-final-project/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	task, err := db.GetTask(id)
	if err != nil {
		log.Println(err.Error())
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if len(task.Repeat) == 0 {
		err = db.DeleteTask(id)
		if err != nil {
			log.Println(err.Error())
			writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
	} else {
		next, err := NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			log.Println(err.Error())
			writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		if err = db.UpdateDate(next, id); err != nil {
			log.Println(err.Error())
			writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
	}

	writeJson(w, http.StatusOK, struct{}{})
}
