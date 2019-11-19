package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type Todo struct {
	Id          int       `orm:"column(id);auto"`
	Todo_body   string    `orm:"column(Todo_body);size(45);null" description:"'todo的正文'"`
	Add_time    time.Time `orm:"column(add_time);type(datetime);null" description:"创建时间"`
	Update_time time.Time `orm:"column(update_time);type(datetime);null" description:"更新时间"`
	Status      int       `orm:"column(status);null" description:"状态"`
}

func (t *Todo) TableName() string {
	return "todo"
}

// 如果没有todo表，就新建一个
func init() {
	orm.RegisterModel(new(Todo))
	errRunSyncdb := orm.RunSyncdb("default", false, true)
	if errRunSyncdb != nil {
		logs.Info(" RegisterDriver err!" + errRunSyncdb.Error())
	}
}
