package session

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

var (
	Sess = sessions.New(sessions.Config{Cookie:"nextsessionid"})
)


func SessionMiddleware(ctx iris.Context) {
	session := Sess.Start(ctx)
	session.Get("user")
}