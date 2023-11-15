package persistence

import (
	"context"
	"github.com/coderlewin/ucenter/internal/infrastructure/entity"
)

type baseDao[T any] interface {
	Insert(ctx context.Context, data T) (int64, error)
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (T, error)
}

type UserDAO interface {
	baseDao[entity.User]
	FindByAccount(ctx context.Context, account string) (entity.User, error)
	Count(ctx context.Context, col string, val any) (int64, error)
	SelectPageByUsername(ctx context.Context, username string, offset, limit int) ([]entity.User, int64, error)
}
