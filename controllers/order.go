package controllers
//
//import (
//	"fmt"
//	"net/http"
//	"encoding/json"
//	"time"
//
//	"github.com/gin-gonic/gin"
//
//	"GinAPI/models"
//)
//
//
///////////Add Order////////////
//func OrderAdd(c *gin.Context) {
//	c.Header("Content-Type", "application/json")
//	c.Header("Access-Control-Allow-Origin", "*")
//
//	var order models.Orders
//	c.BindJSON(&order)
//
//	t := time.Now()
//	formattedDateNow := fmt.Sprintf("%d%02d%02d%d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
//	orderId := "OD-" + formattedDateNow
//	addOrder := models.Orders{
//		CreateDtm: t, OrderId: orderId,
//		Phone: order.Phone, Name: order.Name,
//		Address: order.Address, Menu: order.Menu,
//		TotalItem: order.TotalItem, Pay: order.Pay}
//	if err := models.DB.Create(&addOrder).Error; err != nil {
//		fmt.Printf("error add Order: %3v \n", err)
//		c.AbortWithStatus(http.StatusInternalServerError)
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"phone":    order.Phone,
//		"order_id": orderId,
//	})
//}
//
///////////Delete Order///////////
//func OrderDelete(c *gin.Context) {
//	c.Header("Content-Type", "application/json")
//	c.Header("Access-Control-Allow-Origin", "*")
//
//	var order models.Orders
//	c.BindJSON(&order)
//
//	if err := models.DB.Where("phone = ? AND order_id = ?", order.Phone, order.OrderId).Delete(&models.Orders{}).Error;
//	err != nil {
//		fmt.Printf("error delete order: %3v \n", err)
//		c.AbortWithStatus(http.StatusInternalServerError)
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"phone": order.Phone,
//		"message": "delete ok",
//	})
//}
//
//
///////////List Order by date///////////
//func OrderShowByDate(c *gin.Context) {
//	c.Header("Content-Type", "application/json")
//	c.Header("Access-Control-Allow-Origin", "*")
//
//	var query models.Query
//	c.BindJSON(&query)
//
//	var orders []models.Orders
//
//	if err := models.DB.Raw("SELECT * from orders where phone = ? AND date(create_dtm) = to_date(?, 'DD-Mon-YYYY')",
//		query.Phone, query.Date).Scan(&orders).Error;
//	err != nil {
//		fmt.Printf("error show order by date: %3v \n", err)
//		c.AbortWithStatus(http.StatusInternalServerError)
//		return
//	}
//
//	if orders != nil {
//		c.JSON(http.StatusOK, orders)
//	} else {
//		c.JSON(http.StatusOK, json.RawMessage(`[]`))
//	}
//}
//
///////////List Order by phone user number///////////
//func OrderShowByPhone(c *gin.Context) {
//	c.Header("Content-Type", "application/json")
//	c.Header("Access-Control-Allow-Origin", "*")
//
//	phone := c.Query("phone")
//	var orders []models.Orders
//
//	if err := models.DB.Raw("SELECT * from orders where phone = ?", phone).Scan(&orders).Error; err != nil {
//		fmt.Printf("error show order by phone: %3v \n", err)
//		c.AbortWithStatus(http.StatusInternalServerError)
//		return
//	}
//
//	if (orders != nil) {
//		c.JSON(http.StatusOK, orders)
//	} else {
//		c.JSON(http.StatusOK, json.RawMessage(`[]`))
//	}
//}
//
///////////List Order by order id///////////
//func OrderShowByID(c *gin.Context) {
//	c.Header("Content-Type", "application/json")
//	c.Header("Access-Control-Allow-Origin", "*")
//
//	order_id := c.Param("orderid")
//
//	var orders []models.Orders
//
//	if err := models.DB.Raw("SELECT * from orders where order_id = ?", order_id).Scan(&orders).Error; err != nil {
//		fmt.Printf("error show order by id: %3v \n", err)
//		c.AbortWithStatus(http.StatusInternalServerError)
//		return
//	}
//
//	if (orders != nil) {
//		c.JSON(http.StatusOK, orders)
//	} else {
//		c.JSON(http.StatusOK, json.RawMessage(`[]`))
//	}
//}
//
///////////Edit Order///////////
//func OrderEdit(c *gin.Context) {
//	c.Header("Content-Type", "application/json")
//	c.Header("Access-Control-Allow-Origin", "*")
//
//	var order models.Orders
//	c.BindJSON(&order)
//
//	if err := models.DB.Model(&order).Where("phone = ? AND order_id = ?", order.Phone, order.OrderId).Updates(models.Orders{
//		Menu: order.Menu, TotalItem: order.TotalItem, Pay: order.Pay}).Error;
//	err != nil {
//		fmt.Printf("error update sales payoff : %3v \n", err)
//		c.AbortWithStatus(http.StatusInternalServerError)
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"phone": order.Phone,
//		"message": "order edit success",
//	})
//}