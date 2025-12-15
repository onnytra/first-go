package dto

import "github.com/go-playground/validator/v10"

type HealthRequest struct {
	Status string `json:"status" validate:"required"`
}

func (l HealthRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type HealthResponse struct {
	Status string `json:"status"`
}
