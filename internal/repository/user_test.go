package repository

import (
	"context"
	"github.com/coderlewin/ucenter/internal/domain"
	"github.com/coderlewin/ucenter/internal/infrastructure/persistence"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_userRepository_GetOneById(t *testing.T) {
	testCases := []struct {
		name string
		// 返回 mock 的 UserDAO 和 UserCache
		mock func(ctrl *gomock.Controller) persistence.UserDAO

		// 输入
		ctx context.Context
		id  int64

		// 预期输出
		wantUser domain.User
		wantErr  error
	}{
		//{
		//	name: "找到了用户，直接命中缓存",
		//	mock: func(ctrl *gomock.Controller) persistence.UserDAO {
		//		d := daomocks.NewMockUserDAO(ctrl)
		//		c := cachemocks.NewMockUserCache(ctrl)
		//		// 注意这边，我们传入的是 int64，
		//		// 所以要做一个显式的转换，不然默认 12 是 int 类型
		//		c.EXPECT().Get(gomock.Any(), int64(12)).
		//			// 模拟缓存命中
		//			Return(domain.User{
		//				Id:       12,
		//				Email:    "123@qq.com",
		//				Password: "123456",
		//				Phone:    "15212345678",
		//				Ctime:    now,
		//			}, nil)
		//		return d, c
		//	},
		//
		//	ctx: context.Background(),
		//	id:  12,
		//
		//	wantUser: domain.User{
		//		Id:       12,
		//		Email:    "123@qq.com",
		//		Password: "123456",
		//		Phone:    "15212345678",
		//		Ctime:    now,
		//	},
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			d := tc.mock(ctrl)
			repo := NewUserRepository(d)
			u, err := repo.GetOneById(tc.ctx, tc.id)
			assert.Equal(t, tc.wantUser, u)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
