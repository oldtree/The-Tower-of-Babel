package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

func Debug(infos ...interface{}) {
	if true {
		fmt.Printf("DEBUG: "+fmt.Sprintf("%s\n", infos[0]), infos[1:]...)
	}
}

type NestPreparer interface {
	NestPrepare()
}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Data["PageStartTime"] = time.Now()
	this.Data["context"] = "There is no way to the hell!"

	this.Data["local"] = "Battle Field 3"
	this.Data["Title"] = "½ø»¯Õß"

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (this *BaseController) NestPreparer() {
	islogin := this.Ctx.GetCookie("login")

	if islogin != "true" {
		this.Redirect("/login", 302)
		return
	}
}

func (this *BaseController) Get() {

}

func (this *BaseController) Post() {

}

func (this *BaseController) Delete() {

}

func (this *BaseController) Head() {

}

func (this *BaseController) Put() {

}

func (this *BaseController) Patch() {

}

func (this *BaseController) Options() {

}

func (this *BaseController) Finish() {

}
