package services

import (
	"github.com/projects/COVID_Database/src/database"
	"github.com/projects/COVID_Database/src/models"
	"github.com/projects/COVID_Database/src/validators"
	"gorm.io/gorm"
)

type HospitalService struct{}

func (HospitalService) GetHospitals() ([]models.Hospital, error) {

	db := database.GetDB()

	hospitals := make([]models.Hospital, 0)
	err := db.Find(&hospitals).Error

	return hospitals, err
}

func (HospitalService) GetHospital(id int) (models.Hospital, error) {

	db := database.GetDB()

	hospital := models.Hospital{}
	err := db.First(&hospital, id).Error

	return hospital, err
}

func (HospitalService) AddHospital(request validators.HospitalRegisterRequest) (models.Hospital, error) {

	db := database.GetDB()

	hospital := validators.RegisterRequestToHospitalModel(request)

	err := db.Model(models.Hospital{}).Create(&hospital).Error

	return hospital, err
}

func (HospitalService) DeleteHospital(id int) error {

	db := database.GetDB()

	err := db.Transaction(func(tx *gorm.DB) error {

		if err := tx.First(&models.Hospital{}, id).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.Hospital{}, id).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	return err
}
