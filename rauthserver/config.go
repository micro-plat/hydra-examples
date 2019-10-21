package main

import "github.com/micro-plat/hydra/conf"

func init() {
	app.Conf.API.SetMain(conf.NewAPIServerConf(":9999"))

	////基本验证
	// app.Conf.API.SetAuthes(conf.NewAuthes().
	// 	WithServiceAuth(conf.
	// 		NewServiceAuth("/single/apiserver/md5/auth@authserver.sas_debug")))

	////指定请求参与验证，其它不验证
	// app.Conf.API.SetAuthes(conf.NewAuthes().
	// 	WithServiceAuth(conf.
	// 		NewServiceAuth("/single/apiserver/md5/auth@authserver.sas_debug", "/order/request")))

	//保存签名secret的用户编号为可以通过merid获得
	app.Conf.API.SetAuthes(conf.NewAuthes().
		WithServiceAuth(conf.
			NewServiceAuth("/single/apiserver/md5/auth@authserver.sas_debug", "/order/request").
			WithUIDAlias("merid")))

	// //设置必须参数
	// app.Conf.API.SetAuthes(conf.NewAuthes().
	// 	WithServiceAuth(conf.
	// 		NewServiceAuth("/single/apiserver/md5/auth@authserver.sas_debug", "/order/request").
	// 		WithUIDAlias("merid").
	// 		WithRequired("timestamp", "sign", "merid")))

	// //多个请求使用不同的签名方式
	// app.Conf.API.SetAuthes(conf.NewAuthes().
	// 	WithServiceAuth(
	// 		conf.NewServiceAuth("/single/apiserver/md5/auth@authserver.sas_debug", "/order/request").
	// 			WithUIDAlias("merid").
	// 			WithRequired("timestamp", "sign", "merid"),
	// 		conf.NewServiceAuth("/single/apiserver/md5/auth@authserver.sas_debug", "/order/query").
	// 			WithUIDAlias("merid")))

}
