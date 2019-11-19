package controllers

import (
	"beego_vue/backend/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Show() {
	c.TplName = "index.html"
}

type ShowController struct {
	beego.Controller
}

func (c *ShowController) ShowTodo() {
	o := orm.NewOrm()
	var todos []models.Todo
	// 读取所有数据
	wholeResult := map[string]interface{}{"error_num": 400, "msg": "查询出错"}
	num, errOrmRead := o.QueryTable("todo").All(&todos)
	if errOrmRead != nil {
		logs.Info(" errOrmRead err!" + errOrmRead.Error())
	} else {
		for _, todo := range todos {
			logs.Info(todo)
		}
		logs.Info(num)
		wholeResult["list"] = todos
		wholeResult["msg"] = "success"
		wholeResult["error_num"] = 0
		c.Data["json"] = wholeResult
		c.ServeJSON()
	}
}
