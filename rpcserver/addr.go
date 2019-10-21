package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/micro-plat/lib4go/types"
)

func GetAddr(addr string, serverName string, platName string) (raddr string, faddr string, sName string, pName string, err error) {
	//拆分服务名，服务器名，平台名
	blocks := getBlocks(addr)
	if len(blocks) != 3 {
		err = fmt.Errorf("地址格式错误:%s", addr)
		return
	}
	//处理未匹配到服务名，平台名根据传入值进行设置
	sName = types.GetString(blocks[1], serverName)
	pName = types.GetString(blocks[2], platName)
	b, names := needCheckNames(blocks[0])
	if !b {
		faddr = blocks[0]
		raddr = blocks[0]
		return
	}

	//解析服务名为
	var raddrs = make([]string, 0, 3)
	var faddrs = make([]string, 0, 3)
	for _, name := range names {
		cnames := getNames(name)
		switch len(cnames) {
		case 1:
			faddrs = append(faddrs, cnames[0])
			raddrs = append(raddrs, cnames[0])
		case 2:
			faddrs = append(faddrs, cnames[0])
			raddrs = append(raddrs, cnames[1])
		}
	}
	faddr = "/" + strings.Join(faddrs, "/")
	raddr = "/" + strings.Join(raddrs, "/")
	return

}

func needCheckNames(request string) (bool, []string) {
	if !strings.Contains(request, ":") {
		return false, nil
	}
	lst := strings.Split(strings.Trim(request, "/"), "/")
	return true, lst
}

func getNames(addr string) []string {
	brackets := regexp.MustCompile(`^(:\w+)\[(\w+)\]$`)
	result := brackets.FindStringSubmatch(addr)
	if len(result) > 0 {
		return result[1:]
	}
	result = regexp.MustCompile(`^(\w+)$`).FindStringSubmatch(addr)
	if len(result) > 0 {
		return result[1:]
	}
	return nil
}

func getBlocks(addr string) []string {
	brackets := regexp.MustCompile(`^([^@]+)[@]?([\w]*)[.]?([\w]*)$`)
	result := brackets.FindStringSubmatch(addr)
	if len(result) == 0 {
		return result
	}
	return result[1:]
}
