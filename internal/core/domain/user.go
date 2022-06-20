package domain

import (
	"github.com/stefanowolf/go-expense-manager/internal/core/dto"
	"github.com/stefanowolf/go-expense-manager/internal/infra/errs"
)

type User struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	CreatedAt string `db:"created_at"`
}

// ToDto Converts domain obj into a DTO.
func (u User) ToDto() dto.UserResponse {
	return dto.UserResponse{
		Id:   u.Id,
		Name: u.Name,
	}

}

type UserRepo interface {
	FindAll() ([]User, *errs.AppError)
	Save(user User) (*User, *errs.AppError)
}
