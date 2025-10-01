package usecase

import (
	"context"
	"fmt"
)

func (uc *UseCase) Create(original string) (string, error) {
	if uc.urlRepo == nil {
		return "", fmt.Errorf("UseCase.urlRepo is nil")
	}
	return uc.urlRepo.Save(context.Background(), original)
}
