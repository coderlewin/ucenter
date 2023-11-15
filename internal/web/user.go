package web

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/coderlewin/ucenter/internal/constants"
	"github.com/coderlewin/ucenter/internal/domain"
	"github.com/coderlewin/ucenter/internal/service"
	"github.com/coderlewin/ucenter/internal/web/dto"
	"github.com/coderlewin/ucenter/internal/web/middleware"
	"github.com/coderlewin/ucenter/internal/web/vo"
	"github.com/coderlewin/ucenter/pkg/core"
	"github.com/coderlewin/ucenter/pkg/errno"
	"github.com/duke-git/lancet/v2/slice"
	"strings"
)

type UserHandler struct {
	userSvc service.UserService
}

func NewUserHandler(userSvc service.UserService) *UserHandler {
	return &UserHandler{userSvc: userSvc}
}

// ConfigRoutes 配置路由
func (u *UserHandler) ConfigRoutes(h *route.RouterGroup) {
	group := h.Group("/user")
	{
		group.POST("/register", u.register)
		group.POST("/login", u.login)
		group.GET("/current", u.getCurrentUser)
		group.POST("/logout", u.logout)
		group.Use(middleware.NewCheckRoleMiddlewareBuilder(constants.AdminRole).Build())
		group.GET("/search", u.search)
		group.DELETE("/:id", u.delete)
	}
}

// search 用户列表搜索
func (u *UserHandler) search(ctx context.Context, c *app.RequestContext) {
	var req dto.UserSearchQuery
	if err := c.BindAndValidate(&req); err != nil {
		core.SendResponse(c, errno.ErrParameterInvalid.SetDescription(err.Error()), nil)
		return
	}

	list, total, err := u.userSvc.List(ctx, req.Username, req.Current, req.Size)
	if err != nil {
		core.SendResponse(c, err, nil)
		return
	}
	// 将数据转换为 VO
	records := slice.Map(list, func(index int, user domain.User) *vo.UserVO {
		return u.domainToUserVO(user)
	})

	core.SendResponse(c, nil, vo.PageResult{
		Records: records,
		Total:   total,
	})
}

// getCurrentUser 获取当前用户信息
func (u *UserHandler) getCurrentUser(ctx context.Context, c *app.RequestContext) {
	value, exists := c.Get(constants.LoginUser)
	loginUser := value.(*vo.UserVO)
	if !exists {
		core.SendResponse(c, errno.ErrUnauthorization, nil)
		return
	}
	user, err := u.userSvc.GetCurrentUser(ctx, loginUser.ID)
	if err != nil {
		core.SendResponse(c, err, nil)
		return
	}
	core.SendResponse(c, nil, u.domainToUserVO(user))
}

// delete 删除用户
func (u *UserHandler) delete(ctx context.Context, c *app.RequestContext) {
	var req dto.IdInPathDTO
	if err := c.BindAndValidate(&req); err != nil {
		core.SendResponse(c, errno.ErrParameterInvalid.SetDescription(err.Error()), nil)
		return
	}

	value, exists := c.Get(constants.LoginUser)
	loginUser := value.(*vo.UserVO)
	if !exists {
		core.SendResponse(c, errno.ErrUnauthorization, nil)
		return
	}

	if loginUser.ID == req.ID {
		core.SendResponse(c, errno.ErrParameterInvalid.SetDescription("不能删除本人"), nil)
		return
	}

	err := u.userSvc.Delete(ctx, req.ID)
	if err != nil {
		core.SendResponse(c, err, false)
		return
	}

	core.SendResponse(c, nil, true)
}

// logout 用户退出
func (u *UserHandler) logout(ctx context.Context, c *app.RequestContext) {
	err := u.userSvc.Logout(ctx, c)
	if err != nil {
		core.SendResponse(c, err, nil)
		return
	}
	core.SendResponse(c, nil, true)
}

// login 用户登录
func (u *UserHandler) login(ctx context.Context, c *app.RequestContext) {
	var req dto.UserLoginDTO
	if err := c.BindAndValidate(&req); err != nil {
		core.SendResponse(c, errno.ErrParameterInvalid.SetDescription(err.Error()), nil)
		return
	}
	user, err := u.userSvc.Login(ctx, domain.User{
		UserAccount:  req.Account,
		UserPassword: req.Password,
	})
	if err != nil {
		core.SendResponse(c, err, nil)
		return
	}
	userVO := u.domainToUserVO(user)
	err = core.SetUserLoginState(c, &userVO)
	if err != nil {
		core.SendResponse(c, err, nil)
		return
	}
	core.SendResponse(c, nil, userVO)
}

// register 用户注册
func (u *UserHandler) register(ctx context.Context, c *app.RequestContext) {
	var req dto.UserRegisterDTO
	if err := c.BindAndValidate(&req); err != nil {
		core.SendResponse(c, errno.ErrParameterInvalid.SetDescription(err.Error()), nil)
		return
	}
	id, err := u.userSvc.Register(ctx, domain.User{
		Username:      strings.ToUpper(req.Account),
		UserAccount:   req.Account,
		AvatarURL:     constants.DefaultAvatar,
		UserPassword:  req.Password,
		CheckPassword: req.CheckPassword,
		PlanetCode:    req.PlanetCode,
	})
	if err != nil {
		core.SendResponse(c, err, nil)
		return
	}
	core.SendResponse(c, nil, id)
}

// domainToUserVO 领域模型转视图模型
func (u *UserHandler) domainToUserVO(user domain.User) *vo.UserVO {
	return &vo.UserVO{
		ID:          user.ID,
		Username:    user.Username,
		UserAccount: user.UserAccount,
		AvatarURL:   user.AvatarURL,
		Gender:      user.Gender,
		Phone:       user.Phone,
		Email:       user.Email,
		UserStatus:  user.UserStatus,
		CreateTime:  user.CreateTime,
		UserRole:    user.UserRole,
		PlanetCode:  user.PlanetCode,
	}
}
