package userModules

import (
	"abana/controllers"
	"abana/models"
)

type UserController struct {
	controllers.BaseController
}

type UserUpdateReq struct {
	NickName  string `json:"nick_name"`
	Gender    int    `json:"gender"`
	AvatarUrl string `json:"avatar_url"`
	City      string `json:"city"`
	Province  string `json:"province"`
}

type nearList struct {
	Page int `json:"page"`
}

func (c *UserController) Update() {
	var updateData models.User
	req := &UserUpdateReq{}
	c.GetParams(req)

	userId := c.Ctx.Input.GetData("user_id").(int64)

	updateData = models.User{
		AvatarUrl: req.AvatarUrl,
		City:      req.City,
		Gender:    req.Gender,
		NickName:  req.NickName,
		Province:  req.Province,
		UserId:    userId,
	}

	updateRes := models.UserUpdate(userId, updateData)
	if updateRes > 0 {
		c.RenderJson(updateRes)
	} else {
		c.RenderJsonErr("", "保存失败")
	}
}

func (c *UserController) NearList() {
	req := &nearList{}
	c.GetParams(req)
	list, _ := models.UserNearList(req.Page)

	c.RenderJson(list)
}
