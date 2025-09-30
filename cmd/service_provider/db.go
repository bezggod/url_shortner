package service_provider

import (
	"context"
	"log"
	"url_shortner/internal/config"
)

func (s *ServiceProvider) getDbCluster(ctx context.Context) *config.Cluster {
	if s.dbCluster == nil {
		dbCluster, err := config.NewCluster(ctx)
		if err != nil {
			log.Fatal(err)
		}
		s.dbCluster = dbCluster
	}
	return s.dbCluster
}
