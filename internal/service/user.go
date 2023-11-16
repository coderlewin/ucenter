package service

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/coderlewin/ucenter/internal/domain"
	"github.com/coderlewin/ucenter/internal/repository"
	"github.com/coderlewin/ucenter/pkg/core"
	"github.com/coderlewin/ucenter/pkg/errno"
	"gorm.io/gorm"
)

//go:generate mockgen -source=./user.go -package=svcmocks -destination=./mocks/user.mock.go UserService
type UserService interface {
	Register(ctx context.Context, ud domain.User) (int64, error)
	Login(ctx context.Context, ud domain.User) (domain.User, error)
	Logout(ctx context.Context, c *app.RequestContext) error
	GetCurrentUser(ctx context.Context, id int64) (domain.User, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, username string, current int, size int) ([]domain.User, int64, error)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

type userService struct {
	userRepo repository.UserRepository
}

func (svc *userService) List(ctx context.Context, username string, current int, size int) ([]domain.User, int64, error) {
	return svc.userRepo.ListByUsername(ctx, username, (current-1)*size, size)
}

func (svc *userService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errno.ErrParameterInvalid
	}

	return svc.userRepo.Delete(ctx, id)
}

func (svc *userService) GetCurrentUser(ctx context.Context, id int64) (domain.User, error) {
	user, err := svc.userRepo.GetOneById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, errno.ErrEntityNull
		}
		return domain.User{}, errno.ErrDBFailed
	}
	return user, nil
}

func (svc *userService) Logout(ctx context.Context, c *app.RequestContext) error {
	return core.RemoveUserLoginState(c)
}

func (svc *userService) Login(ctx context.Context, ud domain.User) (domain.User, error) {
	// 校验登录参数是否合法
	if err := ud.ValidateLoginParameters(); err != nil {
		return domain.User{}, err
	}

	// 密码加密
	ud.EncryptPassword()

	user, err := svc.userRepo.FindByAccount(ctx, ud.UserAccount)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, errno.ErrEntityNull.SetDescription("用户不存在")
		}
		return domain.User{}, errno.ErrDBFailed
	}

	// 对比密码
	if ud.UserPassword != user.UserPassword {
		return domain.User{}, errno.ErrEntityNull.SetDescription("账号和密码不匹配")
	}

	// 是否被冻结
	if user.IsFreeze() {
		return domain.User{}, errno.ErrForbidden.SetDescription("账号已被冻结")
	}

	return user, nil
}

func (svc *userService) Register(ctx context.Context, ud domain.User) (int64, error) {
	// 校验注册参数是否合法
	if err := ud.ValidateRegisterParameters(); err != nil {
		return 0, err
	}

	// 判断账号是否已注册
	count, err := svc.userRepo.CountByAccount(ctx, ud.UserAccount)
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, errno.ErrEntityExists.SetDescription("账号已存在")
	}

	// 判断星球编号是否已注册
	count, err = svc.userRepo.CountByPlanetCode(ctx, ud.PlanetCode)
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, errno.ErrEntityExists.SetDescription("星球编号已存在")
	}

	// 密码加密
	ud.EncryptPassword()
	// 保存用户
	return svc.userRepo.Create(ctx, ud)
}
