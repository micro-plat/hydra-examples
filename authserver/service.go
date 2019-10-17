package main

import "github.com/micro-plat/hydra/context"

func query(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------query------")
	return "success"
}

func request(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------request------")
	return "success"
}
