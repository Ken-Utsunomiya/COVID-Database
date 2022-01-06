package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/projects/COVID_Database/src/handlers"
)

func HospitalRoutes(rg *gin.RouterGroup) {

	hospital := rg.Group("/hospitals")
	{
		hospital.GET("/", handlers.HospitalList)
		hospital.GET("/:hospital_id", handlers.HospitalRetrieve)
		hospital.POST("/", handlers.HospitalRegister)
		hospital.DELETE("/:hospital_id", handlers.HospitalDelete)
		hospital.GET("/:hospital_id/doctors", handlers.HospitalDoctorList)
		hospital.GET("/:hospital_id/doctors/:doctor_id", handlers.HospitalDoctorRetrieve)
	}
}
