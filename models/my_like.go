package models

import (
	"github.com/astaxie/beego/orm"
)

type MyLike struct {
	Id        int   `json:"id"`
	ArticleId int   `json:"article_id"`
	UserId    int64 `json:"user_id"`
	CT        int64 `json:"c_t"`
}

func (t *MyLike) TableName() string {
	return TableName("my_like")
}

func MyLikeAdd(t *MyLike) (data int64, err error) {
	o := orm.NewOrm()
	id, err := o.Insert(t)

	return id, err
}

func MyLikeGetById(userId int64, articleId int) (data MyLike, err error) {
	o := orm.NewOrm()
	_ = o.QueryTable("my_like").Filter("user_id", userId).Filter("article_id", articleId).One(&data)
	return data, nil
}

func MyLikeList(pageNum int, pageSetNum int, id int, userId int64) (list *[]MyLike, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("my_like").Filter("user_id", userId).Filter("status", 1)
	pageSetNum = 10
	evaluates := new([]MyLike)
	_, _ = qs.OrderBy("-id").Limit(pageSetNum, pageSetNum*(pageNum-1)).All(evaluates)

	return evaluates, nil
}
