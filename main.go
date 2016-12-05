package main


import (
	"github.com/kataras/iris"
	"./config"
	"github.com/minhlucvan/gotodo/controllers/home"
	"github.com/minhlucvan/gotodo/middwares/i18n"
	"github.com/minhlucvan/gotodo/middwares/logger"
	"github.com/kataras/go-template/html"
	"github.com/minhlucvan/gotodo/API/authapi"
	"github.com/minhlucvan/gotodo/middwares/auth"
	"github.com/minhlucvan/gotodo/exceptions"
)

func init(){
	
}

func setMiddlewares(){
	iris.Use(logger.Handler)
	iris.Use(i18n.Handler)
}

func setControllers(){
	iris.Handle("GET",  "/", home.Index())
}

func setAPI(){
	iris.Handle("POST", "/login/", authapi.Login{})
	admin := iris.Party("/admin")
	admin.Use(auth.RequireRole("admin"))
	admin.Get("/", func(ctx *iris.Context) {
		ctx.WriteString("this is admin dashboard.")
	})
}

func setConfig(){
	iris.Config.IsDevelopment = config.DEV_MODE
	iris.Static("/assets", config.ASSETSPATH, 1)
	iris.UseTemplate(html.New(html.Config{Layout: config.TPL_LAYOUT})).Directory(config.TPL_LOCATION, config.TPL_FORMAT)
	exceptions.CustomHTTPErrors()
}

func startServer(){
	iris.Listen(config.SERVER_PORT)
}

func main(){
	setConfig()
	setMiddlewares()
	setControllers()
	setAPI()
	startServer()
}
