// Eva1 project main.go
package main

import (
	"Eva1/common_layer"
	"Eva1/routers"
	"Eva1/utils"
	"fmt"
	"github.com/astaxie/beego"
)

const (
	APP_VERSION = "0.1.1.1"
)

func init() {
	utils.LoadConfig()
}
func main() {
	if utils.IsProMode {
		beego.Info("Product mode enabled")
	} else {
		beego.Info("Develment mode enabled")
	}
	beego.Info(beego.AppName, utils.AppVersion, utils.AppUrl)
	common_layer.Using_cache()
	common_layer.Mongo_main()
	beego.Router("/", &routers.MainController{}, "get:FirstGet")
	beego.Router("/login", &routers.MainController{}, "get:Get;post:Post")

	beego.Router("/register", &routers.RegisterController{}, "post:Post")

	beego.Router("/:id([0-9]+)/home", &routers.HomeController{}, "get:Get;post:Post")
	beego.Router("/:id([0-9]+)/edit_self", &routers.HomeController{}, "get:Get;post:Post;put:Put;delete:Delete")
	//beego.Router("/home", &routers.HomeController{}, "get:Get")
	fmt.Println("Hello World!")

	beego.Run()
}
