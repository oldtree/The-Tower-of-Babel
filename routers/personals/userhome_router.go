package personals

import (
	"fmt"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	fmt.Println("进化者")
	this.Data["title"] = "进化者"
	this.TplNames = "home.html"
}
