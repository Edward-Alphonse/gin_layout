package loader

import (
	"context"

	"github.com/pkg/errors"

	"gin_layout/internal/service/domain1"
	"gin_layout/internal/service/domain1/api"
)

type GetHelloLoader struct {
	ctx context.Context
	id  uint64

	Result string
}

func NewGetHelloLoader(ctx context.Context, id uint64) *GetHelloLoader {
	return &GetHelloLoader{
		ctx: ctx,
		id:  id,
	}
}

func (l *GetHelloLoader) Load() error {
	req := &api.GetHelloRequest{
		Id: l.id,
	}

	s := domain1.Service{}
	resp, err := s.GetHello(l.ctx, req)
	if err != nil {
		return errors.Wrap(err, "GetHello")
	}

	l.Result = resp.Text
	return nil
}
