package repository

import (
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(req entity.UserRegister) (*entity.User, error)
	Login(req entity.UserLogin) (*entity.User, error)
	Update(username string, model entity.User) (*entity.User, error)
}

type userrepos struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userrepos{
		db: db,
	}
}

func (u *userrepos) Register(req entity.UserRegister) (*entity.User, error) {
	var user entity.User
	user.Email = req.Email
	user.Password = req.Password
	user.Username = req.Username

	findEmail := u.db.Where("email = ?", req.Email).Find(&user)
	if findEmail.RowsAffected != 0 {
		return nil, helper.NotFound("sorry this email has been used")
	}

	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userrepos) Login(req entity.UserLogin) (*entity.User, error) {
	var user entity.User
	if err := u.db.Where("email = ?", req.Email).Find(&user); err.RowsAffected == 0 {
		return nil, helper.NotFound("email not found")
	}

	return &user, nil
}

func (u *userrepos) Update(username string, model entity.User) (*entity.User, error) {
	var user entity.User
	err := u.db.Model(&model).Where("username = ?", username)
	if err.RowsAffected == 0 {
		return nil, helper.NotFound("sorry user with this username not found")
	}

	errs := err.Updates(&user).Error
	if errs != nil {
		return nil, helper.InternalServerError("sorry cannot update user with this username")
	}

	return &user, nil

}
