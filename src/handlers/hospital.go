package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/projects/COVID_Database/src/middlewares"
	"github.com/projects/COVID_Database/src/services"
	"github.com/projects/COVID_Database/src/validators"
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

func HospitalRegister(c *gin.Context) {

	hospitalService := services.HospitalService{}

	registerRequest := validators.HospitalRegisterRequest{}
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		_ = c.Error(errors.New(middlewares.BadRequest)).SetType(gin.ErrorTypePublic)
		return
	}

	hospital, err := hospitalService.AddHospital(registerRequest)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusCreated, hospital)
}

func HospitalDelete(c *gin.Context) {

	hospitalService := services.HospitalService{}

	hospitalId, _ := strconv.Atoi(c.Param("hospital_id"))
	if err := hospitalService.DeleteHospital(hospitalId); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.Status(http.StatusNoContent)
}
