package main

import (
	"net/http"

	"db.com/modules/customer"
	"db.com/modules/dataview"
	"db.com/modules/employee"
	"db.com/modules/primarysubstance"
	"db.com/modules/product"
	"db.com/modules/supplier"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	customer.AddRoutes(r)
	dataview.AddRoutes(r)
	employee.AddRoutes(r)
	primarysubstance.AddRoutes(r)
	product.AddRoutes(r)
	supplier.AddRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000")
}
