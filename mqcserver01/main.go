package main

import (
	"fmt"

	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra"
	_ "github.com/micro-plat/lib4go/mq/lmq"
)

const queueName = `hydra-examples:order-request`

var app = hydra.NewApp(
	hydra.WithPlatName("hydra-examples"),
	hydra.WithSystemName("mqcserver01"),
	hydra.WithServerTypes("api-mqc"),
	hydra.WithClusterName("test"))

func main() {
	app.Conf.MQC.SetServer(conf.NewLMQConf())
	app.Conf.Plat.SetQueue(conf.NewLMQConf())
	app.Conf.MQC.SetQueues(conf.NewQueues().Append(queueName, "/order/query"))

	app.Flow("/order/query", query)
	app.Micro("/order/request", request)
	app.Start()
}

func request(ctx *context.Context) (r interface{}) {
	err := ctx.GetContainer().GetRegularQueue().Push(queueName, fmt.Sprintf(`{"sysid":"%s"}`,
		ctx.GetContainer().GetClusterNodes().GetCurrent().GetClusterID()))
	if err != nil {
		return err
	}
	return "success"
}

func query(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------query------")
	ctx.Log.Info(ctx.Request.GetString("sysid"))
	return "success"
}
