package controllers

import (
	"github.com/astaxie/beego"
	"gbuyapi/util"
)

type GroupbuyController struct {
	beego.Controller
}


// @Title AddGroupbuy
// @Description 添加团购
// @Param	body	body	models.Groupbuy	true	"团购信息"
// @Success 200	{int} models.Groupbuy.Id
// @Failure 210 请求错误
// @router / [post]
func (this *GroupbuyController) AddGroupbuy() {

}


// @Title DeleteGroupbuy
// @Description 删除团购
// @Param	gid	path	int	true	"团购"
// @Success 200 {object} models.Groupbuy
// @Failure 210 请求错误
// @router /:gid [delete]
func (this *GroupbuyController) DeleteGroupbuy() {

}


// @Title GetGroupbuy
// @Description 根据团购id
// @Param	group_id	path	int	true	"团购id"
// @Success 200 {object} models.Groupbuy
// @Failure 210 请求失败
// @router /:gid [get]
func (this *GroupbuyController) GetGroupbuy() {

}


// @Title UpdateGroupbuy
// @Description 更新团购信息
// @Param	group_id	path	int	true	"团购id"
// @Param	body	body	models.Groupbuy	true	"团购信息"
// @Success 200 {object} models.Groupbuy
// @Failure 210 请求失败
// @router / [put]
func (this *GroupbuyController) UpdateGroupbuy() {

}


// @Title BindGroupbuyGoods
// @Description 绑定团购商品
// @Param	gid	path	int	true	"团购id"
// @Param	body	body	params.BindGroupbuyGoodsRequest	true	"绑定信息"
// @Success 200
// @Failure 210 请求失败
// @router /:gid/goods [post]
func (this *GroupbuyController) BindGroupbuyGoods() {

}