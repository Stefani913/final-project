package api

import "net/http"

const dataType = "20060102"

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
}
