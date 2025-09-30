package service_provider

import (
	"url_shortner/internal/config"
	"url_shortner/internal/controller"
	"url_shortner/internal/storage"
	"url_shortner/internal/usecase"
)

type ServiceProvider struct {
	urlRepo storage.Storage

	urlUseCase *usecase.UseCase

	urlController *controller.Controller

	dbCluster *config.Cluster
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
