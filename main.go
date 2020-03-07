package main

import (
	"abana/components/redis"
	_ "abana/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqlUserName := beego.AppConfig.String("mysqluser")
	mysqlPass := beego.AppConfig.String("mysqlpass")
	mysqlIp := beego.AppConfig.String("mysqlurls")
	mysqlPort := beego.AppConfig.String("mysqlport")
	mysqlDb := beego.AppConfig.String("mysqldb")
	dataSource := mysqlUserName + ":" + mysqlPass +"@tcp(" + mysqlIp +":" + mysqlPort +")/" + mysqlDb +"?charset=utf8"
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", dataSource)

	redis.RedisClient("redis")
}

func main() {
	beego.Run()
}
