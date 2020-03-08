package controllers

import (
	"abana/components/common"
	"abana/components/redis"
	"abana/enum"
	"abana/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

/**
这是小程序api项目的基础类
*/

var (
	reqeust_start_time = time.Now()
	request_end_time   = time.Now()
)

// Predefined controller error values.
var (
	err404 = &ControllerError{404, 404, "page not found", "page not found", ""}
)

type BaseController struct {
	beego.Controller
}

// Predefined const error strings.
const (
	ErrInputData    = "数据输入错误"
	ErrDatabase     = "数据库操作错误"
	ErrDupUser      = "用户信息已存在"
	ErrNoUser       = "用户信息不存在"
	ErrPass         = "密码不正确"
	ErrNoUserPass   = "用户信息不存在或密码不正确"
	ErrNoUserChange = "用户信息不存在或数据未改变"
	ErrInvalidUser  = "用户信息不正确"
	ErrOpenFile     = "打开文件出错"
	ErrWriteFile    = "写文件出错"
	ErrSystem       = "操作系统错误"
)

// ControllerError is controller error info structer.
type ControllerError struct {
	Status   int    `json:"status"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	DevInfo  string `json:"dev_info"`
	MoreInfo string `json:"more_info"`
}

func (c *BaseController) Prepare() {
	var userId int64
	token := c.Ctx.Input.Header("token")
	open_id := c.Ctx.Input.Header("open_id")
	avatar := c.Ctx.Input.Header("avatar")
	nick_name := c.Ctx.Input.Header("nick_name")

	//校验token
	client := redis.GetCommonRedis()
	res := client.Get(enum.REDIS_KEY_USER_INFO + token)
	if res == "" {
		//token过期 重新生成token 这数值不会变
		user, _ := models.UserGetByOpenId(open_id)
		id := user.UserId

		token = common.Md5V(strconv.Itoa(int(id)))
		userId = id

		key := enum.REDIS_KEY_USER_INFO + token
		userSet, _ := json.Marshal(user)
		client.Set(key, userSet, enum.REDIS_EXPIRE_TIME_BY_EIGHT_HOUR)
	} else {
		var cacheData models.User
		_ = json.Unmarshal([]byte(res), &cacheData)
		userId = cacheData.UserId
	}


	//下文可直接从ctx中获取数据
	c.Ctx.Input.SetData("token", token)
	c.Ctx.Input.SetData("open_id", open_id)
	c.Ctx.Input.SetData("user_id", userId)
	c.Ctx.Input.SetData("avatar", avatar)
	c.Ctx.Input.SetData("nick_name", nick_name)

}

func (c *BaseController) GetParams(params interface{}) {
	req := c.Ctx.Input.RequestBody
	data := json.Unmarshal(req, params)
	if data != nil {
		//panic("error")
		//fmt.Println("请求参数有误:", data)
	}
}

func (c *BaseController) LogsInfo(data interface{}) {
	beego.Info("file", `{"filename":"runtime/test.log"}`)
}

func (c *BaseController) LogsError(data interface{}) {
	beego.Error("file", `{"filename":"runtime/test.log"}`)
}

func (c *BaseController) RenderJson(data interface{}) {
	request_end_time = time.Now()
	timeNow := request_end_time.Sub(reqeust_start_time)
	str := map[string]interface{}{
		"ret":      1,
		"log_time": timeNow,
		"data":     data,
	}
	c.Data["json"] = str
	c.ServeJSON()
	return
}

func (c *BaseController) RenderJsonErr(data interface{}, err interface{}) {
	request_end_time = time.Now()
	timeNow := request_end_time.Sub(reqeust_start_time)
	str := map[string]interface{}{
		"ret":       0,
		"log_time":  timeNow,
		"data":      data,
		"error_msg": err,
	}
	c.Data["json"] = str
	c.ServeJSON()
	return
}
