package routers

import (
	"fmt"
	"github.com/astaxie/beego"
)

func Debug(infos ...interface{}) {
	if true {
		fmt.Printf("DEBUG: "+fmt.Sprintf("%s\n", infos[0]), infos[1:]...)
	}
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["title"] = "login"
	this.TplNames = "first.html"
}
func (this *MainController) Post() {
	this.Data["title"] = "login"
	this.TplNames = "first.html"
}
