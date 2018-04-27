package services

import (
	. "gbuyapi/models"
	"gbuyapi/util"
	"fmt"
)

type GoodsService struct {
	BaseService
}

func (s *GoodsService) GetGoodsList(page, page_size int) (goods []Goods) {
	if page <= 0 {
		page = 1
	}

	if page_size <= 0 || page_size > 60 {
		page_size = 30
	}

	offset := (page - 1) * page_size
	_, err := s.Orm().QueryTable("goods").Offset(offset).Limit(page_size).OrderBy("-create_time").All(&goods)
	if err != nil {
		panic(util.NewError(util.SYS_ERR, err.Error()))
	}

	return goods
}

func (s *GoodsService) GetGoodsById(goods_id int) (goods *Goods) {
	qs := s.Orm().QueryTable("goods")
	goods = new(Goods)
	errors := qs.Filter("id__eq", goods_id).One(goods)
	if errors != nil {
		panic(util.NewError(util.GOODS_NOT_EXISTS))
	}

	return goods
}

func (s *GoodsService) GetSkusBuGoodsId(goods_id int) (skus []Sku) {
	o := s.Orm().QueryTable("sku")
	o.Filter("goods_id__eq", goods_id).All(&skus)

	return skus
}

func (s *GoodsService) GetSkuimgByGoodsId(goods_id int) (imgs []SkuImage) {
	o := s.Orm().QueryTable("sku_image")
	o.Filter("goods_id__eq", goods_id).All(&imgs)
	return imgs
}

func (s *GoodsService) AddGoods(goods Goods, skus []Sku, imgs []SkuImage) (goods_id int) {
	id, err := s.Orm().Insert(&goods)
	fmt.Println(goods.Title)
	if err != nil {
		panic(util.NewError(util.SYS_ERR, err.Error()))
	}

	for idx := range skus {
		skus[idx].GoodsId = int(id)
	}

	_, err2 := s.Orm().InsertMulti(len(skus), &skus)
	if err2 != nil {
		panic(util.NewError(util.SYS_ERR, err.Error()))
	}

	for idx := range imgs {
		imgs[idx].GoodsId = int(id)
		imgs[idx].SkuId = skus[idx].Id
	}

	_, err3 := s.Orm().InsertMulti(len(imgs), &imgs)
	if err3 != nil {
		panic(util.NewError(util.SYS_ERR, err.Error()))
	}

	return int(id)
}

func (s *GoodsService) UpdateGoods(goods Goods, skus []Sku, imgs []SkuImage) {

}

func (s *GoodsService) DeleteGoods(goods_id int) {

	s.Orm().QueryTable("goods").Filter("goods_id__eq", goods_id).Delete()
	s.Orm().QueryTable("sku").Filter("goods_id__eq", goods_id).Delete()
	s.Orm().QueryTable("sku_image").Filter("goods_id__eq", goods_id).Delete()

}
