package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/nutsdo/go-next/responses"
	"time"
)


//验证用户，并返回token
func LoginHandler(ctx iris.Context) {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.StandardClaims{
		Audience:"user",
		ExpiresAt:time.Now().Unix()+3600,
		Id:"1",
		Issuer:"www.next.com",
		IssuedAt:time.Now().Unix(),
		//NotBefore:0,
		Subject:"APP",
	})
	mySigningKey:=[]byte("dKe580xiDDtIhttpClnsBXigizsR9xys")
	tokenString, err := jwtToken.SignedString(mySigningKey)

	if err != nil {
		errResp:=&responses.ErrResponseJson{
			StatusCode:0,
			ErrorMsg:err.Error(),
		}
		ctx.JSON(errResp)
		return
	}
	resp:=map[string]interface{}{
		"status":"ok",
		"status_code":0,
		"data":tokenString,
	}
	ctx.JSON(resp)
}

func ThirdLoginHandler(ctx iris.Context){

}


