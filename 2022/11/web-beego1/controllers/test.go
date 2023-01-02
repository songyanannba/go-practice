package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Test() {
	fmt.Println("Test")
}
