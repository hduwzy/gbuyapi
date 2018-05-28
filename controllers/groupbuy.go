package controllers

import (
	"github.com/astaxie/beego"
	"gbuyapi/util"
	"gbuyapi/models/params"
	"gbuyapi/models"
	"time"
	"gbuyapi/services"
	"gbuyapi/models/results"
)

type GroupbuyController struct {
	beego.Controller
}

// @Title AddGroupbuy
// @Description 添加团购
// @Param	body	body	params.AddGroupbuyParam	true	"团购信息"
// @Success 200	{int} models.Groupbuy.Id "创建成功的团购id"
// @Failure 210 请求错误
// @router / [post]
func (c *GroupbuyController) AddGroupbuy() {
	var param params.AddGroupbuyParam
	util.ParamObject(c.Ctx, &param)
	s := services.GroupbuyService{}
	format := "2006-01-02 15:04:05"
	ts, err := time.Parse(format, param.Gbuy.StartTime)
	te, err2 := time.Parse(format, param.Gbuy.EndTime)

	if err != nil || err2 != nil {
		panic(util.NewError(util.SYS_ERR, err.Error()+err2.Error()))
	}
	now := time.Now()
	gbuy := models.Groupbuy{
		UserId:s.GetUserInfo().Id,
		Title:       param.Gbuy.Title,
		ImgUrl:      param.Gbuy.ImgUrl,
		Detail:      param.Gbuy.Detail,
		ReceiveType: param.Gbuy.ReceiveType,
		StartTime:   ts,
		EndTime:     te,
		CreateTime:  now,
		UpdateTime:  now,
	}

	var details []models.GroupbuyDetail

	for _, det := range param.Details {
		i := models.GroupbuyDetail{}
		i.CreateTime = now
		i.UpdateTime = now
		i.GoodsId = det.GoodsId
		i.SkuId = det.SkuId
		i.Price = det.Price
		i.Amount = det.Amount
		details = append(details, i)
	}


	gid := s.AddGroupbuy(gbuy, details)
	c.Data["json"] = gid
	c.ServeJSON()
}

// @Title DeleteGroupbuy
// @Description 删除团购
// @Param	group_id	path	int	true	"团购"
// @Success 200 {int}	id	"删除的团购id"
// @Failure 210 请求错误
// @router /:group_id [delete]
func (c *GroupbuyController) DeleteGroupbuy() {
	gid := util.ParamInt(c.Ctx, ":group_id")
	s := services.GroupbuyService{}
	s.Orm()
	s.DeleteGroupbuy(s.GetUserInfo().Id, gid)
	c.Data["json"] = gid
	c.ServeJSON()
}

// @Title GetGroupbuy
// @Description 根据团购id
// @Param	page	path	int	true	"列表页码"
// @Param	page_size	path	int	true	"每页数量"
// @Success 200 {object} results.GetGroupbuyListRespongse	"我发布的团购列表"
// @Failure 210 请求失败
// @router /page/:page/page_size/:page_size [get]
func (c *GroupbuyController) GetGroupbuyList() {
	page := util.ParamInt(c.Ctx, ":page")
	page_size := util.ParamInt(c.Ctx, ":page_size")

	s := services.GroupbuyService{}
	userInfo := s.GetUserInfo()
	list := s.GetGroupbuyList(userInfo.Id, page, page_size)
	c.Data["json"] = list
	c.ServeJSON()
}

// @Title GetGroupbuy
// @Description 根据团购id
// @Param	group_id	path	int	true	"团购id"
// @Success 200 {object} results.GetGroupbuyResponse	"团购信息"
// @Failure 210 请求失败
// @router /:group_id [get]
func (c *GroupbuyController) GetGroupbuy() {
	gid := util.ParamInt(c.Ctx, ":group_id")
	s := services.GroupbuyService{}
	gbuy, details := s.GetGroupbuyById(gid)
	c.Data["json"] = results.GetGroupbuyResponse{
		Gbuy:gbuy,
		GbuyDetail:details,
	}
	c.ServeJSON()
}

// @Title UpdateGroupbuy
// @Description 更新团购信息
// @Param	gid	path	int	true	"团购id"
// @Param	body	body	models.Groupbuy	true	"团购信息"
// @Success 200 {object} models.Groupbuy	"团购信息"
// @Failure 210 {object} util.ApiError 请求失败
// @router /:gid [put]
func (c *GroupbuyController) UpdateGroupbuy() {

}

// @Title BindGroupbuyGoods
// @Description 绑定团购商品
// @Param	gid	path	int	true	"团购id"
// @Param	body	body	params.BindGroupbuyGoodsRequest	true	"绑定信息"
// @Success 200	{int} id	"绑定id"
// @Failure 210 请求失败
// @router /:gid/goods [post]
func (c *GroupbuyController) BindGroupbuyGoods() {

}
