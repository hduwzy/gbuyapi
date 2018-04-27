package services

import (
	"github.com/astaxie/beego/orm"
	"gbuyapi/models"
)


type BaseService struct {

}

func (s *BaseService) Orm() orm.Ormer {
	return orm.NewOrm()
}

func (s *BaseService) GetUserInfo() models.UserInfo {
	return models.UserInfo{UserId:2001}
}