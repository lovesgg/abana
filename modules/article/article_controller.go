package articleModules

import (
	"abana/controllers"
	"abana/models"
	"fmt"
	"time"
)

type ArticleController struct {
	controllers.BaseController
}

type addReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type userListReq struct {
	UserId int64 `json:"user_id"`
	Page   int   `json:"page"`
}

type searchReq struct {
	Text string `json:"text"`
}

func (c *ArticleController) Publish() {
	var article *models.Article
	req := &addReq{}
	c.GetParams(req)

	if req.Content == "" || req.Title == "" {
		c.RenderJsonErr("参数", "不能为空")
	}

	userId := c.Ctx.Input.GetData("user_id").(int64)

	article = &models.Article{
		Title:       req.Title,
		Content:     req.Content,
		LikeNum:     0,
		ShareNum:    0,
		EvaluateNum: 0,
		UserId:      userId,
		CT:          time.Now().Unix(),
		UT:          time.Now().Unix(),
		Type:        0,
		Status:      1, //默认1 展示
		Avatar:      c.Ctx.Input.GetData("avatar").(string),
		NickName:    c.Ctx.Input.GetData("nick_name").(string),
	}

	ret, err := models.ArticleAdd(article)
	if err != nil {
		fmt.Println(err)
		c.RenderJsonErr("", "添加有误")
	}

	c.RenderJson(ret)
}

func (c *ArticleController) UserList() {
	req := &userListReq{}
	c.GetParams(req)
	list, _ := models.ArticleSingleUser(req.Page, req.UserId)

	c.RenderJson(list)
}

func (c *ArticleController) Search() {
	req := &searchReq{}
	c.GetParams(req)
	list, _ := models.ArticleSearch(req.Text)

	c.RenderJson(list)
}
