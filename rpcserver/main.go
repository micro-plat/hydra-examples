package main

import (
	"fmt"
	"time"

	"github.com/micro-plat/hydra/registry"
	_ "github.com/micro-plat/hydra/registry/zookeeper"
	"github.com/micro-plat/lib4go/logger"
)

func main() {
	///single/apiserver/md5/auth@authserver.
	start := time.Now()
	var err error
	watch := &Watcher{plat: "sas_debug", server: "authserver", service: "/single/apiserver/sha1/auth"}
	watch.client, err = registry.NewRegistryWithAddress("zk://192.168.0.101", logger.New("t"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(watch.initialize())
	fmt.Println(time.Since(start))

}

type Watcher struct {
	client  registry.IRegistry
	plat    string
	server  string
	path    string
	service string
}
