package main

import (
	"github.com/micro-plat/hydra/context"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/hydra"
)

var app = hydra.NewApp(
	hydra.WithPlatName("hydra-examples"),
	hydra.WithSystemName("rpcserver01"),
	hydra.WithServerTypes("api-rpc"),
	hydra.WithClusterName("test"))

func main() {
	app.Micro("/order/query", query)
	app.Micro("/order/request", component.NewRPCSerivce("/order/query@rpcserver01.hydra-examples", map[string]string{
		"sysid": "abc",
	}))
	app.Micro("/order/request/ip", component.NewRPCSerivce("/order/query@192.168.4.121:8081", map[string]string{
		"sysid": "abc",
	}))

	app.Start()
}
func query(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------query------")
	return ctx.Request.GetString("sysid")
}
