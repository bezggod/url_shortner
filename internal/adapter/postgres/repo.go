package postgres

import (
	"context"
	"errors"
	"url_shortner/internal/config"
)

var NotFound = errors.New("url not found")

type Repo struct {
	cluster *config.Cluster
}

func NewRepo(cluster *config.Cluster) *Repo {
	return &Repo{
		cluster: cluster,
	}
}

func (r *Repo) Close() error {
	return r.cluster.Conn.Close(context.Background())
}
