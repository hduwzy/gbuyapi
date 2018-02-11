package services

import (
	. "gbuyapi/models"
	"gbuyapi/util"
)

type GoodsService struct {
	BaseService
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