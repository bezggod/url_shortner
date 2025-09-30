package controller

import (
	"net/http"
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

func (c *Controller) Routes(mux *http.ServeMux) {
	mux.HandleFunc("POST /shorten", c.Create)
	mux.HandleFunc("/", c.Resolve)
}
