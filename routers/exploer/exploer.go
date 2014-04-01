package explorer

import (
	//	"Eva/models"
	"fmt"
	//	"github.com/astaxie/beego/validation"
	//	"strings"
	"Eva/routers"
)

type ExploerController struct {
	routers.BaseController
}

func (this *ExploerController) Get() {
	fmt.Println("???бе??")
	this.Data["title"] = "explorer"
	this.TplNames = "exploer.html"
}
func (this *ExploerController) Put() {
	this.Data["title"] = "explorer - search resualt"
	this.TplNames = "exploer.html"
}
