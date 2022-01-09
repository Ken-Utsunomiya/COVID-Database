package validators

type DoctorRegisterRequest struct {
	Specialization string `json:"specialization" binding:"max=40"`
}
