package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	helpers2 "orders/backend/helpers"
	models2 "orders/backend/models"
	"time"
)

func UpdateOrder(g *gin.Context) {
	orderId := g.Params.ByName("id")
	if helpers2.IsNumber(orderId) {

	} else {
		g.JSON(http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": "Order ID passed is not a valid number."})
		return
	}

	log.Println("Updating a datails of an order Id" + orderId)
	var query_order = "SELECT * FROM orders WHERE order_id=" + orderId
	var order models2.Order
	var orderReq models2.NewOrderReq
	err := dbmap.SelectOne(&order, query_order)

	if err != nil || (models2.Order{}) == order || len(order.CustomerName) == 0 {
		g.JSON(http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "error": "No order with requested order ID. Invalid ID"})
	} else {
		var orderReq models2.NewOrderReq
		g.Bind(&orderReq)
		if orderReq.CustomerName == order.CustomerName {
			order.UpdatedAt = time.Now().UnixNano()
		} else {
			order.CustomerName = orderReq.CustomerName
			order.UpdatedAt = time.Now().UnixNano()
		}
	}
	_, err = dbmap.Update(&order)
	var orderProducts []models2.OrderProduct
	var query_order_products = "SELECT * FROM order_products where order_id=" + orderId + " ORDER BY order_product_id"
	_, dberror := dbmap.Select(&orderProducts, query_order_products)

	if dberror != nil || len(orderProducts) == 0 {
		g.JSON(http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "error": "No products found with requested order ID."})
	} else if len(orderProducts) != len(orderReq.Products) {
		g.JSON(http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": "v1 Update API supports only product details updation in the existing order. New products addition and existing products deletion from existing order will be supported in future API ver. Aborting!"})
	} else {
		for index, orderProduct := range orderProducts {
			product := orderReq.Products[index]
			orderProduct.CustomerName = orderReq.CustomerName
			orderProduct.ProductId = product.Id
			orderProduct.ProductEan = product.EanBarcode
			orderProduct.UpdatedAt = time.Now().UnixNano()

			_, err := dbmap.Update(&orderProduct)
			helpers2.CheckErr(err, "Updating order_product mapping failed in order_products table")
		}
		g.JSON(http.StatusOK,
			gin.H{"status": http.StatusOK, "message": "Order Updated successfully", "resourceId": order.Id})
	}

}
