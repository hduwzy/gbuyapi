package controllers

import "github.com/astaxie/beego"

type OrderController struct {
	beego.Controller
}


// @Title CreateOrder
// @Description 用户下单，创建订单
// @Param	body	body	models.Orders	true	"订单信息"
// @Success 200 {int} models.Orders.Id	"新创建的订单id"
// @Failure 210 请求出错
// @router / [post]
func (this *OrderController) CreateOrder() {

}


// @Title GetOrderByOrderId
// @Description 获取单个订单详情
// @Param	order_id	path	string	true	"订单ID"
// @Success 200 {object} models.Orders	"订单信息"
// @Failure 210 请求出错
// @router /:order_id [get]
func (this *OrderController) GetOrderByOrderId() {

}


// @Title GetOrderList
// @Description 获取订单列表
// @Param	page	path	int	true	"当前页数"
// @Param	page_size	path	int	true	"每页展示条数"
// @Param	user_id	path	int	true	"用户ID"
// @Success 200 {object} results.GetOrderListResult	"订单列表"
// @Failure 210 请求出错
// @router /uid/:user_id/page/:page/psize/:page_size [get]
func (this *OrderController) GetOrderList() {

}


// @Title CancelOrder
// @Description 更新订单状态
// @Param	order_id	path	string	true	"订单id"
// @Param	status	path	int	true	"订单状态"
// @Success 200 {object} models.Orders	"取消的订单信息"
// @Failure 210 请求出错
// @router /:order_id/status/:status [put]
func (this *OrderController) CancelOrder() {

}
