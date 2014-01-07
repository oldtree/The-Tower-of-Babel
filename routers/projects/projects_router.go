package projects

import (
	"fmt"
)

type ProjectsController struct {
	BaseController
}

func (this *ProjectsController) Get() {
	fmt.Println("进化者")
	this.Data["title"] = "进化者"
	this.TplNames = "test.html"
}
