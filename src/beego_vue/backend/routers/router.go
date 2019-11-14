package routers

import (
	"beego_vue/backend/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// Get方法对应MainController对象里的Show
    beego.Router("/", &controllers.MainController{},"get:Show")
}
