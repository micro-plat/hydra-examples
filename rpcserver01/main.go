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
	app.Micro("/order/request/2", component.NewRPCSerivce("/order/query", map[string]string{
		"sysid": "abc",
	}))
	app.Micro("/order/request/ip", component.NewRPCCtxSerivce("/order/query@192.168.4.121:8081", queryMap))
	app.Start()
}

func queryMap(ctx *context.Context) (r interface{}) {
	return map[string]string{
		"sysid": ctx.Service,
	}
}

func query(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------query------")
	return map[string]string{
		"sysid": ctx.Request.GetString("sysid"),
	}
}
