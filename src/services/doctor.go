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
