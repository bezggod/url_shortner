package service_provider

import (
	"url_shortner/internal/controller/http_controller"
)

func (s *ServiceProvider) getURLController() *http_controller.Controller {
	if s.urlController == nil {
		s.urlController = http_controller.NewController(s.getUseCase())
	}
	return s.urlController
}
