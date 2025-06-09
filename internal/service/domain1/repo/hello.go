package repo

import (
	"context"

	"gorm.io/gorm"

	"gin_layout/internal/service/domain1/model"
)

type HelloRepo struct {
	db *gorm.DB
}

func NewHelloRepo(db *gorm.DB) *HelloRepo {
	return &HelloRepo{
		db: db,
	}
}

func (repo *HelloRepo) GetText(ctx context.Context, id uint64) (*model.Hello, error) {
	res := new(model.Hello)
	err := repo.db.
		WithContext(ctx).
		Model(&model.Hello{}).
		Where("id = ?", id).
		First(res).Error
	return res, err
}
