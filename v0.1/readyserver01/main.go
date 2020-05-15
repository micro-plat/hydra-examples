package main

import "github.com/micro-plat/hydra/hydra"

import "fmt"

var app = hydra.NewApp(
	hydra.WithSystemName("apiserver01"),
	hydra.WithServerTypes("api"),
	hydra.WithClusterName("test"))

func main() {
	app.Micro("/order/request", request)
	app.Ready(func() error {
		app.PlatName = "hydra-examples-" + app.PlatName
		return nil
	})
	app.Ready(func() error {
		app.Conf.API.SetApp(fmt.Sprintf(`{"plat_name":"%s"}`, app.PlatName))
		return nil
	})
	app.Start()
}
func request(ctx *hydra.Context) (r interface{}) {
	return "success"
}
