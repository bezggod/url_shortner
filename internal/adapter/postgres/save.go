package postgres

import (
	"context"
	"fmt"
	"url_shortner/internal/pkg/short"
)

func (r *Repo) Save(ctx context.Context, original string) (string, error) {
	var existing string
	err := r.cluster.Conn.QueryRow(ctx,
		`SELECT short_url FROM url_mappings WHERE original_url = $1`,
		original,
	).Scan(&existing)

	if err == nil {
		return existing, nil
	}

	code, err := short.NewCode()
	if err != nil {
		return "", err
	}
	_, err = r.cluster.Conn.Exec(ctx,
		`INSERT INTO url_mappings (original_url, short_url) VALUES ($1, $2)`,
		original, code,
	)
	if err != nil {
		return "", fmt.Errorf("failed to insert url mapping: %w", err)
	}

	return code, nil
}
