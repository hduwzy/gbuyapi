package controllers

import (
	"github.com/astaxie/beego"
	"gbuyapi/util"
	"gbuyapi/models/params"
	"gbuyapi/services"
	"gbuyapi/models"
	"gbuyapi/models/results"
	"encoding/json"
)

type UserController struct {
	beego.Controller
}

// @Title GetUserinfo
// @Description 获取单个商品信息
// @Param	code	body	params.UserLoginRequest	true	"用户登录信息"
// @Success 200 {string}	登录结果
// @Failure 210 请求出错
// @router /user/login/:code [get]
func (this *UserController) Login() {
	defer util.ErrHandle(this.Ctx)

	if nil == services.Sess(this.Ctx) {
		panic(util.NewError(210, "用户已登录"))
	}

	loginfo := params.UserLoginRequest{}
	util.ParamObject(this.Ctx, &loginfo)

	s := services.NewWechatService()
	sid := s.LoginSubject(loginfo.Code, loginfo.AvatUrl, loginfo.Nickname)
	res := results.LoginResult{
		IsSuccess: true,
		Sid: sid,
	}

	this.Data["json"] = res
	this.ServeJSON()
}

// @Title GetUserinfo
// @Description 获取当前用户信息
// @Success 200 {object} models.Users	"当前用户信息"
// @Failure 210 请求出错
// @router /:goods_id [get]
func (this *UserController) GetUserinfo() {
	defer util.ErrHandle(this.Ctx)

	sess := services.Sess(this.Ctx)
	uinfo, err := sess.GetString("user_info")
	if err != nil {
		panic(util.NewError(210, err.Error()))
	}
	u := models.Users{}
	json.Unmarshal([]byte(uinfo), &u)

	this.Data["json"] = u
	this.ServeJSON()
}
