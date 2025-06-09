package handler

import (
	"context"

	"gin_layout/internal/service/domain2/api"
)

type GetWorldHandler struct {
	ctx context.Context
	req *api.GetWorldRequest

	Resp *api.GetWorldResponse
}

func NewGetWorldHandler(ctx context.Context, req *api.GetWorldRequest) *GetWorldHandler {
	return &GetWorldHandler{
		ctx: ctx,
		req: req,
	}
}

func (h *GetWorldHandler) Handle() error {
	h.Resp = &api.GetWorldResponse{
		Text: "world",
	}
	return nil
}
