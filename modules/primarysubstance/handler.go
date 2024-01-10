package primarysubstance

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

	dataview.RegisterDataView("primary_substances", "SELECT * FROM primary_substances")
	r.POST("/primarysubstances", CreatePrimarySubstance)
	r.GET("/primarysubstances/:id", GetPrimarySubstance)
	r.PUT("/primarysubstances/:id", UpdatePrimarySubstance)
	r.DELETE("/primarysubstances/:id", DeletePrimarySubstance)

}

func CreatePrimarySubstance(c *gin.Context) {
	var primarySubstance PrimarySubstance
	if err := c.ShouldBindJSON(&primarySubstance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use raw SQL to insert a new Primary Substance
	result := db.Exec("INSERT INTO primary_substances (title) VALUES (?)", primarySubstance.Title)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, primarySubstance)
}

func GetPrimarySubstance(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var primarySubstance PrimarySubstance
	// Use raw SQL to fetch a Primary Substance by ID
	result := db.Raw("SELECT * FROM primary_substances WHERE id = ?", id).Scan(&primarySubstance)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Primary Substance not found"})
		return
	}

	c.JSON(http.StatusOK, primarySubstance)
}

func UpdatePrimarySubstance(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var primarySubstance PrimarySubstance
	if err := c.ShouldBindJSON(&primarySubstance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use raw SQL to update a Primary Substance by ID
	result := db.Exec("UPDATE primary_substances SET title = ? WHERE id = ?", primarySubstance.Title, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Primary Substance not found"})
		return
	}

	c.JSON(http.StatusOK, primarySubstance)
}

func DeletePrimarySubstance(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Use raw SQL to delete a Primary Substance by ID
	result := db.Exec("DELETE FROM primary_substances WHERE id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Primary Substance not found"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
