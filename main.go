package quantfront

import (
	"github.com/astaxie/beego"
	_ "quantfront/routers"
)

func Run() {
	beego.Run()
}
