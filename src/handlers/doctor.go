package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/projects/COVID_Database/src/services"
)

func DoctorList(c *gin.Context) {

	doctorService := services.DoctorService{}

	doctorList, err := doctorService.GetDoctors()
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"doctors": doctorList,
	})
}
