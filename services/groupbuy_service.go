package services

import (
	"gbuyapi/models"
	"gbuyapi/util"
)

const (
	RECEIVE_TYPE_POST = 1 // 快递
	RECEIVE_TYPE_SEND = 2 // 送货上门
	RECEIVE_TYPE_STOP = 3 // 站点自取
)

type GroupbuyService struct {
	BaseService
}

func (s *GroupbuyService) AddGroupbuy(gbuy models.Groupbuy, gbuyDetail []models.GroupbuyDetail) (gid int) {
	id, err := s.Orm().Insert(&gbuy)
	if err != nil {
		panic(util.NewError(util.SYS_ERR, err.Error()))
	}

	for idx := range gbuyDetail {
		gbuyDetail[idx].GroupbuyId = int(id)
	}

	_, err2 := s.Orm().InsertMulti(len(gbuyDetail), &gbuyDetail)
	if err2 != nil {
		panic(util.NewError(util.SYS_ERR, err2.Error()))
	}

	return int(id)
}

func (s *GroupbuyService) UpdateGroupbuy() {

}

func (s *GroupbuyService) GetGroupbuyById(gid int) (gbuy models.Groupbuy, details []models.GroupbuyDetail) {

	err := s.Orm().QueryTable("groupbuy").Filter("id__eq", gid).One(&gbuy)
	if err != nil {
		panic(util.NewError(util.SYS_ERR, err.Error()))
	}

	_, err2 := s.Orm().QueryTable("groupbuy_detail").Filter("groupbuy_id__eq", gid).All(&details)
	if err2 != nil {
		panic(util.NewError(util.SYS_ERR, err2.Error()))
	}

	return gbuy,details
}


func (s *GroupbuyService) GetGroupbuyList(userId int, page, page_size int) (list []models.Groupbuy) {
	if page <= 0 {
		page = 1
	}

	if page_size <= 0 || page_size > 60 {
		page_size = 30
	}


	offset := (page - 1) * page_size
	_, err := s.Orm().QueryTable("groupbuy").Offset(offset).Limit(page_size).OrderBy("-id").All(&list)
	if err != nil {
		panic(util.NewError(util.SYS_ERR, err.Error()))
	}

	return list
}

func (s *GroupbuyService) GetGroupbuyListByIds(ids []int, page, page_size int) (list []models.Groupbuy) {

	return
}


func (s *GroupbuyService) DeleteGroupbuy(userId, gid int) {
	num, err := s.Orm().QueryTable("groupbuy").Filter("id__eq", gid).Filter("user_id__eq", userId).Delete()
	if num != 1 || err != nil {
		return
	}
	s.Orm().QueryTable("groupbuy_detail").Filter("groupbuy_id__eq", gid).Delete()
}

func (s *GroupbuyService) JoinGroupbuy(user_id int, gid int) {

}