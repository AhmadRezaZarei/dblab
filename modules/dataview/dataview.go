package dataview

import (
	"context"
	"fmt"

	"db.com/modules/dbutil"
	"github.com/gin-gonic/gin"
)

var dataviewMap map[string]string

func init() {
	dataviewMap = make(map[string]string)
}

func RegisterDataView(key string, q string) {
	dataviewMap[key] = q
}

func AddRoutes(r *gin.Engine) {
	r.GET("/dataview/:key", Get)
}

func Get(c *gin.Context) {

	db, err := dbutil.GormDB(context.Background())

	key := c.Param("key")
	q, ok := dataviewMap[key]

	fmt.Println("here1", key)
	if !ok {
		c.JSON(404, gin.H{
			"error": "not found",
		})
		return
	}

	fmt.Println("here")
	results := make([]map[string]interface{}, 0)

	err = db.Raw(q).Scan(&results).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, results)
}
