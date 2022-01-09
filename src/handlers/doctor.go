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

func HospitalDoctorList(c *gin.Context) {

	hospitalService := services.HospitalService{}
	doctorService := services.DoctorService{}

	hospitalId, _ := strconv.Atoi(c.Param("hospital_id"))
	_, err := hospitalService.GetHospital(hospitalId)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	doctorList, err := doctorService.GetHospitalDoctors(hospitalId)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"doctors": doctorList,
	})
}

func HospitalDoctorRetrieve(c *gin.Context) {

	hospitalService := services.HospitalService{}
	doctorService := services.DoctorService{}

	hospitalId, _ := strconv.Atoi(c.Param("hospital_id"))
	doctorId, _ := strconv.Atoi(c.Param("doctor_id"))

	_, err := hospitalService.GetHospital(hospitalId)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	doctor, err := doctorService.GetHospitalDoctor(hospitalId, doctorId)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, doctor)
}

func HospitalDoctorRegister(c *gin.Context) {

	hospitalService := services.HospitalService{}
	doctorService := services.DoctorService{}

	hospitalId, _ := strconv.Atoi(c.Param("hospital_id"))

	_, err := hospitalService.GetHospital(hospitalId)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	registerRequest := validators.DoctorRegisterRequest{}
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		_ = c.Error(errors.New(middlewares.BadRequest)).SetType(gin.ErrorTypePublic)
		return
	}

	doctor, err := doctorService.AddDoctor(registerRequest, hospitalId)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusCreated, doctor)
}

func HospitalDoctorDelete(c *gin.Context) {

	hospitalService := services.HospitalService{}
	doctorService := services.DoctorService{}

	hospitalId, _ := strconv.Atoi(c.Param("hospital_id"))
	doctorId, _ := strconv.Atoi(c.Param("doctor_id"))

	_, err := hospitalService.GetHospital(hospitalId)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	if err := doctorService.DeleteHospitalDoctor(hospitalId, doctorId); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.Status(http.StatusNoContent)
}
