package routers

import (
	"abana/controllers"
	articleModules "abana/modules/article"
	userModules "abana/modules/user"
	wechatModules "abana/modules/wechat"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/user/user/update", &userModules.UserController{}, "*:Update")
	beego.Router("/user/user/near-list", &userModules.UserController{}, "*:NearList")

	beego.Router("/wechat/wechat/register-user", &wechatModules.WechatController{}, "*:RegisterUser")

	beego.Router("/article/article/publish", &articleModules.ArticleController{}, "*:Publish")
	beego.Router("/article/article/list", &wechatModules.WechatController{}, "*:List")
	beego.Router("/article/article/user-list", &articleModules.ArticleController{}, "*:UserList")
	beego.Router("/article/article/detail", &wechatModules.WechatController{}, "*:Detail")
	beego.Router("/article/article/search", &articleModules.ArticleController{}, "*:Search")

}
