package service_provider

import (
	"url_shortner/internal/config"
	"url_shortner/internal/controller/grpc_controller"
	"url_shortner/internal/controller/http_controller"
	"url_shortner/internal/storage"
	"url_shortner/internal/usecase"
)

type ServiceProvider struct {
	urlRepo storage.Storage

	urlUseCase *usecase.UseCase

	urlController  *http_controller.Controller
	grpcController *grpc_controller.Controller

	dbCluster *config.Cluster
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
