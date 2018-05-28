package models

import "time"

type UserAddress struct {
	Id           int
	ReceiveName  string
	ReceivePhone string
	Province     string
	City         string
	Area         string
	Street       string
	Detail       string
	CreateTime   time.Time
	UpdateTime   time.Time
}

func (this *UserAddress) TableName() string {
	return "user_address"
}


type Users struct {
	Id int	`json:"id"`
	OpenId string	`json:"open_id"`
	Nickname string	`json:"nick_name"`
	AvaUrl string	`json:"ava_url"`
	AddTime time.Time	`json:"add_time"`
	UpdateTime time.Time	`json:"update_time"`
}


func (this *Users) TableName() string {
	return "users"
}