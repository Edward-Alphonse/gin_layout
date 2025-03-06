package handler

import (
	"context"

	"gin_layout/internal/service/domain1/entity"
)

type GetHelloHandler struct {
	ctx context.Context
	req *entity.GetHelloRequest

	Resp *entity.GetHelloResponse
}

func NewGetHelloHandler(ctx context.Context, req *entity.GetHelloRequest) *GetHelloHandler {
	return &GetHelloHandler{
		ctx: ctx,
		req: req,
	}
}

func (handler *GetHelloHandler) Handle() error {
	panic("implement me")
}
