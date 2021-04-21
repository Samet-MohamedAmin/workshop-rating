package main

import (
	"rating/middlewares/cors"
	v1 "rating/routers/v1"
	mongoService "rating/services/mongo"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	mongoService.InitDB()

	router := gin.Default()
	router.Use(cors.CORSMiddleware())
	v1.InitRoutes(router)
	router.Run()
}
