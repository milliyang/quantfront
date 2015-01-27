package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

const (
	BaseControllerDebugOn = true
)

type BaseController struct {
	beego.Controller
	baseForm *antCommonForm
}

// Prepare runs after Init before request function execution.
// If url don't need login, override Prepare()
func (this *BaseController) Prepare() {
	if BaseControllerDebugOn {
		fmt.Print("", this.Ctx.Input.Uri())
		fmt.Printf(" " + this.Ctx.Request.Method)
		fmt.Printf(" OnPrepare\n")
	}
	// parse header key-value
	form := antCommonForm{}
	this.ParseForm(&form)
	this.baseForm = &form
}

//
// Release resource here, just make it easier for GC
//
func (this *BaseController) Finish() {
	this.baseForm = nil
	this.Data["json"] = nil
	this.Data = nil
}

//
// After looking at func ParseForm(xxx) in [beego templatefunc.go],
// i thought that, it won't cost much, event we put all together.
// and also, we can reuse & have a better overview.
//
type antCommonForm struct {
	Vt       string `form:"vt"`
	Project  string `form:"project"`
	Strategy string `form:"strategy"`
	Stock    string `form:"stock"`
}

type AntResponse struct {
	Code int    `respond code`
	Msg  string `message`
	Data string `data fron device`
}
