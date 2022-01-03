package models

import (
	valid "github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	ARN           string
	Subscriptions []TopicSubscription `gorm:"foreignKey:TopicARN"`
}

func (Topic) TableName() string {
	return "topics"
}

func (t *Topic) Create(db *gorm.DB) error {
	if result, err := valid.ValidateStruct(t); result {
		return err
	} else {
		db.Create(&t)
		return nil
	}
}
