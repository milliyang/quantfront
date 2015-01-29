package quantfront

import (
	"github.com/astaxie/beego"
	_ "quantfront/routers"
)

func Run() {
	// WorkAround for:
	// beego:can't find templatefile in the path:tables.html
	beego.BuildTemplate("../quantfront/views")
	beego.Run()
}
