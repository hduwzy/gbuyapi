package controllers

import "github.com/astaxie/beego"
import . "gbuyapi/services"
type GoodsController struct {
	beego.Controller
}

// @Title GetGoods
// @Description 获取单个商品信息
// @Param	goods_id	path	int	true	"商品ID"
// @Success 200 {object} models.Goods
// @Failure 210 商品id不能为空
// @Failure 211 商品不存在
// @router /:goods_id [get]
func (this *GoodsController) GetGoods() {

	s := GoodsService{}
	s.GetGoodsById()
	this.ServeJSON()
}

// @Title AddGoods
// @Description 添加新的商品
// @Param	body	body	models.Goods	true	"商品信息"
// @Success 200 {int} models.Goods.Id
// @Failure 210 商品信息不全
// @router / [post]
func (this *GoodsController) AddGoods() {

}

// @Title UpdateGoods
// @Description 更新商品信息
// @Param	body	body	models.Goods	true	"商品信息"
// @Success 200 {object} models.Goods
// @Failure 210 商品信息不全
// @Failure 211 商品不存在
// @router / [put]
func (this *GoodsController) UpdateGoods() {

}


// @Title DeleteGoods
// @Description 删除商品
// @Param	goods_id	path	int	true	"商品ID"
// @Success 200 {object} models.Goods
// @Failure 211 商品不存在
// @router /:goods_id [delete]
func (this *GoodsController) DeleteGoods() {

}