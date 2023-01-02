package routers

import (
	"github.com/astaxie/beego"
	"web/controllers"
)

func RouterInit() {
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.TestController{})
}
