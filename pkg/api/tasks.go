package api

import (
	"go-final-project/pkg/db"
	"log"
	"net/http"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50) // в параметре максимальное количество записей
	if err != nil {
		log.Println(err.Error())
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}
	if tasks == nil {
		writeJson(w, TasksResp{
			Tasks: []*db.Task{},
		})
	} else {
		writeJson(w, TasksResp{
			Tasks: tasks,
		})
	}
}
