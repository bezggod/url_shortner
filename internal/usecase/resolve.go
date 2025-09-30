package usecase

import "context"

func (uc *UseCase) Resolve(code string) (string, error) {
	return uc.urlRepo.Get(context.Background(), code)
}
