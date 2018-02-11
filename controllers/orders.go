package controllers

import "github.com/astaxie/beego"

type OrderController struct {
	beego.Controller
}


// @Title CreateOrder
// @Description 用户下单，创建订单
// @Param	body	body	models.Orders	true	"订单信息"
// @Success 200 {int} models.Orders.Id
// @Failure 210 请求出错
// @router / [post]
func (this *OrderController) CreateOrder() {

}