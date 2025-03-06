package loader

import (
	"context"

	"github.com/pkg/errors"

	"gin_layout/internal/db"
	"gin_layout/internal/infra/model"
	"gin_layout/internal/infra/repo"
)

type GetHelloLoader struct {
	ctx context.Context
	Id  uint64

	Resp *model.Hello
}

func NewGetHelloLoader(ctx context.Context, id uint64) *GetHelloLoader {
	return &GetHelloLoader{
		ctx: ctx,
		Id:  id,
	}
}

func (l *GetHelloLoader) Load() error {
	r := repo.NewHelloRepo(db.GetDB())
	res, err := r.GetText(l.ctx, l.Id)
	if err != nil {
		return errors.Wrap(err, "get hello")
	}
	l.Resp = res
	return nil
}
