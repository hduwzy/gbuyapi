package models

import "github.com/astaxie/beego/orm"

func init() {
	models := []interface{}{
		// 商品
		&Goods{},&Sku{},&SkuImage{},
		// 订单
		&Orders{}, &SubOrders{},
		// 团购
		&Groupbuy{}, &GroupbuyDetail{}, &GroupbuyPartake{},
		// 用户
		&UserAddress{}, &Users{},
	}
	orm.RegisterModel(models...)
}