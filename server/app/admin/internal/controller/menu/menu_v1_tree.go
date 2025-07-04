package menu

import (
	"context"

	v1 "server/app/admin/api/menu/v1"
	"server/app/admin/internal/logic/menu"
)

func (c *ControllerV1) GetTree(ctx context.Context, req *v1.GetTreeReq) (res *v1.GetTreeRes, err error) {
	return menu.New().GetTree(ctx, *req)
}
