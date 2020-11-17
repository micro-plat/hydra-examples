package main

import (
	"github.com/micro-plat/hydra/hydra"
)

func main() {
	app := hydra.NewApp(
		hydra.WithPlatName("myplat"),     //平台名称
		hydra.WithSystemName("demo"),     //系统名称
		hydra.WithClusterName("test"),    //集群名称
		hydra.WithServerTypes("api-rpc"), //多个服务器类型，使用"-"分隔
		hydra.WithRegistry("fs://../"))   //使用本地文件系统作为注册中心

	app.API("/api", api)
	app.RPC("/rpc", rpc) //注册为rpc服务
	app.Start()
}

func rpc(ctx *hydra.Context) (r interface{}) {
	return "hello world "
}
func api(ctx *hydra.Context) (r interface{}) {
	_, result, _, err := ctx.RPC.Request("/rpc", nil, nil, true)
	if err != nil {
		return err
	}
	return result
}
