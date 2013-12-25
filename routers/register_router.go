package routers

import (
	"Eva1/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	fmt.Println("into Get function ")
	this.Data["title"] = "注册"
	this.TplNames = "register.html"

}
func (this *RegisterController) Post() {
	fmt.Println("into post function ")
	this.Data["title"] = "注册"
	this.TplNames = "first.html"

	form := models.RegisterForm{}
	err := this.ParseForm(&form)
	fmt.Println(err)
	fmt.Println(1)
	if err != nil {
		beego.Error("ParseForm RegisterForm Error")
		return
	}
	fmt.Println(err)
	fmt.Println(2)
	valid := validation.Validation{}
	ok, _ := valid.Valid(form)
	if !ok {
		err := valid.ErrorMap()
		//todo 处理表单错误信息
		beego.Error(err)
		return
	}
	fmt.Println(err)
	fmt.Println(3)
	user := new(models.User)

	err = models.RegisterUser(user, form)

	if err == nil {
		models.LoginUser(user, &this.Controller)
		this.Redirect("/home", 302)
	} else {
		this.Redirect("/register", 302)
	}
	return

}
