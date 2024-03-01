package service

import (
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
	"github.com/rulanugrh/eirene/src/internal/repository"
)

type UserService interface {
	Register(req entity.UserRegister) (*helper.UserRegister, error)
	Login(req entity.UserLogin) (*helper.UserLogin, error)
}

type userservice struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userservice{
		repo: repo,
	}
}

func (u *userservice) Register(req entity.UserRegister) (*helper.UserRegister, error) {
	data, err := u.repo.Register(req)
	if err != nil {
		return nil, helper.InternalServerError("sorry cannt create user")
	}

	response := helper.UserRegister{
		Email:    data.Email,
		Username: data.Username,
	}

	return &response, nil
}
func (u *userservice) Login(req entity.UserLogin) (*helper.UserLogin, error) {
	data, err := u.repo.Login(req)
	if err != nil {
		return nil, err
	}

	response := helper.UserLogin{
		Token: data.Avatar,
	}

	return &response, nil
}
