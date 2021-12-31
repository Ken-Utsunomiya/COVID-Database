package main

import (
	"github.com/gin-gonic/gin"
	"github.com/projects/COVID_Database/src/database"
	"github.com/projects/COVID_Database/src/middlewares"
	"github.com/projects/COVID_Database/src/routers"
)

type routes struct {
	router *gin.Engine
}

func main() {

	r := routes{
		router: gin.Default(),
	}

	r.router.Use(middlewares.ErrorMiddleware())

	apiEngine := r.router.Group("/api")
	{
		v1 := apiEngine.Group("/v1")

		routers.HospitalRoutes(v1)
		routers.TestLabRoutes(v1)
	}

	db := database.Init()
	database.AutoMigrate(db)

	r.router.Run(":3000")

}
