package models

import (
	"github.com/astaxie/beego/orm"
	"math"
)

type Activity struct {
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

type ActivityListData struct {
	List      *[]Activity `json:"list"`
	Count     int64      `json:"count"`
	PageCount float64    `json:"page_count"`
	PageNum   int        `json:"page_num"`
}

func (t *Activity) TableName() string {
	return TableName("activity")
}

func ActivityAdd(t *Activity) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(t)

	return id, err
}

func ActivityList(pageNum int, pageSetNum int) (list ActivityListData, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("activity").Filter("status", 1)
	count, _ := qs.Count()
	pageSetNum = 10
	pageCount := math.Ceil((float64(count) / float64(pageSetNum)))
	//存储分页数据的切片
	articles := new([]Activity)
	//获取分页数据
	_, _ = qs.OrderBy("-id").Limit(pageSetNum, pageSetNum*(pageNum-1)).All(articles)

	return ActivityListData{
		List:      articles,
		Count:     count,
		PageCount: pageCount,
		PageNum:   pageNum,
	}, nil
}

func ActivitySingleUser(pageNum int, userId int64) (list ActivityListData, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("activity").Filter("status", 1).Filter("user_id", userId)
	count, _ := qs.Count()
	pageSetNum := 10
	pageCount := math.Ceil((float64(count) / float64(pageSetNum)))
	//存储分页数据的切片
	articles := new([]Activity)
	//获取分页数据
	_, _ = qs.OrderBy("-id").Limit(pageSetNum, pageSetNum*(pageNum-1)).All(articles)

	return ActivityListData{
		List:      articles,
		Count:     count,
		PageCount: pageCount,
		PageNum:   pageNum,
	}, nil
}

func ActivityDetail(id int) (article Activity, err error) {
	o := orm.NewOrm()
	article = Activity{Id: id}
	err = o.Read(&article, "id")

	return article, nil
}

func ActivitySearch(text string) (article *[]Activity, err error) {
	articles := new([]Activity)

	o := orm.NewOrm()
	qs := o.QueryTable("activity").Filter("status", 1).Filter("title__contains", text)
	//获取分页数据
	_, _ = qs.OrderBy("-id").Limit(20).All(articles)

	return articles, nil

}

func ActivityEvaluateIncr(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("activity").Filter("id", id).Update(orm.Params{
		"evaluate_num": orm.ColValue(orm.ColAdd, 1),
	})

	return num, err
}

func ActivityLikeIncrNum(id int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("activity").Filter("id", id).Update(orm.Params{
		"like_num": orm.ColValue(orm.ColAdd, 1),
	})

	return num, err
}
