package repository

import (
	"context"
	"github.com/coderlewin/ucenter/internal/domain"
	"github.com/coderlewin/ucenter/internal/infrastructure/entity"
	"github.com/coderlewin/ucenter/internal/infrastructure/persistence"
	daomocks "github.com/coderlewin/ucenter/internal/infrastructure/persistence/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func Test_userRepository_GetOneById(t *testing.T) {
	testCases := []struct {
		name string
		// 返回 mock 的 UserDAO
		mock func(ctrl *gomock.Controller) persistence.UserDAO

		// 输入
		ctx context.Context
		id  int64

		// 预期输出
		wantUser domain.User
		wantErr  error
	}{
		{
			name: "找到了用户",
			mock: func(ctrl *gomock.Controller) persistence.UserDAO {
				d := daomocks.NewMockUserDAO(ctrl)
				// 注意这边，我们传入的是 int64，
				// 所以要做一个显式的转换，不然默认 12 是 int 类型
				d.EXPECT().GetByID(gomock.Any(), int64(2)).Return(entity.User{
					ID:       2,
					Username: "Lewin",
				}, nil)
				return d
			},

			ctx: context.Background(),
			id:  2,

			wantUser: domain.User{
				ID:       2,
				Username: "Lewin",
			},
			wantErr: nil,
		},
		{
			name: "未找到用户",
			mock: func(ctrl *gomock.Controller) persistence.UserDAO {
				d := daomocks.NewMockUserDAO(ctrl)
				// 注意这边，我们传入的是 int64，
				// 所以要做一个显式的转换，不然默认 12 是 int 类型
				d.EXPECT().GetByID(gomock.Any(), int64(32)).Return(entity.User{}, gorm.ErrRecordNotFound)
				return d
			},

			ctx: context.Background(),
			id:  32,

			wantUser: domain.User{},
			wantErr:  gorm.ErrRecordNotFound,
		},
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
