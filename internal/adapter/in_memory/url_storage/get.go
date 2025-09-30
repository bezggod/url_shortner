package url_storage

import "context"

func (r *Repo) Get(_ context.Context, code string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if orig, ok := r.byShort[code]; ok {
		return orig, nil
	}
	return "", ErrNotFound
}
