package usecase

import (
	"log"
	"url_shortner/internal/storage"
)

type CreateURLRequest struct {
	OriginalUrl string
}

type UseCase struct {
	urlRepo storage.Storage
}

func NewUseCase(r storage.Storage) *UseCase {
	if r == nil {
		log.Fatal("!!! NewUseCase  called with NIL repo")
	}
	return &UseCase{
		urlRepo: r,
	}
}
