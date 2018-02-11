package controllers

import "github.com/astaxie/beego"

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
// @router /:gid [post]
func (this *GroupbuyController) GetGroupbuy() {

}


// @Title UpdateGroupbuy
// @Description 更新团购信息
// @Param	group_id	path	int	true	"团购id"
// @Param	body	body	models.Groupbuy	true	"团购信息"
// @Success 200 {object} models.Groupbuy
// @Failure 210 请求失败
// @router / [post]
func (this *GroupbuyController) UpdateGroupbuy() {

}