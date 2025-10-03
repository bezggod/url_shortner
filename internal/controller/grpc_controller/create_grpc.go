package grpc_controller

import (
	"context"
	urlpb "url_shortner/internal/pb/url_shortner"
)

func (c *Controller) Shorten(ctx context.Context, req *urlpb.ShortenReq) (*urlpb.ShortenResp, error) {
	code, err := c.urlsUseCase.Create(ctx, req.OriginalUrl)
	if err != nil {
		return nil, err
	}
	return &urlpb.ShortenResp{ShortUrl: code}, nil
}
