package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/projects/COVID_Database/src/services"
	"net/http"
	"strconv"
)

func HospitalList(c *gin.Context) {

	hospitalService := services.HospitalService{}

	hospitalList, err := hospitalService.GetHospitals()
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"hospitals": hospitalList,
	})

}

func HospitalRetrieve(c *gin.Context) {

	hospitalService := services.HospitalService{}

	hospitalId, _ := strconv.Atoi(c.Param("hospital_id"))
	hospital, err := hospitalService.GetHospital(hospitalId)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, hospital)
}
