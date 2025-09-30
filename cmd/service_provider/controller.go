package service_provider

import "url_shortner/internal/controller"

func (s *ServiceProvider) GetURLController() *controller.Controller {
	if s.urlController == nil {
		s.urlController = controller.NewController(s.urlUseCase)
	}
	return s.urlController
}
