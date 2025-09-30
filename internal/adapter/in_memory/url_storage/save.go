package url_storage

import (
	"context"
	"url_shortner/internal/pkg/short"
)

func (r *Repo) Save(_ context.Context, original string) (string, error) {
	r.mu.RLock()
	if sc, ok := r.byOrig[original]; ok {
		r.mu.RUnlock()
		return sc, nil
	}
	r.mu.RUnlock()

	code, err := short.NewCode()
	if err != nil {
		return "", err
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	r.byOrig[original] = code
	r.byShort[code] = original

	return original, nil
}
