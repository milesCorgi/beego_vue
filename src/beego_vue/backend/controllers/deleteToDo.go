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
	// 读取所有数据
	deleteId, _ := strconv.Atoi(c.GetString("Todo_id"))
	logs.Info("------------", deleteId)
	wholeResult := map[string]interface{}{"error_num": 400, "msg": "查询出错"}
	_, errDelete := o.Delete(&models.Todo{Id: deleteId})
	if errDelete != nil {
		logs.Info(" errDelete err!" + errDelete.Error())
	} else {
		//logs.Info(num)
		wholeResult["msg"] = "success"
		wholeResult["error_num"] = 0
		c.Data["json"] = wholeResult
		c.ServeJSON()
	}
}
