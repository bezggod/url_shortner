package service_provider

import (
	"context"
	"os"
	"url_shortner/internal/adapter/in_memory"
	"url_shortner/internal/adapter/postgres"
	"url_shortner/internal/storage"
)

func (s *ServiceProvider) getURLRepo() storage.Storage {
	if s.urlRepo == nil {
		switch os.Getenv("STORAGE_MODE") {
		case "postgres":
			s.urlRepo = postgres.NewRepo(s.getDbCluster(context.Background()))
		default:
			s.urlRepo = in_memory.NewRepo()
		}
	}
	return s.urlRepo
}
