package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

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
