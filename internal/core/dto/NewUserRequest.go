package dto

import (
	"github.com/stefanowolf/go-expense-manager/internal/infra/errs"
)

type NewUserRequest struct {
	Name string `json:"name"`
}

func (r NewUserRequest) Validate() *errs.AppError {
	if r.Name == "" {
		return errs.NewValidationError("The field 'name' is required")
	}
	return nil
}
