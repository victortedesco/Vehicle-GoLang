package main

import (
	"vehicle/api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/vehicle", controllers.GetAllVehicles)
	router.GET("/vehicle/:id", controllers.GetVehicleById)
	router.POST("/vehicle", controllers.AddVehicle)
	router.PUT("/vehicle/:id", controllers.UpdateVehicle)
	router.DELETE("/vehicle/:id", controllers.DeleteVehicle)

	router.Run("localhost:8080")
}
