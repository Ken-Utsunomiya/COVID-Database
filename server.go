package main

import (
	"github.com/gin-gonic/gin"
	"github.com/projects/COVID_Database/src/database"
	"github.com/projects/COVID_Database/src/routers"
)

type routes struct {
	router *gin.Engine
}

func main() {

	r := routes{
		router: gin.Default(),
	}

	apiEngine := r.router.Group("/api")
	{
		v1 := apiEngine.Group("/v1")

		routers.HospitalRoutes(v1)
	}

	database.Init()

	r.router.Run(":3000")

}
