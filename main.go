package main

import (
	"icedelect/avito-test-adv/controllers"
	"icedelect/avito-test-adv/database"
	"icedelect/avito-test-adv/routes"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	AdvRoutes     routes.Adv
	AdvController controllers.Adv
)

func main() {
	// TODO: Add env config
	db := database.Connect()

	server = gin.Default()
	router := server.Group("/api")

	AdvController = controllers.NewAdv(db)
	AdvRoutes = routes.NewAdv(AdvController)

	AdvRoutes.AdvRoute(router)

	server.Run(":8080")

	// TODO: Graceful shutdown for gin
	// TODO: Graceful shutdown for gorm
}
