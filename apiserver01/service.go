package main

import (
	"fmt"
	"time"

	"github.com/micro-plat/hydra/context"
)

func request(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------query------")
	switch ctx.Request.Param.GetInt("tp") {
	case 1:
		return "success"
	case 2:
		return 100
	case 3:
		return time.Now().String()
	case 4:
		return float32(100.20)
	case 5:
		return true
	case 6:
		type order struct {
			ID string `json:"id"`
		}
		type result struct {
			Name   string   `json:"name"`
			Age    int      `json:"age"`
			Orders []*order `json:"order"`
		}
		return &result{Name: "colin", Age: 8, Orders: []*order{&order{ID: "897776666"}}}
	case 7:
		return `{"name":"colin","age":8}`
	case 8:
		return map[string]string{
			"order": "123456",
		}
	case 9:
		return map[string]interface{}{
			"product": map[string]string{
				"price": "100",
			},
		}
	case 10:
		return `<?xml><name>colin</name><age>8</age></xml>`
	default:
		return fmt.Errorf("值错误，请传入1-10")
	}
}
