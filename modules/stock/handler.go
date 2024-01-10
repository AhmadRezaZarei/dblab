package stock

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

	dataview.RegisterDataView("stocks", "SELECT * FROM stocks")

	r.POST("/stocks", CreateStock)
	r.GET("/stocks/:id", GetStock)
	r.PUT("/stocks/:id", UpdateStock)
	r.DELETE("/stocks/:id", DeleteStock)

}

func CreateStock(c *gin.Context) {
	var stock Stock
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use raw SQL to insert a new Stock
	result := db.Exec("INSERT INTO stocks (title) VALUES (?)", stock.Title)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, stock)
}

func GetStock(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var stock Stock
	// Use raw SQL to fetch a Stock by ID
	result := db.Raw("SELECT * FROM stocks WHERE id = ?", id).Scan(&stock)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	c.JSON(http.StatusOK, stock)
}

func UpdateStock(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var stock Stock
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use raw SQL to update a Stock by ID
	result := db.Exec("UPDATE stocks SET title = ? WHERE id = ?", stock.Title, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	c.JSON(http.StatusOK, stock)
}

func DeleteStock(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Use raw SQL to delete a Stock by ID
	result := db.Exec("DELETE FROM stocks WHERE id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
