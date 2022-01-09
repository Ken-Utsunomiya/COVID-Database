package validators

import "github.com/projects/COVID_Database/src/models"

type DoctorRegisterRequest struct {
	Specialization string `json:"specialization" binding:"max=40"`
}

func RegisterRequestToDoctorModel(request DoctorRegisterRequest, hospitalId int) models.Doctor {

	doctor := models.Doctor{}
	doctor.HospitalID = hospitalId
	doctor.Specialization = request.Specialization

	return doctor
}
