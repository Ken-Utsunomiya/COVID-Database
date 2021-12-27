package routers

import (
	"COVID-Database/src/handlers"
	"github.com/gin-gonic/gin"
)

func HospitalRoutes(rg *gin.RouterGroup) {

	hospital := rg.Group("/hospitals")
	{
		hospital.GET("/", handlers.HospitalList)
	}
}
