package vo

import (
	"time"
)

type UserVO struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	UserAccount string    `json:"user_account"`
	AvatarURL   string    `json:"avatar_url"`
	Gender      int32     `json:"gender"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	UserStatus  int32     `json:"user_status"`
	CreateTime  time.Time `json:"create_time"`
	UserRole    int32     `json:"user_role"`
	PlanetCode  string    `json:"planet_code"`
}
