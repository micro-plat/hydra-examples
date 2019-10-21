package main

import (
	"fmt"
	"testing"
)

func TestGetAddr(t *testing.T) {
	fmt.Println(GetAddr("/order/request", "100bm", "cn"))
	fmt.Println(GetAddr("/order/request@100bm", "100bm", "cn"))
	fmt.Println(GetAddr("/order/request@100bm.cn", "100bm", "cn"))
	fmt.Println(GetAddr("/order/:ident[request]/:type[md5]/auth", "100bm", "cn"))
	fmt.Println(GetAddr("/order/:ident[request]/:type[md5]/auth@100bm", "100bm", "cn"))
	fmt.Println(GetAddr("/order/:ident[request]/:type[md5]/auth@100bm.cn", "100bm", "cn"))
	t.Error("err")
}

func TestGetBlocks(t *testing.T) {
	fmt.Println(getBlocks("/order/request@100bm.cn"))
	fmt.Println(getBlocks("/order/request@100bm"))
	fmt.Println(getBlocks("/order/:ident[apiserver]@100bm.cn"))
	fmt.Println(getBlocks("/order/request@"))
	fmt.Println(getBlocks("/order/request"))

	t.Error("err")
}

func TestGetRequest(t *testing.T) {
	fmt.Println(getNames(":ident[apiserver]"))
	fmt.Println(getNames("apiserver"))
	fmt.Println(getNames("/apiserver"))
	fmt.Println(getNames("apiserver/"))
	fmt.Println(getNames("apiserver[apiserver]"))
	t.Error("err")
}
