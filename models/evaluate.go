package models

import (
	"github.com/astaxie/beego/orm"
)

type Evaluate struct {
	Id        int64  `json:"id"`
	ArticleId int64  `json:"article_id"`
	UserId    int64  `json:"user_id"`
	Detail    string `json:"detail"`
	CT        int64  `json:"c_t"`
	Status    int8   `json:"status"`
}

func (t *Evaluate) TableName() string {
	return TableName("evaluate")
}

func EvaluateAdd(t *Evaluate) (data int64, err error) {
	o := orm.NewOrm()
	id, err := o.Insert(t)

	return id, err
}

func EvaluateList(pageNum int, pageSetNum int, id int64) (list *[]Evaluate, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("evaluate").Filter("article_id", id).Filter("status", 1)
	pageSetNum = 10
	evaluates := new([]Evaluate)
	_, _ = qs.OrderBy("-id").Limit(pageSetNum, pageSetNum*(pageNum-1)).All(evaluates)

	return evaluates, nil
}
