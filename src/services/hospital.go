package services

import (
	"github.com/projects/COVID_Database/src/database"
	"github.com/projects/COVID_Database/src/models"
)

type HospitalService struct{}

func (HospitalService) GetHospitals() ([]models.Hospital, error) {

	db := database.GetDB()

	hospitals := make([]models.Hospital, 0)
	err := db.Find(&hospitals).Error

	return hospitals, err
}
