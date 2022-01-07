package app

import (
	"github.com/Nagateja1034845/microservices/controllers"
)

func mapUrls() {
	router.POST("/employees", controllers.EmpsController.Create)
	router.GET("/employee/:id", controllers.EmpsController.Get)
	//router.PATCH()
}
