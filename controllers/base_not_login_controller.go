package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
)

type BaseNotController struct {
	beego.Controller
}

func (c *BaseNotController) GetParams(params interface{}) {
	req := c.Ctx.Input.RequestBody
	data := json.Unmarshal(req, params)
	if data != nil {
		//panic("error")
	}
}

func (c *BaseNotController) LogsInfo(data interface{}) {
	beego.Info("file", `{"filename":"runtime/test.log"}`)
}

func (c *BaseNotController) LogsError(data interface{}) {
	beego.Error("file", `{"filename":"runtime/test.log"}`)
}

func (c *BaseNotController) RenderJson(data interface{}) {
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

func (c *BaseNotController) RenderJsonErr(data interface{}, err interface{}) {
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
