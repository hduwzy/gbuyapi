package controllers

import "github.com/astaxie/beego"
import . "gbuyapi/services"
import (
	. "gbuyapi/util"
	"gbuyapi/models/params"
	"gbuyapi/models"
	"time"
	"gbuyapi/models/results"
)

type GoodsController struct {
	beego.Controller
}

// @Title GetGoods
// @Description 获取单个商品信息
// @Param	goods_id	path	int	true	"商品ID"
// @Success 200 {object} results.GetGoodsRsponse	"商品信息"
// @Failure 210 请求出错
// @Failure 21001 商品不存在
// @router /:goods_id [get]
func (this *GoodsController) GetGoods() {
	defer ErrHandle(this.Ctx)
	goods_id := ParamInt(this.Ctx, ":goods_id")
	s := GoodsService{}
	goods := s.GetGoodsById(goods_id)
	sku := s.GetSkusBuGoodsId(goods_id)
	img := s.GetSkuimgByGoodsId(goods_id)

	var skus []results.GetGoodsSkuRsponse

	for _, s := range sku {
		ss := results.GetGoodsSkuRsponse{}
		ss.Sku = s
		for _, i := range img {
			if i.SkuId == s.Id {
				ss.Img = i.Url
				break
			}
		}
		skus = append(skus, ss)
	}

	resp := results.GetGoodsRsponse{
		Goods:*goods,
		Skus:skus,
	}

	this.Data["json"] = resp
	this.ServeJSON()
}

// @Title GetGoodsList
// @Description 获取商品信息列表
// @Param	page	path	int	true	"页码"
// @Param	page_size	path	int	true	"一页大小"
// @Success 200 {object} []models.Goods	"商品信息列表"
// @Failure 210 请求出错
// @Failure 21001 商品不存在
// @router /page/:page/page_size/:page_size [get]
func (this *GoodsController) GetGoodsList() {
	defer ErrHandle(this.Ctx)

	var resp []models.Goods

	page := ParamInt(this.Ctx, ":page")
	page_size := ParamInt(this.Ctx, ":page_size")

	s := GoodsService{}
	resp = s.GetGoodsList(page, page_size)
	this.Data["json"] = resp
	this.ServeJSON()
}

// @Title AddGoods
// @Description 添加新的商品
// @Param	body	body	params.AddGoodsParams	true	"商品信息"
// @Success 200 {int} models.Goods.Id	"添加成功的商品id"
// @Failure 210 请求出错
// @Failure 21002 商品信息不全
// @router / [post]
func (this *GoodsController) AddGoods() {
	defer ErrHandle(this.Ctx)
	param := params.AddGoodsParams{}
	ParamObject(this.Ctx, &param)
	goods := models.Goods{}
	var skus []models.Sku
	var skuImgs []models.SkuImage
	s := GoodsService{}
	now := time.Now()
	goods.CateId = param.CateId
	goods.Title = param.Title
	goods.Detail = param.Detail
	goods.UserId = 2001
	goods.CreateTime = now
	goods.UpdateTime = now

	for _, skuinfo := range param.Skus {
		sku := models.Sku{}
		skuimg := models.SkuImage{}

		sku.Price = skuinfo.Price
		sku.SkuTitle = skuinfo.SkuTitle
		sku.Unit = skuinfo.Unit
		sku.Stock = skuinfo.Stock
		sku.CreateTime = now
		sku.UpdateTime = now

		skus = append(skus, sku)

		skuimg.IsDelete = 0
		skuimg.Url = skuinfo.ImgUrl
		skuimg.CreateTime = now
		skuimg.UpdateTime = now

		skuImgs = append(skuImgs, skuimg)
	}
	goods_id := s.AddGoods(goods, skus, skuImgs)

	this.Data["json"] = goods_id
	this.ServeJSON()
}

// @Title UpdateGoods
// @Description 更新商品信息
// @Param	goods_id	path	int	true	"商品ID"
// @Param	body	body	params.UpdateGoodsParams	true	"商品信息"
// @Success 200 {object} results.GetGoodsRsponse	"更新后的商品信息"
// @Failure 210 请求出错
// @Failure 21001 商品不存在
// @router /:goods_id [put]
func (this *GoodsController) UpdateGoods() {
	defer ErrHandle(this.Ctx)
}

// @Title DeleteGoods
// @Description 删除商品
// @Param	goods_id	path	int	true	"商品ID"
// @Success 200 {object} models.Goods	"被删除的商品信息"
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
