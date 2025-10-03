package grpc_controller

import (
	"context"
	urlpb "url_shortner/internal/pb/url_shortner"
)

func (c *Controller) Resolve(ctx context.Context, req *urlpb.ResolveReq) (*urlpb.ResolveResp, error) {
	orig, err := c.urlsUseCase.Resolve(ctx, req.ShortUrl)
	if err != nil {
		return nil, err
	}
	return &urlpb.ResolveResp{OriginalUrl: orig}, nil
}
