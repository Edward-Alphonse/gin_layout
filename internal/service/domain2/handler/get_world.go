package handler

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"

	"gin_layout/internal/infra/loader"
	"gin_layout/internal/service/domain2/entity"
)

type GetWorldHandler struct {
	ctx context.Context
	req *entity.GetWorldRequest

	Resp *entity.GetWorldResponse
}

func NewGetWorldHandler(ctx context.Context, req *entity.GetWorldRequest) *GetWorldHandler {
	return &GetWorldHandler{
		ctx: ctx,
		req: req,
	}
}

func (h *GetWorldHandler) Handle() error {
	group := make([]func() error, 0)
	getHelloLoader := loader.NewGetHelloLoader(h.ctx, h.req.Id)
	group = append(group, getHelloLoader.Load)
	if err := mr.Finish(group...); err != nil {
		return errors.Wrap(err, "GetWorldHandler")
	}
	h.Resp = &entity.GetWorldResponse{
		Text: getHelloLoader.Resp.Text,
	}
	return nil
}
