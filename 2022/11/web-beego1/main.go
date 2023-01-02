package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"web/routers"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	fmt.Println("22222")
}

func main() {

	routers.RouterInit()
	beego.Run()
}
