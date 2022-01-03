package models

import (
	valid "github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Price struct {
	Amount string `yaml:"amount" valid:"string"`
	Pair   string `yaml:"pair" valid:"string"`
	gorm.Model
}

func (Price) TableName() string {
	return "prices"
}

func (p *Price) Create(db *gorm.DB) error {
	if result, err := valid.ValidateStruct(p); result {
		return err
	} else {
		db.Create(&p)
	}
	return nil
}
