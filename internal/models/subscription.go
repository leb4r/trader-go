package models

import (
	valid "github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type TopicSubscription struct {
	gorm.Model
	ARN      string `yaml:"arn" valid:"string"`
	TopicARN string
	Endpoint string `yaml:"endpoint" valid:"string"`
	Protocol string `yaml:"protocol" valid:"string"`
}

func (TopicSubscription) TableName() string {
	return "topic_subscriptions"
}

func (s *TopicSubscription) Create(db *gorm.DB) error {
	if result, err := valid.ValidateStruct(s); result {
		return err
	} else {
		db.Create(&s)
		return nil
	}
}
