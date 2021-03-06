package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(
		new(User),
		new(Article),
		new(Evaluate),
		new(MyLike),
		new(Activity),
	)
}

func TableName(name string) string {
	return beego.AppConfig.String("mysqlprefix") + name
}
