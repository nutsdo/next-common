package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/nutsdo/go-next/models"
	"github.com/nutsdo/go-next/responses"
	"gopkg.in/go-playground/validator.v9"
)

func RegisterHandler(ctx iris.Context) {

	//username:=ctx.FormValueDefault("username","")
	//nickname:=ctx.FormValueDefault("nickname","")
	phone:=ctx.FormValueDefault("phone","")

	validate:=validator.New()
	user:=new(models.User)
	user.Phone = phone
	err:= validate.Struct(user)

	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
			return
		}
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(responses.SuccessResponse(101001, "验证失败",err))
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println()
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
	}else {

		u,err := models.CreateUser(user);
		if err!=nil{
			ctx.StatusCode(iris.StatusOK)
			ctx.JSON(responses.ErrResponse(200,100002, err.Error()))
			return
		}

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(responses.SuccessResponse(0 ,"注册成功",u))
	}

}

