package models

import "time"

type Goods struct {
	Id         int       `orm:"auto;pk;column(id)" json:"id"`
	UserId     int       `orm:";column(user_id)" json:"user_id"`
	Title      string    `orm:";column(title);size(256)" json:"title"`
	CateId     int       `orm:";column(cate_id);default(0)" json:"cate_id"`
	Detail     string    `orm:";column(detail);default('');type(text)" json:"detail"`
	CreateTime time.Time `orm:";auto_now_add;column(create_time);default(current_timestamp())" json:"create_time"`
	UpdateTime time.Time `orm:";auto_now;column(update_time);default(current_timestamp())" json:"update_time"`
}

func (this *Goods) TableName() string {
	return "goods"
}

type Sku struct {
	Id         int       `orm:"auto;pk;column(id)" json:"id"`
	GoodsId    int       `orm:";column(goods_id)" json:"goods_id"`
	SkuTitle   string    `orm:";column(sku_title);size(256)" json:"sku_title"`
	Price      int       `orm:";column(price);default(0)" json:"price"`
	Unit       string    `orm:";column(unit);size(32)" json:"unit"`
	Stock      int       `orm:";column(stock);default(0)" json:"stock"`
	CreateTime time.Time `orm:";auto_now_add;column(create_time);default(current_timestamp())" json:"create_time"`
	UpdateTime time.Time `orm:";auto_now;column(update_time);default(current_timestamp())" json:"update_time"`
}

func (this *Sku) TableName() string {
	return "sku"
}

type SkuImage struct {
	Id         int       `orm:"auto;pk;column(id)" json:"id"`
	GoodsId    int       `orm:";column(goods_id)" json:"goods_id"`
	SkuId      int       `orm:";column(sku_id)" json:"sku_id"`
	Url        string    `orm:";column(url);size(128)" json:"url"`
	IsDelete   int8      `orm:";column(is_delete);default(0)" json:"is_delete"`
	CreateTime time.Time `orm:";auto_now_add;column(create_time);default(current_timestamp())" json:"create_time"`
	UpdateTime time.Time `orm:";auto_now;column(update_time);default(current_timestamp())" json:"update_time"`
}

func (this *SkuImage) TableName() string {
	return "sku_image"
}
