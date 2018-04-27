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


type UserInfo struct {
	UserId int
}