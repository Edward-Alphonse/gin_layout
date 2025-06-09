package handler

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"

	"gin_layout/api"
	"gin_layout/internal/loader"
)

type HelloHandler struct {
	ctx  context.Context
	req  *api.GetHelloRequest
	Resp *api.GetHelloResponse
}

func NewHelloHandler(ctx context.Context, req *api.GetHelloRequest) *HelloHandler {
	return &HelloHandler{
		ctx: ctx,
		req: req,
	}
}

func (h *HelloHandler) Handle() error {
	group := make([]func() error, 0)
	getHelloLoader := loader.NewGetHelloLoader(h.ctx, 123)
	group = append(group, getHelloLoader.Load)
	getWorldLoader := loader.NewGetWorldLoader(h.ctx, "123")
	group = append(group, getWorldLoader.Load)
	if err := mr.Finish(group...); err != nil {
		return errors.Wrap(err, "mr.Finish")
	}

	h.Resp = &api.GetHelloResponse{
		Text1: getHelloLoader.Result,
		Text2: getWorldLoader.Result,
	}
	return nil
}
