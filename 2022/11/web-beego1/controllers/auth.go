package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

//负责 认证 控制器
type AuthController struct {
	beego.Controller
}

func (c *AuthController) Login() {
	fmt.Println("login")
	//get 直接加载页面
	//post 验证
	if c.Ctx.Input.IsPost() {
		return
	}

	c.TplName = "auth/login.html"

}
