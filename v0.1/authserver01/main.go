package main

import "github.com/micro-plat/hydra/hydra"

var app = hydra.NewApp(
	hydra.WithPlatName("hydra-examples"),
	hydra.WithSystemName("authserver01"),
	hydra.WithServerTypes("api"),
	hydra.WithClusterName("test"),
	hydra.WithDebug())

func main() {
	app.Micro("/order/request", request)
	app.Micro("/order/query", query)
	app.Start()
}
