package models

import "time"

type UserAddress struct {
	Id           int       `orm:"auto;pk;column(id)" json:"id"` //自增id
	ReceiveName  string    `orm:";column(receive_name);size(128)" json:"receive_name"`
	ReceivePhone string    `orm:";column(receive_phone);size(32)" json:"receive_phone"`
	Province     string    `orm:";column(province);size(32)" json:"province"`
	City         string    `orm:";column(city);size(32)" json:"city"`
	Area         string    `orm:";column(area);size(32)" json:"area"`
	Street       string    `orm:";column(street);size(128)" json:"street"`
	Detail       string    `orm:";column(detail);size(256)" json:"detail"`
	CreateTime   time.Time `orm:";auto_now_add;column(create_time);default(current_timestamp())" json:"create_time"`
	UpdateTime   time.Time `orm:";auto_now;column(update_time);default(current_timestamp())" json:"update_time"`
}

func (this *UserAddress) TableName() string {
	return "user_address"
}
