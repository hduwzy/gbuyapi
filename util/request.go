package util

import (
	"github.com/astaxie/beego"
	"encoding/json"
)

func ParamInt(ctl beego.Controller, key string, val *int) {

}

func ParamString(ctl beego.Controller, key string, val *string) {

}

func ParamObject(ctl beego.Controller, val interface{}) {
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, val)
	if err != nil {
		panic(err)
	}
}
