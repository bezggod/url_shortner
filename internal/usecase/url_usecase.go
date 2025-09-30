package usecase

import (
	"url_shortner/internal/storage"
)

type CreateURLRequest struct {
	OriginalUrl string
}

type UseCase struct {
	urlRepo storage.Storage
}

func NewUseCase(repo storage.Storage) *UseCase {
	return &UseCase{
		urlRepo: repo,
	}
}
