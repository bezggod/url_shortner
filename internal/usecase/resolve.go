package usecase

import "context"

func (uc *UseCase) Resolve(ctx context.Context, code string) (string, error) {
	return uc.urlRepo.Get(ctx, code)
}
