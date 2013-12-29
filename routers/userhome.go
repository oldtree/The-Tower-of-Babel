package routers

import (
	//	"Eva/models"
	"fmt"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	fmt.Println("进化者")
	this.Data["title"] = "进化者"
	this.TplNames = "home.html"
}
