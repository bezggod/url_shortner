package http_controller

import (
	"url_shortner/internal/usecase"
)

type Controller struct {
	uc *usecase.UseCase
}

func NewController(uc *usecase.UseCase) *Controller {
	return &Controller{
		uc: uc,
	}
}
