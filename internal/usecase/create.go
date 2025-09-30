package usecase

import "context"

func (uc *UseCase) Create(original string) (string, error) {
	return uc.urlRepo.Save(context.Background(), original)
}
