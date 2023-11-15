package mysql

import (
	"context"
	"fmt"
	"github.com/coderlewin/ucenter/internal/infrastructure/entity"
	"github.com/coderlewin/ucenter/internal/infrastructure/persistence"
	"gorm.io/gorm"
)

func NewUserDao(db *gorm.DB) persistence.UserDAO {
	return &userDao{db: db}
}

type userDao struct {
	db *gorm.DB
}

func (u *userDao) GetByID(ctx context.Context, id int64) (entity.User, error) {
	var user entity.User
	err := u.db.WithContext(ctx).First(&user, id).Error
	return user, err
}

func (u *userDao) SelectPageByUsername(ctx context.Context, username string, offset, limit int) ([]entity.User, int64, error) {
	tx := u.db.WithContext(ctx).Model(&entity.User{})
	if len(username) != 0 {
		tx = tx.Where("username LIKE ?", "%"+username+"%")
	}
	var list []entity.User
	var total int64
	err := tx.Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func (u *userDao) Count(ctx context.Context, col string, val any) (int64, error) {
	var count int64
	err := u.db.WithContext(ctx).Model(&entity.User{}).Where(fmt.Sprintf("%s = ?", col), val).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *userDao) FindByAccount(ctx context.Context, account string) (entity.User, error) {
	var user entity.User
	err := u.db.WithContext(ctx).Where("user_account = ?", account).First(&user).Error
	return user, err
}

func (u *userDao) Delete(ctx context.Context, id int64) error {
	return u.db.WithContext(ctx).Delete(&entity.User{}, id).Error
}

func (u *userDao) Insert(ctx context.Context, data entity.User) (int64, error) {
	err := u.db.WithContext(ctx).Create(&data).Error
	return data.ID, err
}
