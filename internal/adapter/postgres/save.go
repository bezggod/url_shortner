package postgres

import (
	"context"
	"fmt"
	"url_shortner/internal/pkg/short"
)

func (r *Repo) Save(ctx context.Context, original string) (string, error) {

	code, err := short.NewCode()
	if err != nil {
		return "", err
	}
	var shortURL string
	err = r.cluster.Conn.QueryRow(ctx,
		`INSERT INTO url_mappings (original_url, short_url) VALUES ($1, $2)
		ON CONFLICT (original_url) DO UPDATE SET short_url = url_mappings.short_url 
		    RETURNING short_url`,
		original, code).Scan(&shortURL)
	if err != nil {
		return "", fmt.Errorf("failed to insert url mapping: %w", err)
	}

	return shortURL, nil
}
