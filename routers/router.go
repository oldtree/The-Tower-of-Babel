package routers

import (
	"Eva1/models"
	"fmt"
	"github.com/astaxie/beego"
	//"strings"
)

func Debug(infos ...interface{}) {
	if true {
		fmt.Printf("DEBUG: "+fmt.Sprintf("%s\n", infos[0]), infos[1:]...)
	}
}

type MainController struct {
	beego.Controller
}

func (this *MainController) FirstGet() {
	this.Redirect("/login", 302)
	fmt.Println("into First Get function")
}
func (this *MainController) Get() {
	this.Data["title"] = "EVA"
	this.TplNames = "first.html"
}
func (this *MainController) Post() {
	//this.Data["title"] = "Home"

	//this.TplNames = "first.html"
	user_email := this.GetString("email")
	user_password := this.GetString("password")
	fmt.Println(user_email)
	fmt.Println(user_password)

	//form := models.UserloginForm{}
	//err := this.ParseForm(&form)
	//fmt.Println(form.Email)
	//fmt.Println(form.Password)

	//if parse(user_email) {
	//	beego.Error("User count is error!")
	//	return
	//}
	//if parse(user_password) {
	//	beego.Error("User user_password is error!")
	//	return
	//}
	user_info := models.User{}
	user_info.User_email = user_email
	user_info.User_password = user_password
	statu := models.User_login(&user_info, &this.Controller)
	fmt.Println(statu)
	if !statu {
		fmt.Println("/register")
		this.Redirect("/register", 302)
	} else {
		path := "/123456789/home" // + []byte(user_info.Id)
		//path := "/home"
		this.Redirect(path, 302)
		fmt.Println("/home")
	}

	return
}
