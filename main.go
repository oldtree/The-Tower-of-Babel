// Eva1 project main.go
package main

import (
	"Eva1/common_layer"
	"Eva1/routers"
	"Eva1/routers/personals"
	"Eva1/utils"
	"fmt"
	"github.com/astaxie/beego"
)

const (
	APP_VERSION = "0.1.1.1"
)

func init() {
	utils.LoadConfig()

	beego.Info(beego.AppName, utils.AppVersion, utils.AppUrl)

	init_cache()
	init_login_routers()
	init_user_home_routers()
	init_user_project_routers()

	//beego.Router("/home", &routers.HomeController{}, "get:Get")
	fmt.Println("Hello World!")
}

func init_cache() {
	common_layer.Using_cache()
	//common_layer.Mongo_main()
}

func init_login_routers() {
	beego.Router("/", &routers.MainController{}, "get:FirstGet")
	beego.Router("/login", &routers.MainController{}, "get:Get;post:Post")
	beego.Router("/register", &routers.RegisterController{}, "post:Post")
}

func init_user_home_routers() {
	beego.Router("/:id([A-z]+|[0-9]+)/home", &personals.HomeController{}, "get:Get;post:Post")
	beego.Router("/:id([A-z]+|[0-9]+)/edit_self", &personals.HomeController{}, "get:Get;post:Post;put:Put;delete:Delete")
}

func init_user_project_routers() {
	beego.Router("/:id([A-z]+)/home", &personals.HomeController{}, "get:Get;post:Post")
	beego.Router("/:id([A-z]+)/edit_self", &personals.HomeController{}, "get:Get;post:Post;put:Put;delete:Delete")
}

func main() {
	if utils.IsProMode {
		beego.Info("Product mode enabled")
	} else {
		beego.Info("Develment mode enabled")
	}

	beego.Run()
}
