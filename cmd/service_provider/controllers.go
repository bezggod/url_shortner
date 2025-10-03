package service_provider

import (
	"url_shortner/internal/controller/grpc_controller"
	"url_shortner/internal/controller/http_controller"
)

func (s *ServiceProvider) getURLHTTPController() *http_controller.Controller {
	if s.urlController == nil {
		s.urlController = http_controller.NewController(s.getUseCase())
	}
	return s.urlController
}

func (s *ServiceProvider) getURLGRPCController() *grpc_controller.Controller {
	if s.grpcController == nil {
		s.grpcController = grpc_controller.NewController(s.getUseCase())
	}
	return s.grpcController
}
