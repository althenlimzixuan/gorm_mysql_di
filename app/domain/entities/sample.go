package entities

import "gorm.io/gorm"

type SampleEntities struct {
	gorm.Model
	Description string
}

func (*SampleEntities) TableName() string {
	return "samples_table"
}
