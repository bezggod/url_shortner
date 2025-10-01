package service_provider

import "net/http"

func (s *ServiceProvider) GetRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	ctrl := s.getURLController()
	mux.HandleFunc("POST /shorten", ctrl.Create)
	mux.HandleFunc("GET /{code}", ctrl.Resolve)
	return mux
}
