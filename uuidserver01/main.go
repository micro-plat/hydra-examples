package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra"
)

var app = hydra.NewApp(
	hydra.WithPlatName("hydra-examples"),
	hydra.WithSystemName("uuidserver01"),
	hydra.WithServerTypes("api"),
	hydra.WithClusterName("test"))

func main() {
	app.Micro("/order/request", request)
	app.Micro("/order/query", NewOrderHandler)
	app.Start()
}
func request(ctx *hydra.Context) (r interface{}) {

	ctx.Log.Info(context.NewUUID(ctx.GetContainer()).Get())
	ctx.Log.Info(context.NewUUID(ctx.GetContainer()).GetString("R"))

	return "success"
}

type OrderHandler struct {
	container component.IContainer
}

func NewOrderHandler(c component.IContainer) (*OrderHandler, error) {
	return &OrderHandler{container: c}, nil
}

//Handle
func (o *OrderHandler) Handle(ctx *context.Context) interface{} {
	ctx.Log.Info(context.NewUUID(o.container).Get())
	ctx.Log.Info(context.NewUUID(o.container).GetString("B"))

	return "success"
}
