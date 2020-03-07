package wechatModules

import (
	"abana/components/common"
	"abana/components/redis"
	"abana/controllers"
	"abana/enum"
	"abana/models"
	"github.com/astaxie/beego"
	"math/rand"
	"strconv"
	"time"
)

type WechatController struct {
	controllers.BaseNotController
}

type RegisterReq struct {
	Code string `json:"code"`
}

type RegisterResponse struct {
	SessionKey string `json:"session_key"`
	OpenId     string `json:"openid"`
}

/**
根据code获取open_id判断是否已经注册
*/
func (c *WechatController) RegisterUser() {
	var registerRes RegisterResponse
	var userId int64
	var token string

	registerRes = c.getOpenId(registerRes)
	user, _ := models.UserGetByOpenId(registerRes.OpenId)

	if user.Id == 0 {
		//注册新用户
		rand.Seed(time.Now().UnixNano())
		id := strconv.Itoa(rand.Intn(100)) + strconv.Itoa(int(time.Now().Unix())) + strconv.Itoa(rand.Intn(100)) + strconv.Itoa(rand.Intn(100)) + strconv.Itoa(rand.Intn(100))
		newId, _ := strconv.ParseInt(id, 10, 64)
		models.UserAdd(&models.User{
			OpenId:       registerRes.OpenId,
			RegisterTime: time.Now().Unix(),
			UserId:       newId,
		})
		userId = newId
	} else {
		userId = user.UserId
	}
	token = common.Md5V(strconv.Itoa(int(userId)))

	key := enum.REDIS_KEY_USER_INFO + strconv.FormatInt(userId, 10)
	client := redis.GetCommonRedis()
	client.Set(key, user, enum.REDIS_EXPIRE_TIME_BY_EIGHT_HOUR)

	//返回必要的参数用来验证用户是否授权
	c.RenderJson(map[string]interface{}{
		"token":     token,
		"open_id":   registerRes.OpenId,
		"nick_name": user.NickName,
		"avatar":    user.AvatarUrl,
		"phone":     user.Phone,
	})
}

func (c *WechatController) getOpenId(res RegisterResponse) (registerRes RegisterResponse) {
	req := &RegisterReq{}
	c.GetParams(req)

	appId := beego.AppConfig.String("mall_app_id")
	appSecrete := beego.AppConfig.String("mall_app_secrete")
	code := req.Code

	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + appId + "&secret=" + appSecrete + "&js_code=" + code + "&grant_type=authorization_code"
	_ = common.Get(url, &res)
	return res
}
