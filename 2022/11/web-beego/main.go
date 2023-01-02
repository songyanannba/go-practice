package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

type InputController struct {
	beego.Controller
}

func (c *InputController) QueryParams() {
	c.Ctx.Request.ParseForm()
	fmt.Println(c.Ctx.Request.Form)
	c.Ctx.WriteString(" ")
}

func main() {

	/*beego.Get("/" , function(ctx *context.Context) {
		name := ctx.Input.Query("name")
		ctx.Output.Context.WriteString(fmt.Sprintf("name is  :%s" , name))

	})*/
	beego.AutoRouter(&InputController{})
	beego.Run()
}
