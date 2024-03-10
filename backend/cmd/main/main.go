package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	service2 "orders/backend/service"
	"os"
)

func main() {
	setUpLogger()
	router := setUpRouter()
	err := router.Run()
	if err != nil {
		return
	}

}

func setUpLogger() {
	f, _ := os.Create("log/gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.SetOutput(gin.DefaultWriter)
	log.Println("Logger is setup for microservice")
}

func setUpRouter() *gin.Engine {

	router := gin.Default()

	router.Static("/frontend", "./frontend")
	v1 := router.Group("/v1/order")
	{
		v1.POST("/", service2.CreateOrder)
		v1.HEAD("/", service2.Ping)
		v1.GET("/:id", service2.FetchOrder)
		v1.PUT("/:id", service2.UpdateOrder)
	}
	return router

}
