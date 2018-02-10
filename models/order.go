package models

import "time"

type Orders struct {
	Id          int       `orm:"auto;pk;column(id)" json:"id"`
	UserId      int       `orm:";column(user_id)" json:"user_id"`
	OrderNo     string    `orm:";column(order_no);size(64)" json:"order_no"`
	KindNum     int       `orm:";column(kind_num);default(0)" json:"kind_num"`
	TotalPrice  int       `orm:";column(total_price);default(0)" json:"total_price"`
	Status      int8      `orm:";column(status);default(0)" json:"status"`
	DeliverType int8      `orm:";column(deliver_type);default(0)" json:"deliver_type"`
	AddressId   int       `orm:";column(address_id)" json:"address_id"`
	BuyerRemark string    `orm:";column(buyer_remark);default('');size(256)" json:"buyer_remark"`
	CreateTime  time.Time `orm:";auto_now_add;column(create_time);" json:"create_time"`
	UpdateTime  time.Time `orm:";auto_now;column(update_time);" json:"update_time"`
}

func (this *Orders) TableName() string {
	return "orders"
}

type SubOrders struct {
	Id         int       `orm:"auto;pk;column(id)" json:"id"`
	UserId     int       `orm:";column(user_id)" json:"user_id"`
	OrderId    int       `orm:";column(order_id)" json:"order_id"`
	GoodsId    int       `orm:";column(goods_id)" json:"goods_id"`
	SkuId      int       `orm:";column(sku_id)" json:"sku_id"`
	Amount     int       `orm:";column(amount);default(0)" json:"amount"`
	Price      int       `orm:";column(price);default(0)" json:"price"`
	TotalPrice int       `orm:";column(total_price);default(0)" json:"total_price"`
	CreateTime time.Time `orm:";auto_now_add;column(create_time);" json:"create_time"`
	UpdateTime time.Time `orm:";auto_now;column(update_time);" json:"update_time"`
}

func (this *SubOrders) TableName() string {
	return "sub_orders"
}
