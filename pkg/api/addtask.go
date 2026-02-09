package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"go-final-project/pkg/db"
)

func writeJson(w http.ResponseWriter, status int, data any) {
	resp, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	w.Write(resp)
}

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(dateLayout)
	}
	t, err := time.Parse(dateLayout, task.Date)
	if err != nil {
		return err
	}

	if afterNow(now, t) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format(dateLayout)
		} else {
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return err
			}
			task.Date = next
		}
	}
	return nil
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		log.Println(err.Error())
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if err := json.Unmarshal(buf.Bytes(), &task); err != nil {
		log.Println(err.Error())
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if task.Title == "" {
		err := errors.New("Поле не может быть пустым")
		log.Println(err.Error())
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	err = checkDate(&task)
	if err != nil {
		log.Println(err.Error())
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		log.Println(err.Error())
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJson(w, http.StatusOK, map[string]string{"id": strconv.FormatInt(id, 10)})
}
