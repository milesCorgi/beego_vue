package main

import (
	// 载入这个，会自动跑routers的初始化
	_ "beego_vue/backend/routers"
	"github.com/astaxie/beego"
	// 载入这个，会自动跑model的初始化
	_ "beego_vue/backend/models"
)

func main() {
	beego.Run()
}
