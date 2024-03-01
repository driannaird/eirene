package util

import (
	"github.com/rulanugrh/eirene/src/config"
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	conf := config.GetConfig()
	password, err := bcrypt.GenerateFromPassword([]byte(conf.Admin.Password), 14)
	if err != nil {
		return helper.InternalServerError("cannot generate hash password")
	}

	admin := entity.User{
		Username: "admin",
		Email:    conf.Admin.Email,
		Avatar:   "https://i.pinimg.com/564x/b5/63/d4/b563d4d23d75c344f6927b76f8b40645.jpg",
		Password: string(password),
	}

	err = db.Create(&admin).Error
	if err != nil {
		return helper.InternalServerError("cannot create admin user")
	}

	return nil
}
