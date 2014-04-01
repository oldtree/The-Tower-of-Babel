package personals

import (
	"Eva/routers"
	"fmt"
)

type HomeController struct {
	routers.BaseController
}

func (this *HomeController) Get() {
	fmt.Println("进化者")
	user_name := this.Ctx.GetCookie("user_name")
	user_email := this.Ctx.GetCookie("user_email")
	user_login := this.Ctx.GetCookie("user_login")
	fmt.Println("get cookie")
	fmt.Println(user_name)
	fmt.Println(user_email)
	fmt.Println(user_login)
	if user_login != "true" {
		this.Redirect("/login", 302)
	}
	this.Data["title"] = "进化者"
	this.Data["username"] = user_name
	this.Data["useremail"] = user_email
	this.TplNames = "home.html"
}

func (this *HomeController) Post() {
	user_name := this.Ctx.GetCookie("user_name")
	user_email := this.Ctx.GetCookie("user_email")
	user_login := this.Ctx.GetCookie("user_login")
	fmt.Println("get cookie")
	fmt.Println(user_name)
	fmt.Println(user_email)
	fmt.Println(user_login)
	if user_login != "true" {
		this.Redirect("/login", 302)
	}
	this.Data["title"] = "进化者"
	this.Data["username"] = user_name
	this.Data["useremail"] = user_email
	this.TplNames = "home.html"
}
