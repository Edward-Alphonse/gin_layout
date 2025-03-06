package handler

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"

	"gin_layout/api/app1"
	"gin_layout/internal/service/domain1/entity"
	domain1service "gin_layout/internal/service/domain1/service"
	"gin_layout/internal/service/domain2/loader"
)

type HelloHandler struct {
	ctx  context.Context
	req  *app1.HelloRequest
	Resp *app1.HelloResponse
}

func NewHelloHandler(ctx context.Context, req *app1.HelloRequest) *HelloHandler {
	return &HelloHandler{
		ctx: ctx,
		req: req,
	}
}

func (h *HelloHandler) Handle() error {
	req := &entity.GetHelloRequest{}
	service := domain1service.Service{}
	resp, err := service.GetHello(h.ctx, req)
	if err != nil {
		return errors.Wrap(err, "domain1service.GetHello")
	}
	group := make([]func() error, 0)
	getWorldLoader := loader.NewGetWorldLoader(h.ctx, "123")
	group = append(group, getWorldLoader.Load)
	if err := mr.Finish(group...); err != nil {
		return errors.Wrap(err, "mr.Finish")
	}

	h.Resp = &app1.HelloResponse{
		Text1: resp.Text,
		Text2: getWorldLoader.Result,
	}
	return nil
}
