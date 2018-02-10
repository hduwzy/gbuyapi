package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["gbuyapi/controllers:GoodsController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "AddGoods",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:GoodsController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "UpdateGoods",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:GoodsController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "GetGoods",
			Router: `/:goods_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:GoodsController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "DeleteGoods",
			Router: `/:goods_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
