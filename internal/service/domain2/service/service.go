package service

import (
	"context"

	"gin_layout/internal/service/domain2/entity"
	"gin_layout/internal/service/domain2/handler"
)

type Service struct {
}

func (s *Service) GetWorld(ctx context.Context, req *entity.GetWorldRequest) (*entity.GetWorldResponse, error) {
	h := handler.NewGetWorldHandler(ctx, req)
	if err := h.Handle(); err != nil {
		return nil, err
	}
	return h.Resp, nil
}
