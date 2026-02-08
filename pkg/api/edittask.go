package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-final-project/pkg/db"
	"log"
	"net/http"
)

func editTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		log.Println(err.Error())
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	if err := json.Unmarshal(buf.Bytes(), &task); err != nil {
		log.Println(err.Error())
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	if task.Title == "" {
		err := errors.New("Поле не может быть пустым")
		log.Println(err.Error())
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	err = checkDate(&task)
	if err != nil {
		log.Println(err.Error())
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		log.Println(err.Error())
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	writeJson(w, struct{}{})
}
