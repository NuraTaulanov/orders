package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	helpers2 "orders/backend/helpers"
	models2 "orders/backend/models"
	"time"
)

const (
	received    = "RECEIVED"
	in_progress = "IN PROGRESS"
	shipped     = "SHIPPED"
	delivered   = "DELIVERED"
	cancelled   = "CANCELLED"
)

var dbmap = ConnectToDb()

func CreateOrder(g *gin.Context) {
	var orderReq models2.NewOrderReq
	g.Bind(&orderReq)

	if helpers2.IsEmpty(orderReq.CustomerName) {
		g.JSON(http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": "Customer name cannot be empty. Pass a valid string value"})
		return
	}
	if len(orderReq.Products) == 0 {
		g.JSON(http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": "Products cannot be empty. An order must be at least 1 product"})
		return
	}
	order := &models2.Order{
		CustomerName: orderReq.CustomerName,
		Status:       received,
		CreatedAt:    time.Now().UnixNano(),
		UpdatedAt:    time.Now().UnixNano(),
	}

	insertError := dbmap.Insert(order)

	helpers2.CheckErr(insertError, "Add new order failed in orders table")

	for _, product := range orderReq.Products {

		//Validate if the product_ean is a valid EAN-13 string
		if helpers2.IsEan(product.EanBarcode) {

			orderProduct := &models2.OrderProduct{
				OrderId:      order.Id,
				ProductId:    product.Id,
				ProductEan:   product.EanBarcode,
				CustomerName: orderReq.CustomerName,
				CreatedAt:    time.Now().UnixNano(),
				UpdatedAt:    time.Now().UnixNano(),
			}
			err := dbmap.Insert(orderProduct)
			fmt.Println(err)

			helpers2.CheckErr(err, "Add new order_product mapping failed in order_products table")
		} else {
			err := "product EAN is incorrect " + product.EanBarcode
			g.JSON(http.StatusBadRequest,
				gin.H{"status": http.StatusBadRequest, "error": err})
			return
		}

	}
	g.JSON(http.StatusCreated,
		gin.H{"status": 200, "message": "Order created successfully", "resourceId": order.Id})
}
