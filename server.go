package main

import (
	"COVID-Database/src/routers"
	"github.com/gin-gonic/gin"
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

	r.router.Run(":3000")

}
