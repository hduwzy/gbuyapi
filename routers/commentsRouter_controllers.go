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
			Method: "GetGoods",
			Router: `/:goods_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:GoodsController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "UpdateGoods",
			Router: `/:goods_id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:GoodsController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "DeleteGoods",
			Router: `/:goods_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:GoodsController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "GetGoodsList",
			Router: `/page/:page/page_size/:page_size`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:GroupbuyController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:GroupbuyController"],
		beego.ControllerComments{
			Method: "AddGroupbuy",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:GroupbuyController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:GroupbuyController"],
		beego.ControllerComments{
			Method: "UpdateGroupbuy",
			Router: `/:gid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:GroupbuyController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:GroupbuyController"],
		beego.ControllerComments{
			Method: "BindGroupbuyGoods",
			Router: `/:gid/goods`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:GroupbuyController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:GroupbuyController"],
		beego.ControllerComments{
			Method: "DeleteGroupbuy",
			Router: `/:group_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:GroupbuyController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:GroupbuyController"],
		beego.ControllerComments{
			Method: "GetGroupbuy",
			Router: `/:group_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:GroupbuyController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:GroupbuyController"],
		beego.ControllerComments{
			Method: "GetGroupbuyList",
			Router: `/page/:page/page_size/:page_size`,
			AllowHTTPMethods: []string{"get"},
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

	beego.GlobalControllerRouter["gbuyapi/controllers:OrderController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:OrderController"],
		beego.ControllerComments{
			Method: "CreateOrder",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:OrderController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:OrderController"],
		beego.ControllerComments{
			Method: "GetOrderByOrderId",
			Router: `/:order_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:OrderController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:OrderController"],
		beego.ControllerComments{
			Method: "CancelOrder",
			Router: `/:order_id/status/:status`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:OrderController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:OrderController"],
		beego.ControllerComments{
			Method: "GetOrderList",
			Router: `/uid/:user_id/page/:page/psize/:page_size`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:UserAddressController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:UserAddressController"],
		beego.ControllerComments{
			Method: "AddUserAddress",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:UserAddressController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:UserAddressController"],
		beego.ControllerComments{
			Method: "DeleteUserAddress",
			Router: `/:addr_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:UserAddressController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:UserAddressController"],
		beego.ControllerComments{
			Method: "GetUserAddress",
			Router: `/:addr_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gbuyapi/controllers:UserAddressController"] = append(beego.GlobalControllerRouter["gbuyapi/controllers:UserAddressController"],
		beego.ControllerComments{
			Method: "GetUserAddressList",
			Router: `/page/:page/page_size/:page_size`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
