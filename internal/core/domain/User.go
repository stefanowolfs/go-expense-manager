package domain

import (
	"github.com/stefanowolf/go-expense-manager/internal/core/dto"
	"github.com/stefanowolf/go-expense-manager/internal/infra/errs"
)

type User struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

// ToDto Converts domain obj into a DTO.
func (c User) ToDto() dto.UserResponse {
	return dto.UserResponse{
		Id:   c.Id,
		Name: c.Name,
	}
}

type UserRepo interface {
	FindAll() ([]User, *errs.AppError)
}
