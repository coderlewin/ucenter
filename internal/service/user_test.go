package service

import (
	"context"
	"github.com/coderlewin/ucenter/internal/domain"
	"github.com/coderlewin/ucenter/internal/repository"
	repomocks "github.com/coderlewin/ucenter/internal/repository/mocks"
	"github.com/coderlewin/ucenter/pkg/errno"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_userService_Register(t *testing.T) {
	testCases := []struct {
		name string

		mock func(ctrl *gomock.Controller) repository.UserRepository

		// 输入
		ctx           context.Context
		account       string
		password      string
		checkPassword string
		planetCode    string

		// 预期中的输出
		wantErr    error
		wantResult int64
	}{
		{
			name: "参数校验错误",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				return repo
			},
			ctx:     context.Background(),
			account: "",
			// 这是原始的密码。然后你用这个密码调用 bcrypt 生成一个加密后的密码
			password:      "12345678",
			checkPassword: "12345678",
			planetCode:    "",
			// 这边这个返回的是，实际上就是在 mock 中返回的
			wantResult: 0,
			wantErr:    errno.ErrParameterInvalid,
		},
		{
			name: "账号长度过短",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				return repo
			},
			ctx:     context.Background(),
			account: "wei",
			// 这是原始的密码。然后你用这个密码调用 bcrypt 生成一个加密后的密码
			password:      "12345678",
			checkPassword: "12345678",
			planetCode:    "1",
			// 这边这个返回的是，实际上就是在 mock 中返回的
			wantResult: 0,
			wantErr:    errno.ErrParameterInvalid,
		},
		{
			name: "密码长度过短",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				return repo
			},
			ctx:     context.Background(),
			account: "lewin",
			// 这是原始的密码。然后你用这个密码调用 bcrypt 生成一个加密后的密码
			password:      "123456",
			checkPassword: "123456",
			planetCode:    "1",
			// 这边这个返回的是，实际上就是在 mock 中返回的
			wantResult: 0,
			wantErr:    errno.ErrParameterInvalid,
		},
		{
			name: "账号含有特殊字符",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				return repo
			},
			ctx:     context.Background(),
			account: "lew in",
			// 这是原始的密码。然后你用这个密码调用 bcrypt 生成一个加密后的密码
			password:      "123456",
			checkPassword: "123456",
			planetCode:    "1",
			// 这边这个返回的是，实际上就是在 mock 中返回的
			wantResult: 0,
			wantErr:    errno.ErrParameterInvalid,
		},
		{
			name: "密码和校验密码不相等",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				return repo
			},
			ctx:     context.Background(),
			account: "lewin",
			// 这是原始的密码。然后你用这个密码调用 bcrypt 生成一个加密后的密码
			password:      "12345678",
			checkPassword: "123456789",
			planetCode:    "1",
			// 这边这个返回的是，实际上就是在 mock 中返回的
			wantResult: 0,
			wantErr:    errno.ErrParameterInvalid,
		},
		{
			name: "星球编号长度大于5",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				return repo
			},
			ctx:     context.Background(),
			account: "lewin",
			// 这是原始的密码。然后你用这个密码调用 bcrypt 生成一个加密后的密码
			password:      "12345678",
			checkPassword: "12345678",
			planetCode:    "1321412412",
			// 这边这个返回的是，实际上就是在 mock 中返回的
			wantResult: 0,
			wantErr:    errno.ErrParameterInvalid,
		},
		{
			name: "账号已被注册",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().CountByAccount(gomock.Any(), "lewin").Return(int64(1), nil)
				return repo
			},
			ctx:     context.Background(),
			account: "lewin",
			// 这是原始的密码。然后你用这个密码调用 bcrypt 生成一个加密后的密码
			password:      "12345678",
			checkPassword: "12345678",
			planetCode:    "1",
			// 这边这个返回的是，实际上就是在 mock 中返回的
			wantResult: 0,
			wantErr:    errno.ErrEntityExists,
		},
		{
			name: "星球编号已存在",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().CountByAccount(gomock.Any(), "lewin").Return(int64(0), nil)
				repo.EXPECT().CountByPlanetCode(gomock.Any(), "1").Return(int64(1), nil)
				return repo
			},
			ctx:     context.Background(),
			account: "lewin",
			// 这是原始的密码。然后你用这个密码调用 bcrypt 生成一个加密后的密码
			password:      "12345678",
			checkPassword: "12345678",
			planetCode:    "1",
			// 这边这个返回的是，实际上就是在 mock 中返回的
			wantResult: 0,
			wantErr:    errno.ErrEntityExists,
		},
		{
			name: "注册成功",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().CountByAccount(gomock.Any(), "lewin").Return(int64(0), nil)
				repo.EXPECT().CountByPlanetCode(gomock.Any(), "1").Return(int64(0), nil)
				repo.EXPECT().
					Create(gomock.Any(), domain.User{
						UserAccount:   "lewin",
						UserPassword:  "9825417a996f1b031543e79ab88ec7ea",
						CheckPassword: "12345678",
						PlanetCode:    "1",
					}).
					Return(int64(1), nil)
				return repo
			},
			ctx:     context.Background(),
			account: "lewin",
			// 这是原始的密码。然后你用这个密码调用 bcrypt 生成一个加密后的密码
			password:      "12345678",
			checkPassword: "12345678",
			planetCode:    "1",
			// 这边这个返回的是，实际上就是在 mock 中返回的
			wantResult: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := tc.mock(ctrl)
			svc := NewUserService(repo)
			result, err := svc.Register(tc.ctx, domain.User{
				UserAccount:   tc.account,
				UserPassword:  tc.password,
				CheckPassword: tc.checkPassword,
				PlanetCode:    tc.planetCode,
			})
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantResult, result)
		})
	}
}
