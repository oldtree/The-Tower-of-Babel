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

func initialize() {
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

	fmt.Println("Hello World!")
	beego.Run()
}
