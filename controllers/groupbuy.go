package controllers

import "github.com/astaxie/beego"

type GroupbuyController struct {
	beego.Controller
}


// @Title
// @Description
// @Param
// @Success
// @Failure
// @router / [post]
func (this *GroupbuyController) AddGroupbuy() {

}


// @Title
// @Description
// @Param
// @Success
// @Failure
// @router / [post]
func (this *GroupbuyController) DeleteGroupbuy() {

}


// @Title
// @Description
// @Param	group_id	path	int	true	"团购id"
// @Success 200 {object} models.Groupbuy
// @Failure
// @router /:group_id [post]
func (this *GroupbuyController) GetGroupbuy() {

}


// @Title
// @Description
// @Param	group_id	path	int	true	"团购id"
// @Param	body	body	models.Groupbuy	true	"团购信息"
// @Success 200 {object} models.Groupbuy
// @Failure 210 请求失败
// @router /:group_id [post]
func (this *GroupbuyController) UpdateGroupbuy() {

}