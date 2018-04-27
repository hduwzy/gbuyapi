package controllers

import "github.com/astaxie/beego"

type UserAddressController struct {
	beego.Controller
}


// @Title AddUserAddress
// @Description 用户添加收货地址
// @Param	body	body	models.UserAddress	true	"收货地址"
// @Success 200 {int} models.UserAddress.Id	"添加成功返回用户地址id"
// @Failure 210 请求出错
// @router / [post]
func (this *UserAddressController) AddUserAddress() {

}

// @Title DeleteUserAddress
// @Description 删除收货地址
// @Param	addr_id	path	int	true	"收货地址Id"
// @Success 200 {int} models.UserAddress	"被删除的用户地址"
// @Failure 210 请求出错
// @router /:addr_id [delete]
func (this *UserAddressController) DeleteUserAddress() {

}


// @Title GetUserAddressList
// @Description 获取用户收货地址L列表
// @Param	page	path	int	true	"收货地址列表，当前页数"
// @Param	page_size	path	int	true	"收货地址列表，当前页展示个数"
// @Success 200 {int} results.GetUserAddressListResult	"用户地址列表"
// @Failure 210 请求出错
// @router /page/:page/page_size/:page_size [get]
func (this *UserAddressController) GetUserAddressList() {

}

// @Title GetUserAddress
// @Description 获取用户收货地址
// @Param	addr_id	path	int	true	"收货地址id"
// @Success 200 {int} models.UserAddress	"获取的单个用户地址"
// @Failure 210 请求出错
// @router /:addr_id [get]
func (this *UserAddressController) GetUserAddress() {

}