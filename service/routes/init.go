package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"scoring_system/service/controller"
)

func Init() *gin.Engine {

	router := gin.Default()
	router.Use(cors.Default())

	Controller := new(controller.Controller)
	router.POST("/scoring_system/api", Controller.Handle)

	return router
}