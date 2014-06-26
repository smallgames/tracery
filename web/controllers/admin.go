package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	this.TplNames = "login.html"
}

func (this *AdminController) Post() {
	fmt.Println("test")
	this.TplNames = "login.html"
}
