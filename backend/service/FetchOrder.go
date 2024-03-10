package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"orders/backend/helpers"
	models2 "orders/backend/models"
)

func FetchOrder(g *gin.Context) {
	orderId := g.Params.ByName("id")
	if helpers.IsNumber(orderId) {

	} else {
		g.JSON(http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": "Passed order id is not valid"})
		return
	}
	log.Println("Fetching an order details for order ID:", orderId)

	var orderProducts []models2.OrderProduct
	var query = "SELECT * FROM order_products WHERE order_id=" + orderId + " ORDER BY order_product_id"

	_, err := dbmap.Select(&orderProducts, query)
	if len(orderProducts) == 0 || err != nil {
		g.JSON(http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "error": "No order with requested order Id. Invalid order ID"})
	} else {
		var orderResArr []models2.OrderResponse
		for _, op := range orderProducts {
			orderRes := models2.OrderResponse{
				OrderId:      op.OrderId,
				ProductId:    op.ProductId,
				ProductEan:   op.ProductEan,
				CustomerName: op.CustomerName,
			}
			orderResArr = append(orderResArr, orderRes)
		}
		g.JSON(http.StatusOK,
			gin.H{"status": http.StatusOK, "message": "Order details fetched successfully", "order": orderResArr})
	}

}
