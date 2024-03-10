package service

import (
	"database/sql"
	"orders/backend/helpers"
)

func InitDb() {
	dbmap := ConnectToDb()
	defer func(Db *sql.DB) {
		err := Db.Close()
		if err != nil {

		}
	}(dbmap.Db)

	err := dbmap.CreateTablesIfNotExists()
	helpers.CheckErr(err, "Create tables failed")
}
