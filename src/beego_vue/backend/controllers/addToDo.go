package controllers

import (
	"beego_vue/backend/models"
	"github.com/astaxie/beego/logs"
	"time"

	//"beego_vue/backend/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/logs"
	//"github.com/astaxie/beego/orm"
)

type AddController struct {
	beego.Controller
}

func (c *AddController) AddTodo() {
	o := orm.NewOrm()
	todo := new(models.Todo)
	// 读取所有数据
	todoBody := c.GetString("Todo_body")
	logs.Info("------------", todoBody)
	todo.Status = 0
	todo.Todo_body = todoBody
	time.Now().Location()
	todo.Add_time = time.Now()
	todo.Update_time = time.Now()
	wholeResult := map[string]interface{}{"error_num": 400, "msg": "查询出错"}
	_, errAdd := o.Insert(todo)
	if errAdd != nil {
		logs.Info(" errDelete err!" + errAdd.Error())
	} else {
		//logs.Info(num)
		wholeResult["msg"] = "success"
		wholeResult["error_num"] = 0
		c.Data["json"] = wholeResult
		c.ServeJSON()
	}
}
