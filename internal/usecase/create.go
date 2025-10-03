package usecase

import (
	"context"
	"fmt"
)

func (uc *UseCase) Create(ctx context.Context, original string) (string, error) {
	if uc.urlRepo == nil {
		return "", fmt.Errorf("UseCase.urlRepo is nil")
	}
	return uc.urlRepo.Save(ctx, original)
}
