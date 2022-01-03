package models

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Price struct {
	gorm.Model
	Amount string `yaml:"amount" valid:"string"`
	Pair   string `yaml:"pair" valid:"string"`
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

func PriceAverage(db *gorm.DB) (string, error) {
	var prices []Price
	var count int64
	db.Find(&prices).Scan(&prices)
	db.Model(&Price{}).Count(&count)
	var total decimal.Decimal
	for _, v := range prices {
		if amount, err := decimal.NewFromString(v.Amount); err != nil {
			return "", err
		} else {
			total = total.Add(amount)
		}
	}

	average := total.Div(decimal.NewFromInt(count))

	return average.String(), nil
}
