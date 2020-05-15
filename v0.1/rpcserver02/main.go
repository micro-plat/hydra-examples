package main

import (
	"github.com/micro-plat/hydra/conf"

	"github.com/micro-plat/hydra/hydra"
)

var app = hydra.NewApp(
	hydra.WithPlatName("hydra-examples"),
	hydra.WithSystemName("rpcserver02"),
	hydra.WithServerTypes("api-rpc"),
	hydra.WithClusterName("test"),
	hydra.WithDebug())

func main() {
	app.Micro("/order/query", query)

	app.Conf.API.SetRouters(conf.NewRouters().AppendRPCProxy("/order/request", "/order/query", map[string]string{
		"id": "abcef",
	}))

	app.Start()
}

type input struct {
	ID string `json:"id"`
}

func query(ctx *hydra.Context) (r interface{}) {
	ctx.Log.Info("------query------")
	var input input
	if err := ctx.Request.Bind(&input); err != nil {
		return err
	}
	return input
}
