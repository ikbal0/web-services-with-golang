package routers

import (
	"web-services/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)
	router.POST("/employee", controllers.SetEmployees)
	router.GET("/employees", controllers.GetEmployees)

	return router
}
