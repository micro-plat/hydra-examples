/*
最简单的apiserver示例
1. 将函数注册为服务
2. 服务支持任意类型数据结构作为返回值
3. 从路由中获取参数的方法
*/

package main

import "github.com/micro-plat/hydra/hydra"

var app = hydra.NewApp(
	hydra.WithPlatName("hydra-examples"),
	hydra.WithSystemName("apiserver01"),
	hydra.WithServerTypes("api"),
	hydra.WithClusterName("test"),
	hydra.WithDebug())

func main() {
	app.Micro("/order/request/:tp", request)
	app.Start()
}
