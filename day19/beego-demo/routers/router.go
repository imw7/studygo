package routers

import (
	"beego-demo/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/register", &controllers.MainController{}, "get:RegisterGet;post:RegisterPost")
	beego.Router("/login", &controllers.MainController{}, "get:LoginGet;post:LoginPost")
}
