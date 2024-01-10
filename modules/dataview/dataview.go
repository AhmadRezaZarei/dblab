package dataview

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var dataviewMap map[string]*gorm.DB

func init() {
	dataviewMap = make(map[string]*gorm.DB)
}

func RegisterDataView(key string, q *gorm.DB) {
	dataviewMap[key] = q
}

func AddRoutes(r *gin.Engine) {
	r.GET("/dataview/:key", Get)
}

func Get(c *gin.Context) {

	key := c.Param("id")
	q, ok := dataviewMap[key]

	if !ok {
		c.JSON(404, gin.H{
			"error": "not found",
		})
		return
	}

	var results []map[string]interface{}

	err := q.Scan(&results).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, results)
}
