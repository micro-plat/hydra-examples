package main

import (
	"github.com/micro-plat/hydra/context"

	"github.com/micro-plat/hydra/component"
)

type OrderHandler struct {
}

func NewOrderHandler(c component.IContainer) (*OrderHandler, error) {
	return &OrderHandler{}, nil
}

//Handle GET,POST,PUT,DELETE缺省执行逻辑
func (o *OrderHandler) Handle(c *context.Context) interface{} {
	return "handle"
}

//GetHandle 请求方式 curl http://192.168.4.121:8090/order/request或 curl -X "GET" http://192.168.4.121:8090/order/request
func (o *OrderHandler) GetHandle(c *context.Context) interface{} {
	return "get.handle"
}

//PostHandle  curl -X "POST" http://192.168.4.121:8090/order/request
func (o *OrderHandler) PostHandle(c *context.Context) interface{} {
	return "post.handle"
}

//PutHandle curl -X "PUT" http://192.168.4.121:8090/order/request
func (o *OrderHandler) PutHandle(c *context.Context) interface{} {
	return "put.handle"
}

//DeleteHandle curl -X "DELETE" http://192.168.4.121:8090/order/request
func (o *OrderHandler) DeleteHandle(c *context.Context) interface{} {
	return "delete.handle"
}

//QueryHandle curl http://192.168.4.121:8090/order/request/query
func (o *OrderHandler) QueryHandle(c *context.Context) interface{} {
	return "query.handle"
}
