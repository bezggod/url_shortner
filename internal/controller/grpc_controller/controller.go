package grpc_controller

import (
	"context"
	urlpb "url_shortner/internal/pb/url_shortner"
)

type urlsUseCase interface {
	Create(ctx context.Context, original string) (string, error)
	Resolve(ctx context.Context, code string) (string, error)
}
type Controller struct {
	urlpb.UnimplementedURLShortnerServer
	urlsUseCase urlsUseCase
	baseUrl     string
}

func NewController(uc urlsUseCase) *Controller {
	return &Controller{
		urlsUseCase: uc,
	}
}
