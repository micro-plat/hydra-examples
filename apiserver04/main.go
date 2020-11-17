package main

import (
	"fmt"

	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra"
)

var app = hydra.NewApp(
	hydra.WithPlatName("hydra-examples"),
	hydra.WithSystemName("apiserver04"),
	hydra.WithServerTypes("api"),
	hydra.WithClusterName("test"),
	hydra.WithDebug())

func main() {
	app.Micro("/order/request", NewOrderHandler)
	app.Micro("/order/query", query)
	app.Start()
}
func query(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------query------")
	return fmt.Errorf("参数错误")
}
