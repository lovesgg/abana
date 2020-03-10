package articleModules

import (
	"abana/controllers"
	"abana/models"
	"time"
)

type EvaluateController struct {
	controllers.BaseController
}

type addEvReq struct {
	Detail string `json:"detail"`
	Id     int64  `json:"id"`
}

func (c *EvaluateController) Add() {
	var evaluate *models.Evaluate
	req := &addEvReq{}
	c.GetParams(req)

	userId := c.Ctx.Input.GetData("user_id").(int64)

	evaluate = &models.Evaluate{
		ArticleId: req.Id,
		UserId:    userId,
		Detail:    req.Detail,
		CT:        time.Now().Unix(),
		Status:    1,
	}

	ret, err := models.EvaluateAdd(evaluate)
	if err != nil {
		c.RenderJsonErr("", "添加有误")
	}
	//文章增肌评论数量
	_, _ = models.ArticleEvaluateIncr(int(req.Id))

	c.RenderJson(ret)
}
