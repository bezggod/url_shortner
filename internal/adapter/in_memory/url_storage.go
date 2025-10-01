package in_memory

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("url not found")

type Repo struct {
	mu      sync.RWMutex
	byOrig  map[string]string
	byShort map[string]string
}

func NewRepo() *Repo {
	return &Repo{
		byOrig:  make(map[string]string),
		byShort: make(map[string]string),
	}
}
func (r *Repo) Close() error {
	return nil
}
