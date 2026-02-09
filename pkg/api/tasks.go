package api

import (
	"log"
	"net/http"

	"go-final-project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {
		log.Println(err.Error())
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if tasks == nil {
		writeJson(w, http.StatusOK, TasksResp{
			Tasks: []*db.Task{},
		})
	} else {
		writeJson(w, http.StatusOK, TasksResp{
			Tasks: tasks,
		})
	}
}
