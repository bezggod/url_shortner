package service_provider

import "net/http"

func (s *ServiceProvider) GetRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", s.GetURLController().Create)
	mux.HandleFunc("GET /{code}", s.GetURLController().Resolve)

	return mux
}
