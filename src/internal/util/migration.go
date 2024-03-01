package util

import (
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.User{}, &entity.Mail{}, &entity.Attachment{})
	if err != nil {
		return helper.InternalServerError("cannot migration table")
	}

	return nil
}
