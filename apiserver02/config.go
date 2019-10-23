package main

import "github.com/micro-plat/hydra/conf"

func init() {
	app.Conf.API.SetResponse(conf.
		NewResponse(`{{if eq .status 200}}{"code":"{{.status}}","msg":"{{.success}}","data":{{.sdata}},"price":"{{.data.product.price}}"}{{else}}{"code":"{{.status}}","msg":"{{.err}}"}{{end}}`, `*`).
		SetParam("success", "充值成功"))
}
