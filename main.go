// webapp project main.go
package main

import (
	"fmt"
	"github.com/itbread/dataweb/routers"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	tmpl := iris.Django("./views", ".html").Reload(true)
	app.RegisterView(tmpl)
	app.StaticWeb("/public", "./public")
	SystemInit(app,"")
	app.Run(iris.Addr(":8090"))
}

func SystemInit(app *iris.Application, fielName string) {
	fmt.Println("sysinit===============starting")
	routers.Configure(app)
	fmt.Println("sysinit===============started")
}
