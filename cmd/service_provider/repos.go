package service_provider

import (
	"context"
	"os"
	"url_shortner/internal/adapter/in_memory/url_storage"
	"url_shortner/internal/adapter/postgres"
	"url_shortner/internal/storage"
)

func (s *ServiceProvider) getURLRepo() storage.Storage {
	if s.urlRepo == nil {
		mode := os.Getenv("STORAGE_MODE")
		if mode == "postgres" {
			s.urlRepo = postgres.NewRepo(s.getDbCluster(context.Background()))
		} else {
			s.urlRepo = url_storage.NewRepo()
		}
	}
	return s.urlRepo
}
