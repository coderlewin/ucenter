package repository

import (
	"context"
	"github.com/coderlewin/ucenter/internal/domain"
	"github.com/coderlewin/ucenter/internal/infrastructure/entity"
	"github.com/coderlewin/ucenter/internal/infrastructure/persistence"
	"github.com/duke-git/lancet/v2/slice"
)

//go:generate mockgen -source=./user.go -package=repomocks -destination=mocks/user.mock.go UserRepository
type UserRepository interface {
	Create(ctx context.Context, user domain.User) (int64, error)
	Delete(ctx context.Context, id int64) error
	GetOneById(ctx context.Context, id int64) (domain.User, error)
	FindByAccount(ctx context.Context, account string) (domain.User, error)
	CountByAccount(ctx context.Context, account string) (int64, error)
	CountByPlanetCode(ctx context.Context, planetCode string) (int64, error)
	ListByUsername(ctx context.Context, username string, offset int, limit int) ([]domain.User, int64, error)
}

func NewUserRepository(userDao persistence.UserDAO) UserRepository {
	return &userRepository{userDao: userDao}
}

type userRepository struct {
	userDao persistence.UserDAO
}

func (u *userRepository) ListByUsername(ctx context.Context, username string, offset int, limit int) ([]domain.User, int64, error) {
	users, total, err := u.userDao.SelectPageByUsername(ctx, username, offset, limit)
	list := slice.Map(users, func(index int, item entity.User) domain.User {
		return u.entityToDomain(item)
	})
	return list, total, err
}

func (u *userRepository) GetOneById(ctx context.Context, id int64) (domain.User, error) {
	user, err := u.userDao.GetByID(ctx, id)
	return u.entityToDomain(user), err
}

func (u *userRepository) CountByAccount(ctx context.Context, account string) (int64, error) {
	return u.userDao.Count(ctx, "user_account", account)
}

func (u *userRepository) CountByPlanetCode(ctx context.Context, planetCode string) (int64, error) {
	return u.userDao.Count(ctx, "planet_code", planetCode)
}

func (u *userRepository) FindByAccount(ctx context.Context, account string) (domain.User, error) {
	user, err := u.userDao.FindByAccount(ctx, account)
	if err != nil {
		return domain.User{}, err
	}
	ud := u.entityToDomain(user)
	return ud, nil
}

func (u *userRepository) Delete(ctx context.Context, id int64) error {
	return u.userDao.Delete(ctx, id)
}

func (u *userRepository) Create(ctx context.Context, user domain.User) (int64, error) {
	return u.userDao.Insert(ctx, u.domainToEntity(user))
}

func (u *userRepository) domainToEntity(user domain.User) entity.User {
	return entity.User{
		ID:           user.ID,
		Username:     user.Username,
		UserAccount:  user.UserAccount,
		AvatarURL:    user.AvatarURL,
		Gender:       user.Gender,
		UserPassword: user.UserPassword,
		Phone:        user.Phone,
		Email:        user.Email,
		UserStatus:   user.UserStatus,
		UserRole:     user.UserRole,
		PlanetCode:   user.PlanetCode,
	}
}

func (u *userRepository) entityToDomain(user entity.User) domain.User {
	return domain.User{
		ID:           user.ID,
		Username:     user.Username,
		UserAccount:  user.UserAccount,
		AvatarURL:    user.AvatarURL,
		Gender:       user.Gender,
		UserPassword: user.UserPassword,
		Phone:        user.Phone,
		Email:        user.Email,
		UserStatus:   user.UserStatus,
		UserRole:     user.UserRole,
		PlanetCode:   user.PlanetCode,
		CreateTime:   user.CreateTime,
		UpdateTime:   user.UpdateTime,
	}
}
