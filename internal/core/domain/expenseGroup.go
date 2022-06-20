package domain

import (
	"github.com/stefanowolf/go-expense-manager/internal/core/dto"
	"github.com/stefanowolf/go-expense-manager/internal/infra/errs"
)

type ExpenseGroup struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

// ToDto Converts domain obj into a DTO.
func (c ExpenseGroup) ToDto() dto.ExpenseGroupResponse {
	return dto.ExpenseGroupResponse{
		Id:   c.Id,
		Name: c.Name,
	}
}

type ExpenseGroupRepo interface {
	FindAll() ([]ExpenseGroup, *errs.AppError)
}
