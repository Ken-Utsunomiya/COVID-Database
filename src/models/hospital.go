package models

type Hospital struct {
	ID       int    `json:"hospital_id" gorm:"column:hospital_id;primaryKey"`
	Address  string `json:"address" gorm:"column:address;unique;not null;size:60"`
	Name     string `json:"name"  gorm:"column:name;unique;not null;size:60"`
	Tel      string `json:"tel"  gorm:"column:tel;unique;not null;size:60"`
	Capacity int    `json:"capacity"  gorm:"column:address;not null"`
}

func (Hospital) TableName() string {
	return "hospital"
}
