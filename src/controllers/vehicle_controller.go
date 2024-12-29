package controllers

import (
	"net/http"
	"vehicle/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var vehicles = []models.Vehicle{
	{
		ID:      "d93568c1-bf9b-47c4-88ab-0273a3f21f6f",
		Name:    "Toyota Corolla 2020",
		Brand:   "Corolla",
		Price:   49_999.99,
		Mileage: 35000,
		Year:    2020,
	},
}

func GetAllVehicles(c *gin.Context) {
	if len(vehicles) == 0 {
		c.Data(http.StatusNoContent, "text/plain; charset=utf-8", []byte(""))
		return
	}
	c.JSON(http.StatusOK, vehicles)
}

func GetVehicleById(c *gin.Context) {
	id := c.Param("id")

	for _, item := range vehicles {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Vehicle not found",
	})
}

func AddVehicle(c *gin.Context) {
	var vehicle models.Vehicle

	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicle.ID = uuid.New().String()
	vehicles = append(vehicles, vehicle)

	location := "/vehicle/" + vehicle.ID

	c.Header("Location", location)
	c.JSON(http.StatusCreated, gin.H{
		"message":  "Vehicle created successfully",
		"location": location,
	})
}

func UpdateVehicle(c *gin.Context) {
	id := c.Param("id")
	var updatedVehicle models.Vehicle

	if err := c.ShouldBindJSON(&updatedVehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, vehicle := range vehicles {
		if vehicle.ID == id {
			vehicles[i].Name = updatedVehicle.Name
			vehicles[i].Brand = updatedVehicle.Brand
			vehicles[i].Price = updatedVehicle.Price
			vehicles[i].Mileage = updatedVehicle.Mileage

			c.JSON(http.StatusOK, gin.H{
				"message": "Vehicle updated successfully",
				"vehicle": vehicles[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
}

func DeleteVehicle(c *gin.Context) {
	id := c.Param("id")
	found := false

	newVehicles := []models.Vehicle{}
	for _, item := range vehicles {
		if item.ID == id {
			found = true
		} else {
			newVehicles = append(newVehicles, item)
		}
	}

	if found {
		vehicles = newVehicles
		c.JSON(http.StatusOK, gin.H{"message": "Vehicle deleted successfully"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
	}
}
