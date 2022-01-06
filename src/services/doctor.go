package services

import (
	"github.com/projects/COVID_Database/src/database"
	"github.com/projects/COVID_Database/src/models"
)

type DoctorService struct{}

func (DoctorService) GetDoctors() ([]models.Doctor, error) {

	db := database.GetDB()

	doctors := make([]models.Doctor, 0)
	err := db.Find(&doctors).Error

	return doctors, err
}

func (DoctorService) GetHospitalDoctors(hospitalId int) ([]models.Doctor, error) {

	db := database.GetDB().
		Model(&models.Hospital{}).
		Where("hospital_id = ?", hospitalId)

	doctors := make([]models.Doctor, 0)
	err := db.Find(&doctors).Error

	return doctors, err
}

func (DoctorService) GetHospitalDoctor(hospitalId int, doctorId int) (models.Doctor, error) {

	db := database.GetDB().
		Model(&models.Hospital{}).
		Where("hospital_id = ?", hospitalId)

	doctor := models.Doctor{}
	err := db.First(&doctor, doctorId).Error

	return doctor, err
}
