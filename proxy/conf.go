package main

import (
	"github.com/micro-plat/hydra"
)

func init() {
	hydra.OnReady(func() {
		hydra.Conf.API(":8080").Proxy(`	
        request := import("request")
        app := import("app")
        text := import("text")
        types :=import("types")

        getUpCluster := func(){
            ip := request.getClientIP()
            current:= app.getCurrentClusterName()
            if text.has_prefix(ip,"192.168."){
                return types.getStringByIndex(types.exclude(app.getAllClusterNames(),current),0)
            }
            return current
        }
        upcluster := getUpCluster()
		`)

	})
}
