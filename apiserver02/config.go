package main

import "github.com/micro-plat/hydra/conf"

func init() {

	// app.Conf.API.SetResponse(conf.
	// NewResponse(`{{if eq .status 200}}{"code":"{{.status}}","msg":"{{.success}}","data":{{.sdata}},"price":"{{.data.product.price}}"}{{else}}{"code":"{{.status}}","msg":"{{.err}}"}{{end}}`, `*`).
	// SetParam("success", "充值成功"))

	// app.Conf.API.SetResponse(conf.
	// 	NewResponseByStatus(`200`, `{{if eq .status 200}}{"code":"{{.status}}","msg":"{{.success}}","data":{{.sdata}},"price":"{{.data.product.price}}"}{{else}}{"code":"{{.status}}","msg":"{{.err}}"}{{end}}`, `*`).
	// 	SetParam("success", "充值成功"))

	app.Conf.API.SetResponse(conf.
		NewResponseByStatus(`{{if eq .status 200}}200{{else}}900{{end}}`, `{{if eq .status 200}}{"code":"{{.status}}","msg":"{{.success}}","data":{{.sdata}},"price":"{{.data.product.price}}"}{{else}}{"code":"{{.status}}","msg":"{{.err}}"}{{end}}`, `*`).
		SetParam("success", "充值成功"))

}
