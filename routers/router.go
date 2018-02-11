// @APIVersion 1.0
// @Title 团购接龙-接口文档
// @Description 团购接龙-接口文档
// @Contact hduwzy@163.com chenhongluan@sina.com
// @TermsOfServiceUrl http://testonly.fun/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gbuyapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/user-addr",
			beego.NSInclude(
				&controllers.UserAddressController{},
			),
		),
		beego.NSNamespace("/order",
			beego.NSInclude(
				&controllers.OrderController{},
			),
		),
		beego.NSNamespace("/goods",
			beego.NSInclude(
				&controllers.GoodsController{},
			),
		),
		beego.NSNamespace("/groupbuy",
			beego.NSInclude(
				&controllers.GroupbuyController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
