package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func nextDayHandler(w http.ResponseWriter, req *http.Request) {
	nowString := req.FormValue("now")
	dstart := req.FormValue("date")
	repeat := req.FormValue("repeat")

	var now time.Time
	if strings.TrimSpace(nowString) == "" {
		now = time.Now()
	} else {
		if nowParsed, err := time.Parse(dataType, nowString); err != nil {
			fmt.Println(err)
			return
		} else {
			now = nowParsed
		}
	}

	resultDate, err := NextDate(now, dstart, repeat)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resultDate))

}

func afterNow(date, now time.Time) bool {
	return date.After(now)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(dataType, dstart)
	if err != nil {
		return "", fmt.Errorf("Время не может быть преобразовано в корректную дату: %w", err)
	}

	if strings.TrimSpace(repeat) == "" {
		return "", errors.New("Интервал дней не может быть пустым")
	}

	num := strings.Split(repeat, " ")
	switch {
	case num[0] != "y" && num[0] != "d":
		return "", errors.New("Неподдерживаемый формат или недопустимый символ")
	case num[0] == "d" && len(num) == 1:
		return "", errors.New("Не указан интервал в днях")
	}

	if num[0] == "y" {
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
	} else {
		numberStr := num[1]
		interval, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Println(err)
		}

		if interval > 400 {
			return "", errors.New("Интервал дней больше максимально допустимого")
		}

		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				break
			}
		}

	}
	return date.Format(dataType), nil
}
