package models

import (
	"github.com/astaxie/beego/orm"
	"math"
)

type User struct {
	Id           int    `json:"id"`
	UserId       int64  `json:"user_id"`
	NickName     string `json:"nick_name"`
	Gender       int    `json:"gender"`
	Phone        string `json:"phone"`
	RegisterTime int64  `json:"register_time"`
	Integral     int    `json:"integral"`
	OpenId       string `json:"open_id"`
	IfDelete     int    `json:"if_delete"`
	VipLevel     int    `json:"vip_level"`
	AvatarUrl    string `json:"avatar_url"`
	City         string `json:"city"`
	Province     string `json:"province"`
	TrueName     string `json:"true_name"`
}

func (t *User) TableName() string {
	return TableName("user_mall_info")
}

func UserAdd(t *User) (id int64) {
	o := orm.NewOrm()
	id, _ = o.Insert(t)

	return id
}

func UserGetByUserId(userId int64) (user User, err error) {
	var data User
	data.UserId = userId
	o := orm.NewOrm()
	err = o.Read(&data)

	if err == orm.ErrNoRows {
		return User{}, err
	}
	return data, nil
}

func UserGetByOpenId(openId string) (user User, err error) {
	o := orm.NewOrm()
	user = User{OpenId: openId}
	err = o.Read(&user, "open_id")

	return user, nil
}

func UserUpdate(userId int64, update User) int64 {
	o := orm.NewOrm()
	user := User{UserId: userId}
	var num int64
	if ret := o.Read(&user, "user_id"); ret == nil {
		user.NickName = update.NickName
		user.Province = update.Province
		user.City = update.City
		user.AvatarUrl = update.AvatarUrl
		user.Gender = update.Gender
		if num, err := o.Update(&user); err == nil {
			return num
		}
	}
	return num
}

func UserNearList(pageNum int) (list *[]User, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("user_mall_info").Filter("if_delete", 0)
	count, _ := qs.Count()
	pageSetNum := 10
	pageCount := math.Ceil((float64(count) / float64(pageSetNum)))
	if pageCount <= 0 {

	}
	users := new([]User)
	//获取分页数据
	_, _ = qs.Limit(pageSetNum, pageSetNum*(pageNum-1)).All(users)

	return users, nil
}
