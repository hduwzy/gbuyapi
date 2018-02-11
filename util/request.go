package util

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"strconv"
)

func ParamInt(ctx *context.Context, key string) int {
	v_str := ctx.Input.Param(key)
	v, err := strconv.Atoi(v_str)
	if err != nil {
		panic(NewError(-1, "参数解析错误:" + key))
	}
	return v
}

func ParamString(ctx *context.Context, key string) string {
	return ctx.Input.Param(key)
}

func ParamObject(ctx *context.Context, val interface{}) {
	err := json.Unmarshal(ctx.Input.RequestBody, val)
	if err != nil {
		panic(NewError(-1, "参数解析错误:"+string(ctx.Input.RequestBody)))
	}
}
