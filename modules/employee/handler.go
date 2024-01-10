package employee

import (
	"context"
	"net/http"

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

	dataview.RegisterDataView("employees", "SELECT * FROM employees")

	r.POST("/employee", createEmployee)
	r.GET("/employees", getEmployees)
	r.PUT("/employee", updateEmployee)
	r.DELETE("/employee/:id", deleteEmployee)

}

func createEmployee(c *gin.Context) {
	var employee Employee
	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Exec("INSERT INTO employees (first_name, last_name, birth_date, rank) VALUES (?, ?, ?, ?)",
		employee.FirstName, employee.LastName, employee.BirthDate, employee.Rank)
	c.JSON(http.StatusOK, gin.H{"message": "Employee created"})
}

func getEmployees(c *gin.Context) {
	var employees []Employee
	db.Raw("SELECT * FROM employees").Scan(&employees)
	c.JSON(http.StatusOK, employees)
}

func updateEmployee(c *gin.Context) {
	var employee Employee
	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Exec("UPDATE employees SET first_name = ?, last_name = ?, birth_date = ?, rank = ? WHERE id = ?",
		employee.FirstName, employee.LastName, employee.BirthDate, employee.Rank, employee.Id)
	c.JSON(http.StatusOK, gin.H{"message": "Employee updated"})
}

func deleteEmployee(c *gin.Context) {
	id := c.Param("id")
	db.Exec("DELETE FROM employees WHERE id = ?", id)
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}
