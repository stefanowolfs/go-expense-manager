package domain

import (
	"github.com/stefanowolf/go-expense-manager/internal/core/dto"
	"github.com/stefanowolf/go-expense-manager/internal/infra/errs"
)

type Activity struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

// ToDto Converts domain obj into a DTO.
func (c Activity) ToDto() dto.ActivityResponse {
	return dto.ActivityResponse{
		Id:   c.Id,
		Name: c.Name,
	}
}

type ActivityRepo interface {
	FindAll() ([]Activity, *errs.AppError)
}
