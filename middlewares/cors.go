package middlewares

import(
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func Crs() context.Handler{
	return func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Content-Type")
		ctx.Header("Access-Control-Max-Age", "1800")
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH")
		ctx.Next()
	}
}