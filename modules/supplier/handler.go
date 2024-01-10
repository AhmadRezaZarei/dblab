package supplier

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

	dataview.RegisterDataView("suppliers", "SELECT * FROM suppliers")

	r.POST("/suppliers", CreateSupplier)
	r.GET("/suppliers/:id", GetSupplier)
	r.PUT("/suppliers/:id", UpdateSupplier)
	r.DELETE("/suppliers/:id", DeleteSupplier)

}

func CreateSupplier(c *gin.Context) {
	var supplier Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use raw SQL to insert a new Supplier
	result := db.Exec("INSERT INTO suppliers (first_name, last_name, phone_number) VALUES (?, ?, ?)", supplier.FirstName, supplier.LastName, supplier.PhoneNumber)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, supplier)
}

func GetSupplier(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var supplier Supplier
	// Use raw SQL to fetch a Supplier by ID
	result := db.Raw("SELECT * FROM suppliers WHERE id = ?", id).Scan(&supplier)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}

	c.JSON(http.StatusOK, supplier)
}

func UpdateSupplier(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var supplier Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use raw SQL to update a Supplier by ID
	result := db.Exec("UPDATE suppliers SET first_name = ?, last_name = ?, phone_number = ? WHERE id = ?", supplier.FirstName, supplier.LastName, supplier.PhoneNumber, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}

	c.JSON(http.StatusOK, supplier)
}

func DeleteSupplier(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Use raw SQL to delete a Supplier by ID
	result := db.Exec("DELETE FROM suppliers WHERE id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
