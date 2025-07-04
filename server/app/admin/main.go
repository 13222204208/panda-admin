package main

import (
	_ "server/app/admin/internal/packed"

	_ "server/app/admin/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"server/app/admin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
