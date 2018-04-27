package main

import (
	_ "gbuyapi/routers"

	"github.com/astaxie/beego"
	"gbuyapi/util"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	util.InitDb()
	beego.Run("0.0.0.0:8899")
}
