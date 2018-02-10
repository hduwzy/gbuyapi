package models

import "github.com/astaxie/beego/orm"

func init() {
	models := []interface{}{
		// 商品
		&Goods{},&Sku{},
		// 订单
		// 团购
		// 用户
	}
	orm.RegisterModel(models...)
}