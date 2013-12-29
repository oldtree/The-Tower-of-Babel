package routers

import (
	"Eva1/models"
	"fmt"
	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/validation"
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
	form := models.RegisterForm{}
	err := this.ParseForm(&form)
	fmt.Println(err)
	fmt.Println(1)
	if err != nil {
		beego.Error("ParseForm RegisterForm Error")
		return
	}
	//fmt.Println(err)
	//fmt.Println(2)
	//valid := validation.Validation{}
	//ok, _ := valid.Valid(form)
	//fmt.Println("8877")
	//fmt.Println(ok)

	user_info := models.User{}
	user_info.User_name = form.Email
	user_info.User_password = form.Password

	statu := models.User_login(&user_info, &this.Controller)
	fmt.Println("fmt.Println(statu)")
	fmt.Println(statu)
	if statu {
		//err := valid.ErrorMap()
		//todo 处理表单错误信息
		beego.Error(err)
		this.Redirect("/register", 302)
		return
	}
	//user := new(models.User)

	err = models.RegisterUser(&user_info, form)
	fmt.Println(err)
	fmt.Println(user_info.User_password)
	if err == nil {
		models.LoginUser(&user_info, &this.Controller)
		this.Redirect("/home", 302)
	} else {
		this.Redirect("/register", 302)
	}
	return

}
