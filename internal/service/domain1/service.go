package domain1

import (
	"context"

	"gin_layout/internal/service/domain1/api"
	"gin_layout/internal/service/domain1/handler"
)

type Service struct {
}

func (s *Service) GetHello(ctx context.Context, req *api.GetHelloRequest) (*api.GetHelloResponse, error) {
	h := handler.NewGetHelloHandler(ctx, req)
	if err := h.Handle(); err != nil {
		return nil, err
	}
	return h.Resp, nil
}
