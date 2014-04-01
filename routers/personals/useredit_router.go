package personals

import (
	//	"Eva/models"
	"fmt"
	//	"github.com/astaxie/beego/validation"
	//	"strings"
	"Eva/routers"
)

type EditController struct {
	routers.BaseController
}

func (this *EditController) Get() {
	fmt.Println("进化者")
	this.Data["title"] = "EDIT self"
	this.TplNames = "editself.html"
}
func (this *EditController) Put() {
	fmt.Println("进化者")
	this.Data["title"] = "EDIT self"
	this.TplNames = "editself.html"
}
func (this *EditController) Delete() {
	fmt.Println("进化者")
	this.Data["title"] = "EDIT self"
	this.TplNames = "editself.html"
}
func (this *EditController) Post() {
	fmt.Println("进化者")
	this.Data["title"] = "EDIT self"
	this.TplNames = "editself.html"
}
