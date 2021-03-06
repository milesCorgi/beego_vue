package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	logs.Info(" models init")
	mysqlUser := beego.AppConfig.String(`mysqluser`)
	mysqlIp := beego.AppConfig.String(`mysqlurls`)
	mysqlPort := beego.AppConfig.String(`mysqlport`)
	mysqlPass := beego.AppConfig.String(`mysqlpass`)
	mysqlDb := beego.AppConfig.String(`mysqldb`)
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", mysqlUser, mysqlPass, mysqlIp, mysqlPort, mysqlDb)
	errRegisterDriver := orm.RegisterDriver("mysql", orm.DRMySQL)
	if errRegisterDriver != nil {
		logs.Info(" RegisterDriver err!" + errRegisterDriver.Error())
	}
	errRegisterDataBase := orm.RegisterDataBase("default", "mysql", dataSource)
	if errRegisterDataBase != nil {
		logs.Info(" RegisterDriver err!" + errRegisterDataBase.Error())
	} else {
		logs.Info(" 数据库连接成功")
	}

}
