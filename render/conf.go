package main

import (
	"github.com/micro-plat/hydra"
)

func init() {
	hydra.OnReady(func() {
		hydra.Conf.API(":8080").Render(`
        request := import("request")
response := import("response")
text := import("text")
types :=import("types")

rc:="<response><code>{@status}</code><msg>{@content}</msg></response>"


getContent := func(){  

    input:={status:response.getStatus(),content:response.getRaw()["id"]}

    if text.has_prefix(request.getPath(),"/tx/request"){
        return [200,types.translate(rc,input)]
    }
     if text.has_prefix(request.getPath(),"/tx/query"){
        return [200,"<json>","application/xml"]
    }
    return [200,response.getContent()]
}

render := getContent()
		`)

	})
}
