package services

import (
	"github.com/projects/COVID_Database/src/database"
	"github.com/projects/COVID_Database/src/models"
)

type TestLabService struct{}

func (TestLabService) GetTestLabs() ([]models.TestLab, error) {

	db := database.GetDB()

	testLabs := make([]models.TestLab, 0)
	err := db.Find(&testLabs).Error

	return testLabs, err
}
