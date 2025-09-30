package service_provider

import "url_shortner/internal/usecase"

func (s *ServiceProvider) UseCase() *usecase.UseCase {
	if s.urlUseCase == nil {
		s.urlUseCase = usecase.NewUseCase(s.urlRepo)
	}
	return s.urlUseCase
}
