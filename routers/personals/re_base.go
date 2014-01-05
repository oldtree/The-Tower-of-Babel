package personals

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func Debug(infos ...interface{}) {
	if true {
		fmt.Printf("DEBUG: "+fmt.Sprintf("%s\n", infos[0]), infos[1:]...)
	}
}
