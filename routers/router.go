package routers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/itbread/dataweb/controllers"
)

//Configure
//配置路由
func Configure(app *iris.Application) {
	route(app,"/",new(controllers.HomeController))
	route(app,"/vehicle",new(controllers.VehicleController))
}

func route(app *iris.Application,partyPath string, controller interface{}){
	mvc.New(app.Party(partyPath)).Handle(controller)
}