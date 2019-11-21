package routers

import (
	"beego_vue/backend/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// Get方法对应MainController对象里的Show
	beego.Router("/", &controllers.MainController{}, "get:Show")
	beego.Router("/api/show_todos", &controllers.ShowController{}, "get:ShowTodo")
	beego.Router("/api/delete_todos", &controllers.DeleteController{}, "post:DeleteTodo")
	beego.Router("/api/add_todos", &controllers.AddController{}, "post:AddTodo")
}
