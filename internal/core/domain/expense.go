package domain

import (
	"github.com/stefanowolf/go-expense-manager/internal/core/dto"
	"github.com/stefanowolf/go-expense-manager/internal/infra/errs"
)

type Expense struct {
	Id         string `db:"id"`
	UserId     string `db:"user_id"`
	ActivityId string `db:"activity_id"`
	Name       string `db:"name"`
	Date       string `db:"billing_date"`
	Cost       string `db:"cost"`
}

// ToDto Converts domain obj into a DTO.
func (c Expense) ToDto() dto.ExpenseResponse {
	return dto.ExpenseResponse{
		Id:   c.Id,
		Name: c.Name,
		Date: c.Date,
		Cost: c.Cost,
	}
}

type ExpenseRepo interface {
	FindAll() ([]Expense, *errs.AppError)
}
