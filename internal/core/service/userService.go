package service

import (
	"github.com/stefanowolf/go-expense-manager/internal/core/domain"
	"github.com/stefanowolf/go-expense-manager/internal/core/dto"
	"github.com/stefanowolf/go-expense-manager/internal/infra/datasource/mysql"
	"github.com/stefanowolf/go-expense-manager/internal/infra/errs"
	"time"
)

type UserService interface {
	GetAll() ([]*dto.UserResponse, *errs.AppError)
	NewUser(dto.NewUserRequest) (*dto.UserResponse, *errs.AppError)
}

type DefaultUserService struct {
	userRepo mysql.UserRepositoryDB
}

func (s DefaultUserService) GetAll() ([]*dto.UserResponse, *errs.AppError) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var response []*dto.UserResponse
	for _, u := range users {
		converted := u.ToDto()
		response = append(response, &converted)
	}

	return response, nil
}

func (s DefaultUserService) NewUser(req dto.NewUserRequest) (*dto.UserResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	u := domain.User{
		Id:        "",
		Name:      req.Name,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	newUser, err := s.userRepo.Save(u)
	if err != nil {
		return nil, err
	}
	response := newUser.ToDto()
	return &response, nil
}

func NewUserService(userRepo mysql.UserRepositoryDB) UserService {
	return DefaultUserService{userRepo}
}
