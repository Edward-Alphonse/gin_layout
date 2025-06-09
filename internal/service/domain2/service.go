package domain2

import (
	"context"

	"gin_layout/internal/service/domain2/api"
	"gin_layout/internal/service/domain2/handler"
)

type Service struct {
}

func (s *Service) GetWorld(ctx context.Context, req *api.GetWorldRequest) (*api.GetWorldResponse, error) {
	h := handler.NewGetWorldHandler(ctx, req)
	if err := h.Handle(); err != nil {
		return nil, err
	}
	return h.Resp, nil
}
