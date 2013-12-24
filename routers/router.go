package routers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["title"] = "this is the end"
	this.Data["path1"] = "F:/forgo/src/Eva/static/img"
	this.Data["path2"] = "F:/forgo/src/Eva/static/img"
	this.Data["path3"] = "F:/forgo/src/Eva/static/img"
	this.TplNames = "first.html"
}
