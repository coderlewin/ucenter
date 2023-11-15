package domain

import (
	"github.com/coderlewin/ucenter/internal/constants"
	"github.com/coderlewin/ucenter/pkg/errno"
	"github.com/coderlewin/ucenter/pkg/utils"
	"github.com/duke-git/lancet/v2/compare"
	"github.com/duke-git/lancet/v2/cryptor"
	"time"
)

type User struct {
	ID            int64
	Username      string
	UserAccount   string
	AvatarURL     string    // 用户头像
	Gender        int32     // 性别
	UserPassword  string    // 密码
	CheckPassword string    // 确认密码
	Phone         string    // 电话
	Email         string    // 邮箱
	UserStatus    int32     // 用户状态 0-正常
	UserRole      int32     // 用户角色 0-普通用户 1-管理员
	PlanetCode    string    // 星球编号
	CreateTime    time.Time // 创建时间
	UpdateTime    time.Time // 更新时间
}

func (u *User) IsAdmin() bool {
	return u.UserRole == constants.AdminRole
}

func (u *User) IsFreeze() bool {
	return u.UserStatus == constants.UserStatusDisabled
}

func (u *User) ValidateLoginParameters() error {
	// 参数不能为空
	if utils.IsAnyStringBlank(u.UserAccount, u.UserPassword) {
		return errno.ErrParameterInvalid
	}

	// 账号长度不小于 4
	if len(u.UserAccount) < 4 {
		return errno.ErrParameterInvalid.SetDescription("账号长度过短")
	}

	// 密码长度不小于 8
	if len(u.UserPassword) < 8 {
		return errno.ErrParameterInvalid.SetDescription("密码长度过短")
	}

	// 账号不能包含特殊字符
	if utils.HasSpecialText(u.UserAccount) {
		return errno.ErrParameterInvalid.SetDescription("账号包含特殊字符")
	}

	return nil
}

func (u *User) ValidateRegisterParameters() error {
	// 参数不能为空
	if utils.IsAnyStringBlank(u.UserAccount, u.UserPassword, u.CheckPassword, u.PlanetCode) {
		return errno.ErrParameterInvalid
	}

	// 账号长度不小于 4
	if len(u.UserAccount) < 4 {
		return errno.ErrParameterInvalid.SetDescription("账号长度过短")
	}

	// 密码长度不小于 8
	if len(u.UserPassword) < 8 || len(u.CheckPassword) < 8 {
		return errno.ErrParameterInvalid.SetDescription("密码长度过短")
	}
	// 账号不能包含特殊字符
	if utils.HasSpecialText(u.UserAccount) {
		return errno.ErrParameterInvalid.SetDescription("账号包含特殊字符")
	}

	// 密码和确认密码相等
	if !compare.Equal(u.UserPassword, u.CheckPassword) {
		return errno.ErrParameterInvalid.SetDescription("密码和校验密码不一致")
	}

	// 星球编号长度不大于 5
	if len(u.PlanetCode) > 5 {
		return errno.ErrParameterInvalid.SetDescription("星球编号长度过长")
	}
	return nil
}

func (u *User) EncryptPassword() {
	u.UserPassword = cryptor.Md5String(constants.PwdSalt + u.UserPassword)
}
