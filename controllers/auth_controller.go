package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/nutsdo/go-next/middlewares"
	"io/ioutil"
	"net/http"
	"time"
)

func TestKongHmacSign(ctx iris.Context){
	requrl := "https://gwapi.gsget.cn/hmactest"
	fmt.Println(requrl)
	username := "gengshengapp001"
	secret := "cjH3B3lXbhplw9PsxpGx4UUleo7mSANd"
	currentTime := time.Now().UTC().Format(http.TimeFormat)
	proxyauthorization := middlewares.KongHmacSignGenerator(username,secret,currentTime)

	req,err := http.NewRequest(http.MethodGet,requrl,nil)
	req.Header = map[string][]string{
		"Date" : []string{currentTime},
		"Proxy-Authorization":[]string{proxyauthorization},
	}
	if err!=nil {
		fmt.Println(err)
	}
	client := http.DefaultClient
	resp,err:=client.Do(req)
	if err !=nil{
		fmt.Println(err)
	}
	body,err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	if err!=nil{
		fmt.Println(err)
	}

	ctx.JSON(string(body))
}
