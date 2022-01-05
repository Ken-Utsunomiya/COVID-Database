package models

type HealthRecord struct {
	ID          int    `json:"health_record_id" gorm:"column:health_record_id;primaryKey;autoIncrement"`
	DrugTaken   string `json:"drug_taken" gorm:"column:drug_taken;size:50"`
	FoodAllergy string `json:"food_allergy"  gorm:"column:food_allergy;size:50"`
}

func (HealthRecord) TableName() string {
	return "health_record"
}
