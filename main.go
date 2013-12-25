// Eva1 project main.go
package main

import (
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

	beego.Router("/", &routers.MainController{})
	beego.Router("/register", &routers.RegisterController{}, "post:Post")
	beego.Router("/login", &routers.MainController{}, "post:Post")
	beego.Router("/home", &routers.MainController{}, "get:Get")
	beego.Router("/home/edit_self", &routers.MainController{}, "post:Post")
	fmt.Println("Hello World!")
	beego.Run()
}
