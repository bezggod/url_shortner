package service_provider

import (
	"google.golang.org/grpc"
	"net/http"
	urlpb "url_shortner/internal/pb/url_shortner"
)

func (s *ServiceProvider) GetHTTPServer() *http.ServeMux {
	mux := http.NewServeMux()
	ctrl := s.getURLHTTPController()
	mux.HandleFunc("POST /shorten", ctrl.Create)
	mux.HandleFunc("GET /{code}", ctrl.Resolve)
	return mux
}

func (s *ServiceProvider) GetGRPCServer() *grpc.Server {
	mux := grpc.NewServer()
	urlpb.RegisterURLShortnerServer(mux, s.getURLGRPCController())
	return mux

}
