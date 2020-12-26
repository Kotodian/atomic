package validate

import (
	"atomic/internal/log"
	"context"
	"github.com/go-playground/validator/v10"
)

func New(ctx context.Context) *validator.Validate {
	validate := validator.New()
	err := validate.RegisterValidation("NameValidationErrors", NameValidationErrors)
	if err != nil {
		log.Error(err.Error(), ctx)
		return nil
	}
	return validate
}

// 名称校验
func NameValidationErrors(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) > 6 && len(fl.Field().String()) < 10
}
