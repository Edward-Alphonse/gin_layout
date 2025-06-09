package handler

import (
	"context"

	"github.com/pkg/errors"

	"gin_layout/internal/libs/db"
	"gin_layout/internal/service/domain1/api"
	"gin_layout/internal/service/domain1/repo"
)

type GetHelloHandler struct {
	ctx context.Context
	req *api.GetHelloRequest

	Resp *api.GetHelloResponse
}

func NewGetHelloHandler(ctx context.Context, req *api.GetHelloRequest) *GetHelloHandler {
	return &GetHelloHandler{
		ctx: ctx,
		req: req,
	}
}

func (h *GetHelloHandler) Handle() error {
	helloRepo := repo.NewHelloRepo(db.GetDB())
	res, err := helloRepo.GetText(h.ctx, h.req.Id)
	if err != nil {
		return errors.Wrap(err, "GetHello")
	}
	h.Resp = &api.GetHelloResponse{
		Text: res.Text,
	}
	return nil
}
