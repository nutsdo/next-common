package controllers

import (
	"fmt"
	"github.com/kataras/iris"
)


func UserInfo(ctx iris.Context) {

	token:= ctx.Header

	fmt.Println(token)

}

