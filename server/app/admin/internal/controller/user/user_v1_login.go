package user

import (
	"context"

	v1 "server/app/admin/api/user/v1"
	"server/app/admin/internal/logic/user"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	return user.New().Login(ctx, *req)
}

func (c *ControllerV1) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	return user.New().RefreshToken(ctx, *req)
}

func (c *ControllerV1) Captcha(ctx context.Context, req *v1.CaptchaReq) (res *v1.CaptchaRes, err error) {
	return user.NewLogin().GetCaptcha(ctx, *req)
}
