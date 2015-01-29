package controllers

import (
	"fmt"
	_ "github.com/astaxie/beego"
	"quant"
	"quant/base/strategy"
	"strconv"
)

type PerformanceController struct {
	BaseController
}

type Respond struct {
	Tables   []string
	Strategy []strategy.IStrategy
}

// http://localhost:8080/p
// http://localhost:8080/p?project=0&stragegy=MyStrategy&stock=STO000002
// http://localhost:8080/p?project=1&strategy=MyStrategy&stock=STO000002
func (this *PerformanceController) Get() {

	rsp := Respond{}

	if this.baseForm.Project == "" {
		this.Data["json"] = quant.DefaultQuant.Projects
		this.ServeJson()

	} else {
		num64, _ := strconv.ParseInt(this.baseForm.Project, 10, 32)

		project := quant.DefaultQuant.Projects[num64]
		oneStrategy, ok := project.MapStrategy[this.baseForm.Strategy+this.baseForm.Stock]
		if ok {
			//fmt.Println("do project draw", project)
			fmt.Println("draw strategy:", oneStrategy.Key())
			rsp.Tables = oneStrategy.DoSvgDrawing()
			rsp.Strategy = append(rsp.Strategy, oneStrategy)
		}
		// for idx, content := range rsp.Tables {
		// 	this.Ctx.WriteString(content)
		// 	this.Ctx.WriteString(fmt.Sprintf("table:%d above.\n", idx))
		// }

		this.TplNames = "tables.html"
		this.Data["Tables"] = rsp.Tables

	}
}
