package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
	"orders/backend/helpers"
	"orders/backend/models"
	"os"
)

func ConnectToDb() *gorp.DbMap {

	//config := loadDBConfigs("config/dbconf.json")

	psqlInfo := "host=localhost port=5432 dbname=dfoodie user=postgres password=password sslmode=disable"

	db, err := sql.Open("postgres", psqlInfo)
	helpers.CheckErr(err, "Error of opening a db")
	err = db.Ping()
	fmt.Println(err)
	helpers.CheckErr(err, "Ping is failed")

	dbMap := &gorp.DbMap{
		Db:      db,
		Dialect: gorp.PostgresDialect{},
	}

	dbMap.AddTableWithName(models.Order{}, "Order").SetKeys(true, "Id")
	dbMap.AddTableWithName(models.Customer{}, "Customer").SetKeys(true, "Id")
	dbMap.AddTableWithName(models.Product{}, "Products").SetKeys(true, "Id")
	dbMap.AddTableWithName(models.OrderProduct{}, "Order_products").SetKeys(true, "Id")
	return dbMap
}

func loadDBConfigs(str string) models.Configuration {

	configFile, err := os.Open(str)
	defer configFile.Close()
	helpers.CheckErr(err, "Error reading DB configs from JSON file")
	jsonParser := json.NewDecoder(configFile)
	config := models.Configuration{}
	err = jsonParser.Decode(&config)
	if err != nil {
		return models.Configuration{}
	}
	return config

}
