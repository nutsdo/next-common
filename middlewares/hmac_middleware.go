package middlewares

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/kataras/iris"
	"github.com/nutsdo/go-next/responses"
	"sort"
	"strings"
)

func KongHmacSignGenerator(username, secret, date string) string {

	var SignStr strings.Builder

	SignStr.Write([]byte("date: "))
	SignStr.Write([]byte(date))

	hmac := hmac.New(sha1.New, []byte(secret))
	hmac.Write([]byte(SignStr.String()))
	sign:=hmac.Sum(nil)

	signStr := fmt.Sprintf("hmac username=\"%s\",algorithm=\"hmac-sha1\",headers=\"date\",signature=\"%s\"",
		username,base64.StdEncoding.EncodeToString(sign))
	return signStr
}

func HmacHandler(ctx iris.Context) {
	fmt.Println(ctx.URLParams())
	params := ctx.URLParams()

	var paramsArray []string
	//参数map 转换位数组
	for k,v := range params{
		if k!="sign" || v==""{
			paramsArray = append(paramsArray, k)
		}
	}
	sort.Strings(paramsArray)

	//拼接参数串
	var str strings.Builder
	for i,v := range paramsArray {
		str.WriteString(v)
		str.WriteString(params[v])
		fmt.Println(i," ",v)
	}
	fmt.Println(str.String())
	key:=[]byte("123456")
	hmac := hmac.New(sha1.New, key)
	hmac.Write([]byte(str.String()))
	sign:=fmt.Sprintf("%x", hmac.Sum(nil))

	fmt.Println("参数签名:", params["sign"])

	if params["sign"] != sign {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(responses.ErrResponse(401,104001,"签名错误"))
		return
	}
	ctx.Next()
}
