package loader

import (
	"context"

	"github.com/pkg/errors"

	"gin_layout/internal/service/domain2/entity"
	"gin_layout/internal/service/domain2/service"
)

type GetWorldLoader struct {
	ctx  context.Context
	text string

	Result string
}

func NewGetWorldLoader(ctx context.Context, text string) *GetWorldLoader {
	return &GetWorldLoader{
		ctx:  ctx,
		text: text,
	}
}

func (l *GetWorldLoader) Load() error {
	req := &entity.GetWorldRequest{}
	s := service.Service{}
	resp, err := s.GetWorld(l.ctx, req)
	if err != nil {
		return errors.Wrap(err, "GetWorld")
	}
	l.Result = resp.Text
	return nil
}
