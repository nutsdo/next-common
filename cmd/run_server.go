package cmd

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/nutsdo/go-next/controllers"
	"github.com/nutsdo/go-next/middlewares"
	"github.com/nutsdo/go-next/session"
	"github.com/nutsdo/go-next/web/auth"
	"gopkg.in/go-playground/validator.v9"
)

func RunServer()  {
	app:=iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())
	//app.UseGlobal(middlewares.HmacHandler)

	app.Use(session.Sess.Handler())

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("Hello World")
	})

	//用户注册
	app.Post("/signup",controllers.RegisterHandler)

	//用户登录
	app.Post("/signin",controllers.LoginHandler)

	//第三方登录
	app.Post("/third/login",controllers.ThirdLoginHandler)

	//账号相关路由
	account := app.Party("/account",middlewares.JwtHandler().Serve)
	//更新用户信息
	//account.Patch("/info")

	//账号绑定
	account.Post("/oauth/bind", controllers.OAuthBind)

	//第三方授权
	account.Post("/oauth/authorize",controllers.ThirdAuthorize)

	//第三方登录回调
	account.Post("/oauth/callback",controllers.OAuthCallback)

	//用户路由
	user := app.Party("/user",middlewares.JwtHandler().Serve)

	user.Get("/info",controllers.UserInfo)

	//sms := app.Party("/sms",middlewares.SMSHandler)
	app.Post("/sms",controllers.SMSHandler)
	//收银台
	cashier := app.Party("/cashier")
	cashier.Party("/payment")

	app.Get("/web/login", auth.Authorize)

	app.Get("/test/auth",controllers.TestKongHmacSign)
	app.Get("/test/session",session.SessionMiddleware)

	app.Run(iris.Addr(":80"), iris.WithoutServerError(iris.ErrServerClosed))
}


var validate *validator.Validate


