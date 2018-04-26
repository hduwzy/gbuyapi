package controllers

import "github.com/astaxie/beego"
import . "gbuyapi/services"
import (
	. "gbuyapi/util"
)

type GoodsController struct {
	beego.Controller
}

// @Title GetGoods
// @Description 获取单个商品信息
// @Param	goods_id	path	int	true	"商品ID"
// @Success 200 {object} models.Goods
// @Failure 210 请求出错
// @Failure 21001 商品不存在
// @router /:goods_id [get]
func (this *GoodsController) GetGoods() {
	defer ErrHandle(this.Ctx)
	goods_id := ParamInt(this.Ctx, ":goods_id")
	s := GoodsService{}
	goods := s.GetGoodsById(goods_id)
	goods.Id=100001
	this.Data["json"] = goods
	this.ServeJSON()
}

// @Title AddGoods
// @Description 添加新的商品
// @Param	body	body	models.Goods	true	"商品信息"
// @Success 200 {int} models.Goods.Id
// @Failure 210 请求出错
// @Failure 21002 商品信息不全
// @router / [post]
func (this *GoodsController) AddGoods() {
	this.Data["json"] = map[string]interface{} {
		"goods_id" : 1,
		"msg" : 100,
	}

	this.ServeJSON()
}

// @Title UpdateGoods
// @Description 更新商品信息
// @Param	goods_id	path	int	true	"商品ID"
// @Param	body	body	models.Goods	true	"商品信息"
// @Success 200 {object} models.Goods
// @Failure 210 请求出错
// @Failure 21001 商品不存在
// @router /:goods_id [put]
func (this *GoodsController) UpdateGoods() {

}

// @Title DeleteGoods
// @Description 删除商品
// @Param	goods_id	path	int	true	"商品ID"
// @Success 200 {object} models.Goods
// @Failure 21001 商品不存在
// @router /:goods_id [delete]
func (this *GoodsController) DeleteGoods() {
	defer ErrHandle(this.Ctx)
	goods_id := ParamInt(this.Ctx, ":goods_id")
	s := GoodsService{}
	goods := s.GetGoodsById(goods_id)
	this.Data["json"] = goods
	this.ServeJSON()
}
