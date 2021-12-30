package validators

import "github.com/projects/COVID_Database/src/models"

type HospitalRegisterRequest struct {
	Address  string `json:"address" binding:"required,max=60"`
	Name     string `json:"name" binding:"required,max=60"`
	Tel      string `json:"tel" binding:"required,max=60"`
	Capacity int    `json:"capacity" binding:"required"`
}

func RegisterRequestToHospitalModel(request HospitalRegisterRequest) models.Hospital {

	hospital := models.Hospital{}
	hospital.Name = request.Name
	hospital.Address = request.Address
	hospital.Tel = request.Tel
	hospital.Capacity = request.Capacity

	return hospital
}
