package postgres

import (
	"context"
)

func (r *Repo) Get(ctx context.Context, code string) (string, error) {
	var original string
	err := r.cluster.Conn.QueryRow(ctx,
		`SELECT original_url FROM url_mappings WHERE short_url = $1`, code).
		Scan(&original)
	if err != nil {
		return "", err
	}
	return original, nil

}
