package models

type Hospital struct {
	ID       int    `json:"hospital_id"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Tel      string `json:"tel"`
	Capacity int    `json:"capacity"`
}

func (Hospital) TableName() string {
	return "hospital"
}
