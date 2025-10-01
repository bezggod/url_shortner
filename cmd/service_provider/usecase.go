package service_provider

import "url_shortner/internal/usecase"

func (s *ServiceProvider) getUseCase() *usecase.UseCase {
	if s.urlUseCase == nil {
		s.urlUseCase = usecase.NewUseCase(s.getURLRepo())
	}
	return s.urlUseCase
}
