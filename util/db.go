package util

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


//dbhost="testonly.fun"
//username="root"
//password="123456"
//port=3306
func InitDb() {
	dbhost := beego.AppConfig.String("dbhost")
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	url := username + ":" + password + "@tcp(" + dbhost + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", url, 3, 10)
}