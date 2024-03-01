package service

import (
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
	"github.com/rulanugrh/eirene/src/internal/repository"
	"golang.org/x/crypto/bcrypt"
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
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return nil, helper.BadRequest("Cannot generate password")
	}

	modelReq := entity.UserRegister{
		Username: req.Username,
		Email:    req.Email,
		Password: string(password),
	}

	data, err := u.repo.Register(modelReq)
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

	compare := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	if compare != nil {
		return nil, helper.Unauthorize("cannot compare password")
	}

	response := helper.UserLogin{
		Token: data.Avatar,
	}

	return &response, nil
}
