package main

import "github.com/micro-plat/hydra/conf"

func init() {
	app.Conf.API.SetCircuitBreaker(conf.NewCircuitBreaker(0).WithForceBreak(true))
}
