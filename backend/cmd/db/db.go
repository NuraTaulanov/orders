package main

import (
	"log"
	"orders/backend/service"
)

func main() {

	// initialize the DB and create required tables
	service.InitDb()
	log.Println("Initialized the DB successfully")
}
