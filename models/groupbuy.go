package models

import "time"

type Groupbuy struct {
	Id          int       `orm:"auto;pk;column(id)" json:"id"`
	Title       string    `orm:";column(title);size(128)" json:"title"`
	Detail      string    `orm:";column(detail);type(text)" json:"detail"`
	ImgUrl      string    `orm:";column(img_url);size(128)" json:"img_url"`
	UserId      int       `orm:";column(user_id)" json:"user_id"`
	ReceiveType int8      `orm:";column(receive_type)" json:"receive_type"`
	IsVisible   int8      `orm:";column(is_visible)" json:"is_visible"`
	StartTime   time.Time `orm:";column(start_time);" json:"start_time"`
	EndTime     time.Time `orm:";column(end_time);" json:"end_time"`
	CreateTime  time.Time `orm:";auto_now_add;column(create_time);" json:"create_time"`
	UpdateTime  time.Time `orm:";auto_now;column(update_time);" json:"update_time"`
}

func (this *Groupbuy) TableName() string {
	return "groupbuy"
}

type GroupbuyDetail struct {
	Id         int       `orm:"auto;pk;column(id)" json:"id"`
	GroupbuyId int       `orm:";column(groupbuy_id)" json:"groupbuy_id"`
	GoodsId    int       `orm:";column(goods_id)" json:"goods_id"`
	SkuId      int       `orm:";column(sku_id)" json:"sku_id"`
	Price      int       `orm:";column(price)" json:"price"`
	Amount     int       `orm:";column(amount)" json:"amount"`
	CreateTime time.Time `orm:";auto_now_add;column(create_time);" json:"create_time"`
	UpdateTime time.Time `orm:";auto_now;column(update_time);" json:"update_time"`
}

func (this *GroupbuyDetail) TableName() string {
	return "groupbuy_detail"
}

type GroupbuyPartake struct {
	Id         int       `orm:"auto;pk;column(id)" json:"id"`
	GroupbuyId int       `orm:";column(groupbuy_id)" json:"groupbuy_id"`
	CreateTime time.Time `orm:";auto_now_add;column(create_time);" json:"create_time"`
}

func (this *GroupbuyPartake) TableName() string {
	return "groupbuy_partake"
}
