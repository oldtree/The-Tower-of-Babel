package routers

import (
	"Eva1/models"
	"fmt"
	//"strings"
	//	"strconv"
)

type MainController struct {
	BaseController
}

func (this *MainController) FirstGet() {
	this.Redirect("/login", 302)
	fmt.Println("into First Get function")
}
func (this *MainController) Get() {
	this.Data["title"] = "进化者"
	this.TplNames = "first.html"
}
func (this *MainController) Post() {
	user_email := this.GetString("email")
	user_password := this.GetString("password")
	fmt.Println(user_email)
	fmt.Println(user_password)
	user_info := models.User{}
	user_info.User_email = user_email
	user_info.User_password = user_password

	statu := models.User_login(&user_info, &this.Controller)
	fmt.Println("fmt.Println(statu)")
	fmt.Println(statu)
	if !statu {
		fmt.Println("/register")
		this.Redirect("/register", 302)
	} else {
		//path := strconv.Itoa(int(user_info.Id))
		path := user_info.User_name
		path = "/" + path + "/home"
		fmt.Println(path)
		//models.Get_user_byname(user_info.User_name)
		fmt.Println(user_info.User_name)
		fmt.Println(user_info.User_email)
		fmt.Println("-----------------------------------------------------------")
		this.Ctx.SetCookie("user_name", user_info.User_name)
		this.Ctx.SetCookie("user_email", user_info.User_email)
		this.Ctx.SetCookie("user_login", "true")
		fmt.Println("-----------------------------------------------------------")
		fmt.Println("begin info ")
		fmt.Println(user_info.User_name)
		fmt.Println(user_info.User_email)
		fmt.Println("end info ")
		this.Redirect(path, 302)
	}

	return
}
