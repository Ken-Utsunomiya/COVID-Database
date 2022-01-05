package models

type Doctor struct {
	ID             int    `json:"doctor_id" gorm:"column:doctor_id;primaryKey;autoIncrement"`
	HospitalID     int    `json:"hospital_id" gorm:"column:hospital_id;primaryKey"`
	Specialization string `json:"specialization" gorm:"column:specialization;max=40"`
}

func (Doctor) TableName() string {
	return "doctor"
}
