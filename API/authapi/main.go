package auth

import (
	//"github.com/minhlucvan/gotodo/middwares/auth"
	"github.com/kataras/iris"
)

type AuthAPI struct {
	
}


func (api AuthAPI) Serve(ctx *iris.Context){
	ctx.Session().Set("token", "secret")
	ctx.JSON(iris.StatusOK, ctx.Session().GetAll())
	
}