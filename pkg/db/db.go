package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

const schema = "CREATE TABLE scheduler (id INTEGER PRIMARY KEY AUTOINCREMENT, date CHAR(8) NULL, comment TEXT NOT NULL DEFAULT '', title VARCHAR(128) NOT NULL DEFAULT '', repeat VARCHAR(128) NOT NULL DEFAULT ''); CREATE INDEX scheduler_date ON scheduler (date);"

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	if install {
		db, err := sql.Open("sqlite", dbFile)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer db.Close()

		_, err = db.Exec(schema)
		if err != nil {
			return err
		}
	}

	return nil

}
