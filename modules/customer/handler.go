package customer

import (
	"context"

	"db.com/modules/dataview"
	"db.com/modules/dbutil"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	mBb, err := dbutil.GormDB(context.Background())
	if err != nil {
		panic(err)
	}
	db = mBb
}

func AddRoutes(r *gin.Engine) {

	dataview.RegisterDataView("customers", "SELECT * FROM customers")

	r.POST("/customers", CreateCustomer)
	r.GET("/customers/:id", GetCustomer)
	r.PUT("/customers/:id", UpdateCustomer)
	r.DELETE("/customers/:id", DeleteCustomer)

}

func CreateCustomer(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Exec("INSERT INTO customers (title, phone_number) VALUES (?, ?)", customer.Title, customer.PhoneNumber)
	c.JSON(200, customer)
}
func GetCustomer(c *gin.Context) {
	var customer Customer
	id := c.Param("id")
	db.Raw("SELECT * FROM customers WHERE id = ?", id).Scan(&customer)
	if customer.Id == 0 {
		c.JSON(404, gin.H{"error": "Customer not found"})
		return
	}
	c.JSON(200, customer)
}

func UpdateCustomer(c *gin.Context) {
	var customer Customer
	id := c.Param("id")
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Exec("UPDATE customers SET title = ?, phone_number = ? WHERE id = ?", customer.Title, customer.PhoneNumber, id)
	c.JSON(200, customer)
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	db.Exec("DELETE FROM customers WHERE id = ?", id)
	c.JSON(200, gin.H{"success": "Customer deleted"})
}
