package controllers

import (
	"beego_vue/backend/models"
	"github.com/astaxie/beego/logs"
	"strconv"

	//"beego_vue/backend/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/logs"
	//"github.com/astaxie/beego/orm"
)

type DeleteController struct {
	beego.Controller
}

func (c *DeleteController) DeleteTodo() {
	o := orm.NewOrm()
	var todos []models.Todo
	// 读取所有数据
	deleteId, _ := strconv.Atoi(c.GetString("Todo_id"))
	logs.Info("------------", deleteId)
	wholeResult := map[string]interface{}{"error_num": 400, "msg": "查询出错"}
	_, errOrmRead := o.QueryTable("todo").All(&todos)
	if errOrmRead != nil {
		logs.Info(" errOrmRead err!" + errOrmRead.Error())
	} else {
		for _, todo := range todos {
			logs.Info(todo)
		}
		//logs.Info(num)
		wholeResult["msg"] = "success"
		wholeResult["error_num"] = 0
		c.Data["json"] = wholeResult
		c.ServeJSON()
	}
}
