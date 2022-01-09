package services

import (
	"github.com/projects/COVID_Database/src/database"
	"github.com/projects/COVID_Database/src/models"
	"github.com/projects/COVID_Database/src/validators"
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
		Model(&models.Doctor{}).
		Where("hospital_id = ?", hospitalId)

	doctors := make([]models.Doctor, 0)
	err := db.Find(&doctors).Error

	return doctors, err
}

func (DoctorService) GetHospitalDoctor(hospitalId int, doctorId int) (models.Doctor, error) {

	db := database.GetDB().
		Model(&models.Doctor{}).
		Where("hospital_id = ?", hospitalId)

	doctor := models.Doctor{}
	err := db.First(&doctor, doctorId).Error

	return doctor, err
}

func (DoctorService) AddDoctor(request validators.DoctorRegisterRequest, hospitalId int) (models.Doctor, error) {

	db := database.GetDB()

	doctor := validators.RegisterRequestToDoctorModel(request, hospitalId)

	err := db.Model(models.Doctor{}).Create(&doctor).Error

	return doctor, err
}

func (DoctorService) DeleteHospitalDoctor(hospitalId int, doctorId int) error {

	db := database.GetDB().
		Model(&models.Doctor{}).
		Where("hospital_id = ?", hospitalId)

	if err := db.Delete(&models.Doctor{}, doctorId).Error; err != nil {
		return err
	}

	return nil
}
