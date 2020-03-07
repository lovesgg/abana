package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

/**
这是小程序api项目的基础类
*/

var (
	reqeust_start_time = time.Now()
	request_end_time   = time.Now()
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

// Predefined controller error values.
var (
	err404 = &ControllerError{404, 404, "page not found", "page not found", ""}
)

func (c *BaseController) checkUser() {
	token := c.GetString("token", "")
	if token == "" {
		panic("get token error")
	}
	fmt.Println(token)
}

func (c *BaseController) GetParams(params interface{}) {
	req := c.Ctx.Input.RequestBody
	data := json.Unmarshal(req, params)
	if data != nil {
		//panic("error")
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
