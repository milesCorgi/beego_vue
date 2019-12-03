package controllers

import (
	"beego_vue/backend/models"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"

	//"beego_vue/backend/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/logs"
	//"github.com/astaxie/beego/orm"
)

type EditController struct {
	beego.Controller
}

func (c *EditController) EditTodo() {
	o := orm.NewOrm()
	editId, _ := strconv.Atoi(c.GetString("Todo_id"))
	todo := &models.Todo{Id: editId}
	// 读取所有数据
	todoBody := c.GetString("Todo_body")
	logs.Info("------------", todoBody)
	todo.Status = 0
	todo.Todo_body = todoBody
	time.Now().Location()
	todo.Update_time = time.Now()

	wholeResult := map[string]interface{}{"error_num": 400, "msg": "查询出错"}
	_, errAdd := o.Update(todo, "Todo_body", "Update_time")
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
