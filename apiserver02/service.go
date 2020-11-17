package main

import (
	"fmt"

	"github.com/micro-plat/hydra/context"
)

func request(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------request------")
	return map[string]interface{}{
		"product": map[string]string{
			"price": "100",
		},
	}
}
func query(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------query------")
	return fmt.Errorf("参数错误")
}
