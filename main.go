// webapp project main.go
package main

import (
	"fmt"

	"github.com/itbread/dataweb/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	//　从　"./views"　目录下加载扩展名是".html" 　的所有模板，
	//　并使用标准的　`html/template`　 包进行解析。
	tmpl := iris.Django("./views", ".html").Reload(true) //是否更新
	app.RegisterView(tmpl)
	app.StaticWeb("/public", "./public") //指定静态文件路径

	pages := mvc.New(app.Party("/list"))          // 一个根路径为 /pages 的组
	pages.Handle(new(controllers.HomeController)) // 定义组的controller
	route(app,"/",new(controllers.HomeController))
	app.Get("/", func(ctx iris.Context) {
		// 绑定： {{.message}}　为　"Hello world!"
		ctx.ViewData("message", "Hello world!")
		// 渲染模板文件：
		ctx.View("index.html")
	})

	app.Get("/login", func(ctx iris.Context) {

		username := ctx.PostValue("username")
		password := ctx.PostValue("password")
		fmt.Printf("username==%v password=%v \n", username, password)
		// 渲染模板文件：
		ctx.View("index.html")
	})

	app.Post("/login", func(ctx iris.Context) {

		username := ctx.PostValue("username")
		password := ctx.PostValue("password")
		username1 := ctx.FormValue("username")
		password1 := ctx.FormValue("password")
		fmt.Printf("username==%v password=%v \n", username, password)
		fmt.Printf("username1==%v password1=%v \n", username1, password1)
		// 渲染模板文件：
		ctx.View("login.html")
	})
	//　使用网络地址启动服务
	app.Run(iris.Addr(":8090"))
}

func route(app *iris.Application,partyPath string, controller interface{}){
	mvc.New(app.Party(partyPath)).Handle(controller)
}
