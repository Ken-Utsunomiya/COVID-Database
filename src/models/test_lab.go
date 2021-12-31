package models

type TestLab struct {
	ID         int    `json:"test_lab_id" gorm:"column:test_lab_id;primaryKey;autoIncrement"`
	Name       string `json:"name" gorm:"column:name;not null;size:30"`
	TestMethod string `json:"test_method" gorm:"column:test_method;not null;size:20"`
}

func (TestLab) TableName() string {
	return "test_lab"
}
