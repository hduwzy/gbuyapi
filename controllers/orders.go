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


// @Title GetOrderByOrderId
// @Description 获取单个订单详情
// @Param	order_id	path	string	true	"订单ID"
// @Success 200 {object} models.Orders
// @Failure 210 请求出错
// @router /:order_id [get]
func (this *OrderController) GetOrderByOrderId() {

}


// @Title GetOrderList
// @Description 获取订单列表
// @Param	page	query	int	true	"当前页数"
// @Param	page_size	query	int	true	"每页展示条数"
// @Param	user_id	query	int	true	"用户ID"
// @Success 200 {object} results.GetOrderListResult
// @Failure 210 请求出错
// @router / [get]
func (this *OrderController) GetOrderList() {

}