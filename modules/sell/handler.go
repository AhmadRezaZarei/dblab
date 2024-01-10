package sell

import (
	"context"
	"net/http"
	"strconv"

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

	dataview.RegisterDataView("sell_transactions", `SELECT c.title as customer_title, p.title as product_title FROM sell_transactions AS st 
		INNER JOIN products AS p ON p.id = st.product_id
		INNER JOIN customers AS c ON c.id = st.customer_id
		`)

	r.POST("/selltransactions", CreateSellTransaction)
	r.GET("/selltransactions/:id", GetSellTransaction)
	r.PUT("/selltransactions/:id", UpdateSellTransaction)
	r.DELETE("/selltransactions/:id", DeleteSellTransaction)

}
func CreateSellTransaction(c *gin.Context) {
	var sellTransaction SellTransaction
	if err := c.ShouldBindJSON(&sellTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use raw SQL to insert a new SellTransaction
	result := db.Exec("INSERT INTO sell_transactions (customer_id, product_id, quantity, price) VALUES (?, ?, ?, ?)",
		sellTransaction.CustomerId, sellTransaction.ProductId, sellTransaction.Quantity, sellTransaction.Price)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, sellTransaction)
}

func GetSellTransaction(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var sellTransaction SellTransaction
	// Use raw SQL to fetch a SellTransaction by ID
	result := db.Raw("SELECT * FROM sell_transactions WHERE id = ?", id).Scan(&sellTransaction)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SellTransaction not found"})
		return
	}

	c.JSON(http.StatusOK, sellTransaction)
}

func UpdateSellTransaction(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var sellTransaction SellTransaction
	if err := c.ShouldBindJSON(&sellTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use raw SQL to update a SellTransaction by ID
	result := db.Exec("UPDATE sell_transactions SET customer_id = ?, product_id = ?, quantity = ?, price = ? WHERE id = ?",
		sellTransaction.CustomerId, sellTransaction.ProductId, sellTransaction.Quantity, sellTransaction.Price, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SellTransaction not found"})
		return
	}

	c.JSON(http.StatusOK, sellTransaction)
}

func DeleteSellTransaction(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Use raw SQL to delete a SellTransaction by ID
	result := db.Exec("DELETE FROM sell_transactions WHERE id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SellTransaction not found"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
