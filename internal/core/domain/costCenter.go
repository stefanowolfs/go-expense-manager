package domain

import (
	"github.com/stefanowolf/go-expense-manager/internal/core/dto"
	"github.com/stefanowolf/go-expense-manager/internal/infra/errs"
)

type CostCenter struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

// ToDto Converts domain obj into a DTO.
func (c CostCenter) ToDto() dto.CostCenterResponse {
	return dto.CostCenterResponse{
		Id:   c.Id,
		Name: c.Name,
	}
}

type CostCenterRepo interface {
	FindAll() ([]CostCenter, *errs.AppError)
}
