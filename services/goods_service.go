package services

import . "gbuyapi/models"

type GoodsService struct {
	BaseService
}

func (s *GoodsService) GetGoodsById(goods_id int) (goods Goods, code int) {
	qs := s.Orm().QueryTable("goods")

	err := qs.Filter("goods_id__eq", goods_id).One(&goods)
	if err != nil {
		return goods, 210
	}

	return goods, 200
}