package models

import (
	"github.com/astaxie/beego/orm"
	"math"
)

type Article struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	LikeNum     int    `json:"like_num"`
	ShareNum    int    `json:"share_num"`
	EvaluateNum int    `json:"evaluate_num"`
	UserId      int64  `json:"user_id"`
	CT          int64  `json:"c_t"`
	UT          int64  `json:"u_t"`
	Type        int    `json:"type"`
	Status      int    `json:"status"`
	Avatar      string `json:"avatar"`
	NickName    string `json:"nick_name"`
}

type ListData struct {
	List      *[]Article `json:"list"`
	Count     int64      `json:"count"`
	PageCount float64    `json:"page_count"`
	PageNum   int        `json:"page_num"`
}

func (t *Article) TableName() string {
	return TableName("article")
}

func ArticleAdd(t *Article) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(t)

	return id, err
}

func ArticleList(pageNum int, pageSetNum int) (list ListData, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("article").Filter("status", 1)
	count, _ := qs.Count()
	pageSetNum = 10
	pageCount := math.Ceil((float64(count) / float64(pageSetNum)))
	//存储分页数据的切片
	articles := new([]Article)
	//获取分页数据
	_, _ = qs.OrderBy("-id").Limit(pageSetNum, pageSetNum*(pageNum-1)).All(articles)

	return ListData{
		List:      articles,
		Count:     count,
		PageCount: pageCount,
		PageNum:   pageNum,
	}, nil
}

func ArticleSingleUser(pageNum int, userId int64) (list ListData, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("article").Filter("status", 1).Filter("user_id", userId)
	count, _ := qs.Count()
	pageSetNum := 10
	pageCount := math.Ceil((float64(count) / float64(pageSetNum)))
	//存储分页数据的切片
	articles := new([]Article)
	//获取分页数据
	_, _ = qs.OrderBy("-id").Limit(pageSetNum, pageSetNum*(pageNum-1)).All(articles)

	return ListData{
		List:      articles,
		Count:     count,
		PageCount: pageCount,
		PageNum:   pageNum,
	}, nil
}

func ArticleDetail(id int) (article Article, err error) {
	o := orm.NewOrm()
	article = Article{Id: id}
	err = o.Read(&article, "id")

	return article, nil
}

func ArticleSearch(text string) (article *[]Article, err error) {
	articles := new([]Article)

	o := orm.NewOrm()
	qs := o.QueryTable("article").Filter("status", 1).Filter("title__contains", text)
	//获取分页数据
	_, _ = qs.OrderBy("-id").Limit(20).All(articles)

	return articles, nil

}

func ArticleEvaluateIncr(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("article").Filter("id", id).Update(orm.Params{
		"evaluate_num": orm.ColValue(orm.ColAdd, 1),
	})

	return num, err
}

func ArticleLikeIncrNum(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("article").Filter("id", id).Update(orm.Params{
		"like_num": orm.ColValue(orm.ColAdd, 1),
	})

	return num, err
}
