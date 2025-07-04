package user

import (
	"context"

	v1 "server/app/admin/api/user/v1"
	"server/app/admin/internal/logic/user"
)

func (c *ControllerV1) GetUserRoutes(ctx context.Context, req *v1.GetUserRoutesReq) (res *v1.GetUserRoutesRes, err error) {
	return user.New().GetUserRoutes(ctx, req)
}
