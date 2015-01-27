package routers

import (
	"github.com/astaxie/beego"
	"quantfront/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/p", &controllers.PerformanceController{})
	beego.Router("/p/*", &controllers.PerformanceController{})

}
