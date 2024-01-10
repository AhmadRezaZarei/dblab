package drivertrip

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

	dataview.RegisterDataView("driver_trips", `SELECT dt.id ,dt.source, dt.destination, CONCAT(s.first_name, s.last_name ) as supplier_title, CONCAT(d.first_name, d.last_name) AS driver_title FROM driver_trips AS dt
		INNER JOIN employees AS d ON dt.driver_id = d.id 
		INNER JOIN suppliers AS s ON s.id = dt.supplier_id
		`)

	r.POST("/drivertrips", CreateDriverTrip)
	r.GET("/drivertrips/:id", GetDriverTrip)
	r.GET("/drivertrips", GetAllDriverTrips)
	r.PUT("/drivertrips/:id", UpdateDriverTrip)
	r.DELETE("/drivertrips/:id", DeleteDriverTrip)
}

func CreateDriverTrip(c *gin.Context) {
	var req DriverTripRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	driverTrip := req.Model

	// Use raw SQL to insert a new DriverTrip
	result := db.Exec("INSERT INTO driver_trips (supplier_id, source, destination, driver_id, date) VALUES (?, ?, ?, ?, ?)",
		driverTrip.SupplierId, driverTrip.Source, driverTrip.Destination, driverTrip.DriverId, driverTrip.Date)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var lastInsertID int64
	db.Raw("SELECT LAST_INSERT_ID()").Scan(&lastInsertID)

	for _, item := range req.Items {
		db.Exec("INSERT INTO driver_trip_items (trip_id, primary_substance_id, quantity, price) VALUES (?, ?, ? , ?)", lastInsertID, item.PrimarySubstanceId, item.Quantity, item.Price)
	}

	c.JSON(http.StatusCreated, driverTrip)
}

func GetDriverTrip(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var driverTrip DriverTrip
	// Use raw SQL to fetch a DriverTrip by ID
	result := db.Raw("SELECT * FROM driver_trips WHERE id = ?", id).Scan(&driverTrip)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "DriverTrip not found"})
		return
	}

	c.JSON(http.StatusOK, driverTrip)
}

func GetAllDriverTrips(c *gin.Context) {
	var driverTrips []DriverTrip
	// Use raw SQL to fetch all DriverTrips
	result := db.Raw("SELECT * FROM driver_trips").Scan(&driverTrips)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, driverTrips)
}

func UpdateDriverTrip(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var driverTrip DriverTrip
	if err := c.ShouldBindJSON(&driverTrip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use raw SQL to update a DriverTrip by ID
	result := db.Exec("UPDATE driver_trips SET supplier_id = ?, source = ?, destination = ?, driver_id = ?, date = ? WHERE id = ?",
		driverTrip.SupplierId, driverTrip.Source, driverTrip.Destination, driverTrip.DriverId, driverTrip.Date, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "DriverTrip not found"})
		return
	}

	c.JSON(http.StatusOK, driverTrip)
}

func DeleteDriverTrip(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Use raw SQL to delete a DriverTrip by ID
	result := db.Exec("DELETE FROM driver_trips WHERE id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "DriverTrip not found"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
