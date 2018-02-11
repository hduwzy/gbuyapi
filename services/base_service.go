package services

import "github.com/astaxie/beego/orm"


type BaseService struct {

}

func (s *BaseService) Orm() orm.Ormer {
	return orm.NewOrm()
}