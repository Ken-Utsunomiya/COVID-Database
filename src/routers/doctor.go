package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/projects/COVID_Database/src/handlers"
)

func DoctorRoutes(rg *gin.RouterGroup) {

	doctor := rg.Group("/doctors")
	{
		doctor.GET("/", handlers.DoctorList)
	}
}
