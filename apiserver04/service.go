package main

import (
	"errors"

	"github.com/micro-plat/hydra/context"

	"github.com/micro-plat/hydra/component"
)

type OrderHandler struct {
}

func NewOrderHandler(c component.IContainer) (*OrderHandler, error) {
	return &OrderHandler{}, nil
}

func (o *OrderHandler) Handle(c *context.Context) interface{} {
	return errors.New("请求出错")
}
func (o *OrderHandler) Fallback(c *context.Context) interface{} {
	return "fallback"
}
