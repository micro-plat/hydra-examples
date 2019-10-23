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
