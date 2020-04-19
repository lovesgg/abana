package articleModules

import (
	"abana/controllers"
	"abana/models"
	"time"
)

type ArticleController struct {
	controllers.BaseController
}

type addReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

type userListReq struct {
	UserId int64 `json:"user_id"`
	Page   int   `json:"page"`
}

type searchReq struct {
	Text string `json:"text"`
}

type addLikeReq struct {
	ArticleId int   `json:"article_id"`
	UserId    int64 `json:"user_id"`
}

func (c *ArticleController) Publish() {
	var article *models.Article
	var activity *models.Activity
	var ret int64
	req := &addReq{}
	c.GetParams(req)

	if req.Content == "" || req.Title == "" || req.Type == "" {
		c.RenderJsonErr("参数", "不能为空")
	}

	userId := c.Ctx.Input.GetData("user_id").(int64)
	user, _ := models.UserGetByUserId(userId)

	if req.Type == "index" {
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
			Avatar:      user.AvatarUrl,
			NickName:    user.NickName,
		}
		_, err := models.ArticleAdd(article)
		if err != nil {
			c.RenderJsonErr("", "添加有误")
		}
	} else {
		activity = &models.Activity{
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
			Avatar:      user.AvatarUrl,
			NickName:    user.NickName,
		}
		_, err := models.ActivityAdd(activity)

		if err != nil {
			c.RenderJsonErr("", "添加有误")
		}
	}

	_, _ = models.UserArticleIncrNum(userId)

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

func (c *ArticleController) AddLike() {
	var add *models.MyLike
	req := &addLikeReq{}
	c.GetParams(req)

	userId := c.Ctx.Input.GetData("user_id").(int64)

	if userId == req.UserId {
		c.RenderJsonErr("", "收藏失败")
	}

	//先校验是否存在
	data, _ := models.MyLikeGetById(userId, req.ArticleId)
	if data.UserId == 0 {
		//可添加
		add = &models.MyLike{
			ArticleId: req.ArticleId,
			UserId:    userId,
			CT:        time.Now().Unix(),
		}
		_, _ = models.MyLikeAdd(add)

		//文章增加喜欢数
		_, _ = models.ArticleLikeIncrNum(req.ArticleId)
		_, _ = models.UserLikeIncrNum(userId)

		c.RenderJson(1)
	} else {
		c.RenderJsonErr("", "已收藏")
	}

}
