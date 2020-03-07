package routers

import (
	"abana/controllers"
	userModules "abana/modules/user"
	wechatModules "abana/modules/wechat"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/user/user/hello", &userModules.UserController{}, "*:Hello")
	beego.Router("/wechat/wechat/register-user", &wechatModules.WechatController{}, "*:RegisterUser")

}
