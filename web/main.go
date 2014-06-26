package main

import (
	"fmt"
	"tracery/web/controllers"

	"github.com/astaxie/beego"
)

func main() {
	//beego.SetStaticPath("/html", "tpl")

	//beego.Get("/", func(ctx *context.Context) {
	//	ctx.Output.Body([]byte("hello world"))
	//})

	fmt.Println(beego.AppConfig.String("mysql.user"))
	fmt.Println(beego.AppConfig.String("mysql.password"))
	fmt.Println(beego.AppConfig.String("mysql.url"))
	fmt.Println(beego.AppConfig.String("mysql.db"))

	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/admin/login", &controllers.AdminController{})

	beego.Router("/", &controllers.HomeController{})

	beego.Run()
}
