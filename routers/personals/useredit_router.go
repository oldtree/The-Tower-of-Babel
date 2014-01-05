package personals

import (
	//	"Eva/models"
	"fmt"
	//	"github.com/astaxie/beego/validation"
	//	"strings"
)

type EditController struct {
	BaseController
}

func (this *EditController) Get() {
	fmt.Println("进化者")
	this.Data["title"] = "进化者"
	this.TplNames = "home.html"
}
