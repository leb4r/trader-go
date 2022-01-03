package migrations

import (
	"github.com/leb4r/trader-go/internal/models"
	"gorm.io/gorm"
)

func Execute(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Price{}, &models.Topic{}, &models.TopicSubscription{}); err != nil {
		return err
	} else {
		return nil
	}
}
