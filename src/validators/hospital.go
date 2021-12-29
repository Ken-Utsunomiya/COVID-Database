package validators

type HospitalRegisterRequest struct {
	Address  string `json:"address" binding:"required,max=60"`
	Name     string `json:"name" binding:"required,max=60"`
	Tel      string `json:"tel" binding:"required,max=60"`
	Capacity int    `json:"capacity" binding:"required"`
}
