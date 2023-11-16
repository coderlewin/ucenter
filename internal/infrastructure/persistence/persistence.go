package persistence

import (
	"context"
	"github.com/coderlewin/ucenter/internal/infrastructure/entity"
)

//go:generate mockgen -source=./persistence.go -package=daomocks -destination=mocks/user.mock.go UserDAO
type UserDAO interface {
	Insert(ctx context.Context, data entity.User) (int64, error)
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (entity.User, error)
	FindByAccount(ctx context.Context, account string) (entity.User, error)
	Count(ctx context.Context, col string, val any) (int64, error)
	SelectPageByUsername(ctx context.Context, username string, offset, limit int) ([]entity.User, int64, error)
}
