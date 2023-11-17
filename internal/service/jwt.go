package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/coderlewin/ucenter/internal/constants"
	"github.com/coderlewin/ucenter/internal/web/dto"
	"github.com/coderlewin/ucenter/pkg/errno"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

type JWTService interface {
	SetJWTToken(ctx *app.RequestContext, ssid string, uid int64) error
	CheckSession(ctx context.Context, ssid string) error
	SetLoginToken(ctx *app.RequestContext, uid int64) error
	ExtractTokenString(ctx *app.RequestContext) string
	ClearToken(c context.Context, ctx *app.RequestContext) error
}

func NewRedisJWTService(cmd redis.Cmdable) JWTService {
	return &redisJWTService{
		cmd:                    cmd,
		refreshTokenExpiration: time.Hour * 24 * 7,
		accessTokenExpiration:  time.Minute * 30,
	}
}

type redisJWTService struct {
	cmd redis.Cmdable
	// refresh token 的过期时间
	refreshTokenExpiration time.Duration
	// token 过期时间
	accessTokenExpiration time.Duration
}

func (r *redisJWTService) SetJWTToken(ctx *app.RequestContext, ssid string, uid int64) error {
	// 在 token 中设置一些参数字段如 用户id等
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dto.UserClaims{
		Id:        uid,
		UserAgent: string(ctx.GetHeader("User-Agent")),
		Ssid:      ssid,
		RegisteredClaims: jwt.RegisteredClaims{
			// 设置 token 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(r.accessTokenExpiration)),
		},
	})
	// 根据密钥生成 token 字符串
	tokenString, err := token.SignedString(constants.AccessTokenKey)
	if err != nil {
		return err
	}
	// 将 token 设置在请求头，前端从请求头获取token
	ctx.Header("x-jwt-token", tokenString)
	return nil
}

func (r *redisJWTService) CheckSession(ctx context.Context, ssid string) error {
	logout, err := r.cmd.Exists(ctx,
		r.key(ssid)).Result()
	if err != nil {
		return err
	}
	if logout > 0 {
		return errno.ErrUnauthorization.SetDescription("用户已经退出登录")
	}
	return nil
}

func (r *redisJWTService) SetLoginToken(ctx *app.RequestContext, uid int64) error {
	ssid := uuid.New().String()
	err := r.SetJWTToken(ctx, ssid, uid)
	if err != nil {
		return err
	}
	err = r.setRefreshToken(ctx, ssid, uid)
	return err
}

func (r *redisJWTService) setRefreshToken(ctx *app.RequestContext, ssid string, uid int64) error {
	rc := dto.RefreshClaims{
		Id:   uid,
		Ssid: ssid,
		RegisteredClaims: jwt.RegisteredClaims{
			// 设置为七天过期
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(r.refreshTokenExpiration)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rc)
	refreshTokenStr, err := refreshToken.SignedString(constants.RefreshTokenKey)
	if err != nil {
		return err
	}
	ctx.Header("x-refresh-token", refreshTokenStr)
	return nil
}

func (r *redisJWTService) ExtractTokenString(ctx *app.RequestContext) string {
	authCode := string(ctx.GetHeader("Authorization"))
	if len(authCode) <= 0 {
		return ""
	}

	authSegments := strings.SplitN(authCode, " ", 2)
	if len(authSegments) != 2 {
		// 格式不对
		return ""
	}
	return authSegments[1]
}

func (r *redisJWTService) ClearToken(c context.Context, ctx *app.RequestContext) error {
	// 设置给请求头字段为空字符串，前端会拿到并保存, 下次请求token就是个空字符串了
	ctx.Header("x-jwt-token", "")
	ctx.Header("x-refresh-token", "")
	uc := ctx.MustGet(constants.UserLoginState).(dto.UserClaims)
	return r.cmd.Set(c, r.key(uc.Ssid), "", r.refreshTokenExpiration).Err()
}

func (r *redisJWTService) key(ssid string) string {
	return fmt.Sprintf("ucenter:users:ssid:%s", ssid)
}
