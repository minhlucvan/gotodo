package exceptions

import "github.com/kataras/iris"

func CustomHTTPErrors(){
	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Write("CUSTOM 500 INTERNAL SERVER ERROR PAGE")
		// or ctx.Render, ctx.HTML any render method you want
		ctx.Log("error: INTERNAL SERVER")
	})
	
	iris.OnError(iris.StatusUnauthorized, func(ctx *iris.Context) {
		ctx.Write("f**k you, wana hack me?")
		// or ctx.Render, ctx.HTML any render method you want
		ctx.Log("error: INTERNAL SERVER")
	})
}
